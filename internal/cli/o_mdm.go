// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOMdmCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mdm",
		Short: "Manage mdm",
	}

	cmd.AddCommand(newOMdmCreateCmd(flags))
	cmd.AddCommand(newOMdmCreateOCmd(flags))
	return cmd
}
