// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newManagedAppConfigurationsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "managed-app-configurations",
		Short:  "Manage managed app configurations",
		Hidden: true,
	}

	cmd.AddCommand(newManagedAppConfigurationsCreateCmd(flags))
	cmd.AddCommand(newManagedAppConfigurationsDeleteCmd(flags))
	cmd.AddCommand(newManagedAppConfigurationsListCmd(flags))
	return cmd
}
