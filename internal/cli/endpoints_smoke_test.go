// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// endpointCase describes one generated command discovered by walking the
// Cobra tree. The pp:* annotations are emitted by the generator on every
// endpoint-backed command and are the contract this smoke test verifies.
type endpointCase struct {
	cmdPath  []string // e.g. ["device-script-assignments", "create"]
	endpoint string   // pp:endpoint, e.g. "device-script-assignments.create"
	method   string   // pp:method
	path     string   // pp:path with {param} placeholders
}

// collectEndpointCases walks the command tree depth-first and returns one
// case per command carrying a pp:endpoint annotation. Hand-written compound
// commands (sync, search, doctor, ...) carry no annotation and are skipped.
func collectEndpointCases(cmd *cobra.Command, prefix []string, out *[]endpointCase) {
	if ep, ok := cmd.Annotations["pp:endpoint"]; ok {
		*out = append(*out, endpointCase{
			cmdPath:  append(append([]string{}, prefix...), cmd.Name()),
			endpoint: ep,
			method:   cmd.Annotations["pp:method"],
			path:     cmd.Annotations["pp:path"],
		})
	}
	for _, sub := range cmd.Commands() {
		collectEndpointCases(sub, append(append([]string{}, prefix...), cmd.Name()), out)
	}
}

// requiredPositionals counts <placeholder> tokens in a command's Use line.
// Generated commands take path parameters as required positional args.
func requiredPositionals(use string) int {
	return strings.Count(use, "<")
}

// pathPattern converts a pp:path template into an anchored regex that
// matches the template with each {param} segment replaced by any non-empty,
// slash-free value. The default BasePath /api/v2 is prepended because the
// client joins BaseURL + BasePath + path.
func pathPattern(t *testing.T, ppPath string) *regexp.Regexp {
	t.Helper()
	var b strings.Builder
	b.WriteString(`^/api/v2`)
	rest := ppPath
	for {
		open := strings.Index(rest, "{")
		if open < 0 {
			b.WriteString(regexp.QuoteMeta(rest))
			break
		}
		closing := strings.Index(rest, "}")
		if closing < open {
			t.Fatalf("malformed pp:path template %q", ppPath)
		}
		b.WriteString(regexp.QuoteMeta(rest[:open]))
		b.WriteString(`[^/]+`)
		rest = rest[closing+1:]
	}
	b.WriteString(`$`)
	return regexp.MustCompile(b.String())
}

// buildArgs assembles the argv for one endpoint command: the command path,
// a dummy value per required positional (path params), and a dummy value
// for every locally-defined string flag so the generated required-flag
// checks pass. The dummy `"x"` is a valid JSON string, so it satisfies both
// raw string flags and flags the generated code round-trips through
// json.Unmarshal.
func buildArgs(c endpointCase, leaf *cobra.Command) []string {
	args := append([]string{}, c.cmdPath[1:]...) // drop the root segment

	for i := 0; i < requiredPositionals(leaf.Use); i++ {
		args = append(args, "smoke-test-value")
	}

	leaf.Flags().VisitAll(func(f *pflag.Flag) {
		if f.Name == "stdin" {
			return
		}
		switch f.Value.Type() {
		case "string":
			args = append(args, "--"+f.Name, `"x"`)
		case "bool":
			args = append(args, "--"+f.Name+"=true")
		case "int", "int32", "int64", "float64":
			args = append(args, "--"+f.Name, "1")
		}
	})

	// Root flags: machine output, skip mutation confirmation, bypass the
	// response cache, force live API reads past the local-store resolver,
	// and lift the rate limiter so 200+ sequential calls stay fast.
	return append(args,
		"--json", "--yes", "--no-cache", "--data-source", "live",
		"--rate-limit", "1000",
	)
}

