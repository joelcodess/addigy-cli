// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"github.com/spf13/cobra"
)

func newOComplianceRulesCmd(flags *rootFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compliance-rules",
		Short: "Manage compliance rules",
	}

	cmd.AddCommand(newOComplianceRulesCreateCmd(flags))
	cmd.AddCommand(newOComplianceRulesDeleteCmd(flags))
	cmd.AddCommand(newOComplianceRulesGetCmd(flags))
	cmd.AddCommand(newOComplianceRulesUpdateCmd(flags))
	return cmd
}
