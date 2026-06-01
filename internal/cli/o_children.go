// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOChildrenCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "children",
		Short: "Manage children",
	}

	cmd.AddCommand(newOChildrenGetCmd(flags))
	return cmd
}
