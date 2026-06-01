// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newOBenchmarksUpdateCmd(flags *rootFlags) *cobra.Command {
	var bodyComplianceRulesIds string
	var bodyId string
	var bodyMaximumOsVersion string
	var bodyMinimumOsVersion string
	var bodyName string
	var bodyTargetOs string
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "update <organization_id>",
		Short:       "Update a benchmark asset. Permission Required: Edit Benchmark.",
		Example:     " addigy-cli o benchmarks update 550e8400-e29b-41d4-a716-446655440000",
		Annotations: map[string]string{"pp:endpoint": "benchmarks.update", "pp:method": "PUT", "pp:path": "/o/{organization_id}/benchmarks"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/o/{organization_id}/benchmarks"
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
				if bodyComplianceRulesIds != "" {
					var parsedComplianceRulesIds any
					if err := json.Unmarshal([]byte(bodyComplianceRulesIds), &parsedComplianceRulesIds); err != nil {
						return fmt.Errorf("parsing --compliance-rules-ids JSON: %w", err)
					}
					body["compliance_rules_ids"] = parsedComplianceRulesIds
				}
				if bodyId != "" {
					body["id"] = bodyId
				}
				if bodyMaximumOsVersion != "" {
					body["maximum_os_version"] = bodyMaximumOsVersion
				}
				if bodyMinimumOsVersion != "" {
					body["minimum_os_version"] = bodyMinimumOsVersion
				}
				if bodyName != "" {
					body["name"] = bodyName
				}
				if bodyTargetOs != "" {
					body["target_os"] = bodyTargetOs
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
					"resource": "benchmarks",
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
	cmd.Flags().StringVar(&bodyComplianceRulesIds, "compliance-rules-ids", "", "List of compliance rule ids")
	cmd.Flags().StringVar(&bodyId, "id", "", "Id")
	cmd.Flags().StringVar(&bodyMaximumOsVersion, "maximum-os-version", "", "Maximum os version")
	cmd.Flags().StringVar(&bodyMinimumOsVersion, "minimum-os-version", "", "Minimum os version")
	cmd.Flags().StringVar(&bodyName, "name", "", "Name")
	cmd.Flags().StringVar(&bodyTargetOs, "target-os", "", "Target os")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
