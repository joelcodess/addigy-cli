// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newMonitoringUpdateCmd(flags *rootFlags) *cobra.Command {
	var bodyCategory string
	var bodyEmails string
	var bodyFact string
	var bodyFactIdentifier string
	var bodyHasScript bool
	var bodyId string
	var bodyInstructions string
	var bodyIsInBlueprint bool
	var bodyLevel string
	var bodyMaxValue float64
	var bodyMinValue float64
	var bodyName string
	var bodyRemediationEnabled bool
	var bodyRemediationTime int
	var bodyScriptId string
	var bodySelector string
	var bodySendTicket bool
	var bodyValue string
	var bodyValueType string
	var bodyVersion int
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "update",
		Short:       "Update a monitoring item.",
		Example:     " addigy-cli monitoring update",
		Annotations: map[string]string{"pp:endpoint": "monitoring.update", "pp:method": "PUT", "pp:path": "/monitoring"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/monitoring"
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
				if bodyCategory != "" {
					body["category"] = bodyCategory
				}
				if bodyEmails != "" {
					var parsedEmails any
					if err := json.Unmarshal([]byte(bodyEmails), &parsedEmails); err != nil {
						return fmt.Errorf("parsing --emails JSON: %w", err)
					}
					body["emails"] = parsedEmails
				}
				if bodyFact != "" {
					body["fact"] = bodyFact
				}
				if bodyFactIdentifier != "" {
					body["fact_identifier"] = bodyFactIdentifier
				}
				if bodyHasScript {
					body["has_script"] = bodyHasScript
				}
				if bodyId != "" {
					body["id"] = bodyId
				}
				if bodyInstructions != "" {
					var parsedInstructions any
					if err := json.Unmarshal([]byte(bodyInstructions), &parsedInstructions); err != nil {
						return fmt.Errorf("parsing --instructions JSON: %w", err)
					}
					body["instructions"] = parsedInstructions
				}
				if bodyIsInBlueprint {
					body["is_in_blueprint"] = bodyIsInBlueprint
				}
				if bodyLevel != "" {
					body["level"] = bodyLevel
				}
				if bodyMaxValue != 0.0 {
					body["max_value"] = bodyMaxValue
				}
				if bodyMinValue != 0.0 {
					body["min_value"] = bodyMinValue
				}
				if bodyName != "" {
					body["name"] = bodyName
				}
				if bodyRemediationEnabled {
					body["remediation_enabled"] = bodyRemediationEnabled
				}
				if bodyRemediationTime != 0 {
					body["remediation_time"] = bodyRemediationTime
				}
				if bodyScriptId != "" {
					body["script_id"] = bodyScriptId
				}
				if bodySelector != "" {
					body["selector"] = bodySelector
				}
				if bodySendTicket {
					body["send_ticket"] = bodySendTicket
				}
				if bodyValue != "" {
					var parsedValue any
					if err := json.Unmarshal([]byte(bodyValue), &parsedValue); err != nil {
						return fmt.Errorf("parsing --value JSON: %w", err)
					}
					body["value"] = parsedValue
				}
				if bodyValueType != "" {
					body["value_type"] = bodyValueType
				}
				if bodyVersion != 0 {
					body["version"] = bodyVersion
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
					"resource": "monitoring",
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
	cmd.Flags().StringVar(&bodyCategory, "category", "", "Category")
	cmd.Flags().StringVar(&bodyEmails, "emails", "", "Emails")
	cmd.Flags().StringVar(&bodyFact, "fact", "", "Fact Name")
	cmd.Flags().StringVar(&bodyFactIdentifier, "fact-identifier", "", "Identifier of the Fact: Custom or Addigy")
	cmd.Flags().BoolVar(&bodyHasScript, "has-script", false, "Possible values:  True if a saved script will be used.  False if a manual script will be indicated")
	cmd.Flags().StringVar(&bodyId, "id", "", "Id for the monitoring item to update")
	cmd.Flags().StringVar(&bodyInstructions, "instructions", "", "Manual script to remediate the issue (has_script=false).")
	cmd.Flags().BoolVar(&bodyIsInBlueprint, "is-in-blueprint", false, "Is in blueprint")
	cmd.Flags().StringVar(&bodyLevel, "level", "", "Possible values:   warning or critical")
	cmd.Flags().Float64Var(&bodyMaxValue, "max-value", 0.0, "Must be used when selector is between")
	cmd.Flags().Float64Var(&bodyMinValue, "min-value", 0.0, "Must be used when selector is between")
	cmd.Flags().StringVar(&bodyName, "name", "", "Name")
	cmd.Flags().BoolVar(&bodyRemediationEnabled, "remediation-enabled", false, "Remediation enabled")
	cmd.Flags().IntVar(&bodyRemediationTime, "remediation-time", 0, "Remediation time")
	cmd.Flags().StringVar(&bodyScriptId, "script-id", "", "Id of the script used to remediate")
	cmd.Flags().StringVar(&bodySelector, "selector", "", "Possible values:   number: >, >=, <, <=,between, changed,  string/date: >, >=, <, <=,between, contains, no...")
	cmd.Flags().BoolVar(&bodySendTicket, "send-ticket", false, "Send ticket")
	cmd.Flags().StringVar(&bodyValue, "value", "", "Value(s) to compare the fact with.")
	cmd.Flags().StringVar(&bodyValueType, "value-type", "", "Possible values:   number, string, date, boolean, list.")
	cmd.Flags().IntVar(&bodyVersion, "version", 0, "Version")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
