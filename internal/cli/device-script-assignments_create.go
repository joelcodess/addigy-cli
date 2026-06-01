// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newDeviceScriptAssignmentsCreateCmd(flags *rootFlags) *cobra.Command {
	var bodyAgentIds string
	var bodyScriptId string
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "create",
		Short:       "Creates a device script assignment in the organization.",
		Example:     "  addigy-cli device-script-assignments create --script-id 550e8400-e29b-41d4-a716-446655440000",
		Annotations: map[string]string{"pp:endpoint": "device-script-assignments.create", "pp:method": "POST", "pp:path": "/device-script-assignments"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if !stdinBody {
				if !cmd.Flags().Changed("agent-ids") && !flags.dryRun {
					return fmt.Errorf("required flag \"%s\" not set", "agent-ids")
				}
				if !cmd.Flags().Changed("script-id") && !flags.dryRun {
					return fmt.Errorf("required flag \"%s\" not set", "script-id")
				}
			}
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/device-script-assignments"
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
				if bodyAgentIds != "" {
					var parsedAgentIds any
					if err := json.Unmarshal([]byte(bodyAgentIds), &parsedAgentIds); err != nil {
						return fmt.Errorf("parsing --agent-ids JSON: %w", err)
					}
					body["agent_ids"] = parsedAgentIds
				}
				if bodyScriptId != "" {
					body["script_id"] = bodyScriptId
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
					"resource": "device-script-assignments",
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
	cmd.Flags().StringVar(&bodyAgentIds, "agent-ids", "", "Agent ids")
	cmd.Flags().StringVar(&bodyScriptId, "script-id", "", "Script id")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
