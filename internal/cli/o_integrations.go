// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOIntegrationsCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "integrations",
		Short: "Manage integrations",
	}

	cmd.AddCommand(newOIntegrationsCreateCmd(flags))
	cmd.AddCommand(newOIntegrationsCreateOCmd(flags))
	cmd.AddCommand(newOIntegrationsCreateO2Cmd(flags))
	cmd.AddCommand(newOIntegrationsCreateO3Cmd(flags))
	cmd.AddCommand(newOIntegrationsDeleteCmd(flags))
	cmd.AddCommand(newOIntegrationsGetCmd(flags))
	cmd.AddCommand(newOIntegrationsGetOCmd(flags))
	cmd.AddCommand(newOIntegrationsGetO2Cmd(flags))
	return cmd
}
