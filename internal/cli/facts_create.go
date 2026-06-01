// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newFactsCreateCmd(flags *rootFlags) *cobra.Command {
	var bodyName string
	var bodyNotes string
	var bodyOsArchitecturesLinuxIsSupported bool
	var bodyOsArchitecturesLinuxLanguage string
	var bodyOsArchitecturesLinuxScript string
	var bodyOsArchitecturesLinuxShebang string
	var bodyOsArchitecturesMacOSIsSupported bool
	var bodyOsArchitecturesMacOSLanguage string
	var bodyOsArchitecturesMacOSScript string
	var bodyOsArchitecturesMacOSShebang string
	var bodyReturnType string
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "create",
		Short:       "Create a custom fact.",
		Example:     " addigy-cli facts create --name example-resource",
		Annotations: map[string]string{"pp:endpoint": "facts.create", "pp:method": "POST", "pp:path": "/facts/custom"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if !stdinBody {
				if !cmd.Flags().Changed("name") && !flags.dryRun {
					return fmt.Errorf("required flag \"%s\" not set", "name")
				}
				if !cmd.Flags().Changed("return-type") && !flags.dryRun {
					return fmt.Errorf("required flag \"%s\" not set", "return-type")
				}
			}
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/facts/custom"
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
				if bodyName != "" {
					body["name"] = bodyName
				}
				if bodyNotes != "" {
					body["notes"] = bodyNotes
				}
				{
					nestedOsArchitectures := map[string]any{}
					{
						nestedOsArchitecturesLinux := map[string]any{}
						if bodyOsArchitecturesLinuxIsSupported {
							nestedOsArchitecturesLinux["is_supported"] = bodyOsArchitecturesLinuxIsSupported
						}
						if bodyOsArchitecturesLinuxLanguage != "" {
							nestedOsArchitecturesLinux["language"] = bodyOsArchitecturesLinuxLanguage
						}
						if bodyOsArchitecturesLinuxScript != "" {
							nestedOsArchitecturesLinux["script"] = bodyOsArchitecturesLinuxScript
						}
						if bodyOsArchitecturesLinuxShebang != "" {
							nestedOsArchitecturesLinux["shebang"] = bodyOsArchitecturesLinuxShebang
						}
						if len(nestedOsArchitecturesLinux) > 0 {
							nestedOsArchitectures["linux"] = nestedOsArchitecturesLinux
						}
					}
					{
						nestedOsArchitecturesMacOS := map[string]any{}
						if bodyOsArchitecturesMacOSIsSupported {
							nestedOsArchitecturesMacOS["is_supported"] = bodyOsArchitecturesMacOSIsSupported
						}
						if bodyOsArchitecturesMacOSLanguage != "" {
							nestedOsArchitecturesMacOS["language"] = bodyOsArchitecturesMacOSLanguage
						}
						if bodyOsArchitecturesMacOSScript != "" {
							nestedOsArchitecturesMacOS["script"] = bodyOsArchitecturesMacOSScript
						}
						if bodyOsArchitecturesMacOSShebang != "" {
							nestedOsArchitecturesMacOS["shebang"] = bodyOsArchitecturesMacOSShebang
						}
						if len(nestedOsArchitecturesMacOS) > 0 {
							nestedOsArchitectures["macOS"] = nestedOsArchitecturesMacOS
						}
					}
					if len(nestedOsArchitectures) > 0 {
						body["os_architectures"] = nestedOsArchitectures
					}
				}
				if bodyReturnType != "" {
					body["return_type"] = bodyReturnType
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
					"resource": "facts",
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
	cmd.Flags().StringVar(&bodyName, "name", "", "Name of the custom fact.")
	cmd.Flags().StringVar(&bodyNotes, "notes", "", "Notes")
	cmd.Flags().BoolVar(&bodyOsArchitecturesLinuxIsSupported, "os-architectures-linux-is-supported", false, "Possible values:   true, false")
	cmd.Flags().StringVar(&bodyOsArchitecturesLinuxLanguage, "os-architectures-linux-language", "", "Possible values:   zsh, bash, python")
	cmd.Flags().StringVar(&bodyOsArchitecturesLinuxScript, "os-architectures-linux-script", "", "Script")
	cmd.Flags().StringVar(&bodyOsArchitecturesLinuxShebang, "os-architectures-linux-shebang", "", "Shebang")
	cmd.Flags().BoolVar(&bodyOsArchitecturesMacOSIsSupported, "os-architectures-mac-os-is-supported", false, "Possible values:   true, false")
	cmd.Flags().StringVar(&bodyOsArchitecturesMacOSLanguage, "os-architectures-mac-os-language", "", "Possible values:   zsh, bash, python")
	cmd.Flags().StringVar(&bodyOsArchitecturesMacOSScript, "os-architectures-mac-os-script", "", "Script")
	cmd.Flags().StringVar(&bodyOsArchitecturesMacOSShebang, "os-architectures-mac-os-shebang", "", "Shebang")
	cmd.Flags().StringVar(&bodyReturnType, "return-type", "", "Possible values:   string, number, boolean, list")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
