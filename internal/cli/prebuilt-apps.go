// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newPrebuiltAppsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "prebuilt-apps",
		Short:  "Manage prebuilt apps",
		Hidden: true,
	}

	cmd.AddCommand(newPrebuiltAppsCreateCmd(flags))
	cmd.AddCommand(newPrebuiltAppsCreatePrebuiltappsCmd(flags))
	cmd.AddCommand(newPrebuiltAppsCreatePrebuiltapps2Cmd(flags))
	cmd.AddCommand(newPrebuiltAppsCreatePrebuiltapps3Cmd(flags))
	cmd.AddCommand(newPrebuiltAppsDeleteCmd(flags))
	cmd.AddCommand(newPrebuiltAppsDeletePrebuiltappsCmd(flags))
	cmd.AddCommand(newPrebuiltAppsGetCmd(flags))
	cmd.AddCommand(newPrebuiltAppsGetPrebuiltappsCmd(flags))
	cmd.AddCommand(newPrebuiltAppsUpdateCmd(flags))
	cmd.AddCommand(newPrebuiltAppsUpdatePrebuiltappsCmd(flags))
	return cmd
}
