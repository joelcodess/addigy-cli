// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOWebhooksCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "webhooks",
		Short: "Manage webhooks",
	}

	cmd.AddCommand(newOWebhooksCreateCmd(flags))
	cmd.AddCommand(newOWebhooksDeleteCmd(flags))
	cmd.AddCommand(newOWebhooksUpdateCmd(flags))
	return cmd
}
