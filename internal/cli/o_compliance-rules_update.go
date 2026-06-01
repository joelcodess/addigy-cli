// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newOComplianceRulesUpdateCmd(flags *rootFlags) *cobra.Command {
	var bodyAgentRemediationScriptId string
	var bodyFilterSets string
	var bodyId string
	var bodyName string
	var bodyRemediationEnabled bool
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "update <organization_id>",
		Short:       "Update a compliance rule. Permission Required: Edit Benchmark.",
		Example:     " addigy-cli o compliance-rules update 550e8400-e29b-41d4-a716-446655440000",
		Annotations: map[string]string{"pp:endpoint": "compliance-rules.update", "pp:method": "PUT", "pp:path": "/o/{organization_id}/compliance-rules"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/o/{organization_id}/compliance-rules"
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
				if bodyAgentRemediationScriptId != "" {
					body["agent_remediation_script_id"] = bodyAgentRemediationScriptId
				}
				if bodyFilterSets != "" {
					var parsedFilterSets any
					if err := json.Unmarshal([]byte(bodyFilterSets), &parsedFilterSets); err != nil {
						return fmt.Errorf("parsing --filter-sets JSON: %w", err)
					}
					body["filter_sets"] = parsedFilterSets
				}
				if bodyId != "" {
					body["id"] = bodyId
				}
				if bodyName != "" {
					body["name"] = bodyName
				}
				if bodyRemediationEnabled {
					body["remediation_enabled"] = bodyRemediationEnabled
				}
			}
			data, statusCode, err := c.Put(path, body)
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
					"action":   "put",
					"resource": "compliance-rules",
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
	cmd.Flags().StringVar(&bodyAgentRemediationScriptId, "agent-remediation-script-id", "", "Agent remediation script id")
	cmd.Flags().StringVar(&bodyFilterSets, "filter-sets", "", "Filter sets")
	cmd.Flags().StringVar(&bodyId, "id", "", "Id")
	cmd.Flags().StringVar(&bodyName, "name", "", "Name")
	cmd.Flags().BoolVar(&bodyRemediationEnabled, "remediation-enabled", false, "Remediation enabled")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
