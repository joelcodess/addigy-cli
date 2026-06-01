// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newOHomescreenCreateCmd(flags *rootFlags) *cobra.Command {
	var bodyAssigned bool
	var bodyDeviceType string
	var bodyDock string
	var bodyPages string
	var bodyPolicyId string
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "create <organization_id>",
		Short:       "Create a policy home screen layout",
		Example:     "  addigy-cli o homescreen create 550e8400-e29b-41d4-a716-446655440000",
		Annotations: map[string]string{"pp:endpoint": "homescreen.create", "pp:method": "POST", "pp:path": "/o/{organization_id}/homescreen"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/o/{organization_id}/homescreen"
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
				if bodyAssigned {
					body["assigned"] = bodyAssigned
				}
				if bodyDeviceType != "" {
					body["device_type"] = bodyDeviceType
				}
				if bodyDock != "" {
					var parsedDock any
					if err := json.Unmarshal([]byte(bodyDock), &parsedDock); err != nil {
						return fmt.Errorf("parsing --dock JSON: %w", err)
					}
					body["dock"] = parsedDock
				}
				if bodyPages != "" {
					var parsedPages any
					if err := json.Unmarshal([]byte(bodyPages), &parsedPages); err != nil {
						return fmt.Errorf("parsing --pages JSON: %w", err)
					}
					body["pages"] = parsedPages
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
					"resource": "homescreen",
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
	cmd.Flags().BoolVar(&bodyAssigned, "assigned", false, "Assigned")
	cmd.Flags().StringVar(&bodyDeviceType, "device-type", "", "Device type")
	cmd.Flags().StringVar(&bodyDock, "dock", "", "Dock")
	cmd.Flags().StringVar(&bodyPages, "pages", "", "Pages")
	cmd.Flags().StringVar(&bodyPolicyId, "policy-id", "", "Policy id")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
