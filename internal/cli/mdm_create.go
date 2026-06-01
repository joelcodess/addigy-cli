// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newMdmCreateCmd(flags *rootFlags) *cobra.Command {
	var bodyPage int
	var bodyPerPage int
	var bodyQueryDaysUntilExpiration int
	var bodyQueryDevices string
	var bodyQueryIssuerCommonName string
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "create",
		Short:       "Paginated request that returns list of installed certificates by mdm devices. Permission Required:...",
		Example:     " addigy-cli mdm create",
		Annotations: map[string]string{"pp:endpoint": "mdm.create", "pp:method": "POST", "pp:path": "/mdm/certificates/query"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/mdm/certificates/query"
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
				if bodyPage != 0 {
					body["page"] = bodyPage
				}
				if bodyPerPage != 0 {
					body["per_page"] = bodyPerPage
				}
				{
					nestedQuery := map[string]any{}
					if bodyQueryDaysUntilExpiration != 0 {
						nestedQuery["days_until_expiration"] = bodyQueryDaysUntilExpiration
					}
					if bodyQueryDevices != "" {
						var parsedQueryDevices any
						if err := json.Unmarshal([]byte(bodyQueryDevices), &parsedQueryDevices); err != nil {
							return fmt.Errorf("parsing --query-devices JSON: %w", err)
						}
						nestedQuery["devices"] = parsedQueryDevices
					}
					if bodyQueryIssuerCommonName != "" {
						nestedQuery["issuer_common_name"] = bodyQueryIssuerCommonName
					}
					if len(nestedQuery) > 0 {
						body["query"] = nestedQuery
					}
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
					"resource": "mdm",
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
	cmd.Flags().IntVar(&bodyPage, "page", 1, "Page")
	cmd.Flags().IntVar(&bodyPerPage, "per-page", 5, "Per page")
	cmd.Flags().IntVar(&bodyQueryDaysUntilExpiration, "query-days-until-expiration", 0, "Number of days until certificate expiration")
	cmd.Flags().StringVar(&bodyQueryDevices, "query-devices", "", "List of devices uuids to filter by")
	cmd.Flags().StringVar(&bodyQueryIssuerCommonName, "query-issuer-common-name", "", "Certificate issuer name to filter by")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
