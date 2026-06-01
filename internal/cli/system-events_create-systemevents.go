// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newSystemEventsCreateSystemeventsCmd(flags *rootFlags) *cobra.Command {
	var bodyFromDateTime string
	var bodyOptionsAggregationIntervalMinutes int
	var bodyOptionsHighlight bool
	var bodyOptionsKeywords bool
	var bodyPage int
	var bodyPerPage int
	var bodyQueries string
	var bodySortDirection string
	var bodyToDateTime string
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "create-systemevents",
		Short:       "Allow to search system events. Permission Required: View System Events.",
		Example:     " addigy-cli system-events create-systemevents",
		Annotations: map[string]string{"pp:endpoint": "system-events.create-systemevents", "pp:method": "POST", "pp:path": "/system-events/search"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/system-events/search"
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
				if bodyFromDateTime != "" {
					body["from_date_time"] = bodyFromDateTime
				}
				{
					nestedOptions := map[string]any{}
					if bodyOptionsAggregationIntervalMinutes != 0 {
						nestedOptions["aggregation_interval_minutes"] = bodyOptionsAggregationIntervalMinutes
					}
					if bodyOptionsHighlight {
						nestedOptions["highlight"] = bodyOptionsHighlight
					}
					if bodyOptionsKeywords {
						nestedOptions["keywords"] = bodyOptionsKeywords
					}
					if len(nestedOptions) > 0 {
						body["options"] = nestedOptions
					}
				}
				if bodyPage != 0 {
					body["page"] = bodyPage
				}
				if bodyPerPage != 0 {
					body["per_page"] = bodyPerPage
				}
				if bodyQueries != "" {
					var parsedQueries any
					if err := json.Unmarshal([]byte(bodyQueries), &parsedQueries); err != nil {
						return fmt.Errorf("parsing --queries JSON: %w", err)
					}
					body["queries"] = parsedQueries
				}
				if bodySortDirection != "" {
					body["sort_direction"] = bodySortDirection
				}
				if bodyToDateTime != "" {
					body["to_date_time"] = bodyToDateTime
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
					"resource": "system-events",
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
	cmd.Flags().StringVar(&bodyFromDateTime, "from-date-time", "", "Search start date")
	cmd.Flags().IntVar(&bodyOptionsAggregationIntervalMinutes, "options-aggregation-interval-minutes", 0, "Aggregation bucket size in minutes")
	cmd.Flags().BoolVar(&bodyOptionsHighlight, "options-highlight", false, "Highlight search matches")
	cmd.Flags().BoolVar(&bodyOptionsKeywords, "options-keywords", false, "Find top search keywords")
	cmd.Flags().IntVar(&bodyPage, "page", 0, "Page")
	cmd.Flags().IntVar(&bodyPerPage, "per-page", 0, "Per page")
	cmd.Flags().StringVar(&bodyQueries, "queries", "", "The search results include the logical AND of the results of each query")
	cmd.Flags().StringVar(&bodySortDirection, "sort-direction", "", "Possible values: ['asc', 'desc']")
	cmd.Flags().StringVar(&bodyToDateTime, "to-date-time", "", "Search end date")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
