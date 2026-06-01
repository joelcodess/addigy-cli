---
name: addigy
description: "Every Addigy v2 endpoint, plus a local fleet mirror with compound queries the web UI can't answer. Trigger phrases: `addigy fleet health`, `find stale macs`, `addigy compliance audit`, `smart software rollout status`, `diff two addigy devices`, `use addigy`, `run addigy`."
author: "joelcodess"
license: "Apache-2.0"
argument-hint: "<command> [args] | install cli|mcp"
allowed-tools: "Read Bash"
metadata:
  addigy:
    requires:
      bins:
        - addigy-cli
---

# Addigy CLI

## Prerequisites: Install the CLI

This skill drives the `addigy-cli` binary. **You must verify the CLI is installed before invoking any command from this skill.** If it is missing, install it first:

1. Install with Go:
   ```bash
   go install github.com/joelcodess/addigy-cli/cmd/addigy-cli@latest
   ```
   Or build from source:
   ```bash
   git clone https://github.com/joelcodess/addigy-cli && cd addigy-cli && make build
   ```
2. Verify: `addigy-cli --version`
3. Ensure `$GOPATH/bin` (or `$HOME/go/bin`) is on `$PATH`.

If `--version` reports "command not found" after install, the install step did not put the binary on `$PATH`. Do not proceed with skill commands until verification succeeds.

Addigy's REST API is rate-limited to 1,000 requests per 10 seconds. This CLI mirrors the API's list resources — custom facts, MDM profiles and configuration definitions, monitoring alerts, OmniAgent benchmarks and compliance rules, static fields, and system updates — into local SQLite so repeated lookups run as one local query instead of paging the live API. Every spec endpoint is a typed Cobra command with `--json`, `--select`, `--csv`, `--dry-run`, and typed exit codes; agents pipe the output to jq without any HTML-scraping detour. Live writes are guarded by a confirmation prompt (bypass with `--yes`/`--agent`).

Note: the Addigy v2 API has no flat list endpoints for devices, policies, or Smart Software, so the compound commands that read them (`compliance`, `devices stale`, `fleet-summary`, `policy-coverage`, `rollout`) can return empty until that data is present in the mirror.

## When to Use This CLI

Use this CLI when an agent or admin needs fleet-shaped Addigy answers: custom-fact mining, monitoring-alert sweeps, MDM-profile inventory, and OmniAgent compliance-rule lookups. Pick the live REST surface for one-off CRUD; pick the compound commands (`compliance`, `rollout`, `fleet-summary`, `drift`) for any question that would otherwise require fan-out across the rate-limited API — these read from the local mirror, so they depend on the relevant data having been synced (see the mirror note above). The local mirror makes repeated lookups far cheaper in API calls.

## Unique Capabilities

These capabilities aren't available in any other tool for this API.

### Local state that compounds
- **`devices stale`** — List devices whose last check-in is older than N days, with optional policy and OS filters.

  _Use this for Monday fleet health passes instead of paging /devices and joining client-side._

  ```bash
  addigy-cli devices stale --days 7 --policy ENG-Standard --json --agent
  ```
- **`compliance`** — Surface devices whose assigned policy rules are unmet, joined against current device facts.

  _Pick this when an agent needs the noncompliant set of devices for a policy without walking the UI per device._

  ```bash
  addigy-cli compliance --policy ENG-Standard --rule require-filevault --json --agent
  ```
- **`rollout`** — Per-device install state for one Smart Software item across the assigned fleet, with success/pending/failed counts.

  _Use this for Friday rollout reports, or to answer 'is this package stuck on any device?' without manual UI walking._

  ```bash
  addigy-cli rollout slack-business-v4 --policy MARKETING --json --agent
  ```
- **`facts search`** — FTS5 across mirrored device facts; --group-by value returns a histogram of values per fact.

  _Reach for this whenever an agent needs to characterize a fact across the fleet (FileVault status, OS version, custom-fact distributions)._

  ```bash
  addigy-cli facts search "FileVault" --group-by value --json --agent
  ```
- **`devices diff`** — Set differences across facts, applications, policies, and Smart-Software install state between two devices.

  _Use this for ticket triage when one device works and another does not._

  ```bash
  addigy-cli devices diff DEV-A-123 DEV-B-456 --json --agent
  ```
