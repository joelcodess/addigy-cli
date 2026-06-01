// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newOaCreateMonitoringCmd(flags *rootFlags) *cobra.Command {
	var bodyExcludedIds string
	var bodyIds string
	var bodyLimit int
	var bodyNameContains string
	var bodySkip int
	var bodySortDirection string
	var bodySortField string
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "create-monitoring",
		Short:       "Query for a list of scheduled alerts with pagination.",
		Example:     "  addigy-cli oa create-monitoring",
		Annotations: map[string]string{"pp:endpoint": "oa.create-monitoring", "pp:method": "POST", "pp:path": "/oa/monitoring/query"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/oa/monitoring/query"
			var body map[string]any
			if stdinBody {
				stdinData, err := io.ReadAll(os.Stdin)
				if err != nil {
					return fmt.Errorf("reading stdin: %w", err)
				}
				var jsonBody map[string]any
				if err := json.Unmarshal(stdinData, &jsonBody); err != nil {
					return fmt.Errorf("parsing stdin JSON: %w", err)
				}
				body = jsonBody
			} else {
				body = map[string]any{}
				if bodyExcludedIds != "" {
					var parsedExcludedIds any
					if err := json.Unmarshal([]byte(bodyExcludedIds), &parsedExcludedIds); err != nil {
						return fmt.Errorf("parsing --excluded-ids JSON: %w", err)
					}
					body["excluded_ids"] = parsedExcludedIds
				}
				if bodyIds != "" {
					var parsedIds any
					if err := json.Unmarshal([]byte(bodyIds), &parsedIds); err != nil {
						return fmt.Errorf("parsing --ids JSON: %w", err)
					}
					body["ids"] = parsedIds
				}
				if bodyLimit != 0 {
					body["limit"] = bodyLimit
				}
				if bodyNameContains != "" {
					body["name_contains"] = bodyNameContains
				}
				if bodySkip != 0 {
					body["skip"] = bodySkip
				}
				if bodySortDirection != "" {
					body["sort_direction"] = bodySortDirection
				}
				if bodySortField != "" {
					body["sort_field"] = bodySortField
				}
			}
			data, statusCode, err := c.Post(path, body)
			if err != nil {
				return classifyAPIError(err, flags)
			}
			if wantsHumanTable(cmd.OutOrStdout(), flags) {
				// Check if response contains an array (directly or wrapped in "data")
				var items []map[string]any
				if json.Unmarshal(data, &items) == nil && len(items) > 0 {
					if err := printAutoTable(cmd.OutOrStdout(), items); err != nil {
						fmt.Fprintf(os.Stderr, "warning: table rendering failed, falling back to JSON: %v\n", err)
					} else {
						return nil
					}
				} else {
					var wrapped struct {
						Data []map[string]any `json:"data"`
					}
					if json.Unmarshal(data, &wrapped) == nil && len(wrapped.Data) > 0 {
						if err := printAutoTable(cmd.OutOrStdout(), wrapped.Data); err != nil {
							fmt.Fprintf(os.Stderr, "warning: table rendering failed, falling back to JSON: %v\n", err)
						} else {
							return nil
						}
					}
				}
			}
			if flags.asJSON || (!isTerminal(cmd.OutOrStdout()) && !flags.csv && !flags.quiet && !flags.plain) {
				if flags.quiet {
					return nil
				}
				// Apply --compact and --select to the API response before wrapping.
				// --select wins when both are set: explicit field choice trumps the
				// generic high-gravity allow-list. Otherwise --compact still applies
				// when --agent is on but the user did not name fields.
				filtered := data
				if flags.selectFields != "" {
					filtered = filterFields(filtered, flags.selectFields)
				} else if flags.compact {
					filtered = compactFields(filtered)
				}
				envelope := map[string]any{
					"action":   "post",
					"resource": "oa",
					"path":     path,
					"status":   statusCode,
					"success":  statusCode >= 200 && statusCode < 300,
				}
				if flags.dryRun {
					envelope["dry_run"] = true
					envelope["status"] = 0
					envelope["success"] = false
				}
				if len(filtered) > 0 {
					var parsed any
					if err := json.Unmarshal(filtered, &parsed); err == nil {
						envelope["data"] = parsed
					}
				}
				envelopeJSON, err := json.Marshal(envelope)
				if err != nil {
					return err
				}
				return printOutput(cmd.OutOrStdout(), json.RawMessage(envelopeJSON), true)
			}
			return printOutputWithFlags(cmd.OutOrStdout(), data, flags)
		},
	}
	cmd.Flags().StringVar(&bodyExcludedIds, "excluded-ids", "", "List of alert IDs to exclude")
	cmd.Flags().StringVar(&bodyIds, "ids", "", "List of alert IDs to include")
	cmd.Flags().IntVar(&bodyLimit, "limit", 0, "Maximum number of alerts to return")
	cmd.Flags().StringVar(&bodyNameContains, "name-contains", "", "String to search for in alert names")
	cmd.Flags().IntVar(&bodySkip, "skip", 0, "Number of alerts to skip")
	cmd.Flags().StringVar(&bodySortDirection, "sort-direction", "", "Direction to sort alerts by")
	cmd.Flags().StringVar(&bodySortField, "sort-field", "", "Field to sort alerts by")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
