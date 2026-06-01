// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newMdmCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "mdm",
		Short:  "Manage mdm",
		Hidden: true,
	}

	cmd.AddCommand(newMdmCreateCmd(flags))
	cmd.AddCommand(newMdmCreateCommandsCmd(flags))
	cmd.AddCommand(newMdmCreateConfigurationsCmd(flags))
	cmd.AddCommand(newMdmCreateConfigurations2Cmd(flags))
	cmd.AddCommand(newMdmCreateConfigurations3Cmd(flags))
	cmd.AddCommand(newMdmCreateDevicesCmd(flags))
	cmd.AddCommand(newMdmCreateProfilesCmd(flags))
	cmd.AddCommand(newMdmDeleteCmd(flags))
	cmd.AddCommand(newMdmDeleteConfigurationsCmd(flags))
	cmd.AddCommand(newMdmDeleteConfigurations2Cmd(flags))
	cmd.AddCommand(newMdmGetCmd(flags))
	cmd.AddCommand(newMdmGetConfigurationsCmd(flags))
	cmd.AddCommand(newMdmGetConfigurations2Cmd(flags))
	cmd.AddCommand(newMdmGetDevicesCmd(flags))
	cmd.AddCommand(newMdmListCmd(flags))
	cmd.AddCommand(newMdmListCommandsCmd(flags))
	cmd.AddCommand(newMdmListConfigurationsCmd(flags))
	cmd.AddCommand(newMdmListConfigurations2Cmd(flags))
	cmd.AddCommand(newMdmListConfigurations3Cmd(flags))
	cmd.AddCommand(newMdmUpdateCmd(flags))
	return cmd
}