- **`drift`** — Diffs the mirror's current rows against the prior snapshot for any entity (devices, facts, policies, software).

  _Pick this when an agent asks 'what changed in this fleet recently?' instead of paging multiple endpoints._

  ```bash
  addigy-cli drift --since 24h --entity devices --json --agent
  ```
- **`policy-coverage`** — Per-policy device counts joined to last-checkin so the user sees both coverage and freshness in one view.

  _Use this to find policies with high coverage but stale devices, or low coverage but healthy devices._

  ```bash
  addigy-cli policy-coverage --policy ENG-Standard --json --agent
  ```

### Agent-native triage
- **`fleet-summary`** — Single command emitting device count, stale fraction, alert count, MDM queue depth, Smart-Software pending count, and policy coverage percent.

  _Use this as the first call in any Addigy triage workflow; everything else narrows from here._

  ```bash
  addigy-cli fleet-summary --json --agent
  ```

## Command Reference

**assets** — Manage assets

- `addigy-cli assets create` — Get a list of Default Alerts.
- `addigy-cli assets create-default` — Get a list of Default Maintenance Jobs.
- `addigy-cli assets create-default-2` — Get a list of Default MDM Configurations.
- `addigy-cli assets create-default-3` — Get a list of Default Self Service Configurations.

**configuration** — Manage configuration

- `addigy-cli configuration` — Get API key's permissions

**device-script-assignments** — Manage device script assignments

- `addigy-cli device-script-assignments create` — Creates a device script assignment in the organization.
- `addigy-cli device-script-assignments delete` — Deletes a device script assignment from the organization.
- `addigy-cli device-script-assignments list` — Get Device Script Assignments available for the organization.

**devices** — Manage devices

- `addigy-cli devices` — Allow to query for a set of devices based on a value that pertains to one of their device facts. <br><b>Permission...

**facts** — Manage facts

- `addigy-cli facts create` — Create a custom fact.
- `addigy-cli facts create-custom` — Assign Custom Facts to policies.
- `addigy-cli facts create-custom-2` — Get a list of Custom Facts filtered by id or name for an organization.
- `addigy-cli facts delete` — Delete a custom fact.
- `addigy-cli facts delete-custom` — Unassign a custom fact from a policy.
- `addigy-cli facts list` — Get all custom facts for the organization.
- `addigy-cli facts update` — Update a custom fact.

**feature-betas** — Manage feature betas

- `addigy-cli feature-betas create` — Enables a Beta Feature in the organization. <br><b>Permission Required: </b>Toggle Feature Betas.
- `addigy-cli feature-betas delete` — Disables the Beta Features from the organization. <br><b>Permission Required: </b>Toggle Feature Betas.
- `addigy-cli feature-betas list` — Get all Beta Features available for the organization. <br><b>Permission Required: </b>Toggle Feature Betas.

**files** — Manage files

- `addigy-cli files` — Get a list of file usages for a list of File IDs.

**impersonation** — Manage impersonation

- `addigy-cli impersonation` — Creates a session for impersonating into a child organization.

**maintenance** — Manage maintenance

- `addigy-cli maintenance create` — Create a maintenance item. <br><b>Permission Required: </b>Create Catalog Maintenance.
- `addigy-cli maintenance create-policy` — Assign polices to a maintenance item. <br><b>Permission Required: </b>Edit Policy Maintenance.
- `addigy-cli maintenance create-query` — Get a list of maintenance items for an organization.
- `addigy-cli maintenance create-staged` — Creates a staged maintenance item from an existing one.
- `addigy-cli maintenance create-staged-2` — Confirm a staged maintenance. This will create a maintenance with the same details as the staged maintenance and...
- `addigy-cli maintenance create-staged-3` — Get a list of maintenance items for an organization.
- `addigy-cli maintenance delete` — Delete a maintenance item.<br><b>Permission Required: </b>Delete Catalog Maintenance.
- `addigy-cli maintenance delete-policy` — Unassign a maintenance item from policy. <br><b>Permission Required: </b>Edit Policy Maintenance.
- `addigy-cli maintenance delete-staged` — Deletes a staged maintenance item.
- `addigy-cli maintenance update` — Update a maintenance item. <br><b>Permission Required: </b>Edit Catalog Maintenance.
- `addigy-cli maintenance update-staged` — Updates a staged maintenance item.

