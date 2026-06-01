// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOUsersCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "Manage users",
	}

	cmd.AddCommand(newOUsersCreateCmd(flags))
	return cmd
}
