// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package mcp

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/joelcodess/addigy-cli/internal/client"
	"github.com/joelcodess/addigy-cli/internal/config"
)

func testClient(baseURL string) *client.Client {
	c := client.New(&config.Config{BaseURL: baseURL, BasePath: ""}, 5*time.Second, 0)
	c.AssumeYes = true // MCP/agent context: bypass the interactive confirm guardrail
	c.NoCache = true
	return c
}

// TestExecEndpointRequest_DeleteSendsQueryParams guards the MCP addigy_execute
// regression: query-param DELETEs must send their target id/key in the query
// string. A bare c.Delete(path) would drop it and the API would get a
// target-less DELETE.
func TestExecEndpointRequest_DeleteSendsQueryParams(t *testing.T) {
	var gotMethod, gotQuery string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotQuery = r.Method, r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	_, err := execEndpointRequest(testClient(srv.URL), "DELETE", "/maintenance",
		map[string]string{"id": "abc123"}, map[string]any{"id": "abc123"})
	if err != nil {
		t.Fatalf("execEndpointRequest DELETE: %v", err)
	}
	if gotMethod != http.MethodDelete {
		t.Errorf("method = %q, want DELETE", gotMethod)
	}
	if gotQuery != "id=abc123" {
		t.Errorf("query = %q, want id=abc123 (the target id must reach the API, not be dropped)", gotQuery)
	}
}

// TestExecEndpointRequest_PostSendsBody confirms the body-carrying methods send
// their params in the request body, not the query string.
func TestExecEndpointRequest_PostSendsBody(t *testing.T) {
	var gotBody string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b := make([]byte, r.ContentLength)
		_, _ = r.Body.Read(b)
		gotBody = string(b)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	_, err := execEndpointRequest(testClient(srv.URL), "POST", "/facts/custom",
		map[string]string{}, map[string]any{"name": "x"})
	if err != nil {
		t.Fatalf("execEndpointRequest POST: %v", err)
	}
	if gotBody == "" || gotBody == "{}" {
		t.Errorf("POST body = %q, want the params marshaled into the body", gotBody)
	}
}