**managed-app-configurations** — Manage managed app configurations

- `addigy-cli managed-app-configurations create` — Requests to create managed app configuration for Apps & Books applications.
- `addigy-cli managed-app-configurations delete` — Requests to delete managed app configuration for Apps & Books applications.
- `addigy-cli managed-app-configurations list` — Gets managed app configuration for Apps & Books applications.

**mdm** — Manage mdm

- `addigy-cli mdm create` — Paginated request that returns list of installed certificates by mdm devices. <br><br><b>Permission Required:...
- `addigy-cli mdm create-commands` — Send MDM command to restart a device
- `addigy-cli mdm create-configurations` — Create MDM configuration profile
- `addigy-cli mdm create-configurations-2` — Assign policies to manifest-based MDM configuration profile
- `addigy-cli mdm create-configurations-3` — Confirm changes to manifest-based MDM configuration profile
- `addigy-cli mdm create-devices` — Deploys profile to list of devices and/or managed users. It is an atomic request meaning that if one error is...
- `addigy-cli mdm create-profiles` — Get MDM profiles assigned to policies
- `addigy-cli mdm delete` — This command allows the server to delete a user that has an active account on the device. Please provide the device...
- `addigy-cli mdm delete-configurations` — Unassign an MDM profile from policies
- `addigy-cli mdm delete-configurations-2` — Delete manifest-based MDM configuration profile
- `addigy-cli mdm get` — Get MDM device details including enrollment profile, APN certificate and last response.
- `addigy-cli mdm get-configurations` — Get MDM configuration profile definition
- `addigy-cli mdm get-configurations-2` — Get manifest-based MDM configuration profile
- `addigy-cli mdm get-devices` — Test MDM response.
- `addigy-cli mdm list` — Get MDM profiles
- `addigy-cli mdm list-commands` — Returns a list of known users that were given to Addigy via the Request User List command.Please provide the device...
- `addigy-cli mdm list-configurations` — Get MDM configuration profile definitions
- `addigy-cli mdm list-configurations-2` — Get manifest-based MDM configuration profiles
- `addigy-cli mdm list-configurations-3` — Get policy profiles by Addigy payload type
- `addigy-cli mdm update` — Update an MDM configuration profile

**monitoring** — Manage monitoring

- `addigy-cli monitoring create` — Create a monitoring item.
- `addigy-cli monitoring create-policy` — Assign monitoring item to policy. <br><b>Permission Required: </b>Edit Policy Monitoring.
- `addigy-cli monitoring create-query` — Get a list of monitoring items for an organization.
- `addigy-cli monitoring delete` — Delete a monitoring item.<br><b>Permission Required: </b>Delete Custom Monitoring.
- `addigy-cli monitoring delete-policy` — Unassign a monitoring item from policy. <br><b>Permission Required: </b>Edit Policy Monitoring.
- `addigy-cli monitoring list` — Get list of received alerts for the organization.
- `addigy-cli monitoring update` — Update a monitoring item.

**o** — Manage o

- `addigy-cli o` — Request additional Azure Conditional Access Connectors for their organization.

**oa** — Manage oa

