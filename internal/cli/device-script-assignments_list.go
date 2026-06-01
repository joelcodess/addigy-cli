// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newDeviceScriptAssignmentsListCmd(flags *rootFlags) *cobra.Command {
	var flagScriptId string
	var flagAgentId string

	cmd := &cobra.Command{
		Use:         "list",
		Short:       "Get Device Script Assignments available for the organization.",
		Example:     "  addigy-cli device-script-assignments list",
		Annotations: map[string]string{"pp:endpoint": "device-script-assignments.list", "pp:method": "GET", "pp:path": "/device-script-assignments", "mcp:read-only": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/device-script-assignments"
			params := map[string]string{}
			if flagScriptId != "" {
				params["script_id"] = fmt.Sprintf("%v", flagScriptId)
			}
			if flagAgentId != "" {
				params["agent_id"] = fmt.Sprintf("%v", flagAgentId)
			}
			data, prov, err := resolveRead(cmd.Context(), c, flags, "device-script-assignments", false, path, params, nil)
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
	cmd.Flags().StringVar(&flagScriptId, "script-id", "", "Script or Command ID of the script to get device script assignments.")
	cmd.Flags().StringVar(&flagAgentId, "agent-id", "", "Agent ID of the device to get device script assignments.")

	return cmd
}
