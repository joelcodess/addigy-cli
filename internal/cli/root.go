// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/joelcodess/addigy-cli/internal/client"
	"github.com/joelcodess/addigy-cli/internal/config"
	"github.com/spf13/cobra"
)

var version = "1.0.0"

// init brands the HTTP client's User-Agent with the real build version and
// platform so every API request is attributable to this CLI. version is set
// via -ldflags -X at build time; those values are in place before init runs.
// Runs for both the CLI (Execute) and the MCP server, which imports this
// package for the command tree.
func init() {
	client.UserAgent = fmt.Sprintf("addigy-cli/%s (%s/%s; +https://github.com/joelcodess/addigy-cli)", version, runtime.GOOS, runtime.GOARCH)
}

type rootFlags struct {
	asJSON        bool
	compact       bool
	csv           bool
	plain         bool
	quiet         bool
	dryRun        bool
	noCache       bool
	noInput       bool
	idempotent    bool
	ignoreMissing bool
	yes           bool
	agent         bool
	selectFields  string
	configPath    string
	profileName   string
	deliverSpec   string
	timeout       time.Duration
	rateLimit     float64
	dataSource    string
	freshnessMeta any

	// deliverBuf captures command output when --deliver is set to a
	// non-stdout sink. Flushed to the sink after Execute returns.
	deliverBuf  *bytes.Buffer
	deliverSink DeliverSink
}

// RootCmd returns the Cobra command tree without executing it. The MCP server
// uses this to mirror every user-facing command as an agent tool.
func RootCmd() *cobra.Command {
	var flags rootFlags
	return newRootCmd(&flags)
}

// Execute runs the CLI in non-interactive mode: never prompts, all values via flags or stdin.
func Execute() error {
	var flags rootFlags
	rootCmd := newRootCmd(&flags)

	err := rootCmd.Execute()
	if err != nil && strings.Contains(err.Error(), "unknown flag") {
		msg := err.Error()
		// Extract the flag name from the error message (e.g., "unknown flag: --foob")
		if idx := strings.Index(msg, "unknown flag: "); idx >= 0 {
			flagStr := strings.TrimSpace(msg[idx+len("unknown flag: "):])
			if suggestion := suggestFlag(flagStr, rootCmd); suggestion != "" {
				err = fmt.Errorf("%w\nhint: did you mean --%s?", err, suggestion)
			}
		}
	}
	if err == nil && flags.deliverBuf != nil {
		if derr := Deliver(flags.deliverSink, flags.deliverBuf.Bytes(), flags.compact); derr != nil {
			fmt.Fprintf(os.Stderr, "warning: deliver to %s:%s failed: %v\n", flags.deliverSink.Scheme, flags.deliverSink.Target, derr)
			return derr
		}
	}
	if err != nil && isCobraUsageError(err) {
		// Cobra/pflag pre-RunE errors (unknown flag, unknown command,
		// missing required, etc.) never flow through usageErr() because
		// they originate inside rootCmd.Execute() before any user RunE
		// runs. Without this wrap, ExitCode() falls through to the
		// default and emits 1 — clobbering the conventional code-2 for
		// usage errors that the helpers.go contract already promises.
		return usageErr(err)
	}
	return err
}

// isCobraUsageError reports whether err matches one of Cobra/pflag's
// pre-RunE usage-error shapes. Detection is by message prefix to match
// the same approach the unknown-flag hint path uses above; neither
// Cobra nor pflag exports typed sentinels for these.
//
// Patterns are anchored to the literal punctuation Cobra and pflag
// emit so an application's own RunE error that happens to contain the
// substring "required flag" or "invalid argument" doesn't get
// misclassified as a usage error.
//
// Patterns covered (Cobra v1.x + pflag v1.x as of 2026-05):
//   - "unknown flag: --foo"                            (pflag)
//   - "unknown shorthand flag: 'x' in -x"              (pflag)
//   - "unknown command \"foo\" for ..."                (Cobra)
//   - "required flag(s) \"foo\" not set"               (Cobra MarkFlagRequired)
//   - "flag needs an argument: --foo"                  (pflag, missing value)
//   - "invalid argument \"x\" for \"--y\" flag: ..."   (pflag, parse failure)
//
// Returns false for nil err.
func isCobraUsageError(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.HasPrefix(msg, "unknown flag") ||
		strings.HasPrefix(msg, "unknown shorthand flag") ||
		strings.HasPrefix(msg, "unknown command") ||
		strings.HasPrefix(msg, `required flag(s) "`) ||
		strings.HasPrefix(msg, "flag needs an argument:") ||
		strings.HasPrefix(msg, `invalid argument "`)
}