- `addigy-cli oa create` — Get a list of benchmark assets for an organization. <br><b>Permission Required: </b>View Benchmarks.
- `addigy-cli oa create-ade` — Get a list of ade tokens assigned to policies.
- `addigy-cli oa create-appsandbooks` — Get a list of apps and books tokens assigned to policies.
- `addigy-cli oa create-compliancerules` — Get a list of compliance rules for an organization. <br><b>Permission Required: </b>View Benchmarks.
- `addigy-cli oa create-devices` — Get devices compliance status. <br><b>Permission Required: </b>View Devices.
- `addigy-cli oa create-files` — Get a list of files for an organization. <br><b>Permission Required:</b> View Files.
- `addigy-cli oa create-identity` — Get a list of identity configurations assigned to policies.
- `addigy-cli oa create-installedapps` — Query installed apps from a device providing some agent IDs. <br><b>Permission Required:</b> View Devices.
- `addigy-cli oa create-integrations` — Get Azure Conditional Access all accounts metadata.
- `addigy-cli oa create-monitoring` — Query for a list of scheduled alerts with pagination.
- `addigy-cli oa create-policies` — Query an organization for all policies or filter to get specific policy info
- `addigy-cli oa create-policies-2` — Gets a list of available assets for the provided location ID (token ID).
- `addigy-cli oa create-prebuiltapps` — Query Prebuilt Apps Configurations
- `addigy-cli oa create-reports` — Get report statuses.
- `addigy-cli oa create-variables` — Get a list of variables for an organization. <br><b>Permission Required: </b>View Variables.
- `addigy-cli oa create-webhooks` — Get a list of webhooks.
- `addigy-cli oa create-webhooks-2` — Get a count of webhooks schedule.
- `addigy-cli oa create-webhooks-3` — Get a list of webhooks status.
- `addigy-cli oa list` — Get a policy home screen layout
- `addigy-cli oa list-benchmarks` — Get pre-built benchmarks. <br><b>Permission Required: </b>View Benchmarks.
- `addigy-cli oa list-compliancerules` — Get pre-built compliance rules. <br><b>Permission Required: </b>View Benchmarks.
- `addigy-cli oa list-compliancerules-2` — Get a compliance rule usage. <br><b>Permission Required: </b>View Benchmarks.
- `addigy-cli oa list-integrations` — Get Azure Conditional Access unique enabled tenants metadata
- `addigy-cli oa list-reports` — Get a report.
- `addigy-cli oa list-reports-2` — Get a list of available reports.
- `addigy-cli oa list-selfservice` — Get the self service configurations by OS for a policy

**prebuilt-apps** — Manage prebuilt apps

- `addigy-cli prebuilt-apps create` — Create a Prebuilt App Version
- `addigy-cli prebuilt-apps create-prebuiltapps` — Create a prebuilt app
- `addigy-cli prebuilt-apps create-prebuiltapps-2` — Query the prebuilt app library
- `addigy-cli prebuilt-apps create-prebuiltapps-3` — Query Prebuilt App Versions
- `addigy-cli prebuilt-apps delete` — Delete a prebuilt app
- `addigy-cli prebuilt-apps delete-prebuiltapps` — Delete a Prebuilt App Version
- `addigy-cli prebuilt-apps get` — Get a prebuilt app
- `addigy-cli prebuilt-apps get-prebuiltapps` — Get a Prebuilt App Version
- `addigy-cli prebuilt-apps update` — Update a prebuilt app
- `addigy-cli prebuilt-apps update-prebuiltapps` — Update a Prebuilt App Version

**self-service-configurations** — Manage self service configurations

- `addigy-cli self-service-configurations` — Creates a new self service configuration in the organization. <br><b>Permission Required: </b>Create Instruction.

**static-fields** — Manage static fields

- `addigy-cli static-fields create` — Creates a new static field in the organization. <br><b>Permission Required: </b>View Devices.
- `addigy-cli static-fields create-staticfields` — Assign static field values to device(s) in the organization. <br><b>Permission Required: </b>View Devices.
- `addigy-cli static-fields delete` — Removes the static field from the organization. <br><b>Permission Required: </b>View Devices.
- `addigy-cli static-fields list` — Gets a list of all static fields available for the organization. <br><b>Permission Required: </b>View Devices.
- `addigy-cli static-fields list-staticfields` — Gets a list of all static fields assigned to devices for the organization. <br><b>Permission Required: </b>View Devices.
- `addigy-cli static-fields update` — Updates the name of an existing static field in the organization. <br><b>Permission Required: </b>View Devices.

**system-events** — Manage system events

- `addigy-cli system-events create` — Allow to query for a set of system events. <br><b>Permission Required: </b>View System Events.
- `addigy-cli system-events create-systemevents` — Allow to search system events. <br><b>Permission Required: </b>View System Events.

**system-updates** — Manage system updates

