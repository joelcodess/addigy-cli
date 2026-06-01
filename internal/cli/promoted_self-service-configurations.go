// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newSelfServiceConfigurationsPromotedCmd(flags *rootFlags) *cobra.Command {
	var bodyAppLogoContentType string
	var bodyAppLogoCreated string
	var bodyAppLogoFilename string
	var bodyAppLogoId string
	var bodyAppLogoMd5Hash string
	var bodyAppLogoOrganizationId string
	var bodyAppLogoProvider string
	var bodyAppLogoSize int
	var bodyAppLogoUserEmail string
	var bodyDockIconContentType string
	var bodyDockIconCreated string
	var bodyDockIconFilename string
	var bodyDockIconId string
	var bodyDockIconMd5Hash string
	var bodyDockIconOrganizationId string
	var bodyDockIconProvider string
	var bodyDockIconSize int
	var bodyDockIconUserEmail string
	var bodyFilevaultPromptText string
	var bodyHideChat bool
	var bodyHomeScreenAddress string
	var bodyHomeScreenCompanyName string
	var bodyHomeScreenConfigureDetails bool
	var bodyHomeScreenDescription string
	var bodyHomeScreenEmail string
	var bodyHomeScreenPhone string
	var bodyHomeScreenShowAddress bool
	var bodyHomeScreenShowDescription bool
	var bodyHomeScreenShowEmail bool
	var bodyHomeScreenShowPhone bool
	var bodyIntegrationIntuneEnabled bool
	var bodyIsInBlueprint bool
	var bodyIsOnboardingConfig bool
	var bodyMaintenancePromptText string
	var bodyMenubarIconContentType string
	var bodyMenubarIconCreated string
	var bodyMenubarIconFilename string
	var bodyMenubarIconId string
	var bodyMenubarIconMd5Hash string
	var bodyMenubarIconOrganizationId string
	var bodyMenubarIconProvider string
	var bodyMenubarIconSize int
	var bodyMenubarIconUserEmail string
	var bodyMsOfficeUpdatesPromptText string
	var bodyName string
	var bodyOsType string
	var bodyScreenviewPromptText string
	var bodyShowDockIcon bool
	var bodyShowInApplications bool
	var bodyShowMenubarIcon bool
	var bodyShowSupport bool
	var bodyUserSentimentPromptText string
	var bodyVersion int

	cmd := &cobra.Command{
		Use:         "self-service-configurations",
		Short:       "Creates a new self service configuration in the organization. Permission Required: Create Instruction.",
		Long:        "Shortcut for 'self-service-configurations create'. Creates a new self service configuration in the organization. Permission Required: Create Instruction.",
		Example:     " addigy-cli self-service-configurations",
		Annotations: map[string]string{"pp:endpoint": "self-service-configurations.create", "pp:method": "POST", "pp:path": "/self-service-configurations"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/self-service-configurations"
			// HasStore + non-GET falls through to a live API call here
			// rather than through resolveRead (GET-only internally); a
			// body-aware cached read helper is filed as #425 for when a
			// second store-backed POST-search consumer ships.
			body := map[string]any{}
			{
				nestedAppLogo := map[string]any{}
				if bodyAppLogoContentType != "" {
					nestedAppLogo["content_type"] = bodyAppLogoContentType
				}
				if bodyAppLogoCreated != "" {
					nestedAppLogo["created"] = bodyAppLogoCreated
				}
				if bodyAppLogoFilename != "" {
					nestedAppLogo["filename"] = bodyAppLogoFilename
				}
				if bodyAppLogoId != "" {
					nestedAppLogo["id"] = bodyAppLogoId
				}
				if bodyAppLogoMd5Hash != "" {
					nestedAppLogo["md5_hash"] = bodyAppLogoMd5Hash
				}
				if bodyAppLogoOrganizationId != "" {
					nestedAppLogo["organization_id"] = bodyAppLogoOrganizationId
				}
				if bodyAppLogoProvider != "" {
					nestedAppLogo["provider"] = bodyAppLogoProvider
				}
				if bodyAppLogoSize != 0 {
					nestedAppLogo["size"] = bodyAppLogoSize
				}
				if bodyAppLogoUserEmail != "" {
					nestedAppLogo["user_email"] = bodyAppLogoUserEmail
				}
				if len(nestedAppLogo) > 0 {
					body["app_logo"] = nestedAppLogo
				}
			}
			{
				nestedDockIcon := map[string]any{}
				if bodyDockIconContentType != "" {
					nestedDockIcon["content_type"] = bodyDockIconContentType
				}
				if bodyDockIconCreated != "" {
					nestedDockIcon["created"] = bodyDockIconCreated
				}
				if bodyDockIconFilename != "" {
					nestedDockIcon["filename"] = bodyDockIconFilename
				}
				if bodyDockIconId != "" {
					nestedDockIcon["id"] = bodyDockIconId
				}
				if bodyDockIconMd5Hash != "" {
					nestedDockIcon["md5_hash"] = bodyDockIconMd5Hash
				}
				if bodyDockIconOrganizationId != "" {
					nestedDockIcon["organization_id"] = bodyDockIconOrganizationId
				}
				if bodyDockIconProvider != "" {
					nestedDockIcon["provider"] = bodyDockIconProvider
				}
				if bodyDockIconSize != 0 {
					nestedDockIcon["size"] = bodyDockIconSize
				}
				if bodyDockIconUserEmail != "" {
					nestedDockIcon["user_email"] = bodyDockIconUserEmail
				}
				if len(nestedDockIcon) > 0 {
					body["dock_icon"] = nestedDockIcon
				}
			}
			if bodyFilevaultPromptText != "" {
				body["filevault_prompt_text"] = bodyFilevaultPromptText
			}
			if bodyHideChat {
				body["hide_chat"] = bodyHideChat
			}
			if bodyHomeScreenAddress != "" {
				body["home_screen_address"] = bodyHomeScreenAddress
			}
			if bodyHomeScreenCompanyName != "" {
				body["home_screen_company_name"] = bodyHomeScreenCompanyName
			}
			if bodyHomeScreenConfigureDetails {
				body["home_screen_configure_details"] = bodyHomeScreenConfigureDetails
			}
			if bodyHomeScreenDescription != "" {
				body["home_screen_description"] = bodyHomeScreenDescription
			}
			if bodyHomeScreenEmail != "" {
				body["home_screen_email"] = bodyHomeScreenEmail
			}
			if bodyHomeScreenPhone != "" {
				body["home_screen_phone"] = bodyHomeScreenPhone
			}
			if bodyHomeScreenShowAddress {
				body["home_screen_show_address"] = bodyHomeScreenShowAddress
			}
			if bodyHomeScreenShowDescription {
				body["home_screen_show_description"] = bodyHomeScreenShowDescription
			}
			if bodyHomeScreenShowEmail {
				body["home_screen_show_email"] = bodyHomeScreenShowEmail
			}
			if bodyHomeScreenShowPhone {
				body["home_screen_show_phone"] = bodyHomeScreenShowPhone
			}
			if bodyIntegrationIntuneEnabled {
				body["integration_intune_enabled"] = bodyIntegrationIntuneEnabled
			}
			if bodyIsInBlueprint {
				body["is_in_blueprint"] = bodyIsInBlueprint
			}
			if bodyIsOnboardingConfig {
				body["is_onboarding_config"] = bodyIsOnboardingConfig
			}
			if bodyMaintenancePromptText != "" {
				body["maintenance_prompt_text"] = bodyMaintenancePromptText
			}
			{
				nestedMenubarIcon := map[string]any{}
				if bodyMenubarIconContentType != "" {
					nestedMenubarIcon["content_type"] = bodyMenubarIconContentType
				}
				if bodyMenubarIconCreated != "" {
					nestedMenubarIcon["created"] = bodyMenubarIconCreated
				}
				if bodyMenubarIconFilename != "" {
					nestedMenubarIcon["filename"] = bodyMenubarIconFilename
				}
				if bodyMenubarIconId != "" {
					nestedMenubarIcon["id"] = bodyMenubarIconId
				}
				if bodyMenubarIconMd5Hash != "" {
					nestedMenubarIcon["md5_hash"] = bodyMenubarIconMd5Hash
				}
				if bodyMenubarIconOrganizationId != "" {
					nestedMenubarIcon["organization_id"] = bodyMenubarIconOrganizationId
				}
				if bodyMenubarIconProvider != "" {
					nestedMenubarIcon["provider"] = bodyMenubarIconProvider
				}
				if bodyMenubarIconSize != 0 {
					nestedMenubarIcon["size"] = bodyMenubarIconSize
				}
				if bodyMenubarIconUserEmail != "" {
					nestedMenubarIcon["user_email"] = bodyMenubarIconUserEmail
				}
				if len(nestedMenubarIcon) > 0 {
					body["menubar_icon"] = nestedMenubarIcon
				}
			}
			if bodyMsOfficeUpdatesPromptText != "" {
				body["ms_office_updates_prompt_text"] = bodyMsOfficeUpdatesPromptText
			}
			if bodyName != "" {
				body["name"] = bodyName
			}
			if bodyOsType != "" {
				body["os_type"] = bodyOsType
			}
			if bodyScreenviewPromptText != "" {
				body["screenview_prompt_text"] = bodyScreenviewPromptText
			}
			if bodyShowDockIcon {
				body["show_dock_icon"] = bodyShowDockIcon
			}
			if bodyShowInApplications {
				body["show_in_applications"] = bodyShowInApplications
			}
			if bodyShowMenubarIcon {
				body["show_menubar_icon"] = bodyShowMenubarIcon
			}
			if bodyShowSupport {
				body["show_support"] = bodyShowSupport
			}
			if bodyUserSentimentPromptText != "" {
				body["user_sentiment_prompt_text"] = bodyUserSentimentPromptText
			}
			if bodyVersion != 0 {
				body["version"] = bodyVersion
			}
			data, _, err := c.Post(path, body)
			prov := attachFreshness(DataProvenance{Source: "live"}, flags)
			if err != nil {
				return classifyAPIError(err, flags)
			}
			// Unwrap API response envelopes (e.g. {"status":"success","data":[...]})
			// so output helpers see the inner data, not the wrapper.
			data = extractResponseData(data)

			// Print provenance to stderr
			{
				var countItems []json.RawMessage
				if json.Unmarshal(data, &countItems) != nil {
					// Single object, not an array
					countItems = []json.RawMessage{data}
				}
				printProvenance(cmd, len(countItems), prov)
			}
			// For JSON output, wrap with provenance envelope. --select wins over
			// --compact when both are set; --compact only runs when no explicit
			// fields were requested. Explicit format flags (--csv, --quiet, --plain)
			// opt out of the auto-JSON path so piped consumers that asked for a
			// non-JSON format reach the standard pipeline below.
			if flags.asJSON || (!isTerminal(cmd.OutOrStdout()) && !flags.csv && !flags.quiet && !flags.plain) {
				filtered := data
				if flags.selectFields != "" {
					filtered = filterFields(filtered, flags.selectFields)
				} else if flags.compact {
					filtered = compactFields(filtered)
				}
				wrapped, wrapErr := wrapWithProvenance(filtered, prov)
				if wrapErr != nil {
					return wrapErr
				}
				return printOutput(cmd.OutOrStdout(), wrapped, true)
			}
			if wantsHumanTable(cmd.OutOrStdout(), flags) {
				var items []map[string]any
				if json.Unmarshal(data, &items) == nil && len(items) > 0 {
					if err := printAutoTable(cmd.OutOrStdout(), items); err != nil {
						return err
					}
					if len(items) >= 25 {
						fmt.Fprintf(os.Stderr, "\nShowing %d results. To narrow: add --limit, --json --select, or filter flags.\n", len(items))
					}
					return nil
				}
			}
			return printOutputWithFlags(cmd.OutOrStdout(), data, flags)
		},
	}
	cmd.Flags().StringVar(&bodyAppLogoContentType, "app-logo-content-type", "", "Content type")
	cmd.Flags().StringVar(&bodyAppLogoCreated, "app-logo-created", "", "Created")
	cmd.Flags().StringVar(&bodyAppLogoFilename, "app-logo-filename", "", "Filename")
	cmd.Flags().StringVar(&bodyAppLogoId, "app-logo-id", "", "Id")
	cmd.Flags().StringVar(&bodyAppLogoMd5Hash, "app-logo-md5-hash", "", "Md5 hash")
	cmd.Flags().StringVar(&bodyAppLogoOrganizationId, "app-logo-organization-id", "", "Organization id")
	cmd.Flags().StringVar(&bodyAppLogoProvider, "app-logo-provider", "", "Provider")
	cmd.Flags().IntVar(&bodyAppLogoSize, "app-logo-size", 0, "Size")
	cmd.Flags().StringVar(&bodyAppLogoUserEmail, "app-logo-user-email", "", "User email")
	cmd.Flags().StringVar(&bodyDockIconContentType, "dock-icon-content-type", "", "Content type")
	cmd.Flags().StringVar(&bodyDockIconCreated, "dock-icon-created", "", "Created")
	cmd.Flags().StringVar(&bodyDockIconFilename, "dock-icon-filename", "", "Filename")
	cmd.Flags().StringVar(&bodyDockIconId, "dock-icon-id", "", "Id")
	cmd.Flags().StringVar(&bodyDockIconMd5Hash, "dock-icon-md5-hash", "", "Md5 hash")
	cmd.Flags().StringVar(&bodyDockIconOrganizationId, "dock-icon-organization-id", "", "Organization id")
	cmd.Flags().StringVar(&bodyDockIconProvider, "dock-icon-provider", "", "Provider")
	cmd.Flags().IntVar(&bodyDockIconSize, "dock-icon-size", 0, "Size")
	cmd.Flags().StringVar(&bodyDockIconUserEmail, "dock-icon-user-email", "", "User email")
	cmd.Flags().StringVar(&bodyFilevaultPromptText, "filevault-prompt-text", "", "Filevault prompt text")
	cmd.Flags().BoolVar(&bodyHideChat, "hide-chat", false, "Hide chat")
	cmd.Flags().StringVar(&bodyHomeScreenAddress, "home-screen-address", "", "Home screen address")
	cmd.Flags().StringVar(&bodyHomeScreenCompanyName, "home-screen-company-name", "", "Home screen company name")
	cmd.Flags().BoolVar(&bodyHomeScreenConfigureDetails, "home-screen-configure-details", false, "Home screen configure details")
	cmd.Flags().StringVar(&bodyHomeScreenDescription, "home-screen-description", "", "Home screen description")
	cmd.Flags().StringVar(&bodyHomeScreenEmail, "home-screen-email", "", "Home screen email")
	cmd.Flags().StringVar(&bodyHomeScreenPhone, "home-screen-phone", "", "Home screen phone")
	cmd.Flags().BoolVar(&bodyHomeScreenShowAddress, "home-screen-show-address", false, "Home screen show address")
	cmd.Flags().BoolVar(&bodyHomeScreenShowDescription, "home-screen-show-description", false, "Home screen show description")
	cmd.Flags().BoolVar(&bodyHomeScreenShowEmail, "home-screen-show-email", false, "Home screen show email")
	cmd.Flags().BoolVar(&bodyHomeScreenShowPhone, "home-screen-show-phone", false, "Home screen show phone")
	cmd.Flags().BoolVar(&bodyIntegrationIntuneEnabled, "integration-intune-enabled", false, "Integration intune enabled")
	cmd.Flags().BoolVar(&bodyIsInBlueprint, "is-in-blueprint", false, "Is in blueprint")
	cmd.Flags().BoolVar(&bodyIsOnboardingConfig, "is-onboarding-config", false, "Is onboarding config")
	cmd.Flags().StringVar(&bodyMaintenancePromptText, "maintenance-prompt-text", "", "Maintenance prompt text")
	cmd.Flags().StringVar(&bodyMenubarIconContentType, "menubar-icon-content-type", "", "Content type")
	cmd.Flags().StringVar(&bodyMenubarIconCreated, "menubar-icon-created", "", "Created")
	cmd.Flags().StringVar(&bodyMenubarIconFilename, "menubar-icon-filename", "", "Filename")
	cmd.Flags().StringVar(&bodyMenubarIconId, "menubar-icon-id", "", "Id")
	cmd.Flags().StringVar(&bodyMenubarIconMd5Hash, "menubar-icon-md5-hash", "", "Md5 hash")
	cmd.Flags().StringVar(&bodyMenubarIconOrganizationId, "menubar-icon-organization-id", "", "Organization id")
	cmd.Flags().StringVar(&bodyMenubarIconProvider, "menubar-icon-provider", "", "Provider")
	cmd.Flags().IntVar(&bodyMenubarIconSize, "menubar-icon-size", 0, "Size")
	cmd.Flags().StringVar(&bodyMenubarIconUserEmail, "menubar-icon-user-email", "", "User email")
	cmd.Flags().StringVar(&bodyMsOfficeUpdatesPromptText, "ms-office-updates-prompt-text", "", "Ms office updates prompt text")
	cmd.Flags().StringVar(&bodyName, "name", "", "Name")
	cmd.Flags().StringVar(&bodyOsType, "os-type", "", "Possible values:   macOS or iOS")
	cmd.Flags().StringVar(&bodyScreenviewPromptText, "screenview-prompt-text", "", "Screenview prompt text")
	cmd.Flags().BoolVar(&bodyShowDockIcon, "show-dock-icon", false, "Show dock icon")
	cmd.Flags().BoolVar(&bodyShowInApplications, "show-in-applications", false, "Show in applications")
	cmd.Flags().BoolVar(&bodyShowMenubarIcon, "show-menubar-icon", false, "Show menubar icon")
	cmd.Flags().BoolVar(&bodyShowSupport, "show-support", false, "Show support")
	cmd.Flags().StringVar(&bodyUserSentimentPromptText, "user-sentiment-prompt-text", "", "User sentiment prompt text")
	cmd.Flags().IntVar(&bodyVersion, "version", 0, "Version")

	// Wire sibling endpoints and sub-resources as subcommands

	return cmd
}
