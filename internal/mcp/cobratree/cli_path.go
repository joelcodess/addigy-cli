// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cobratree

import (
	"os"
	"os/exec"
	"path/filepath"
)

// SiblingCLIPath resolves the companion CLI via sibling-of-executable,
// ADDIGY_CLI_PATH env var, then PATH.
func SiblingCLIPath() (string, error) {
	const cliName = "addigy-cli"
	if exe, err := os.Executable(); err == nil {
		candidate := filepath.Join(filepath.Dir(exe), cliName)
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
	}
	if v := os.Getenv("ADDIGY_CLI_PATH"); v != "" {
		return v, nil
	}
	return exec.LookPath(cliName)
}
