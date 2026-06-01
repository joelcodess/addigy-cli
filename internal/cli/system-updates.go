// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newSystemUpdatesCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "system-updates",
		Short:  "Manage system updates",
		Hidden: true,
	}

	cmd.AddCommand(newSystemUpdatesCreateCmd(flags))
	cmd.AddCommand(newSystemUpdatesCreateSystemupdatesCmd(flags))
	cmd.AddCommand(newSystemUpdatesCreateSystemupdates2Cmd(flags))
	cmd.AddCommand(newSystemUpdatesCreateSystemupdates3Cmd(flags))
	cmd.AddCommand(newSystemUpdatesCreateSystemupdates4Cmd(flags))
	cmd.AddCommand(newSystemUpdatesCreateSystemupdates5Cmd(flags))
	cmd.AddCommand(newSystemUpdatesCreateSystemupdates6Cmd(flags))
	cmd.AddCommand(newSystemUpdatesCreateSystemupdates7Cmd(flags))
	cmd.AddCommand(newSystemUpdatesCreateSystemupdates8Cmd(flags))
	cmd.AddCommand(newSystemUpdatesCreateSystemupdates9Cmd(flags))
	cmd.AddCommand(newSystemUpdatesListCmd(flags))
	cmd.AddCommand(newSystemUpdatesListSystemupdatesCmd(flags))
	cmd.AddCommand(newSystemUpdatesListSystemupdates2Cmd(flags))
	cmd.AddCommand(newSystemUpdatesListSystemupdates3Cmd(flags))
	cmd.AddCommand(newSystemUpdatesListSystemupdates4Cmd(flags))
	cmd.AddCommand(newSystemUpdatesListSystemupdates5Cmd(flags))
	cmd.AddCommand(newSystemUpdatesListSystemupdates6Cmd(flags))
	return cmd
}
