// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newAssetsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "assets",
		Short:  "Manage assets",
		Hidden: true,
	}

	cmd.AddCommand(newAssetsCreateCmd(flags))
	cmd.AddCommand(newAssetsCreateDefaultCmd(flags))
	cmd.AddCommand(newAssetsCreateDefault2Cmd(flags))
	cmd.AddCommand(newAssetsCreateDefault3Cmd(flags))
	return cmd
}
