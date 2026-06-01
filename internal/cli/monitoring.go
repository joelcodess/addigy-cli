// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newMonitoringCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "monitoring",
		Short:  "Manage monitoring",
		Hidden: true,
	}

	cmd.AddCommand(newMonitoringCreateCmd(flags))
	cmd.AddCommand(newMonitoringCreatePolicyCmd(flags))
	cmd.AddCommand(newMonitoringCreateQueryCmd(flags))
	cmd.AddCommand(newMonitoringDeleteCmd(flags))
	cmd.AddCommand(newMonitoringDeletePolicyCmd(flags))
	cmd.AddCommand(newMonitoringListCmd(flags))
	cmd.AddCommand(newMonitoringUpdateCmd(flags))
	return cmd
}
