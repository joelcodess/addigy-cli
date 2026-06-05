// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joelcodess/addigy-cli/internal/cli"
	"github.com/joelcodess/addigy-cli/internal/client"
	"github.com/joelcodess/addigy-cli/internal/config"
	"github.com/joelcodess/addigy-cli/internal/mcp/cobratree"
	"github.com/joelcodess/addigy-cli/internal/store"
	mcplib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all API operations as MCP tools.
func RegisterTools(s *server.MCPServer) {
	// Code-orchestration mode — the full surface is covered by two tools
	// (<api>_search + <api>_execute). Endpoint-mirror tools are suppressed.
	RegisterCodeOrchestrationTools(s)
	// Search tool — faster than iterating list endpoints for finding specific items
	s.AddTool(
		mcplib.NewTool("search",
			mcplib.WithDescription("Full-text search across all synced data. Faster than paginating list endpoints. Requires sync first."),
			mcplib.WithString("query", mcplib.Required(), mcplib.Description("Search query (supports FTS5 syntax: AND, OR, NOT, quotes for phrases)")),
			mcplib.WithNumber("limit", mcplib.Description("Max results (default 25)")),
			mcplib.WithReadOnlyHintAnnotation(true),
			mcplib.WithDestructiveHintAnnotation(false),
		),
		handleSearch,
	)
	// SQL tool — ad-hoc analysis on synced data without API calls
	s.AddTool(
		mcplib.NewTool("sql",
			mcplib.WithDescription("Run read-only SQL against local database. Use for ad-hoc analysis, aggregations, and joins across synced resources. Requires sync first."),
			mcplib.WithString("query", mcplib.Required(), mcplib.Description("SQL query (SELECT or WITH...SELECT). Tables match resource names.")),
			mcplib.WithReadOnlyHintAnnotation(true),
			mcplib.WithDestructiveHintAnnotation(false),
		),
		handleSQL,
	)

	// Context tool — front-loaded domain knowledge for agents.
	// Call this first to understand the API taxonomy, query patterns, and capabilities.
	s.AddTool(
		mcplib.NewTool("context",
			mcplib.WithDescription("Get API domain context: resource taxonomy, auth requirements, query tips, and unique capabilities. Call this first."),
			mcplib.WithReadOnlyHintAnnotation(true),
			mcplib.WithDestructiveHintAnnotation(false),
		),
		handleContext,
	)

	// Runtime Cobra-tree mirror — exposes every user-facing command that is
	// not already covered by a typed endpoint or framework MCP tool.
	cobratree.RegisterAll(s, cli.RootCmd(), cobratree.SiblingCLIPath)
}

func newMCPClient() (*client.Client, error) {
	home, _ := os.UserHomeDir()
	cfgPath := filepath.Join(home, ".config", "addigy-cli", "config.toml")
	cfg, err := config.Load(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}
	c := client.New(cfg, 30*time.Second, 0)
	// Agents calling through MCP need fresh data every call. The on-disk
	// response cache survives across MCP server invocations, so a
	// DELETE/PATCH followed by a GET would otherwise return the
	// pre-mutation snapshot for up to the cache TTL. The interactive CLI
	// constructs its own client and is unaffected.
	c.NoCache = true
	return c, nil
}

func dbPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".local", "share", "addigy-cli", "data.db")
}

// Note: MCP tools use their own dbPath() because they are in a separate package (main, not cli).
// The CLI's defaultDBPath() in the cli package uses the same canonical path.

