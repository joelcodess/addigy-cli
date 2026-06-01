# Security Policy

## Reporting a vulnerability

Please report security vulnerabilities privately rather than opening a public
issue. Use GitHub's **[Report a vulnerability](https://github.com/joelcodess/addigy-cli/security/advisories/new)**
flow (Security tab → Report a vulnerability) so the report stays confidential
until a fix is available.

Include where possible:

- the affected version (`addigy-cli version`) and platform,
- a description of the issue and its impact,
- steps to reproduce or a proof of concept.

You can expect an acknowledgement within a few business days. Please allow time
for a fix before any public disclosure.

## Handling credentials

This CLI authenticates with an Addigy API key (`ADDIGY_DOCUMENTATION_API_KEY`),
sent as the `x-api-key` header.

- The key is stored at `~/.config/addigy-cli/config.toml` with `0600`
  permissions and is never written to logs.
- Never paste your key into issues, pull requests, or screenshots. The CLI
  redacts key-shaped strings in error output, but treat the key as a secret.
- Rotate the key immediately at <https://app.addigy.com/integrations> if you
  believe it has been exposed.
