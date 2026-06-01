// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newOPrebuiltAppsCreateCmd(flags *rootFlags) *cobra.Command {
	var bodyAutoUpdate bool
	var bodyPolicyIds string
	var bodyPrebuiltAppId string
	var bodyPrebuiltAppVersionId string
	var bodyRunRemovalScript bool
	var bodyUserDeferral int
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "create <organization_id>",
		Short:       "Create a prebuilt app configuration and assign it to one or more policies",
		Example:     "  addigy-cli o prebuilt-apps create 550e8400-e29b-41d4-a716-446655440000",
		Annotations: map[string]string{"pp:endpoint": "prebuilt-apps.create", "pp:method": "POST", "pp:path": "/o/{organization_id}/prebuilt-apps/configurations"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/o/{organization_id}/prebuilt-apps/configurations"
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
				if bodyAutoUpdate {
					body["auto_update"] = bodyAutoUpdate
				}
				if bodyPolicyIds != "" {
					var parsedPolicyIds any
					if err := json.Unmarshal([]byte(bodyPolicyIds), &parsedPolicyIds); err != nil {
						return fmt.Errorf("parsing --policy-ids JSON: %w", err)
					}
					body["policy_ids"] = parsedPolicyIds
				}
				if bodyPrebuiltAppId != "" {
					body["prebuilt_app_id"] = bodyPrebuiltAppId
				}
				if bodyPrebuiltAppVersionId != "" {
					body["prebuilt_app_version_id"] = bodyPrebuiltAppVersionId
				}
				if bodyRunRemovalScript {
					body["run_removal_script"] = bodyRunRemovalScript
				}
				if bodyUserDeferral != 0 {
					body["user_deferral"] = bodyUserDeferral
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
					"resource": "prebuilt-apps",
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
	cmd.Flags().BoolVar(&bodyAutoUpdate, "auto-update", false, "Whether to auto update the configuration")
	cmd.Flags().StringVar(&bodyPolicyIds, "policy-ids", "", "Policy IDs to apply to the configuration to")
	cmd.Flags().StringVar(&bodyPrebuiltAppId, "prebuilt-app-id", "", "Prebuilt App ID to create the configuration for")
	cmd.Flags().StringVar(&bodyPrebuiltAppVersionId, "prebuilt-app-version-id", "", "Prebuilt App Version ID to create the configuration for")
	cmd.Flags().BoolVar(&bodyRunRemovalScript, "run-removal-script", false, "Whether to run the removal script")
	cmd.Flags().IntVar(&bodyUserDeferral, "user-deferral", 0, "Number of days to defer the installation")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