func handleSearch(ctx context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	args := req.GetArguments()
	query, ok := args["query"].(string)
	if !ok || query == "" {
		return mcplib.NewToolResultError("query is required"), nil
	}

	limit := 25
	if v, ok := args["limit"].(float64); ok && v > 0 {
		limit = int(v)
	}

	db, err := store.OpenReadOnly(dbPath())
	if err != nil {
		return mcplib.NewToolResultError(fmt.Sprintf("opening database: %v", err)), nil
	}
	defer db.Close()

	results, err := db.Search(query, limit)
	if err != nil {
		return mcplib.NewToolResultError(fmt.Sprintf("search failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(results, "", "  ")
	return mcplib.NewToolResultText(string(data)), nil
}

// validateReadOnlyQuery gates the MCP sql tool. The agent contract advertised
// to the host is ReadOnlyHintAnnotation(true); a false annotation on a
// mutating tool lets MCP hosts auto-approve writes and is treated as a real
// bug per the project's agent-native security model.
//
// The gate is an allowlist (SELECT or WITH only) applied AFTER stripping the
// leading whitespace, line comments, block comments, and semicolons that
// SQLite itself ignores before parsing. A naive HasPrefix check on a
// keyword blocklist is bypassable by prefixing the dangerous statement with
// "/* x */" or "-- x\n" — TrimSpace strips outer whitespace but does not
// understand SQL comment syntax. Combined with the empirical fact that
// modernc.org/sqlite's mode=ro does NOT block VACUUM INTO (writes a snapshot
// to a new file) or ATTACH DATABASE (opens a separate writable handle),
// such a bypass produces silent exfiltration to an attacker-chosen path.
//
// SELECT and WITH are the only allowed leading keywords. WITH supports
// SELECT-form CTEs; CTE-wrapped writes ("WITH x AS (...) INSERT ...") are
// caught by OpenReadOnly's mode=ro one layer down. PRAGMA, ATTACH, VACUUM,
// and every other DDL/DML keyword fail at this gate before reaching SQLite.
func validateReadOnlyQuery(query string) error {
	upper := strings.ToUpper(stripLeadingSQLNoise(query))
	if !strings.HasPrefix(upper, "SELECT") && !strings.HasPrefix(upper, "WITH") {
		return fmt.Errorf("only SELECT queries are allowed")
	}
	return nil
}

// stripLeadingSQLNoise removes leading whitespace, SQL line comments
// (-- to end of line), block comments (/* ... */), and statement
// separators (;) from query. SQLite skips these before parsing the first
// keyword, so a security gate that does not strip them mismatches what the
// driver actually executes.
func stripLeadingSQLNoise(query string) string {
	for {
		query = strings.TrimLeft(query, " \t\r\n;")
		switch {
		case strings.HasPrefix(query, "--"):
			if idx := strings.IndexByte(query, '\n'); idx >= 0 {
				query = query[idx+1:]
				continue
			}
			return ""
		case strings.HasPrefix(query, "/*"):
			if idx := strings.Index(query[2:], "*/"); idx >= 0 {
				query = query[2+idx+2:]
				continue
			}
			return ""
		default:
			return query
		}
	}
}

func handleSQL(ctx context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	args := req.GetArguments()
	query, ok := args["query"].(string)
	if !ok || query == "" {
		return mcplib.NewToolResultError("query is required"), nil
	}

	if err := validateReadOnlyQuery(query); err != nil {
		return mcplib.NewToolResultError(err.Error()), nil
	}

	db, err := store.OpenReadOnly(dbPath())
	if err != nil {
		return mcplib.NewToolResultError(fmt.Sprintf("opening database: %v", err)), nil
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return mcplib.NewToolResultError(fmt.Sprintf("query failed: %v", err)), nil
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	var results []map[string]any
	for rows.Next() {
		values := make([]any, len(cols))
		ptrs := make([]any, len(cols))
		for i := range values {
			ptrs[i] = &values[i]
		}
		rows.Scan(ptrs...)
		row := make(map[string]any)
		for i, col := range cols {
			row[col] = values[i]
		}
		results = append(results, row)
	}

	data, _ := json.MarshalIndent(results, "", "  ")
	return mcplib.NewToolResultText(string(data)), nil
}

func handleContext(_ context.Context, _ mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	ctx := map[string]any{
		"api":         "addigy",
		"description": "Every Addigy v2 endpoint, plus a local fleet mirror with compound queries the web UI can't answer.",
		"archetype":   "device-management",
		"tool_count":  203,
		// tool_surface tells agents which surface a capability lives on.
		"tool_surface": "MCP exposes typed endpoint tools plus a runtime mirror of user-facing CLI commands. Endpoint tools keep typed schemas; command-mirror tools shell out to the companion addigy-cli binary.",
		"auth": map[string]any{
			"type": "api_key",
			"env_vars": []map[string]any{
				{
					"name":        "ADDIGY_DOCUMENTATION_API_KEY",
					"kind":        "per_call",
					"required":    true,
					"sensitive":   true,
					"description": "Set to your API credential.",
				},
			},
			"key_url": "https://app.addigy.com/integrations",
		},
		"resources": []map[string]any{
			{
				"name":        "assets",
				"description": "Manage assets",
				"endpoints":   []string{"create", "create-default", "create-default-2", "create-default-3"},
				"searchable":  true,
			},
			{
				"name":        "configuration",
				"description": "Manage configuration",
				"endpoints":   []string{"list"},
				"syncable":    true,
			},
			{
				"name":        "device-script-assignments",
				"description": "Manage device script assignments",
				"endpoints":   []string{"create", "delete", "list"},
				"syncable":    true,
				"searchable":  true,
			},
			{
				"name":        "devices",
				"description": "Manage devices",
				"endpoints":   []string{"create"},
				"searchable":  true,
			},
			{
				"name":        "facts",
				"description": "Manage facts",
				"endpoints":   []string{"create", "create-custom", "create-custom-2", "delete", "delete-custom", "list", "update"},
				"syncable":    true,
				"searchable":  true,
			},
			{
				"name":        "feature-betas",
				"description": "Manage feature betas",
				"endpoints":   []string{"create", "delete", "list"},
				"syncable":    true,
				"searchable":  true,
			},
			{
				"name":        "files",
				"description": "Manage files",
				"endpoints":   []string{"create"},
			},
			{
				"name":        "maintenance",
				"description": "Manage maintenance",
				"endpoints":   []string{"create", "create-policy", "create-query", "create-staged", "create-staged-2", "create-staged-3", "delete", "delete-policy", "delete-staged", "update", "update-staged"},
				"searchable":  true,
			},
			{
				"name":        "managed-app-configurations",
				"description": "Manage managed app configurations",
				"endpoints":   []string{"create", "delete", "list"},
				"searchable":  true,
			},
			{
				"name":        "mdm",
				"description": "Manage mdm",
				"endpoints":   []string{"create", "create-commands", "create-configurations", "create-configurations-2", "create-configurations-3", "create-devices", "create-profiles", "delete", "delete-configurations", "delete-configurations-2", "get", "get-configurations", "get-configurations-2", "get-devices", "list", "list-commands", "list-configurations", "list-configurations-2", "list-configurations-3", "update"},
				"syncable":    true,
				"searchable":  true,
			},
			{
				"name":        "monitoring",
				"description": "Manage monitoring",
				"endpoints":   []string{"create", "create-policy", "create-query", "delete", "delete-policy", "list", "update"},
				"syncable":    true,
				"searchable":  true,
			},
			{
				"name":        "o",
				"description": "Manage o",
				"endpoints":   []string{"create"},
			},
			{
				"name":        "oa",
				"description": "Manage oa",
				"endpoints":   []string{"create", "create-ade", "create-appsandbooks", "create-compliancerules", "create-devices", "create-files", "create-identity", "create-installedapps", "create-integrations", "create-monitoring", "create-policies", "create-policies-2", "create-prebuiltapps", "create-reports", "create-variables", "create-webhooks", "create-webhooks-2", "create-webhooks-3", "list", "list-benchmarks", "list-compliancerules", "list-compliancerules-2", "list-integrations", "list-reports", "list-reports-2", "list-selfservice"},
				"syncable":    true,
				"searchable":  true,
			},
			{
				"name":        "prebuilt-apps",
				"description": "Manage prebuilt apps",
				"endpoints":   []string{"create", "create-prebuiltapps", "create-prebuiltapps-2", "create-prebuiltapps-3", "delete", "delete-prebuiltapps", "get", "get-prebuiltapps", "update", "update-prebuiltapps"},
				"searchable":  true,
			},
			{
				"name":        "self-service-configurations",
				"description": "Manage self service configurations",
				"endpoints":   []string{"create"},
				"searchable":  true,
			},
			{
				"name":        "static-fields",
				"description": "Manage static fields",
				"endpoints":   []string{"create", "create-staticfields", "delete", "list", "list-staticfields", "update"},
				"syncable":    true,
				"searchable":  true,
			},
			{
				"name":        "system-events",
				"description": "Manage system events",
				"endpoints":   []string{"create", "create-systemevents"},
				"searchable":  true,
			},
			{
				"name":        "system-updates",
				"description": "Manage system updates",
				"endpoints":   []string{"create", "create-systemupdates", "create-systemupdates-2", "create-systemupdates-3", "create-systemupdates-4", "create-systemupdates-5", "create-systemupdates-6", "create-systemupdates-7", "create-systemupdates-8", "create-systemupdates-9", "list", "list-systemupdates", "list-systemupdates-2", "list-systemupdates-3", "list-systemupdates-4", "list-systemupdates-5", "list-systemupdates-6"},
				"syncable":    true,
				"searchable":  true,
			},
			{
				"name":        "users",
				"description": "Manage users",
				"endpoints":   []string{"delete", "update"},
				"searchable":  true,
			},
		},
		"query_tips": []string{
			"Pagination uses cursor-based paging. Pass after parameter for subsequent pages.",
			"Control page size with the per_page parameter (default 100).",
			"Use the sql tool for ad-hoc analysis on synced data. Run sync first to populate the local database.",
			"Use the search tool for full-text search across all synced resources. Faster than iterating list endpoints.",
			"Prefer sql/search over repeated API calls when the data is already synced.",
		},
		// Command-mirror capabilities are exposed through MCP by shelling out
		// to the companion CLI binary.
		"command_mirror_capabilities": []map[string]string{
			{"name": "Stale-device aggregation", "command": "devices stale", "description": "List devices whose last check-in is older than N days, with optional policy and OS filters.", "rationale": "Aggregating last-checkin across the fleet against the live API would burn the 1,000-req/10s budget; this runs as a...", "via": "mcp-command-mirror"},
			{"name": "Cross-fleet compliance audit", "command": "compliance", "description": "Surface devices whose assigned policy rules are unmet, joined against current device facts.", "rationale": "No API endpoint evaluates rule compliance per device — the UI does it client-side, one device at a time. This...", "via": "mcp-command-mirror"},
			{"name": "Smart Software rollout tracker", "command": "rollout", "description": "Per-device install state for one Smart Software item across the assigned fleet, with success/pending/failed counts.", "rationale": "The UI lacks a bulk rollout view; this aggregates per-device state across the fleet in one shot via the local mirror.", "via": "mcp-command-mirror"},
			{"name": "Custom-facts fleet search", "command": "facts search", "description": "FTS5 across mirrored device facts; --group-by value returns a histogram of values per fact.", "rationale": "The live API has no full-text fact search and no bulk fact aggregation; this is local-only leverage.", "via": "mcp-command-mirror"},
			{"name": "Device-to-device diff", "command": "devices diff", "description": "Set differences across facts, applications, policies, and Smart-Software install state between two devices.", "rationale": "No single API call diffs two devices; today this is a manual two-tabs eyeball.", "via": "mcp-command-mirror"},
			{"name": "Fleet drift since timestamp", "command": "drift", "description": "Diffs the mirror's current rows against the prior snapshot for any entity (devices, facts, policies, software).", "rationale": "Temporal diffs are impossible against a live API that does not retain history; the mirror's snapshot table makes...", "via": "mcp-command-mirror"},
			{"name": "Policy coverage + sync health", "command": "policy-coverage", "description": "Per-policy device counts joined to last-checkin so the user sees both coverage and freshness in one view.", "rationale": "Aggregates over the fleet without re-issuing /devices per policy; impossible to compute live within rate budget for...", "via": "mcp-command-mirror"},
			{"name": "Fleet summary one-shot", "command": "fleet-summary", "description": "Single command emitting device count, stale fraction, alert count, MDM queue depth, Smart-Software pending count,...", "rationale": "Six different API resources collapsed into one local aggregate so triage agents can ask 'what is the state of this...", "via": "mcp-command-mirror"},
		},
		"playbook": []map[string]string{
			{"topic": "Stale-device aggregation", "insight": "Aggregating last-checkin across the fleet against the live API would burn the 1,000-req/10s budget; this runs as a single local SQLite query."},
			{"topic": "Cross-fleet compliance audit", "insight": "No API endpoint evaluates rule compliance per device — the UI does it client-side, one device at a time. This joins devices x policy_rules x device_facts in local SQLite."},
			{"topic": "Smart Software rollout tracker", "insight": "The UI lacks a bulk rollout view; this aggregates per-device state across the fleet in one shot via the local mirror."},
			{"topic": "Custom-facts fleet search", "insight": "The live API has no full-text fact search and no bulk fact aggregation; this is local-only leverage."},
			{"topic": "Device-to-device diff", "insight": "No single API call diffs two devices; today this is a manual two-tabs eyeball."},
			{"topic": "Fleet drift since timestamp", "insight": "Temporal diffs are impossible against a live API that does not retain history; the mirror's snapshot table makes this a one-line query."},
			{"topic": "Policy coverage + sync health", "insight": "Aggregates over the fleet without re-issuing /devices per policy; impossible to compute live within rate budget for fleets > a few hundred devices."},
			{"topic": "Fleet summary one-shot", "insight": "Six different API resources collapsed into one local aggregate so triage agents can ask 'what is the state of this fleet right now?' in one call."},
		},
	}
	data, _ := json.MarshalIndent(ctx, "", "  ")
	return mcplib.NewToolResultText(string(data)), nil
}

// RegisterNovelFeatureTools is kept as a compatibility no-op for older MCP
// mains. New generated mains call RegisterTools only; RegisterTools now
// includes the runtime Cobra-tree mirror.
func RegisterNovelFeatureTools(s *server.MCPServer) {
	_ = s
}
