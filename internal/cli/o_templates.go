// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOTemplatesCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "templates",
		Short: "Manage templates",
	}

	cmd.AddCommand(newOTemplatesCreateCmd(flags))
	cmd.AddCommand(newOTemplatesCreateOCmd(flags))
	return cmd
}
