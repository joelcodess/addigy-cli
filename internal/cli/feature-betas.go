// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newFeatureBetasCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "feature-betas",
		Short:  "Manage feature betas",
		Hidden: true,
	}

	cmd.AddCommand(newFeatureBetasCreateCmd(flags))
	cmd.AddCommand(newFeatureBetasDeleteCmd(flags))
	cmd.AddCommand(newFeatureBetasListCmd(flags))
	return cmd
}
