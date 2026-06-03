// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

// newODevicesCommandsCmd groups the device command-execution endpoints:
// POST /o/{organization_id}/devices/commands/run and
// GET  /o/{organization_id}/devices/{agent_id}/commands/{action_id}/output.
func newODevicesCommandsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commands",
		Short: "Run commands on devices and read their output",
	}

	cmd.AddCommand(newODevicesCommandsRunCmd(flags))
	cmd.AddCommand(newODevicesCommandsOutputCmd(flags))
	return cmd
}
