// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOBenchmarksCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "benchmarks",
		Short: "Manage benchmarks",
	}

	cmd.AddCommand(newOBenchmarksCreateCmd(flags))
	cmd.AddCommand(newOBenchmarksCreateOCmd(flags))
	cmd.AddCommand(newOBenchmarksDeleteCmd(flags))
	cmd.AddCommand(newOBenchmarksUpdateCmd(flags))
	return cmd
}
