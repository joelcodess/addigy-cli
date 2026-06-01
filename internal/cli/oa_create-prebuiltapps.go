// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newOaCreatePrebuiltappsCmd(flags *rootFlags) *cobra.Command {
	var bodyIds string
	var bodyPage int
	var bodyPerPage int
	var bodyPrebuiltAppIds string
	var bodyPrebuiltAppVersionIds string
	var bodySortDirection string
	var bodySortField string
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "create-prebuiltapps",
		Short:       "Query Prebuilt Apps Configurations",
		Example:     "  addigy-cli oa create-prebuiltapps",
		Annotations: map[string]string{"pp:endpoint": "oa.create-prebuiltapps", "pp:method": "POST", "pp:path": "/oa/prebuilt-apps/configurations/query"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/oa/prebuilt-apps/configurations/query"
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
				if bodyIds != "" {
					var parsedIds any
					if err := json.Unmarshal([]byte(bodyIds), &parsedIds); err != nil {
						return fmt.Errorf("parsing --ids JSON: %w", err)
					}
					body["ids"] = parsedIds
				}
				if bodyPage != 0 {
					body["page"] = bodyPage
				}
				if bodyPerPage != 0 {
					body["per_page"] = bodyPerPage
				}
				if bodyPrebuiltAppIds != "" {
					var parsedPrebuiltAppIds any
					if err := json.Unmarshal([]byte(bodyPrebuiltAppIds), &parsedPrebuiltAppIds); err != nil {
						return fmt.Errorf("parsing --prebuilt-app-ids JSON: %w", err)
					}
					body["prebuilt_app_ids"] = parsedPrebuiltAppIds
				}
				if bodyPrebuiltAppVersionIds != "" {
					var parsedPrebuiltAppVersionIds any
					if err := json.Unmarshal([]byte(bodyPrebuiltAppVersionIds), &parsedPrebuiltAppVersionIds); err != nil {
						return fmt.Errorf("parsing --prebuilt-app-version-ids JSON: %w", err)
					}
					body["prebuilt_app_version_ids"] = parsedPrebuiltAppVersionIds
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
	cmd.Flags().StringVar(&bodyIds, "ids", "", "Ids")
	cmd.Flags().IntVar(&bodyPage, "page", 1, "Page")
	cmd.Flags().IntVar(&bodyPerPage, "per-page", 50, "Per page")
	cmd.Flags().StringVar(&bodyPrebuiltAppIds, "prebuilt-app-ids", "", "Prebuilt app ids")
	cmd.Flags().StringVar(&bodyPrebuiltAppVersionIds, "prebuilt-app-version-ids", "", "Prebuilt app version ids")
	cmd.Flags().StringVar(&bodySortDirection, "sort-direction", "asc", "Sort direction")
	cmd.Flags().StringVar(&bodySortField, "sort-field", "id", "Sort field")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
