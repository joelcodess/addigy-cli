// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newOChildrenGetCmd(flags *rootFlags) *cobra.Command {
	var flagPage string
	var flagPerPage int
	var flagSortField string
	var flagSortDirection string
	var flagSearchString string
	var flagChildOrganizationId string
	var flagAll bool

	cmd := &cobra.Command{
		Use:         "get <organization_id>",
		Short:       "Get a list of child organizations belonging to the provided organization",
		Example:     "  addigy-cli o children get 550e8400-e29b-41d4-a716-446655440000",
		Annotations: map[string]string{"pp:endpoint": "children.get", "pp:method": "GET", "pp:path": "/o/{organization_id}/children/query", "mcp:read-only": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			if cmd.Flags().Changed("sort-direction") {
				allowedSortDirection := []string{"asc", "desc"}
				validSortDirection := false
				for _, v := range allowedSortDirection {
					if flagSortDirection == v {
						validSortDirection = true
						break
					}
				}
				if !validSortDirection {
					fmt.Fprintf(os.Stderr, "warning: --%s %q not in allowed set %v\n", "sort-direction", flagSortDirection, allowedSortDirection)
				}
			}
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/o/{organization_id}/children/query"
			path = replacePathParam(path, "organization_id", args[0])
			data, prov, err := resolvePaginatedRead(cmd.Context(), c, flags, "children", path, map[string]string{
				"page":                  fmt.Sprintf("%v", flagPage),
				"per_page":              fmt.Sprintf("%v", flagPerPage),
				"sort_field":            fmt.Sprintf("%v", flagSortField),
				"sort_direction":        fmt.Sprintf("%v", flagSortDirection),
				"search_string":         fmt.Sprintf("%v", flagSearchString),
				"child_organization_id": fmt.Sprintf("%v", flagChildOrganizationId),
			}, nil, flagAll, "", "", "")
			if err != nil {
				return classifyAPIError(err, flags)
			}
			// Print provenance to stderr for human-facing output
			{
				var countItems []json.RawMessage
				_ = json.Unmarshal(data, &countItems)
				printProvenance(cmd, len(countItems), prov)
			}
			// For JSON output, wrap with provenance envelope before passing through flags.
			// --select wins over --compact when both are set; --compact only runs when
			// no explicit fields were requested. Explicit format flags (--csv, --quiet,
			// --plain) opt out of the auto-JSON path so piped consumers that asked for
			// a non-JSON format reach the standard pipeline below.
			if flags.asJSON || (!isTerminal(cmd.OutOrStdout()) && !flags.csv && !flags.quiet && !flags.plain) {
				filtered := data
				if flags.selectFields != "" {
					filtered = filterFields(filtered, flags.selectFields)
				} else if flags.compact {
					filtered = compactFields(filtered)
				}
				wrapped, wrapErr := wrapWithProvenance(filtered, prov)
				if wrapErr != nil {
					return wrapErr
				}
				return printOutput(cmd.OutOrStdout(), wrapped, true)
			}
			// For all other output modes (table, csv, plain, quiet), use the standard pipeline
			if wantsHumanTable(cmd.OutOrStdout(), flags) {
				var items []map[string]any
				if json.Unmarshal(data, &items) == nil && len(items) > 0 {
					if err := printAutoTable(cmd.OutOrStdout(), items); err != nil {
						return err
					}
					if len(items) >= 25 {
						fmt.Fprintf(os.Stderr, "\nShowing %d results. To narrow: add --limit, --json --select, or filter flags.\n", len(items))
					}
					return nil
				}
			}
			return printOutputWithFlags(cmd.OutOrStdout(), data, flags)
		},
	}
	cmd.Flags().StringVar(&flagPage, "page", "1", "Requested Page")
	cmd.Flags().IntVar(&flagPerPage, "per-page", 50, "Number of items in the response.")
	cmd.Flags().StringVar(&flagSortField, "sort-field", "companyName", "Field used for sorting")
	cmd.Flags().StringVar(&flagSortDirection, "sort-direction", "asc", "Sort Direction (one of: asc, desc)")
	cmd.Flags().StringVar(&flagSearchString, "search-string", "", "Search string for organization name")
	cmd.Flags().StringVar(&flagChildOrganizationId, "child-organization-id", "", "Filter result by orgid for specific child organization")
	cmd.Flags().BoolVar(&flagAll, "all", false, "Fetch all pages")

	return cmd
}
