// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newMaintenanceCreateCmd(flags *rootFlags) *cobra.Command {
	var bodyDay string
	var bodyEnabled bool
	var bodyExpectedRemediationTime int
	var bodyFrequency string
	var bodyIsInBlueprint bool
	var bodyLocalTime bool
	var bodyMaxTryCount int
	var bodyName string
	var bodyPromptUser bool
	var bodyScheduledTime string
	var bodyScripts string
	var bodyTimeoutSeconds int
	var bodyVersion int
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "create",
		Short:       "Create a maintenance item. Permission Required: Create Catalog Maintenance.",
		Example:     " addigy-cli maintenance create",
		Annotations: map[string]string{"pp:endpoint": "maintenance.create", "pp:method": "POST", "pp:path": "/maintenance"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/maintenance"
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
				if bodyDay != "" {
					body["day"] = bodyDay
				}
				if bodyEnabled {
					body["enabled"] = bodyEnabled
				}
				if bodyExpectedRemediationTime != 0 {
					body["expected_remediation_time"] = bodyExpectedRemediationTime
				}
				if bodyFrequency != "" {
					body["frequency"] = bodyFrequency
				}
				if bodyIsInBlueprint {
					body["is_in_blueprint"] = bodyIsInBlueprint
				}
				if bodyLocalTime {
					body["local_time"] = bodyLocalTime
				}
				if bodyMaxTryCount != 0 {
					body["max_try_count"] = bodyMaxTryCount
				}
				if bodyName != "" {
					body["name"] = bodyName
				}
				if bodyPromptUser {
					body["prompt_user"] = bodyPromptUser
				}
				if bodyScheduledTime != "" {
					body["scheduled_time"] = bodyScheduledTime
				}
				if bodyScripts != "" {
					var parsedScripts any
					if err := json.Unmarshal([]byte(bodyScripts), &parsedScripts); err != nil {
						return fmt.Errorf("parsing --scripts JSON: %w", err)
					}
					body["scripts"] = parsedScripts
				}
				if bodyTimeoutSeconds != 0 {
					body["timeout_seconds"] = bodyTimeoutSeconds
				}
				if bodyVersion != 0 {
					body["version"] = bodyVersion
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
					"resource": "maintenance",
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
	cmd.Flags().StringVar(&bodyDay, "day", "", "Possible values:   For Monthly: a number between 1 to 27,  For Daily: do not set a value,   For Weekly:...")
	cmd.Flags().BoolVar(&bodyEnabled, "enabled", false, "Enabled")
	cmd.Flags().IntVar(&bodyExpectedRemediationTime, "expected-remediation-time", 0, "Time in minutes needed to remediate")
	cmd.Flags().StringVar(&bodyFrequency, "frequency", "", "Possible values:   Monthly, Daily or Weekly")
	cmd.Flags().BoolVar(&bodyIsInBlueprint, "is-in-blueprint", false, "Is in blueprint")
	cmd.Flags().BoolVar(&bodyLocalTime, "local-time", false, "Local time")
	cmd.Flags().IntVar(&bodyMaxTryCount, "max-try-count", 0, "Max try count")
	cmd.Flags().StringVar(&bodyName, "name", "", "Name")
	cmd.Flags().BoolVar(&bodyPromptUser, "prompt-user", false, "Prompt user")
	cmd.Flags().StringVar(&bodyScheduledTime, "scheduled-time", "", "Possible values:  A number between 0 to 23")
	cmd.Flags().StringVar(&bodyScripts, "scripts", "", "Scripts")
	cmd.Flags().IntVar(&bodyTimeoutSeconds, "timeout-seconds", 0, "Timeout seconds")
	cmd.Flags().IntVar(&bodyVersion, "version", 0, "Version")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
