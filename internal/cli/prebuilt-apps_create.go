// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newPrebuiltAppsCreateCmd(flags *rootFlags) *cobra.Command {
	var bodyAppId string
	var bodyConditionScript string
	var bodyFiles string
	var bodyInstallScript string
	var bodyNotification string
	var bodyProfiles string
	var bodyPublishedDate string
	var bodyReleaseNotesUrl string
	var bodyRemoveScript string
	var bodyVariables string
	var bodyVersion string
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "create",
		Short:       "Create a Prebuilt App Version",
		Example:     "  addigy-cli prebuilt-apps create --app-id 550e8400-e29b-41d4-a716-446655440000",
		Annotations: map[string]string{"pp:endpoint": "prebuilt-apps.create", "pp:method": "POST", "pp:path": "/prebuilt-apps/versions"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if !stdinBody {
				if !cmd.Flags().Changed("app-id") && !flags.dryRun {
					return fmt.Errorf("required flag \"%s\" not set", "app-id")
				}
				if !cmd.Flags().Changed("condition-script") && !flags.dryRun {
					return fmt.Errorf("required flag \"%s\" not set", "condition-script")
				}
				if !cmd.Flags().Changed("files") && !flags.dryRun {
					return fmt.Errorf("required flag \"%s\" not set", "files")
				}
				if !cmd.Flags().Changed("install-script") && !flags.dryRun {
					return fmt.Errorf("required flag \"%s\" not set", "install-script")
				}
				if !cmd.Flags().Changed("remove-script") && !flags.dryRun {
					return fmt.Errorf("required flag \"%s\" not set", "remove-script")
				}
				if !cmd.Flags().Changed("version") && !flags.dryRun {
					return fmt.Errorf("required flag \"%s\" not set", "version")
				}
			}
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/prebuilt-apps/versions"
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
				if bodyAppId != "" {
					body["app_id"] = bodyAppId
				}
				if bodyConditionScript != "" {
					body["condition_script"] = bodyConditionScript
				}
				if bodyFiles != "" {
					var parsedFiles any
					if err := json.Unmarshal([]byte(bodyFiles), &parsedFiles); err != nil {
						return fmt.Errorf("parsing --files JSON: %w", err)
					}
					body["files"] = parsedFiles
				}
				if bodyInstallScript != "" {
					body["install_script"] = bodyInstallScript
				}
				if bodyNotification != "" {
					body["notification"] = bodyNotification
				}
				if bodyProfiles != "" {
					var parsedProfiles any
					if err := json.Unmarshal([]byte(bodyProfiles), &parsedProfiles); err != nil {
						return fmt.Errorf("parsing --profiles JSON: %w", err)
					}
					body["profiles"] = parsedProfiles
				}
				if bodyPublishedDate != "" {
					body["published_date"] = bodyPublishedDate
				}
				if bodyReleaseNotesUrl != "" {
					body["release_notes_url"] = bodyReleaseNotesUrl
				}
				if bodyRemoveScript != "" {
					body["remove_script"] = bodyRemoveScript
				}
				if bodyVariables != "" {
					var parsedVariables any
					if err := json.Unmarshal([]byte(bodyVariables), &parsedVariables); err != nil {
						return fmt.Errorf("parsing --variables JSON: %w", err)
					}
					body["variables"] = parsedVariables
				}
				if bodyVersion != "" {
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
	cmd.Flags().StringVar(&bodyAppId, "app-id", "", "App id")
	cmd.Flags().StringVar(&bodyConditionScript, "condition-script", "", "Condition script")
	cmd.Flags().StringVar(&bodyFiles, "files", "", "Files")
	cmd.Flags().StringVar(&bodyInstallScript, "install-script", "", "Install script")
	cmd.Flags().StringVar(&bodyNotification, "notification", "", "Notification")
	cmd.Flags().StringVar(&bodyProfiles, "profiles", "", "Profiles")
	cmd.Flags().StringVar(&bodyPublishedDate, "published-date", "", "Published date")
	cmd.Flags().StringVar(&bodyReleaseNotesUrl, "release-notes-url", "", "Release notes url")
	cmd.Flags().StringVar(&bodyRemoveScript, "remove-script", "", "Remove script")
	cmd.Flags().StringVar(&bodyVariables, "variables", "", "Variables")
	cmd.Flags().StringVar(&bodyVersion, "version", "", "Version")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
