// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func newODevicesCommandsRunCmd(flags *rootFlags) *cobra.Command {
	var bodyAgentIds string
	var bodyCommand string
	var bodyBackground bool
	var bodyRunAsClient bool
	var stdinBody bool

	cmd := &cobra.Command{
		Use:   "run <organization_id>",
		Short: "Run a shell command on one or more devices.",
		Long: "Run a shell command on one or more devices via the Addigy agent.\n\n" +
			"This is a device-impacting write: the command executes on every agent in\n" +
			"--agent-ids. Preview with --dry-run, and only pass --yes once the target\n" +
			"devices, the command, and its effects are clear. The response includes an\n" +
			"action_id; pass it to 'o devices commands output' to read the result.",
		Example: "  addigy-cli o devices commands run 550e8400-e29b-41d4-a716-446655440000 \\\n" +
			"    --agent-ids 89435d48-7f42-4020-97ae-de62134f56cc --command 'whoami'",
		Annotations: map[string]string{"pp:endpoint": "devices.commands.run", "pp:method": "POST", "pp:path": "/o/{organization_id}/devices/commands/run"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			if !stdinBody {
				if !cmd.Flags().Changed("agent-ids") && !flags.dryRun {
					return fmt.Errorf("required flag \"%s\" not set", "agent-ids")
				}
				if !cmd.Flags().Changed("command") && !flags.dryRun {
					return fmt.Errorf("required flag \"%s\" not set", "command")
				}
			}
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/o/{organization_id}/devices/commands/run"
			path = replacePathParam(path, "organization_id", args[0])
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
					// Accept either a JSON array (["id1","id2"]) for parity with the
					// other array-valued flags in this CLI, or a bare comma-separated
					// list (id1,id2) so a single agent id needs no shell quoting.
					var parsedAgentIds any
					if err := json.Unmarshal([]byte(bodyAgentIds), &parsedAgentIds); err == nil {
						body["agent_ids"] = parsedAgentIds
					} else {
						parts := strings.Split(bodyAgentIds, ",")
						ids := make([]string, 0, len(parts))
						for _, p := range parts {
							if s := strings.TrimSpace(p); s != "" {
								ids = append(ids, s)
							}
						}
						body["agent_ids"] = ids
					}
				}
				if bodyCommand != "" {
					body["command"] = bodyCommand
				}
				// Booleans are only sent when explicitly set so the request mirrors
				// the user's intent rather than always pinning the API defaults.
				if cmd.Flags().Changed("background") {
					body["background"] = bodyBackground
				}
				if cmd.Flags().Changed("run-as-client") {
					body["run_as_client"] = bodyRunAsClient
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
					"resource": "devices",
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
	cmd.Flags().StringVar(&bodyAgentIds, "agent-ids", "", "Agent ids to run the command on: JSON array or comma-separated list")
	cmd.Flags().StringVar(&bodyCommand, "command", "", "Shell command to run on the devices")
	cmd.Flags().BoolVar(&bodyBackground, "background", false, "Run the command in the background")
	cmd.Flags().BoolVar(&bodyRunAsClient, "run-as-client", false, "Run the command as the logged-in client user instead of root")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
