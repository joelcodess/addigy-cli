// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newOWebhooksCreateCmd(flags *rootFlags) *cobra.Command {
	var bodyActionUrl string
	var bodyName string
	var bodyTriggerActionEntityIdentifier string
	var bodyTriggerActionEntityName string
	var bodyTriggerActionEntityType string
	var bodyTriggerActionName string
	var bodyTriggerReceiverIdentifier string
	var bodyTriggerReceiverName string
	var bodyTriggerReceiverType string
	var bodyTriggerSenderIdentifier string
	var bodyTriggerSenderName string
	var bodyTriggerSenderType string
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "create <organization_id>",
		Short:       "Create a webhook.",
		Example:     "  addigy-cli o webhooks create 550e8400-e29b-41d4-a716-446655440000",
		Annotations: map[string]string{"pp:endpoint": "webhooks.create", "pp:method": "POST", "pp:path": "/o/{organization_id}/webhooks"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/o/{organization_id}/webhooks"
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
				{
					nestedAction := map[string]any{}
					if bodyActionUrl != "" {
						nestedAction["url"] = bodyActionUrl
					}
					if len(nestedAction) > 0 {
						body["action"] = nestedAction
					}
				}
				if bodyName != "" {
					body["name"] = bodyName
				}
				{
					nestedTrigger := map[string]any{}
					if bodyTriggerActionEntityIdentifier != "" {
						nestedTrigger["action_entity_identifier"] = bodyTriggerActionEntityIdentifier
					}
					if bodyTriggerActionEntityName != "" {
						nestedTrigger["action_entity_name"] = bodyTriggerActionEntityName
					}
					if bodyTriggerActionEntityType != "" {
						nestedTrigger["action_entity_type"] = bodyTriggerActionEntityType
					}
					if bodyTriggerActionName != "" {
						nestedTrigger["action_name"] = bodyTriggerActionName
					}
					if bodyTriggerReceiverIdentifier != "" {
						nestedTrigger["receiver_identifier"] = bodyTriggerReceiverIdentifier
					}
					if bodyTriggerReceiverName != "" {
						nestedTrigger["receiver_name"] = bodyTriggerReceiverName
					}
					if bodyTriggerReceiverType != "" {
						nestedTrigger["receiver_type"] = bodyTriggerReceiverType
					}
					if bodyTriggerSenderIdentifier != "" {
						nestedTrigger["sender_identifier"] = bodyTriggerSenderIdentifier
					}
					if bodyTriggerSenderName != "" {
						nestedTrigger["sender_name"] = bodyTriggerSenderName
					}
					if bodyTriggerSenderType != "" {
						nestedTrigger["sender_type"] = bodyTriggerSenderType
					}
					if len(nestedTrigger) > 0 {
						body["trigger"] = nestedTrigger
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
					"resource": "webhooks",
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
	cmd.Flags().StringVar(&bodyActionUrl, "action-url", "", "The callback url for when webhook is triggered.")
	cmd.Flags().StringVar(&bodyName, "name", "", "The name of the webhook.")
	cmd.Flags().StringVar(&bodyTriggerActionEntityIdentifier, "trigger-action-entity-identifier", "", "The action entity identifier for the webhook trigger.")
	cmd.Flags().StringVar(&bodyTriggerActionEntityName, "trigger-action-entity-name", "", "The action entity name for the webhook trigger.")
	cmd.Flags().StringVar(&bodyTriggerActionEntityType, "trigger-action-entity-type", "", "The action entity type for the webhook trigger.")
	cmd.Flags().StringVar(&bodyTriggerActionName, "trigger-action-name", "", "The action name for the webhook trigger.")
	cmd.Flags().StringVar(&bodyTriggerReceiverIdentifier, "trigger-receiver-identifier", "", "The receiver identifier for the webhook trigger.")
	cmd.Flags().StringVar(&bodyTriggerReceiverName, "trigger-receiver-name", "", "The receiver name for the webhook trigger.")
	cmd.Flags().StringVar(&bodyTriggerReceiverType, "trigger-receiver-type", "", "The receiver type for the webhook trigger.")
	cmd.Flags().StringVar(&bodyTriggerSenderIdentifier, "trigger-sender-identifier", "", "The sender identifier for the webhook trigger.")
	cmd.Flags().StringVar(&bodyTriggerSenderName, "trigger-sender-name", "", "The sender type for the webhook trigger.")
	cmd.Flags().StringVar(&bodyTriggerSenderType, "trigger-sender-type", "", "The sender type for the webhook trigger.")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
