// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newFactsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "facts",
		Short:  "Manage facts",
		Hidden: true,
	}

	cmd.AddCommand(newFactsCreateCmd(flags))
	cmd.AddCommand(newFactsCreateCustomCmd(flags))
	cmd.AddCommand(newFactsCreateCustom2Cmd(flags))
	cmd.AddCommand(newFactsDeleteCmd(flags))
	cmd.AddCommand(newFactsDeleteCustomCmd(flags))
	cmd.AddCommand(newFactsListCmd(flags))
	cmd.AddCommand(newFactsUpdateCmd(flags))
	return cmd
}
