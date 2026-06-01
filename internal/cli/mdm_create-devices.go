// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newMdmCreateDevicesCmd(flags *rootFlags) *cobra.Command {
	var bodyDevicesUserIds string
	var bodyDevicesUuid string
	var bodyPayloadsGroupId string
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "create-devices",
		Short:       "Deploys profile to list of devices and/or managed users. It is an atomic request meaning that if one error is...",
		Example:     "  addigy-cli mdm create-devices --payloads-group-id 550e8400-e29b-41d4-a716-446655440000",
		Annotations: map[string]string{"pp:endpoint": "mdm.create-devices", "pp:method": "POST", "pp:path": "/mdm/devices/profile/deploy"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if !stdinBody {
				if !cmd.Flags().Changed("payloads-group-id") && !flags.dryRun {
					return fmt.Errorf("required flag \"%s\" not set", "payloads-group-id")
				}
			}
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/mdm/devices/profile/deploy"
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
				if bodyDevicesUserIds != "" {
					var parsedDevicesUserIds any
					if err := json.Unmarshal([]byte(bodyDevicesUserIds), &parsedDevicesUserIds); err != nil {
						return fmt.Errorf("parsing --devices-user-ids JSON: %w", err)
					}
					body["devices_user_ids"] = parsedDevicesUserIds
				}
				if bodyDevicesUuid != "" {
					var parsedDevicesUuid any
					if err := json.Unmarshal([]byte(bodyDevicesUuid), &parsedDevicesUuid); err != nil {
						return fmt.Errorf("parsing --devices-uuid JSON: %w", err)
					}
					body["devices_uuid"] = parsedDevicesUuid
				}
				if bodyPayloadsGroupId != "" {
					body["payloads_group_id"] = bodyPayloadsGroupId
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
					"resource": "mdm",
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
	cmd.Flags().StringVar(&bodyDevicesUserIds, "devices-user-ids", "", "Devices user ids")
	cmd.Flags().StringVar(&bodyDevicesUuid, "devices-uuid", "", "Devices uuid")
	cmd.Flags().StringVar(&bodyPayloadsGroupId, "payloads-group-id", "", "Payloads group id")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
