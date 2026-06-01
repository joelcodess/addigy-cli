# Addigy CLI — Agent Guide

`addigy-cli` exposes every Addigy v2 API endpoint as a typed command, plus a local
SQLite mirror and compound queries. There's also an MCP server (`addigy-mcp`) that
surfaces the same surface as tools. This guide is written for AI agents driving the
CLI — how to use it well, and how to report problems.

## Start from runtime truth, not assumptions

Don't guess at commands or rely on a memorized list — ask the CLI:

```bash
addigy-cli doctor --json          # auth, base URL, API reachability, mirror freshness
addigy-cli agent-context --json   # machine-readable description of this CLI for agents
addigy-cli which "<capability>"   # find the command that does a thing
addigy-cli <command> --help       # exact flags, args, and the endpoint it calls
```

`doctor` verifies your credentials against the API — if it reports auth invalid,
fix that before anything else (most "it doesn't work" reports are auth or rate-limit,
not bugs).

## Use agent mode for scripting

Add `--agent` to any command. It sets JSON output, compact fields, non-interactive
defaults, no color, and auto-confirms prompts — everything you want for piping to `jq`:

```bash
addigy-cli <command> --agent | jq '...'
```

Individual flags also exist: `--json`, `--compact`, `--select id,name,status`, `--csv`,
`--plain`, `--quiet`.

## Read the exit code — it's typed

Branch on the exit code instead of scraping stderr:

| Code | Meaning | Agent action |
|-----:|---------|--------------|
| `0` | success | continue |
| `2` | usage error (bad flag/args) | fix the invocation |
| `3` | not found | check the id / list first |
| `4` | auth error (401/403) | re-check the API key / permissions |
| `5` | API error (other 4xx/5xx) | inspect the response, may be upstream |
| `7` | rate limited (429) | back off; the key is locked up to 24h if abused |
| `10` | config error | fix `~/.config/addigy-cli/config.toml` or env |

## Mutations are guarded — preview first

Any write (create/update/delete) prompts for confirmation. In non-interactive mode it
**refuses** rather than send unless you pass `--yes`. Always preview first:

```bash
addigy-cli <command> --help            # understand the side effects
addigy-cli <command> --dry-run --agent # see the exact request, nothing sent
addigy-cli <command> --yes --agent     # only once target + args + effects are clear
```

Never run device-impacting commands (erase, lock, restart, MDM command push) unless
the user has explicitly asked for that specific action on that specific device.

## Know the limits

The Addigy v2 API has **no flat list endpoints for devices, policies, or Smart
Software**, so the local mirror (`sync`) does not contain them. Compound commands that
read device/policy data (`compliance`, `devices stale`, `fleet-summary`,
`policy-coverage`, `rollout`) can return empty until that data is present. See the
**Known limitations** section of `README.md`. For install, auth, and examples, read
`README.md` and `SKILL.md`.

## Hit a problem? Open an issue or a PR 🛠️

This CLI is generated from a spec and not every endpoint is battle-tested against a
real tenant. **If something is broken, wrong, or missing, please report it — that's how
it gets fixed.**

**First, rule out the non-bugs:** an exit code of `4` (auth), `7` (rate limit), or `3`
(not found) usually means configuration/usage, not a defect. Run `addigy-cli doctor`.

**If it's a real bug, open an issue:**

```bash
gh issue create --repo joelcodess/addigy-cli \
  --title "<command>: <one-line summary>" \
  --body "$(cat <<'EOF'
**Command:** addigy-cli <command> <args>
**Version:** <output of `addigy-cli version`>
**Expected:** ...
**Actual:** ... (exit code, and `--json` output — REDACT api keys, device serials, user emails, org ids)
**Repro:** the smallest command that shows it (prefer `--dry-run` output so nothing real leaks)
EOF
)"
```

Or open it in a browser: <https://github.com/joelcodess/addigy-cli/issues/new>

**Even better — send a fix as a PR.** The repo is standard Go; the gate is
`go build ./... && go vet ./... && golangci-lint run && go test ./...` (CI runs all of
these on every PR):

```bash
gh repo fork joelcodess/addigy-cli --clone
# make the fix + add/adjust a test
go test ./... && golangci-lint run ./...
gh pr create --repo joelcodess/addigy-cli --fill
```

Good first-PR targets: a wrong/dropped request parameter, a command whose `--help` or
example is inaccurate, a missing flag, or a doc fix. Keep the change focused and include
a test that fails before and passes after.

**Security issues** (anything involving the API key or credential handling) must NOT go
in a public issue — follow `SECURITY.md` to report privately.

## Contributing to this repo

When you modify the CLI, keep it reviewable: a focused commit message explaining what
and why, and a short inline comment at any non-obvious change site so intent survives in
the source. Diffs live in git; there's no separate patch ledger to maintain. Run the
full gate above before pushing.
