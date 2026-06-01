// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOCommunityCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "community",
		Short: "Manage community",
	}

	cmd.AddCommand(newOCommunityCreateCmd(flags))
	return cmd
}
