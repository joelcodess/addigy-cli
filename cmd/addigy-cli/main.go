// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package main

import (
	"os"

	"github.com/joelcodess/addigy-cli/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		// Propagate the typed exit code (auth=4, not-found=3, usage=2, rate-limit=7,
		// api=5, config=10, …) computed by the command layer, not a flat 1, so
		// agents and scripts can branch on the failure class. See README exit codes.
		os.Exit(cli.ExitCode(err))
	}
}