func newRootCmd(flags *rootFlags) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "addigy-cli",
		Short: `Addigy CLI — Every Addigy v2 endpoint, plus a local fleet mirror with compound queries the web UI can't answer.`,
		Long: `Addigy CLI — Every Addigy v2 endpoint, plus a local fleet mirror with compound queries the web UI can't answer.

Highlights (not in the official API docs):
  • devices stale   List devices whose last check-in is older than N days, with optional policy and OS filters.
  • compliance   Surface devices whose assigned policy rules are unmet, joined against current device facts.
  • rollout   Per-device install state for one Smart Software item across the assigned fleet, with success/pending/failed counts.
  • facts search   FTS5 across mirrored device facts; --group-by value returns a histogram of values per fact.
  • devices diff   Set differences across facts, applications, policies, and Smart-Software install state between two devices.
  • drift   Diffs the mirror's current rows against the prior snapshot for any entity (devices, facts, policies, software).
  • policy-coverage   Per-policy device counts joined to last-checkin so the user sees both coverage and freshness in one view.
  • fleet-summary   Single command emitting device count, stale fraction, alert count, MDM queue depth, Smart-Software pending count, and policy coverage percent.

Agent mode: add --agent to any command for JSON output + non-interactive mode.
Health check: run 'addigy-cli doctor' to verify auth and connectivity.
See README.md or the bundled SKILL.md for recipes.`,
		SilenceUsage: true,
		Version:      version,
	}
	rootCmd.SetVersionTemplate("addigy-cli {{ .Version }}\n")

	rootCmd.PersistentFlags().BoolVar(&flags.asJSON, "json", false, "Output as JSON")
	rootCmd.PersistentFlags().BoolVar(&flags.compact, "compact", false, "Return only key fields (id, name, status, timestamps) for minimal token usage")
	rootCmd.PersistentFlags().BoolVar(&flags.csv, "csv", false, "Output as CSV (table and array responses)")
	rootCmd.PersistentFlags().BoolVar(&flags.plain, "plain", false, "Output as plain tab-separated text")
	rootCmd.PersistentFlags().BoolVar(&flags.quiet, "quiet", false, "Bare output, one value per line")
	rootCmd.PersistentFlags().StringVar(&flags.configPath, "config", "", "Config file path")
	rootCmd.PersistentFlags().DurationVar(&flags.timeout, "timeout", 30*time.Second, "Request timeout")
	rootCmd.PersistentFlags().BoolVar(&flags.dryRun, "dry-run", false, "Show request without sending")
	rootCmd.PersistentFlags().BoolVar(&flags.noCache, "no-cache", false, "Bypass response cache")
	rootCmd.PersistentFlags().BoolVar(&flags.noInput, "no-input", false, "Disable all interactive prompts (for CI/agents)")
	rootCmd.PersistentFlags().BoolVar(&flags.idempotent, "idempotent", false, "Treat already-existing create results as a successful no-op")
	rootCmd.PersistentFlags().BoolVar(&flags.ignoreMissing, "ignore-missing", false, "Treat missing delete targets as a successful no-op")
	rootCmd.PersistentFlags().StringVar(&flags.selectFields, "select", "", "Comma-separated fields to include in output (e.g. --select id,name,status)")
	rootCmd.PersistentFlags().BoolVar(&flags.yes, "yes", false, "Skip confirmation prompts (for agents and scripts)")
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "Disable colored output")
	rootCmd.PersistentFlags().BoolVar(&humanFriendly, "human-friendly", false, "Enable colored output and rich formatting")
	rootCmd.PersistentFlags().BoolVar(&flags.agent, "agent", false, "Set all agent-friendly defaults (--json --compact --no-input --no-color --yes)")
	rootCmd.PersistentFlags().StringVar(&flags.dataSource, "data-source", "auto", "Data source for read commands: auto (live with local fallback), live (API only), local (synced data only)")
	rootCmd.PersistentFlags().StringVar(&flags.profileName, "profile", "", "Apply values from a saved profile (see 'addigy-cli profile list')")
	rootCmd.PersistentFlags().StringVar(&flags.deliverSpec, "deliver", "", "Route output to a sink: stdout (default), file:<path>, webhook:<url>")
	rootCmd.PersistentFlags().Float64Var(&flags.rateLimit, "rate-limit", 0, "Max requests per second (0 to disable)")

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if flags.deliverSpec != "" {
			sink, err := ParseDeliverSink(flags.deliverSpec)
			if err != nil {
				return err
			}
			flags.deliverSink = sink
			if sink.Scheme != "stdout" && sink.Scheme != "" {
				flags.deliverBuf = &bytes.Buffer{}
				cmd.SetOut(io.MultiWriter(os.Stdout, flags.deliverBuf))
			}
		}
		if flags.profileName != "" {
			profile, err := GetProfile(flags.profileName)
			if err != nil {
				return err
			}
			if profile == nil {
				available := ListProfileNames()
				if len(available) == 0 {
					return fmt.Errorf("profile %q not found (no profiles saved yet; run '%s profile save <name> --<flag> <value>')", flags.profileName, cmd.Root().Name())
				}
				return fmt.Errorf("profile %q not found; available: %s", flags.profileName, strings.Join(available, ", "))
			}
			if err := ApplyProfileToFlags(cmd, profile); err != nil {
				return err
			}
		}
		if flags.agent {
			if !cmd.Flags().Changed("json") {
				flags.asJSON = true
			}
			if !cmd.Flags().Changed("compact") {
				flags.compact = true
			}
			if !cmd.Flags().Changed("no-input") {
				flags.noInput = true
			}
			if !cmd.Flags().Changed("yes") {
				flags.yes = true
			}
			if !cmd.Flags().Changed("no-color") {
				noColor = true
			}
		}
		switch flags.dataSource {
		case "auto", "live", "local":
			// valid
		default:
			return fmt.Errorf("invalid --data-source value %q: must be auto, live, or local", flags.dataSource)
		}
		return nil
	}
	rootCmd.AddCommand(newAssetsCmd(flags))
	rootCmd.AddCommand(newDeviceScriptAssignmentsCmd(flags))
	rootCmd.AddCommand(newFactsCmd(flags))
	rootCmd.AddCommand(newFeatureBetasCmd(flags))
	rootCmd.AddCommand(newMaintenanceCmd(flags))
	rootCmd.AddCommand(newManagedAppConfigurationsCmd(flags))
	rootCmd.AddCommand(newMdmCmd(flags))
	rootCmd.AddCommand(newMonitoringCmd(flags))
	rootCmd.AddCommand(newOaCmd(flags))
	rootCmd.AddCommand(newPrebuiltAppsCmd(flags))
	rootCmd.AddCommand(newStaticFieldsCmd(flags))
	rootCmd.AddCommand(newSystemEventsCmd(flags))
	rootCmd.AddCommand(newSystemUpdatesCmd(flags))
	rootCmd.AddCommand(newUsersCmd(flags))
	rootCmd.AddCommand(newDoctorCmd(flags))
	rootCmd.AddCommand(newAuthCmd(flags))
	rootCmd.AddCommand(newAgentContextCmd(rootCmd))
	rootCmd.AddCommand(newProfileCmd(flags))
	rootCmd.AddCommand(newFeedbackCmd(flags))
	rootCmd.AddCommand(newWhichCmd(flags))
	rootCmd.AddCommand(newExportCmd(flags))
	rootCmd.AddCommand(newImportCmd(flags))
	rootCmd.AddCommand(newSearchCmd(flags))
	rootCmd.AddCommand(newSyncCmd(flags))
	rootCmd.AddCommand(newTailCmd(flags))
	rootCmd.AddCommand(newAnalyticsCmd(flags))
	rootCmd.AddCommand(newWorkflowCmd(flags))
	rootCmd.AddCommand(newAPICmd(flags))
	rootCmd.AddCommand(newConfigurationPromotedCmd(flags))
	rootCmd.AddCommand(newDevicesPromotedCmd(flags))
	rootCmd.AddCommand(newFilesPromotedCmd(flags))
	rootCmd.AddCommand(newImpersonationPromotedCmd(flags))
	rootCmd.AddCommand(newOPromotedCmd(flags))
	rootCmd.AddCommand(newSelfServiceConfigurationsPromotedCmd(flags))
	rootCmd.AddCommand(newVersionCliCmd())

	// Hand-authored compound commands (see internal/cli/compound_commands.go).
	registerCompoundCommands(rootCmd, flags)

	return rootCmd
}

