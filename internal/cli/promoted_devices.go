// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newDevicesPromotedCmd(flags *rootFlags) *cobra.Command {
	var bodyDesiredFactIdentifiers string
	var bodyPage int
	var bodyPerPage int
	var bodyQueryFilters string
	var bodyQueryPolicyId string
	var bodyQuerySearchAny string
	var bodySortDirection string
	var bodySortField string

	cmd := &cobra.Command{
		Use:         "devices",
		Short:       "Allow to query for a set of devices based on a value that pertains to one of their device facts. Permission...",
		Long:        "Shortcut for 'devices create'. Allow to query for a set of devices based on a value that pertains to one of their device facts. Permission...",
		Example:     " addigy-cli devices",
		Annotations: map[string]string{"pp:endpoint": "devices.create", "pp:method": "POST", "pp:path": "/devices"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/devices"
			// HasStore + non-GET falls through to a live API call here
			// rather than through resolveRead (GET-only internally); a
			// body-aware cached read helper is filed as #425 for when a
			// second store-backed POST-search consumer ships.
			body := map[string]any{}
			if bodyDesiredFactIdentifiers != "" {
				var parsedDesiredFactIdentifiers any
				if err := json.Unmarshal([]byte(bodyDesiredFactIdentifiers), &parsedDesiredFactIdentifiers); err != nil {
					return fmt.Errorf("parsing --desired-fact-identifiers JSON: %w", err)
				}
				body["desired_fact_identifiers"] = parsedDesiredFactIdentifiers
			}
			if bodyPage != 0 {
				body["page"] = bodyPage
			}
			if bodyPerPage != 0 {
				body["per_page"] = bodyPerPage
			}
			{
				nestedQuery := map[string]any{}
				if bodyQueryFilters != "" {
					var parsedQueryFilters any
					if err := json.Unmarshal([]byte(bodyQueryFilters), &parsedQueryFilters); err != nil {
						return fmt.Errorf("parsing --query-filters JSON: %w", err)
					}
					nestedQuery["filters"] = parsedQueryFilters
				}
				if bodyQueryPolicyId != "" {
					nestedQuery["policy_id"] = bodyQueryPolicyId
				}
				if bodyQuerySearchAny != "" {
					nestedQuery["search_any"] = bodyQuerySearchAny
				}
				if len(nestedQuery) > 0 {
					body["query"] = nestedQuery
				}
			}
			if bodySortDirection != "" {
				body["sort_direction"] = bodySortDirection
			}
			if bodySortField != "" {
				body["sort_field"] = bodySortField
			}
			data, _, err := c.Post(path, body)
			prov := attachFreshness(DataProvenance{Source: "live"}, flags)
			if err != nil {
				return classifyAPIError(err, flags)
			}
			// Unwrap API response envelopes (e.g. {"status":"success","data":[...]})
			// so output helpers see the inner data, not the wrapper.
			data = extractResponseData(data)

			// Print provenance to stderr
			{
				var countItems []json.RawMessage
				if json.Unmarshal(data, &countItems) != nil {
					// Single object, not an array
					countItems = []json.RawMessage{data}
				}
				printProvenance(cmd, len(countItems), prov)
			}
			// For JSON output, wrap with provenance envelope. --select wins over
			// --compact when both are set; --compact only runs when no explicit
			// fields were requested. Explicit format flags (--csv, --quiet, --plain)
			// opt out of the auto-JSON path so piped consumers that asked for a
			// non-JSON format reach the standard pipeline below.
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
	cmd.Flags().StringVar(&bodyDesiredFactIdentifiers, "desired-fact-identifiers", "", "Optional field. Limits response to only provide the facts within the passed array. If no desired fact identifiers...")
	cmd.Flags().IntVar(&bodyPage, "page", 0, "Page")
	cmd.Flags().IntVar(&bodyPerPage, "per-page", 0, "Per page")
	cmd.Flags().StringVar(&bodyQueryFilters, "query-filters", "", "Filters")
	cmd.Flags().StringVar(&bodyQueryPolicyId, "query-policy-id", "", "Policy id")
	cmd.Flags().StringVar(&bodyQuerySearchAny, "query-search-any", "", "Search any")
	cmd.Flags().StringVar(&bodySortDirection, "sort-direction", "", "Sort direction")
	cmd.Flags().StringVar(&bodySortField, "sort-field", "", "Sort field")

	// Wire sibling endpoints and sub-resources as subcommands

	return cmd
}
