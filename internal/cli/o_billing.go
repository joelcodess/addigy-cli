// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOBillingCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "billing",
		Short: "Manage billing",
	}

	cmd.AddCommand(newOBillingCreateCmd(flags))
	cmd.AddCommand(newOBillingDeleteCmd(flags))
	cmd.AddCommand(newOBillingGetCmd(flags))
	cmd.AddCommand(newOBillingGetOCmd(flags))
	cmd.AddCommand(newOBillingGetO2Cmd(flags))
	cmd.AddCommand(newOBillingGetO3Cmd(flags))
	cmd.AddCommand(newOBillingGetO4Cmd(flags))
	cmd.AddCommand(newOBillingGetO5Cmd(flags))
	cmd.AddCommand(newOBillingUpdateCmd(flags))
	cmd.AddCommand(newOBillingUpdateOCmd(flags))
	return cmd
}