func ExitCode(err error) int {
	var codeErr *cliError
	if As(err, &codeErr) {
		return codeErr.code
	}
	return 1
}

func (f *rootFlags) newClient() (*client.Client, error) {
	cfg, err := config.Load(f.configPath)
	if err != nil {
		return nil, configErr(err)
	}
	c := client.New(cfg, f.timeout, f.rateLimit)
	c.DryRun = f.dryRun
	c.NoCache = f.noCache
	c.AssumeYes = f.yes
	c.NoInput = f.noInput
	return c, nil
}

func (f *rootFlags) printTable(w *cobra.Command, headers []string, rows [][]string) error {
	if f.asJSON {
		return fmt.Errorf("use printJSON for JSON output")
	}
	tw := tabwriter.NewWriter(w.OutOrStdout(), 2, 4, 2, ' ', 0)
	header := ""
	for i, h := range headers {
		if i > 0 {
			header += "\t"
		}
		header += h
	}
	fmt.Fprintln(tw, header)
	for _, row := range rows {
		line := ""
		for i, cell := range row {
			if i > 0 {
				line += "\t"
			}
			line += cell
		}
		fmt.Fprintln(tw, line)
	}
	return tw.Flush()
}

func newVersionCliCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("addigy-cli %s\n", version)
		},
	}
}
