// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOIdentityCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "identity",
		Short: "Manage identity",
	}

	cmd.AddCommand(newOIdentityUpdateCmd(flags))
	return cmd
}
