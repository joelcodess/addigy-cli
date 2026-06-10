// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newOPromotedCmd(flags *rootFlags) *cobra.Command {

	cmd := &cobra.Command{
		Use:         "o",
		Short:       "Request additional Azure Conditional Access Connectors for their organization.",
		Long:        "Shortcut for 'o create'. Request additional Azure Conditional Access Connectors for their organization.",
		Example:     "  addigy-cli o",
		Annotations: map[string]string{"pp:endpoint": "o.create", "pp:method": "POST", "pp:path": "/o/integrations/azure-ca/tenants/expand-request"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/o/integrations/azure-ca/tenants/expand-request"
			// HasStore + non-GET falls through to a live API call here
			// rather than through resolveRead (GET-only internally); a
			// body-aware cached read helper is filed as #425 for when a
			// second store-backed POST-search consumer ships.
			body := map[string]any{}
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

	// Wire sibling endpoints and sub-resources as subcommands
	{
		sub := newOBenchmarksCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOChildrenCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOCommunityCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOComplianceRulesCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newODdmUpdatesCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newODeviceCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newODevicesCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOHomescreenCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOIdentityCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOIntegrationsCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOMdmCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOMonitoringCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOPoliciesCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOPolicyCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOPrebuiltAppsCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOReportsCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOScriptsCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOTemplatesCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOUsersCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOVariablesCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}
	{
		sub := newOWebhooksCmd(flags)
		sub.Hidden = false // unhide: the raw parent is hidden but these are useful under the promoted command
		cmd.AddCommand(sub)
	}

	return cmd
}
