# Changelog

All notable changes to this project are documented here. The format is based on
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and this project aims to
follow [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-05-30

Initial public release: every Addigy v2 endpoint exposed as a typed Cobra
command, a local SQLite mirror with `sync`/`search`, compound query commands,
and an MCP server (`addigy-mcp`).

### Added
- Live-mutation confirmation guardrail: any write (non-`/query` POST/PUT/PATCH/
  DELETE) prompts before sending; `--yes`/`--agent` bypass, and non-interactive
  invocations without `--yes` refuse rather than send.
- Default API base URL (`https://api.addigy.com`) so the CLI connects with only
  an API key set; overridable via `ADDIGY_BASE_URL` / `ADDIGY_BASE_PATH`.
- Authenticated `doctor` health check that verifies credentials against the API
  and exits non-zero (4) on 401/403.
- Versioned `User-Agent` (`addigy-cli/<version> (<os>/<arch>)`) on every request.
- Tests for the mutation guardrail, config base-URL resolution, and the MCP
  execute path's delete query-parameter handling.
- CI (build/vet/test/lint/govulncheck) and tag-driven GoReleaser release
  workflows; `SECURITY.md`, `CONTRIBUTING.md`, `CODE_OF_CONDUCT.md`.

### Fixed
- Delete commands now transmit the target id/key as a query parameter (both the
  typed CLI commands and the MCP `addigy_execute` path); previously the
  identifier was dropped and the API rejected the request.

### Changed
- Typed exit codes are propagated by the CLI entrypoint (auth=4, not-found=3,
  usage=2, rate-limit=7, api=5, config=10) instead of a flat `1`.
- Stripped raw HTML (`<br>`, `<b>`) from command help text.
- Removed two sync resources that require path params the bulk syncer cannot
  supply (they always returned HTTP 400); added ID-field overrides so MDM
  profiles persist to the mirror.
- Response cache and config files are written owner-only (`0600`); error output
  redacts credential-shaped strings.
- Documentation corrected: the mirror does not include devices, policies, or
  Smart Software (the Addigy v2 API has no flat list endpoints for them); added
  a Known Limitations section.

### Security
- Bumped `golang.org/x/sys` to clear advisory GO-2026-5024.
