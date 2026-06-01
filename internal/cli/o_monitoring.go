// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOMonitoringCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitoring",
		Short: "Manage monitoring",
	}

	cmd.AddCommand(newOMonitoringCreateCmd(flags))
	cmd.AddCommand(newOMonitoringCreateOCmd(flags))
	cmd.AddCommand(newOMonitoringDeleteCmd(flags))
	cmd.AddCommand(newOMonitoringUpdateCmd(flags))
	return cmd
}
