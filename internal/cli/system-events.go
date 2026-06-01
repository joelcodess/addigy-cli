// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newSystemEventsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "system-events",
		Short:  "Manage system events",
		Hidden: true,
	}

	cmd.AddCommand(newSystemEventsCreateCmd(flags))
	cmd.AddCommand(newSystemEventsCreateSystemeventsCmd(flags))
	return cmd
}
