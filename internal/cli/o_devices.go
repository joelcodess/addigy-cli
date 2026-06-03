// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newODevicesCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "devices",
		Short: "Manage devices",
	}

	cmd.AddCommand(newODevicesCreateCmd(flags))
	cmd.AddCommand(newODevicesCommandsCmd(flags))
	return cmd
}