- `addigy-cli system-updates create` — Requests available system updates for a device via MDM command.<br><br><b>Permission Required: </b>View Device List,...
- `addigy-cli system-updates create-systemupdates` — Requests a system updates scan for a device via MDM command.<br><br><b>Permission Required: </b>View Device List,...
- `addigy-cli system-updates create-systemupdates-2` — Requests the schedule of system updates via MDM command.<br><br><b>Permission Required: </b>View Device List,...
- `addigy-cli system-updates create-systemupdates-3` — Requests to create or update system updates settings for a policy.<br><br><b>Permission Required: </b>Create System...
- `addigy-cli system-updates create-systemupdates-4` — Requests system updates statuses for a device via MDM command.<br><br><b>Permission Required: </b>View Device List,...
- `addigy-cli system-updates create-systemupdates-5` — Gets available updates reported for multiple devices
- `addigy-cli system-updates create-systemupdates-6` — Gets installed system updates reported for multiple devices
- `addigy-cli system-updates create-systemupdates-7` — Requests to schedule system updates (on-demand) for devices via MDM command.<br><br><b>Permission Required:...
- `addigy-cli system-updates create-systemupdates-8` — Requests to schedule system updates (on-demand) for policy devices via MDM command.<br><br><b>Permission Required:...
- `addigy-cli system-updates create-systemupdates-9` — Requests to send installed system updates reported for policy devices to user email.<br><br><b>Permission Required:...
- `addigy-cli system-updates list` — Gets available system updates reported for a device.<br><br><b>Permission Required: </b>View Device List.
- `addigy-cli system-updates list-systemupdates` — Get latest system updates available by os type
- `addigy-cli system-updates list-systemupdates-2` — Gets system updates settings for a policy.<br><br><b>Permission Required: </b>View System Updates Settings.
- `addigy-cli system-updates list-systemupdates-3` — Gets current system updates statuses reported for a device.<br><br><b>Permission Required: </b>View Device List,...
- `addigy-cli system-updates list-systemupdates-4` — Gets available system updates reported for a device, with their current installation statuses.<br><br><b>Permission...
- `addigy-cli system-updates list-systemupdates-5` — Gets device system updates statuses via ddm status report.<br><br><b>Permission Required: </b>View Device List.
- `addigy-cli system-updates list-systemupdates-6` — Gets installed system updates reported for a device.<br><br><b>Permission Required: </b>View System Updates Settings.

**users** — Manage users

- `addigy-cli users delete` — Deletes a user from the organization. <br><b>Permission Required: </b>Remove User.
- `addigy-cli users update` — Update a user. <br><b>Permission Required: </b>Edit User.


### Finding the right command

When you know what you want to do but not which command does it, ask the CLI directly:

```bash
addigy-cli which "<capability in your own words>"
```

`which` resolves a natural-language capability query to the best matching command from this CLI's curated feature index. Exit code `0` means at least one match; exit code `2` means no confident match — fall back to `--help` or use a narrower query.

## Recipes


### Monday fleet health pass

```bash
addigy-cli sync --since 24h && addigy-cli fleet-summary --json --agent
```

Refresh the mirror's deltas, then emit the one-shot triage view ready to pipe into jq.

### Find devices stuck on macOS 13

```bash
addigy-cli devices stale --days 14 --os '13.*' --json --agent --select "id,name,last_checkin,os_version"
```

Cross-filter the stale view by OS regex; --select trims the payload to the four fields an agent needs.

### Ticket triage: why is device X different from device Y?

```bash
addigy-cli devices diff DEV-A-123 DEV-B-456 --json --agent
```

Set diff across facts, applications, policies, and Smart Software state for two device IDs.

### Smart Software rollout report for one customer

```bash
addigy-cli rollout slack-business-v4 --policy MARKETING --csv
```

Per-device install state in CSV; pipe to a file or paste into a customer status email.

### Fleet drift since yesterday

```bash
addigy-cli drift --since 24h --entity devices --json --agent --select "id,name,changed_fields"
```

What changed in the last 24h; --select narrows to the audit-trail fields agents actually consume from a deeply nested diff payload.

## Auth Setup

Authenticate by exporting an Addigy API token (`export ADDIGY_DOCUMENTATION_API_KEY=...`) generated at https://app.addigy.com/integrations. The CLI sends it as the `x-api-key` header on every request. `addigy-cli doctor` will confirm reachability, auth, and the resolved base URL. The CLI defaults to `https://api.addigy.com`; set `ADDIGY_BASE_URL` to target a different host.

Run `addigy-cli doctor` to verify setup.

