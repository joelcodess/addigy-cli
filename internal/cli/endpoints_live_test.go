// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// TestEndpointLiveSmoke exercises the read-only (GET) endpoint commands
// against the real Addigy API using whatever credentials the local CLI is
// already configured with. It never issues a mutation: only commands whose
// pp:method is GET are run, and POST "query" endpoints are deliberately
// excluded so the test cannot write anything no matter what the server does.
//
// Opt-in only: set ADDIGY_LIVE_SMOKE=1 to enable, otherwise it skips. Set
// ADDIGY_LIVE_ORG_ID to also cover /o/{organization_id}/... endpoints.
// Commands needing other path params (device_uuid, id, ...) are skipped —
// they require real resource IDs that this test does not discover.
//
// Outcome policy per command — only unambiguous CLI-side bugs fail:
//   - success                  → pass
//   - "required flag" refusal  → skip (needs a real resource ID we can't invent)
//   - 404 / not-found          → pass (valid request shape; resource absent for org)
//   - other HTTP 4xx/5xx       → logged + tolerated (org-dependent: unconfigured
//     features, server-side required params, upstream bugs — e.g. billing
//     endpoints 500 for orgs with no billing data)
//   - 401 unauthorized, transport errors, panics, anything else → fail
func TestEndpointLiveSmoke(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live smoke in -short mode")
	}
	if os.Getenv("ADDIGY_LIVE_SMOKE") != "1" {
		t.Skip("set ADDIGY_LIVE_SMOKE=1 to run against the real API")
	}
	orgID := os.Getenv("ADDIGY_LIVE_ORG_ID")

	var cases []endpointCase
	collectEndpointCases(RootCmd(), nil, &cases)

	var ran, passed, notFound, tolerated, needsID, skipped int
	for _, tc := range cases {
		if tc.method != "GET" {
			continue
		}
		args, ok := liveArgs(tc, orgID)
		if !ok {
			skipped++
			continue
		}

		t.Run(strings.Join(tc.cmdPath[1:], "/"), func(t *testing.T) {
			ran++
			root := RootCmd()
			var out, errBuf bytes.Buffer
			root.SetArgs(args)
			root.SetOut(&out)
			root.SetErr(&errBuf)
			err := root.Execute()
			switch {
			case err == nil:
				passed++
			case strings.HasPrefix(err.Error(), "required flag"):
				needsID++
				t.Skipf("needs a real resource ID: %v", err)
			case strings.Contains(err.Error(), "HTTP 401"):
				t.Errorf("authentication failed — key invalid or expired: %v", err)
			case ExitCode(err) == 3:
				notFound++
				t.Logf("tolerated not-found: %v", err)
			case strings.Contains(err.Error(), "returned HTTP "):
				tolerated++
				t.Logf("tolerated org-dependent API response: %v", err)
			default:
				t.Errorf("live call failed: %v", err)
			}
		})
	}
	t.Logf("live smoke: %d GET commands ran (%d ok, %d not-found, %d org-dependent 4xx/5xx, %d need real IDs), %d skipped (unfillable path params)",
		ran, passed, notFound, tolerated, needsID, skipped)
	if ran == 0 {
		t.Fatal("no live commands ran — gating or enumeration is broken")
	}
}

// liveArgs builds argv for a read-only live call. It fills {organization_id}
// from orgID and refuses (ok=false) any command whose path needs a param we
// cannot supply. Optional query flags are left at their defaults — the point
// is to validate real request/response handling, not to exercise filters.
func liveArgs(c endpointCase, orgID string) (args []string, ok bool) {
	args = append([]string{}, c.cmdPath[1:]...)

	params := pathParams(c.path)
	for _, p := range params {
		if p == "organization_id" && orgID != "" {
			continue
		}
		return nil, false
	}
	// Generated commands take path params as positional args in template
	// order; organization_id is the only one we ever fill here.
	for range params {
		args = append(args, orgID)
	}

	return append(args, "--json", "--no-cache", "--data-source", "live"), true
}

// pathParams returns the {param} names in a pp:path template, in order.
func pathParams(tmpl string) []string {
	var out []string
	rest := tmpl
	for {
		open := strings.Index(rest, "{")
		if open < 0 {
			return out
		}
		closing := strings.Index(rest[open:], "}")
		if closing < 0 {
			return out
		}
		out = append(out, rest[open+1:open+closing])
		rest = rest[open+closing+1:]
	}
}
