// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOHomescreenCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "homescreen",
		Short: "Manage homescreen",
	}

	cmd.AddCommand(newOHomescreenCreateCmd(flags))
	cmd.AddCommand(newOHomescreenCreateOCmd(flags))
	cmd.AddCommand(newOHomescreenDeleteCmd(flags))
	return cmd
}
