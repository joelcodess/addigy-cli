// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newODdmUpdatesCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ddm-updates",
		Short: "Manage ddm updates",
	}

	cmd.AddCommand(newODdmUpdatesGetCmd(flags))
	return cmd
}
