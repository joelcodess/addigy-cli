// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newOPoliciesCreateOCmd(flags *rootFlags) *cobra.Command {
	var bodyAutoRemove bool
	var bodyDisabled bool
	var bodyFilters string
	var bodyPolicyId string
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "create-o <organization_id>",
		Short:       "Add assignment rule to policy. Permission Required: Automatic Policy Assignments.",
		Example:     " addigy-cli o policies create-o 550e8400-e29b-41d4-a716-446655440000",
		Annotations: map[string]string{"pp:endpoint": "policies.create-o", "pp:method": "POST", "pp:path": "/o/{organization_id}/policies/rule"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/o/{organization_id}/policies/rule"
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
				if bodyAutoRemove {
					body["auto_remove"] = bodyAutoRemove
				}
				if bodyDisabled {
					body["disabled"] = bodyDisabled
				}
				if bodyFilters != "" {
					var parsedFilters any
					if err := json.Unmarshal([]byte(bodyFilters), &parsedFilters); err != nil {
						return fmt.Errorf("parsing --filters JSON: %w", err)
					}
					body["filters"] = parsedFilters
				}
				if bodyPolicyId != "" {
					body["policy_id"] = bodyPolicyId
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
					"resource": "policies",
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
	cmd.Flags().BoolVar(&bodyAutoRemove, "auto-remove", false, "Unassign devices that no longer matches filter set. Manually assigned devices will also be removed")
	cmd.Flags().BoolVar(&bodyDisabled, "disabled", false, "Status for the auto assignment rule.")
	cmd.Flags().StringVar(&bodyFilters, "filters", "", "Filters")
	cmd.Flags().StringVar(&bodyPolicyId, "policy-id", "", "Policy id for the rule assignment.")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
