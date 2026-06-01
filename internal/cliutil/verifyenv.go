// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cliutil

import "os"

// VerifyEnvVar is the env var the verifier sets in every mock-mode
// subprocess. Commands that perform visible side effects (open browser
// tabs, send notifications, dial out to OS handlers) MUST short-circuit
// when this env var is "1" to avoid spamming the user's environment
// during verify runs.
const VerifyEnvVar = "ADDIGY_CLI_VERIFY"

// IsVerifyEnv reports whether the current process is running under the
// verifier in mock mode. Commands with side
// effects pair this check with print-by-default + explicit opt-in
// (--launch, --send, --play) so a verify pass on a fresh CLI does not
// pop browser tabs or fire off real notifications.
//
// Defense-in-depth: even if the verifier's heuristic side-effect
// classifier misses a command, this env-var short-circuit catches it.
//
//	if cliutil.IsVerifyEnv() {
//	    fmt.Fprintln(cmd.OutOrStdout(), "would launch:", url)
//	    return nil
//	}
func IsVerifyEnv() bool {
	return os.Getenv(VerifyEnvVar) == "1"
}
