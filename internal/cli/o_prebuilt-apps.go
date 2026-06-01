// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOPrebuiltAppsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prebuilt-apps",
		Short: "Manage prebuilt apps",
	}

	cmd.AddCommand(newOPrebuiltAppsCreateCmd(flags))
	cmd.AddCommand(newOPrebuiltAppsDeleteCmd(flags))
	cmd.AddCommand(newOPrebuiltAppsDeleteOCmd(flags))
	cmd.AddCommand(newOPrebuiltAppsUpdateCmd(flags))
	cmd.AddCommand(newOPrebuiltAppsUpdateOCmd(flags))
	return cmd
}
