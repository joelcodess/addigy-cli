// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOPoliciesCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policies",
		Short: "Manage policies",
	}

	cmd.AddCommand(newOPoliciesCreateCmd(flags))
	cmd.AddCommand(newOPoliciesCreateOCmd(flags))
	cmd.AddCommand(newOPoliciesDeleteCmd(flags))
	cmd.AddCommand(newOPoliciesDeleteOCmd(flags))
	cmd.AddCommand(newOPoliciesDeleteO2Cmd(flags))
	cmd.AddCommand(newOPoliciesGetCmd(flags))
	cmd.AddCommand(newOPoliciesUpdateCmd(flags))
	cmd.AddCommand(newOPoliciesUpdateOCmd(flags))
	return cmd
}