// TestEndpointSmoke executes every generated endpoint command in-process
// against a local mock API and asserts each one issues exactly the HTTP
// method and path its pp:* annotations promise. It needs no credentials and
// performs no network I/O beyond the loopback server.
func TestEndpointSmoke(t *testing.T) {
	type recorded struct {
		method string
		path   string
	}
	var mu sync.Mutex
	var last *recorded

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		last = &recorded{method: r.Method, path: r.URL.Path}
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data": []}`))
	}))
	defer srv.Close()

	t.Setenv("ADDIGY_CONFIG", t.TempDir()+"/config.toml") // never touch the user's config
	t.Setenv("ADDIGY_BASE_URL", srv.URL)
	t.Setenv("ADDIGY_DOCUMENTATION_API_KEY", "smoke-test-key")

	var cases []endpointCase
	collectEndpointCases(RootCmd(), nil, &cases)
	if len(cases) < 150 {
		t.Fatalf("expected 150+ generated endpoint commands, found %d — tree walk is broken", len(cases))
	}
	t.Logf("discovered %d endpoint commands", len(cases))

	for _, tc := range cases {
		t.Run(strings.Join(tc.cmdPath[1:], "/"), func(t *testing.T) {
			root := RootCmd()
			leaf, _, err := root.Find(tc.cmdPath[1:])
			if err != nil {
				t.Fatalf("finding command: %v", err)
			}

			mu.Lock()
			last = nil
			mu.Unlock()

			root.SetArgs(buildArgs(tc, leaf))
			root.SetOut(io.Discard)
			root.SetErr(io.Discard)
			if err := root.Execute(); err != nil {
				t.Fatalf("command failed: %v", err)
			}

			mu.Lock()
			got := last
			mu.Unlock()
			if got == nil {
				t.Fatal("command succeeded but sent no HTTP request")
			}
			if got.method != tc.method {
				t.Errorf("method = %s, want %s (pp:method)", got.method, tc.method)
			}
			if re := pathPattern(t, tc.path); !re.MatchString(got.path) {
				t.Errorf("path = %s, want match for template %s", got.path, tc.path)
			}
		})
	}
}

// TestManifestCoverage asserts the command tree and tools-manifest.json
// describe the same endpoint set in both directions, so manifest drift
// (generator updated one but not the other) fails CI instead of shipping
// stale endpoint IDs to MCP agents.
func TestManifestCoverage(t *testing.T) {
	raw, err := os.ReadFile("../../tools-manifest.json")
	if err != nil {
		t.Fatalf("reading tools-manifest.json: %v", err)
	}
	var manifest struct {
		Tools []struct {
			Name   string `json:"name"`
			Method string `json:"method"`
			Path   string `json:"path"`
		} `json:"tools"`
	}
	if err := json.Unmarshal(raw, &manifest); err != nil {
		t.Fatalf("parsing tools-manifest.json: %v", err)
	}

	manifestSet := make(map[string]bool, len(manifest.Tools))
	for _, tool := range manifest.Tools {
		manifestSet[tool.Name] = true
	}

	var cases []endpointCase
	collectEndpointCases(RootCmd(), nil, &cases)
	// A command is known to the manifest under either its command path
	// joined with underscores (`o billing update-o` ↔ "o_billing_update-o")
	// or its pp:endpoint with dots flattened — single-level commands like
	// `addigy-cli files` carry endpoint "files.create" which the manifest
	// names "files_create".
	treeSet := make(map[string]bool, len(cases)*2)
	for _, c := range cases {
		treeSet[strings.Join(c.cmdPath[1:], "_")] = true
		treeSet[strings.ReplaceAll(c.endpoint, ".", "_")] = true
	}

	for name := range manifestSet {
		if !treeSet[name] {
			t.Errorf("manifest tool %q has no command in the Cobra tree", name)
		}
	}
	for _, c := range cases {
		pathName := strings.Join(c.cmdPath[1:], "_")
		endpointName := strings.ReplaceAll(c.endpoint, ".", "_")
		if !manifestSet[pathName] && !manifestSet[endpointName] {
			t.Errorf("command %q (endpoint %s) is missing from tools-manifest.json", pathName, c.endpoint)
		}
	}
}
