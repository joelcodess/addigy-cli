# Contributing to Addigy CLI

Thanks for helping improve `addigy-cli`. Bug reports, fixes, and docs improvements are all welcome.

## Reporting bugs

Open an issue using the bug template. Please include:

- the command you ran and what you expected,
- `addigy-cli version` output and your OS,
- the exit code and `--json` output — **redact** API keys, device serials, user emails, and org IDs,
- the smallest repro (prefer `--dry-run` output so nothing real leaks).

Before filing, run `addigy-cli doctor` — an exit code of `4` (auth), `7` (rate limit), or `3` (not found) is usually configuration, not a defect.

**Security issues** must not go in a public issue — see [`SECURITY.md`](SECURITY.md).

## Development

Requires Go (see `go.mod` for the version). The full gate — what CI runs on every PR — is:

```bash
go build ./...
go vet ./...
gofmt -l internal cmd          # must print nothing
golangci-lint run ./...        # 0 issues
go test -race ./...
```

Install the linter with `go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest`.

## Pull requests

1. Fork and branch from `main`.
2. Keep the change focused; add or adjust a test that fails before your change and passes after.
3. Most command files under `internal/cli/` are code-generated from the API spec — prefer fixing the root cause and note in the PR when a change is to generated code.
4. Run the full gate above; make sure it's green.
5. Open the PR with a clear description of what changed and why.

## Reporting a wrong/dropped request parameter, bad `--help`, or missing flag

These are great first contributions. Reproduce with `--dry-run` (shows the exact request without sending it), then fix the relevant command in `internal/cli/` and add a test.
