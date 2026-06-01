// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newODeviceCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "device",
		Short: "Manage device",
	}

	cmd.AddCommand(newODeviceGetCmd(flags))
	return cmd
}
