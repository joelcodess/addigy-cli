// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newSystemUpdatesCreateSystemupdates3Cmd(flags *rootFlags) *cobra.Command {
	var bodyIosSettingsAllowBetaUpdatesInDdm bool
	var bodyIosSettingsAllowedDays string
	var bodyIosSettingsDaysAfterRelease int
	var bodyIosSettingsDaysAfterReleaseRsr int
	var bodyIosSettingsDeniedPeriod string
	var bodyIosSettingsEnabled bool
	var bodyIosSettingsHoursAfterRelease int
	var bodyIosSettingsInstallAction string
	var bodyIosSettingsKeepOsUpdated bool
	var bodyIosSettingsMaxOsVersionAllowed string
	var bodyIosSettingsMaxUserDeferrals int
	var bodyIosSettingsMinutesAfterRelease int
	var bodyIosSettingsResendUpdateCommandHour int
	var bodyIpadosSettingsAllowBetaUpdatesInDdm bool
	var bodyIpadosSettingsAllowedDays string
	var bodyIpadosSettingsDaysAfterRelease int
	var bodyIpadosSettingsDaysAfterReleaseRsr int
	var bodyIpadosSettingsDeniedPeriod string
	var bodyIpadosSettingsEnabled bool
	var bodyIpadosSettingsHoursAfterRelease int
	var bodyIpadosSettingsInstallAction string
	var bodyIpadosSettingsKeepOsUpdated bool
	var bodyIpadosSettingsMaxOsVersionAllowed string
	var bodyIpadosSettingsMaxUserDeferrals int
	var bodyIpadosSettingsMinutesAfterRelease int
	var bodyIpadosSettingsResendUpdateCommandHour int
	var bodyMacosSettingsAllowBetaUpdatesInDdm bool
	var bodyMacosSettingsAllowedDays string
	var bodyMacosSettingsDaysAfterRelease int
	var bodyMacosSettingsDaysAfterReleaseRsr int
	var bodyMacosSettingsDeniedPeriod string
	var bodyMacosSettingsEnabled bool
	var bodyMacosSettingsHoursAfterRelease int
	var bodyMacosSettingsInstallAction string
	var bodyMacosSettingsKeepOsUpdated bool
	var bodyMacosSettingsMaxOsVersionAllowed string
	var bodyMacosSettingsMaxUserDeferrals int
	var bodyMacosSettingsMinutesAfterRelease int
	var bodyMacosSettingsResendUpdateCommandHour int
	var bodyPolicyId string
	var bodyScheduleCutOffTime string
	var bodyScheduleEnabled bool
	var bodyScheduleMaintenanceWindow string
	var bodyScheduleStartingTimeHour string
	var bodyScheduleStartingTimeMinute string
	var bodyScheduleWeekDays string
	var bodyTvosSettingsAllowBetaUpdatesInDdm bool
	var bodyTvosSettingsAllowedDays string
	var bodyTvosSettingsDaysAfterRelease int
	var bodyTvosSettingsDaysAfterReleaseRsr int
	var bodyTvosSettingsDeniedPeriod string
	var bodyTvosSettingsEnabled bool
	var bodyTvosSettingsHoursAfterRelease int
	var bodyTvosSettingsInstallAction string
	var bodyTvosSettingsKeepOsUpdated bool
	var bodyTvosSettingsMaxOsVersionAllowed string
	var bodyTvosSettingsMaxUserDeferrals int
	var bodyTvosSettingsMinutesAfterRelease int
	var bodyTvosSettingsResendUpdateCommandHour int
	var stdinBody bool

	cmd := &cobra.Command{
		Use:         "create-systemupdates-3",
		Short:       "Requests to create or update system updates settings for a policy. Permission Required: Create System...",
		Example:     " addigy-cli system-updates create-systemupdates-3",
		Annotations: map[string]string{"pp:endpoint": "system-updates.create-systemupdates-3", "pp:method": "POST", "pp:path": "/system-updates/settings"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := flags.newClient()
			if err != nil {
				return err
			}

			path := "/system-updates/settings"
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
					nestedIosSettings := map[string]any{}
					if bodyIosSettingsAllowBetaUpdatesInDdm {
						nestedIosSettings["allow_beta_updates_in_ddm"] = bodyIosSettingsAllowBetaUpdatesInDdm
					}
					if bodyIosSettingsAllowedDays != "" {
						var parsedIosSettingsAllowedDays any
						if err := json.Unmarshal([]byte(bodyIosSettingsAllowedDays), &parsedIosSettingsAllowedDays); err != nil {
							return fmt.Errorf("parsing --ios-settings-allowed-days JSON: %w", err)
						}
						nestedIosSettings["allowed_days"] = parsedIosSettingsAllowedDays
					}
					if bodyIosSettingsDaysAfterRelease != 0 {
						nestedIosSettings["days_after_release"] = bodyIosSettingsDaysAfterRelease
					}
					if bodyIosSettingsDaysAfterReleaseRsr != 0 {
						nestedIosSettings["days_after_release_rsr"] = bodyIosSettingsDaysAfterReleaseRsr
					}
					if bodyIosSettingsDeniedPeriod != "" {
						var parsedIosSettingsDeniedPeriod any
						if err := json.Unmarshal([]byte(bodyIosSettingsDeniedPeriod), &parsedIosSettingsDeniedPeriod); err != nil {
							return fmt.Errorf("parsing --ios-settings-denied-period JSON: %w", err)
						}
						nestedIosSettings["denied_period"] = parsedIosSettingsDeniedPeriod
					}
					if bodyIosSettingsEnabled {
						nestedIosSettings["enabled"] = bodyIosSettingsEnabled
					}
					if bodyIosSettingsHoursAfterRelease != 0 {
						nestedIosSettings["hours_after_release"] = bodyIosSettingsHoursAfterRelease
					}
					if bodyIosSettingsInstallAction != "" {
						nestedIosSettings["install_action"] = bodyIosSettingsInstallAction
					}
					if bodyIosSettingsKeepOsUpdated {
						nestedIosSettings["keep_os_updated"] = bodyIosSettingsKeepOsUpdated
					}
					if bodyIosSettingsMaxOsVersionAllowed != "" {
						nestedIosSettings["max_os_version_allowed"] = bodyIosSettingsMaxOsVersionAllowed
					}
					if bodyIosSettingsMaxUserDeferrals != 0 {
						nestedIosSettings["max_user_deferrals"] = bodyIosSettingsMaxUserDeferrals
					}
					if bodyIosSettingsMinutesAfterRelease != 0 {
						nestedIosSettings["minutes_after_release"] = bodyIosSettingsMinutesAfterRelease
					}
					if bodyIosSettingsResendUpdateCommandHour != 0 {
						nestedIosSettings["resend_update_command_hour"] = bodyIosSettingsResendUpdateCommandHour
					}
					if len(nestedIosSettings) > 0 {
						body["ios_settings"] = nestedIosSettings
					}
				}
				{
					nestedIpadosSettings := map[string]any{}
					if bodyIpadosSettingsAllowBetaUpdatesInDdm {
						nestedIpadosSettings["allow_beta_updates_in_ddm"] = bodyIpadosSettingsAllowBetaUpdatesInDdm
					}
					if bodyIpadosSettingsAllowedDays != "" {
						var parsedIpadosSettingsAllowedDays any
						if err := json.Unmarshal([]byte(bodyIpadosSettingsAllowedDays), &parsedIpadosSettingsAllowedDays); err != nil {
							return fmt.Errorf("parsing --ipados-settings-allowed-days JSON: %w", err)
						}
						nestedIpadosSettings["allowed_days"] = parsedIpadosSettingsAllowedDays
					}
					if bodyIpadosSettingsDaysAfterRelease != 0 {
						nestedIpadosSettings["days_after_release"] = bodyIpadosSettingsDaysAfterRelease
					}
					if bodyIpadosSettingsDaysAfterReleaseRsr != 0 {
						nestedIpadosSettings["days_after_release_rsr"] = bodyIpadosSettingsDaysAfterReleaseRsr
					}
					if bodyIpadosSettingsDeniedPeriod != "" {
						var parsedIpadosSettingsDeniedPeriod any
						if err := json.Unmarshal([]byte(bodyIpadosSettingsDeniedPeriod), &parsedIpadosSettingsDeniedPeriod); err != nil {
							return fmt.Errorf("parsing --ipados-settings-denied-period JSON: %w", err)
						}
						nestedIpadosSettings["denied_period"] = parsedIpadosSettingsDeniedPeriod
					}
					if bodyIpadosSettingsEnabled {
						nestedIpadosSettings["enabled"] = bodyIpadosSettingsEnabled
					}
					if bodyIpadosSettingsHoursAfterRelease != 0 {
						nestedIpadosSettings["hours_after_release"] = bodyIpadosSettingsHoursAfterRelease
					}
					if bodyIpadosSettingsInstallAction != "" {
						nestedIpadosSettings["install_action"] = bodyIpadosSettingsInstallAction
					}
					if bodyIpadosSettingsKeepOsUpdated {
						nestedIpadosSettings["keep_os_updated"] = bodyIpadosSettingsKeepOsUpdated
					}
					if bodyIpadosSettingsMaxOsVersionAllowed != "" {
						nestedIpadosSettings["max_os_version_allowed"] = bodyIpadosSettingsMaxOsVersionAllowed
					}
					if bodyIpadosSettingsMaxUserDeferrals != 0 {
						nestedIpadosSettings["max_user_deferrals"] = bodyIpadosSettingsMaxUserDeferrals
					}
					if bodyIpadosSettingsMinutesAfterRelease != 0 {
						nestedIpadosSettings["minutes_after_release"] = bodyIpadosSettingsMinutesAfterRelease
					}
					if bodyIpadosSettingsResendUpdateCommandHour != 0 {
						nestedIpadosSettings["resend_update_command_hour"] = bodyIpadosSettingsResendUpdateCommandHour
					}
					if len(nestedIpadosSettings) > 0 {
						body["ipados_settings"] = nestedIpadosSettings
					}
				}
				{
					nestedMacosSettings := map[string]any{}
					if bodyMacosSettingsAllowBetaUpdatesInDdm {
						nestedMacosSettings["allow_beta_updates_in_ddm"] = bodyMacosSettingsAllowBetaUpdatesInDdm
					}
					if bodyMacosSettingsAllowedDays != "" {
						var parsedMacosSettingsAllowedDays any
						if err := json.Unmarshal([]byte(bodyMacosSettingsAllowedDays), &parsedMacosSettingsAllowedDays); err != nil {
							return fmt.Errorf("parsing --macos-settings-allowed-days JSON: %w", err)
						}
						nestedMacosSettings["allowed_days"] = parsedMacosSettingsAllowedDays
					}
					if bodyMacosSettingsDaysAfterRelease != 0 {
						nestedMacosSettings["days_after_release"] = bodyMacosSettingsDaysAfterRelease
					}
					if bodyMacosSettingsDaysAfterReleaseRsr != 0 {
						nestedMacosSettings["days_after_release_rsr"] = bodyMacosSettingsDaysAfterReleaseRsr
					}
					if bodyMacosSettingsDeniedPeriod != "" {
						var parsedMacosSettingsDeniedPeriod any
						if err := json.Unmarshal([]byte(bodyMacosSettingsDeniedPeriod), &parsedMacosSettingsDeniedPeriod); err != nil {
							return fmt.Errorf("parsing --macos-settings-denied-period JSON: %w", err)
						}
						nestedMacosSettings["denied_period"] = parsedMacosSettingsDeniedPeriod
					}
					if bodyMacosSettingsEnabled {
						nestedMacosSettings["enabled"] = bodyMacosSettingsEnabled
					}
					if bodyMacosSettingsHoursAfterRelease != 0 {
						nestedMacosSettings["hours_after_release"] = bodyMacosSettingsHoursAfterRelease
					}
					if bodyMacosSettingsInstallAction != "" {
						nestedMacosSettings["install_action"] = bodyMacosSettingsInstallAction
					}
					if bodyMacosSettingsKeepOsUpdated {
						nestedMacosSettings["keep_os_updated"] = bodyMacosSettingsKeepOsUpdated
					}
					if bodyMacosSettingsMaxOsVersionAllowed != "" {
						nestedMacosSettings["max_os_version_allowed"] = bodyMacosSettingsMaxOsVersionAllowed
					}
					if bodyMacosSettingsMaxUserDeferrals != 0 {
						nestedMacosSettings["max_user_deferrals"] = bodyMacosSettingsMaxUserDeferrals
					}
					if bodyMacosSettingsMinutesAfterRelease != 0 {
						nestedMacosSettings["minutes_after_release"] = bodyMacosSettingsMinutesAfterRelease
					}
					if bodyMacosSettingsResendUpdateCommandHour != 0 {
						nestedMacosSettings["resend_update_command_hour"] = bodyMacosSettingsResendUpdateCommandHour
					}
					if len(nestedMacosSettings) > 0 {
						body["macos_settings"] = nestedMacosSettings
					}
				}
				if bodyPolicyId != "" {
					body["policy_id"] = bodyPolicyId
				}
				{
					nestedSchedule := map[string]any{}
					if bodyScheduleCutOffTime != "" {
						nestedSchedule["cut_off_time"] = bodyScheduleCutOffTime
					}
					if bodyScheduleEnabled {
						nestedSchedule["enabled"] = bodyScheduleEnabled
					}
					if bodyScheduleMaintenanceWindow != "" {
						nestedSchedule["maintenance_window"] = bodyScheduleMaintenanceWindow
					}
					{
						nestedScheduleStartingTime := map[string]any{}
						if bodyScheduleStartingTimeHour != "" {
							nestedScheduleStartingTime["hour"] = bodyScheduleStartingTimeHour
						}
						if bodyScheduleStartingTimeMinute != "" {
							nestedScheduleStartingTime["minute"] = bodyScheduleStartingTimeMinute
						}
						if len(nestedScheduleStartingTime) > 0 {
							nestedSchedule["starting_time"] = nestedScheduleStartingTime
						}
					}
					if bodyScheduleWeekDays != "" {
						var parsedScheduleWeekDays any
						if err := json.Unmarshal([]byte(bodyScheduleWeekDays), &parsedScheduleWeekDays); err != nil {
							return fmt.Errorf("parsing --schedule-week-days JSON: %w", err)
						}
						nestedSchedule["week_days"] = parsedScheduleWeekDays
					}
					if len(nestedSchedule) > 0 {
						body["schedule"] = nestedSchedule
					}
				}
				{
					nestedTvosSettings := map[string]any{}
					if bodyTvosSettingsAllowBetaUpdatesInDdm {
						nestedTvosSettings["allow_beta_updates_in_ddm"] = bodyTvosSettingsAllowBetaUpdatesInDdm
					}
					if bodyTvosSettingsAllowedDays != "" {
						var parsedTvosSettingsAllowedDays any
						if err := json.Unmarshal([]byte(bodyTvosSettingsAllowedDays), &parsedTvosSettingsAllowedDays); err != nil {
							return fmt.Errorf("parsing --tvos-settings-allowed-days JSON: %w", err)
						}
						nestedTvosSettings["allowed_days"] = parsedTvosSettingsAllowedDays
					}
					if bodyTvosSettingsDaysAfterRelease != 0 {
						nestedTvosSettings["days_after_release"] = bodyTvosSettingsDaysAfterRelease
					}
					if bodyTvosSettingsDaysAfterReleaseRsr != 0 {
						nestedTvosSettings["days_after_release_rsr"] = bodyTvosSettingsDaysAfterReleaseRsr
					}
					if bodyTvosSettingsDeniedPeriod != "" {
						var parsedTvosSettingsDeniedPeriod any
						if err := json.Unmarshal([]byte(bodyTvosSettingsDeniedPeriod), &parsedTvosSettingsDeniedPeriod); err != nil {
							return fmt.Errorf("parsing --tvos-settings-denied-period JSON: %w", err)
						}
						nestedTvosSettings["denied_period"] = parsedTvosSettingsDeniedPeriod
					}
					if bodyTvosSettingsEnabled {
						nestedTvosSettings["enabled"] = bodyTvosSettingsEnabled
					}
					if bodyTvosSettingsHoursAfterRelease != 0 {
						nestedTvosSettings["hours_after_release"] = bodyTvosSettingsHoursAfterRelease
					}
					if bodyTvosSettingsInstallAction != "" {
						nestedTvosSettings["install_action"] = bodyTvosSettingsInstallAction
					}
					if bodyTvosSettingsKeepOsUpdated {
						nestedTvosSettings["keep_os_updated"] = bodyTvosSettingsKeepOsUpdated
					}
					if bodyTvosSettingsMaxOsVersionAllowed != "" {
						nestedTvosSettings["max_os_version_allowed"] = bodyTvosSettingsMaxOsVersionAllowed
					}
					if bodyTvosSettingsMaxUserDeferrals != 0 {
						nestedTvosSettings["max_user_deferrals"] = bodyTvosSettingsMaxUserDeferrals
					}
					if bodyTvosSettingsMinutesAfterRelease != 0 {
						nestedTvosSettings["minutes_after_release"] = bodyTvosSettingsMinutesAfterRelease
					}
					if bodyTvosSettingsResendUpdateCommandHour != 0 {
						nestedTvosSettings["resend_update_command_hour"] = bodyTvosSettingsResendUpdateCommandHour
					}
					if len(nestedTvosSettings) > 0 {
						body["tvos_settings"] = nestedTvosSettings
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
					"resource": "system-updates",
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
	cmd.Flags().BoolVar(&bodyIosSettingsAllowBetaUpdatesInDdm, "ios-settings-allow-beta-updates-in-ddm", false, "Allow beta updates in ddm")
	cmd.Flags().StringVar(&bodyIosSettingsAllowedDays, "ios-settings-allowed-days", "", "Allowed days")
	cmd.Flags().IntVar(&bodyIosSettingsDaysAfterRelease, "ios-settings-days-after-release", 0, "Days after release")
	cmd.Flags().IntVar(&bodyIosSettingsDaysAfterReleaseRsr, "ios-settings-days-after-release-rsr", 0, "Days after release rsr")
	cmd.Flags().StringVar(&bodyIosSettingsDeniedPeriod, "ios-settings-denied-period", "", "Denied period")
	cmd.Flags().BoolVar(&bodyIosSettingsEnabled, "ios-settings-enabled", false, "Enabled")
	cmd.Flags().IntVar(&bodyIosSettingsHoursAfterRelease, "ios-settings-hours-after-release", 0, "Hours after release")
	cmd.Flags().StringVar(&bodyIosSettingsInstallAction, "ios-settings-install-action", "", "Install action")
	cmd.Flags().BoolVar(&bodyIosSettingsKeepOsUpdated, "ios-settings-keep-os-updated", false, "Keep os updated")
	cmd.Flags().StringVar(&bodyIosSettingsMaxOsVersionAllowed, "ios-settings-max-os-version-allowed", "", "Max os version allowed")
	cmd.Flags().IntVar(&bodyIosSettingsMaxUserDeferrals, "ios-settings-max-user-deferrals", 0, "Max user deferrals")
	cmd.Flags().IntVar(&bodyIosSettingsMinutesAfterRelease, "ios-settings-minutes-after-release", 0, "Minutes after release")
	cmd.Flags().IntVar(&bodyIosSettingsResendUpdateCommandHour, "ios-settings-resend-update-command-hour", 0, "Resend update command hour")
	cmd.Flags().BoolVar(&bodyIpadosSettingsAllowBetaUpdatesInDdm, "ipados-settings-allow-beta-updates-in-ddm", false, "Allow beta updates in ddm")
	cmd.Flags().StringVar(&bodyIpadosSettingsAllowedDays, "ipados-settings-allowed-days", "", "Allowed days")
	cmd.Flags().IntVar(&bodyIpadosSettingsDaysAfterRelease, "ipados-settings-days-after-release", 0, "Days after release")
	cmd.Flags().IntVar(&bodyIpadosSettingsDaysAfterReleaseRsr, "ipados-settings-days-after-release-rsr", 0, "Days after release rsr")
	cmd.Flags().StringVar(&bodyIpadosSettingsDeniedPeriod, "ipados-settings-denied-period", "", "Denied period")
	cmd.Flags().BoolVar(&bodyIpadosSettingsEnabled, "ipados-settings-enabled", false, "Enabled")
	cmd.Flags().IntVar(&bodyIpadosSettingsHoursAfterRelease, "ipados-settings-hours-after-release", 0, "Hours after release")
	cmd.Flags().StringVar(&bodyIpadosSettingsInstallAction, "ipados-settings-install-action", "", "Install action")
	cmd.Flags().BoolVar(&bodyIpadosSettingsKeepOsUpdated, "ipados-settings-keep-os-updated", false, "Keep os updated")
	cmd.Flags().StringVar(&bodyIpadosSettingsMaxOsVersionAllowed, "ipados-settings-max-os-version-allowed", "", "Max os version allowed")
	cmd.Flags().IntVar(&bodyIpadosSettingsMaxUserDeferrals, "ipados-settings-max-user-deferrals", 0, "Max user deferrals")
	cmd.Flags().IntVar(&bodyIpadosSettingsMinutesAfterRelease, "ipados-settings-minutes-after-release", 0, "Minutes after release")
	cmd.Flags().IntVar(&bodyIpadosSettingsResendUpdateCommandHour, "ipados-settings-resend-update-command-hour", 0, "Resend update command hour")
	cmd.Flags().BoolVar(&bodyMacosSettingsAllowBetaUpdatesInDdm, "macos-settings-allow-beta-updates-in-ddm", false, "Allow beta updates in ddm")
	cmd.Flags().StringVar(&bodyMacosSettingsAllowedDays, "macos-settings-allowed-days", "", "Allowed days")
	cmd.Flags().IntVar(&bodyMacosSettingsDaysAfterRelease, "macos-settings-days-after-release", 0, "Days after release")
	cmd.Flags().IntVar(&bodyMacosSettingsDaysAfterReleaseRsr, "macos-settings-days-after-release-rsr", 0, "Days after release rsr")
	cmd.Flags().StringVar(&bodyMacosSettingsDeniedPeriod, "macos-settings-denied-period", "", "Denied period")
	cmd.Flags().BoolVar(&bodyMacosSettingsEnabled, "macos-settings-enabled", false, "Enabled")
	cmd.Flags().IntVar(&bodyMacosSettingsHoursAfterRelease, "macos-settings-hours-after-release", 0, "Hours after release")
	cmd.Flags().StringVar(&bodyMacosSettingsInstallAction, "macos-settings-install-action", "", "Install action")
	cmd.Flags().BoolVar(&bodyMacosSettingsKeepOsUpdated, "macos-settings-keep-os-updated", false, "Keep os updated")
	cmd.Flags().StringVar(&bodyMacosSettingsMaxOsVersionAllowed, "macos-settings-max-os-version-allowed", "", "Max os version allowed")
	cmd.Flags().IntVar(&bodyMacosSettingsMaxUserDeferrals, "macos-settings-max-user-deferrals", 0, "Max user deferrals")
	cmd.Flags().IntVar(&bodyMacosSettingsMinutesAfterRelease, "macos-settings-minutes-after-release", 0, "Minutes after release")
	cmd.Flags().IntVar(&bodyMacosSettingsResendUpdateCommandHour, "macos-settings-resend-update-command-hour", 0, "Resend update command hour")
	cmd.Flags().StringVar(&bodyPolicyId, "policy-id", "", "Policy id")
	cmd.Flags().StringVar(&bodyScheduleCutOffTime, "schedule-cut-off-time", "", "Cut off time")
	cmd.Flags().BoolVar(&bodyScheduleEnabled, "schedule-enabled", false, "Enabled")
	cmd.Flags().StringVar(&bodyScheduleMaintenanceWindow, "schedule-maintenance-window", "", "Maintenance window")
	cmd.Flags().StringVar(&bodyScheduleStartingTimeHour, "schedule-starting-time-hour", "", "Hour")
	cmd.Flags().StringVar(&bodyScheduleStartingTimeMinute, "schedule-starting-time-minute", "", "Minute")
	cmd.Flags().StringVar(&bodyScheduleWeekDays, "schedule-week-days", "", "Week days")
	cmd.Flags().BoolVar(&bodyTvosSettingsAllowBetaUpdatesInDdm, "tvos-settings-allow-beta-updates-in-ddm", false, "Allow beta updates in ddm")
	cmd.Flags().StringVar(&bodyTvosSettingsAllowedDays, "tvos-settings-allowed-days", "", "Allowed days")
	cmd.Flags().IntVar(&bodyTvosSettingsDaysAfterRelease, "tvos-settings-days-after-release", 0, "Days after release")
	cmd.Flags().IntVar(&bodyTvosSettingsDaysAfterReleaseRsr, "tvos-settings-days-after-release-rsr", 0, "Days after release rsr")
	cmd.Flags().StringVar(&bodyTvosSettingsDeniedPeriod, "tvos-settings-denied-period", "", "Denied period")
	cmd.Flags().BoolVar(&bodyTvosSettingsEnabled, "tvos-settings-enabled", false, "Enabled")
	cmd.Flags().IntVar(&bodyTvosSettingsHoursAfterRelease, "tvos-settings-hours-after-release", 0, "Hours after release")
	cmd.Flags().StringVar(&bodyTvosSettingsInstallAction, "tvos-settings-install-action", "", "Install action")
	cmd.Flags().BoolVar(&bodyTvosSettingsKeepOsUpdated, "tvos-settings-keep-os-updated", false, "Keep os updated")
	cmd.Flags().StringVar(&bodyTvosSettingsMaxOsVersionAllowed, "tvos-settings-max-os-version-allowed", "", "Max os version allowed")
	cmd.Flags().IntVar(&bodyTvosSettingsMaxUserDeferrals, "tvos-settings-max-user-deferrals", 0, "Max user deferrals")
	cmd.Flags().IntVar(&bodyTvosSettingsMinutesAfterRelease, "tvos-settings-minutes-after-release", 0, "Minutes after release")
	cmd.Flags().IntVar(&bodyTvosSettingsResendUpdateCommandHour, "tvos-settings-resend-update-command-hour", 0, "Resend update command hour")
	cmd.Flags().BoolVar(&stdinBody, "stdin", false, "Read request body as JSON from stdin")

	return cmd
}
