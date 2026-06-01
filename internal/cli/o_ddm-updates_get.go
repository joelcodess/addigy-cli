// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newODdmUpdatesGetCmd(flags *rootFlags) *cobra.Command {
	var flagDeviceUdid string

	cmd := &cobra.Command{
		Use:         "get <organization_id>",
		Short:       "Gets device active ddm system updates declaration. Permission Required: View devices.",
		Example:     " addigy-cli o ddm-updates get 550e8400-e29b-41d4-a716-446655440000 --device-udid 550e8400-e29b-41d4-a716-446655440000",
		Annotations: map[string]string{"pp:endpoint": "ddm-updates.get", "pp:method": "GET", "pp:path": "/o/{organization_id}/ddm-updates/devices/declarations/active", "mcp:read-only": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			if !cmd.Flags().Changed("device-udid") && !flags.dryRun {
				return fmt.Errorf("required flag \"%s\" not set", "device-udid")
			}
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/o/{organization_id}/ddm-updates/devices/declarations/active"
			path = replacePathParam(path, "organization_id", args[0])
			params := map[string]string{}
			if flagDeviceUdid != "" {
				params["device_udid"] = fmt.Sprintf("%v", flagDeviceUdid)
			}
			data, prov, err := resolveRead(cmd.Context(), c, flags, "ddm-updates", false, path, params, nil)
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
	cmd.Flags().StringVar(&flagDeviceUdid, "device-udid", "", "Device UDID")

	return cmd
}
