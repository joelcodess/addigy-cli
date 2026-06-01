// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOaCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "oa",
		Short:  "Manage oa",
		Hidden: true,
	}

	cmd.AddCommand(newOaCreateCmd(flags))
	cmd.AddCommand(newOaCreateAdeCmd(flags))
	cmd.AddCommand(newOaCreateAppsandbooksCmd(flags))
	cmd.AddCommand(newOaCreateCompliancerulesCmd(flags))
	cmd.AddCommand(newOaCreateDevicesCmd(flags))
	cmd.AddCommand(newOaCreateFilesCmd(flags))
	cmd.AddCommand(newOaCreateIdentityCmd(flags))
	cmd.AddCommand(newOaCreateInstalledappsCmd(flags))
	cmd.AddCommand(newOaCreateIntegrationsCmd(flags))
	cmd.AddCommand(newOaCreateMonitoringCmd(flags))
	cmd.AddCommand(newOaCreatePoliciesCmd(flags))
	cmd.AddCommand(newOaCreatePolicies2Cmd(flags))
	cmd.AddCommand(newOaCreatePrebuiltappsCmd(flags))
	cmd.AddCommand(newOaCreateReportsCmd(flags))
	cmd.AddCommand(newOaCreateVariablesCmd(flags))
	cmd.AddCommand(newOaCreateWebhooksCmd(flags))
	cmd.AddCommand(newOaCreateWebhooks2Cmd(flags))
	cmd.AddCommand(newOaCreateWebhooks3Cmd(flags))
	cmd.AddCommand(newOaListCmd(flags))
	cmd.AddCommand(newOaListBenchmarksCmd(flags))
	cmd.AddCommand(newOaListCompliancerulesCmd(flags))
	cmd.AddCommand(newOaListCompliancerules2Cmd(flags))
	cmd.AddCommand(newOaListIntegrationsCmd(flags))
	cmd.AddCommand(newOaListReportsCmd(flags))
	cmd.AddCommand(newOaListReports2Cmd(flags))
	cmd.AddCommand(newOaListSelfserviceCmd(flags))
	return cmd
}
