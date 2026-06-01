// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOReportsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reports",
		Short: "Manage reports",
	}

	cmd.AddCommand(newOReportsCreateCmd(flags))
	return cmd
}
