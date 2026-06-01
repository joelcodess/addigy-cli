// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOPolicyCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policy",
		Short: "Manage policy",
	}

	cmd.AddCommand(newOPolicyCreateCmd(flags))
	cmd.AddCommand(newOPolicyCreateOCmd(flags))
	cmd.AddCommand(newOPolicyDeleteCmd(flags))
	cmd.AddCommand(newOPolicyDeleteOCmd(flags))
	cmd.AddCommand(newOPolicyGetCmd(flags))
	return cmd
}
