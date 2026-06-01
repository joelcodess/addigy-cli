// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newDeviceScriptAssignmentsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "device-script-assignments",
		Short:  "Manage device script assignments",
		Hidden: true,
	}

	cmd.AddCommand(newDeviceScriptAssignmentsCreateCmd(flags))
	cmd.AddCommand(newDeviceScriptAssignmentsDeleteCmd(flags))
	cmd.AddCommand(newDeviceScriptAssignmentsListCmd(flags))
	return cmd
}
