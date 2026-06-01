// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newStaticFieldsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "static-fields",
		Short:  "Manage static fields",
		Hidden: true,
	}

	cmd.AddCommand(newStaticFieldsCreateCmd(flags))
	cmd.AddCommand(newStaticFieldsCreateStaticfieldsCmd(flags))
	cmd.AddCommand(newStaticFieldsDeleteCmd(flags))
	cmd.AddCommand(newStaticFieldsListCmd(flags))
	cmd.AddCommand(newStaticFieldsListStaticfieldsCmd(flags))
	cmd.AddCommand(newStaticFieldsUpdateCmd(flags))
	return cmd
}
