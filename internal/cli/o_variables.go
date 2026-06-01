// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOVariablesCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "variables",
		Short: "Manage variables",
	}

	cmd.AddCommand(newOVariablesCreateCmd(flags))
	cmd.AddCommand(newOVariablesCreateOCmd(flags))
	cmd.AddCommand(newOVariablesDeleteCmd(flags))
	cmd.AddCommand(newOVariablesDeleteOCmd(flags))
	cmd.AddCommand(newOVariablesGetCmd(flags))
	cmd.AddCommand(newOVariablesGetOCmd(flags))
	cmd.AddCommand(newOVariablesGetO2Cmd(flags))
	cmd.AddCommand(newOVariablesGetO3Cmd(flags))
	cmd.AddCommand(newOVariablesUpdateCmd(flags))
	return cmd
}
