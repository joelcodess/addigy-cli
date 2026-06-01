// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newMaintenanceCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "maintenance",
		Short:  "Manage maintenance",
		Hidden: true,
	}

	cmd.AddCommand(newMaintenanceCreateCmd(flags))
	cmd.AddCommand(newMaintenanceCreatePolicyCmd(flags))
	cmd.AddCommand(newMaintenanceCreateQueryCmd(flags))
	cmd.AddCommand(newMaintenanceCreateStagedCmd(flags))
	cmd.AddCommand(newMaintenanceCreateStaged2Cmd(flags))
	cmd.AddCommand(newMaintenanceCreateStaged3Cmd(flags))
	cmd.AddCommand(newMaintenanceDeleteCmd(flags))
	cmd.AddCommand(newMaintenanceDeletePolicyCmd(flags))
	cmd.AddCommand(newMaintenanceDeleteStagedCmd(flags))
	cmd.AddCommand(newMaintenanceUpdateCmd(flags))
	cmd.AddCommand(newMaintenanceUpdateStagedCmd(flags))
	return cmd
}