## Agent Mode

Add `--agent` to any command. Expands to: `--json --compact --no-input --no-color --yes`.

- **Pipeable** — JSON on stdout, errors on stderr
- **Filterable** — `--select` keeps a subset of fields. Dotted paths descend into nested structures; arrays traverse element-wise. Critical for keeping context small on verbose APIs:

  ```bash
  addigy-cli configuration --agent --select id,name,status
  ```
- **Previewable** — `--dry-run` shows the request without sending
- **Offline-friendly** — sync/search commands can use the local SQLite store when available
- **Non-interactive** — never prompts, every input is a flag
- **Explicit retries** — use `--idempotent` only when an already-existing create should count as success, and `--ignore-missing` only when a missing delete target should count as success

### Response envelope

Commands that read from the local store or the API wrap output in a provenance envelope:

```json
{
  "meta": {"source": "live" | "local", "synced_at": "...", "reason": "..."},
  "results": <data>
}
```

Parse `.results` for data and `.meta.source` to know whether it's live or local. A human-readable `N results (live)` summary is printed to stderr only when stdout is a terminal — piped/agent consumers get pure JSON on stdout.

## Agent Feedback

When you (or the agent) notice something off about this CLI, record it:

```
addigy-cli feedback "the --since flag is inclusive but docs say exclusive"
addigy-cli feedback --stdin < notes.txt
addigy-cli feedback list --json --limit 10
```

Entries are stored locally at `~/.addigy-cli/feedback.jsonl`. They are never POSTed unless `ADDIGY_FEEDBACK_ENDPOINT` is set AND either `--send` is passed or `ADDIGY_FEEDBACK_AUTO_SEND=true`. Default behavior is local-only.

Write what *surprised* you, not a bug report. Short, specific, one line: that is the part that compounds.

## Output Delivery

Every command accepts `--deliver <sink>`. The output goes to the named sink in addition to (or instead of) stdout, so agents can route command results without hand-piping. Three sinks are supported:

| Sink | Effect |
|------|--------|
| `stdout` | Default; write to stdout only |
| `file:<path>` | Atomically write output to `<path>` (tmp + rename) |
| `webhook:<url>` | POST the output body to the URL (`application/json` or `application/x-ndjson` when `--compact`) |

Unknown schemes are refused with a structured error naming the supported set. Webhook failures return non-zero and log the URL + HTTP status on stderr.

## Named Profiles

A profile is a saved set of flag values, reused across invocations. Use it when a scheduled agent calls the same command every run with the same configuration.

```
addigy-cli profile save briefing --json
addigy-cli --profile briefing configuration
addigy-cli profile list --json
addigy-cli profile show briefing
addigy-cli profile delete briefing --yes
```

Explicit flags always win over profile values; profile values win over defaults. `agent-context` lists all available profiles under `available_profiles` so introspecting agents discover them at runtime.

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 2 | Usage error (wrong arguments) |
| 3 | Resource not found |
| 4 | Authentication required |
| 5 | API error (upstream issue) |
| 7 | Rate limited (wait and retry) |
| 10 | Config error |

## Argument Parsing

Parse `$ARGUMENTS`:

1. **Empty, `help`, or `--help`** → show `addigy-cli --help` output
2. **Starts with `install`** → ends with `mcp` → MCP installation; otherwise → see Prerequisites above
3. **Anything else** → Direct Use (execute as CLI command with `--agent`)

## MCP Server Installation

Install the MCP binary with Go, then register it:

```bash
go install github.com/joelcodess/addigy-cli/cmd/addigy-cli@latest
go install github.com/joelcodess/addigy-cli/cmd/addigy-mcp@latest
claude mcp add addigy -e ADDIGY_DOCUMENTATION_API_KEY=<your-key> -- addigy-mcp
```

Verify: `claude mcp list`

## Direct Use

1. Check if installed: `which addigy-cli`
   If not found, offer to install (see Prerequisites at the top of this skill).
2. Match the user query to the best command from the Unique Capabilities and Command Reference above.
3. Execute with the `--agent` flag:
   ```bash
   addigy-cli <command> [subcommand] [args] --agent
   ```
4. If ambiguous, drill into subcommand help: `addigy-cli <command> --help`.
