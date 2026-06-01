# Addigy CLI

**Every Addigy v2 endpoint, plus a local fleet mirror with compound queries the web UI can't answer.**

> _Unofficial / community project. Not affiliated with or endorsed by Addigy, Inc. "Addigy" is a trademark of its respective owner._

Addigy's REST API is rate-limited to 1,000 requests per 10 seconds. This CLI mirrors the API's list resources — custom facts, MDM profiles and configuration definitions, monitoring alerts, OmniAgent benchmarks and compliance rules, static fields, and system updates — into local SQLite so repeated lookups run as one local query instead of paging the live API. Every spec endpoint is a typed Cobra command with `--json`, `--select`, `--csv`, `--dry-run`, and typed exit codes; agents pipe the output to jq without any HTML-scraping detour. Live writes are guarded by a confirmation prompt (bypass with `--yes`/`--agent`).

> **Compound device/policy commands** (`compliance`, `devices stale`, `fleet-summary`, `policy-coverage`, `rollout`) read from the mirror. The Addigy v2 API exposes devices via a query endpoint (`POST /devices`) rather than a bulk list, and has no Smart Software list endpoint, so these commands can return empty until device/policy data is present in the mirror — see [Known limitations](#known-limitations).

## Install

### Go install

Install the CLI (and, optionally, the MCP server) directly with Go:

```bash
go install github.com/joelcodess/addigy-cli/cmd/addigy-cli@latest
go install github.com/joelcodess/addigy-cli/cmd/addigy-mcp@latest
```

### Build from source

```bash
git clone https://github.com/joelcodess/addigy-cli && cd addigy-cli && make build
```

## Authentication

Authenticate by exporting an Addigy API token (`export ADDIGY_DOCUMENTATION_API_KEY=...`) generated at https://app.addigy.com/integrations. The CLI sends it as the `x-api-key` header on every request. `addigy-cli doctor` will confirm reachability, auth, and the resolved base URL.

The CLI defaults to the production host `https://api.addigy.com` (API base path `/api/v2`), so the API key is the only thing you need to set. To target a different host (self-hosted or test), set `ADDIGY_BASE_URL` (and optionally `ADDIGY_BASE_PATH`).

## Quick Start

```bash
# Confirm the token is loaded and the API is reachable before anything else.
addigy-cli doctor


# Mirror devices, policies, facts, smart software, and applications into local SQLite. ~30s for a 1k-device fleet.
addigy-cli sync --full


# Single-shot triage view — device count, stale %, alert count, MDM queue depth, software pending, policy coverage.
addigy-cli fleet-summary --json --agent


# Find every device that hasn't checked in for a week.
addigy-cli devices stale --days 7 --json --agent


# Cross-join devices x policy_rules x facts to find noncompliant devices.
addigy-cli compliance --policy ENG-Standard --json --agent

```

## Unique Features

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

## Usage

Run `addigy-cli --help` for the full command reference and flag list.

## Commands

### assets

Manage assets

- **`addigy-cli assets create`** - Get a list of Default Alerts.
- **`addigy-cli assets create-default`** - Get a list of Default Maintenance Jobs.
- **`addigy-cli assets create-default-2`** - Get a list of Default MDM Configurations.
- **`addigy-cli assets create-default-3`** - Get a list of Default Self Service Configurations.

### configuration

Manage configuration

- **`addigy-cli configuration list`** - Get API key's permissions

### device-script-assignments

Manage device script assignments

- **`addigy-cli device-script-assignments create`** - Creates a device script assignment in the organization.
- **`addigy-cli device-script-assignments delete`** - Deletes a device script assignment from the organization.
- **`addigy-cli device-script-assignments list`** - Get Device Script Assignments available for the organization.

### devices

Manage devices

- **`addigy-cli devices create`** - Allow to query for a set of devices based on a value that pertains to one of their device facts. <br><b>Permission Required: </b>View Devices.

### facts

Manage facts

- **`addigy-cli facts create`** - Create a custom fact.
- **`addigy-cli facts create-custom`** - Assign Custom Facts to policies.
- **`addigy-cli facts create-custom-2`** - Get a list of Custom Facts filtered by id or name for an organization.
- **`addigy-cli facts delete`** - Delete a custom fact.
- **`addigy-cli facts delete-custom`** - Unassign a custom fact from a policy.
- **`addigy-cli facts list`** - Get all custom facts for the organization.
- **`addigy-cli facts update`** - Update a custom fact.

### feature-betas

Manage feature betas

- **`addigy-cli feature-betas create`** - Enables a Beta Feature in the organization. <br><b>Permission Required: </b>Toggle Feature Betas.
- **`addigy-cli feature-betas delete`** - Disables the Beta Features from the organization. <br><b>Permission Required: </b>Toggle Feature Betas.
- **`addigy-cli feature-betas list`** - Get all Beta Features available for the organization. <br><b>Permission Required: </b>Toggle Feature Betas.

### files

Manage files

- **`addigy-cli files create`** - Get a list of file usages for a list of File IDs.

### impersonation

Manage impersonation

- **`addigy-cli impersonation create`** - Creates a session for impersonating into a child organization.

### maintenance

Manage maintenance

- **`addigy-cli maintenance create`** - Create a maintenance item. <br><b>Permission Required: </b>Create Catalog Maintenance.
- **`addigy-cli maintenance create-policy`** - Assign polices to a maintenance item. <br><b>Permission Required: </b>Edit Policy Maintenance.
- **`addigy-cli maintenance create-query`** - Get a list of maintenance items for an organization.
- **`addigy-cli maintenance create-staged`** - Creates a staged maintenance item from an existing one.
- **`addigy-cli maintenance create-staged-2`** - Confirm a staged maintenance. This will create a maintenance with the same details as the staged maintenance and will send an event. <br><b>Permission Required: </b>Edit Catalog Maintenance.
- **`addigy-cli maintenance create-staged-3`** - Get a list of maintenance items for an organization.
- **`addigy-cli maintenance delete`** - Delete a maintenance item.<br><b>Permission Required: </b>Delete Catalog Maintenance.
- **`addigy-cli maintenance delete-policy`** - Unassign a maintenance item from policy. <br><b>Permission Required: </b>Edit Policy Maintenance.
- **`addigy-cli maintenance delete-staged`** - Deletes a staged maintenance item.
- **`addigy-cli maintenance update`** - Update a maintenance item. <br><b>Permission Required: </b>Edit Catalog Maintenance.
- **`addigy-cli maintenance update-staged`** - Updates a staged maintenance item.

### managed-app-configurations

Manage managed app configurations

- **`addigy-cli managed-app-configurations create`** - Requests to create managed app configuration for Apps & Books applications.
- **`addigy-cli managed-app-configurations delete`** - Requests to delete managed app configuration for Apps & Books applications.
- **`addigy-cli managed-app-configurations list`** - Gets managed app configuration for Apps & Books applications.

### mdm

Manage mdm

- **`addigy-cli mdm create`** - Paginated request that returns list of installed certificates by mdm devices. <br><br><b>Permission Required: </b>View Devices
- **`addigy-cli mdm create-commands`** - Send MDM command to restart a device
- **`addigy-cli mdm create-configurations`** - Create MDM configuration profile
- **`addigy-cli mdm create-configurations-2`** - Assign policies to manifest-based MDM configuration profile
- **`addigy-cli mdm create-configurations-3`** - Confirm changes to manifest-based MDM configuration profile
- **`addigy-cli mdm create-devices`** - Deploys profile to list of devices and/or managed users. It is an atomic request meaning that if one error is encountered no profile will be deployed to any of the devices and/or managed users <br><br/><b>Permission Required: </b><ol><li>View Devices</li><li>Execute commands</li></ol>
- **`addigy-cli mdm create-profiles`** - Get MDM profiles assigned to policies
- **`addigy-cli mdm delete`** - This command allows the server to delete a user that has an active account on the device. Please provide the device agent ID or the device uuid
- **`addigy-cli mdm delete-configurations`** - Unassign an MDM profile from policies
- **`addigy-cli mdm delete-configurations-2`** - Delete manifest-based MDM configuration profile
- **`addigy-cli mdm get`** - Get MDM device details including enrollment profile, APN certificate and last response.
- **`addigy-cli mdm get-configurations`** - Get MDM configuration profile definition
- **`addigy-cli mdm get-configurations-2`** - Get manifest-based MDM configuration profile
- **`addigy-cli mdm get-devices`** - Test MDM response.
- **`addigy-cli mdm list`** - Get MDM profiles
- **`addigy-cli mdm list-commands`** - Returns a list of known users that were given to Addigy via the Request User List command.Please provide the device agent id or the device uuid
- **`addigy-cli mdm list-configurations`** - Get MDM configuration profile definitions
- **`addigy-cli mdm list-configurations-2`** - Get manifest-based MDM configuration profiles
- **`addigy-cli mdm list-configurations-3`** - Get policy profiles by Addigy payload type
- **`addigy-cli mdm update`** - Update an MDM configuration profile

### monitoring

Manage monitoring

- **`addigy-cli monitoring create`** - Create a monitoring item.
- **`addigy-cli monitoring create-policy`** - Assign monitoring item to policy. <br><b>Permission Required: </b>Edit Policy Monitoring.
- **`addigy-cli monitoring create-query`** - Get a list of monitoring items for an organization.
- **`addigy-cli monitoring delete`** - Delete a monitoring item.<br><b>Permission Required: </b>Delete Custom Monitoring.
- **`addigy-cli monitoring delete-policy`** - Unassign a monitoring item from policy. <br><b>Permission Required: </b>Edit Policy Monitoring.
- **`addigy-cli monitoring list`** - Get list of received alerts for the organization.
- **`addigy-cli monitoring update`** - Update a monitoring item.

### o

Manage o

- **`addigy-cli o create`** - Request additional Azure Conditional Access Connectors for their organization.

### oa

Manage oa

- **`addigy-cli oa create`** - Get a list of benchmark assets for an organization. <br><b>Permission Required: </b>View Benchmarks.
- **`addigy-cli oa create-ade`** - Get a list of ade tokens assigned to policies.
- **`addigy-cli oa create-appsandbooks`** - Get a list of apps and books tokens assigned to policies.
- **`addigy-cli oa create-compliancerules`** - Get a list of compliance rules for an organization. <br><b>Permission Required: </b>View Benchmarks.
- **`addigy-cli oa create-devices`** - Get devices compliance status. <br><b>Permission Required: </b>View Devices.
- **`addigy-cli oa create-files`** - Get a list of files for an organization. <br><b>Permission Required:</b> View Files.
- **`addigy-cli oa create-identity`** - Get a list of identity configurations assigned to policies.
- **`addigy-cli oa create-installedapps`** - Query installed apps from a device providing some agent IDs. <br><b>Permission Required:</b> View Devices.
- **`addigy-cli oa create-integrations`** - Get Azure Conditional Access all accounts metadata.
- **`addigy-cli oa create-monitoring`** - Query for a list of scheduled alerts with pagination.
- **`addigy-cli oa create-policies`** - Query an organization for all policies or filter to get specific policy info
- **`addigy-cli oa create-policies-2`** - Gets a list of available assets for the provided location ID (token ID).
- **`addigy-cli oa create-prebuiltapps`** - Query Prebuilt Apps Configurations
- **`addigy-cli oa create-reports`** - Get report statuses.
- **`addigy-cli oa create-variables`** - Get a list of variables for an organization. <br><b>Permission Required: </b>View Variables.
- **`addigy-cli oa create-webhooks`** - Get a list of webhooks.
- **`addigy-cli oa create-webhooks-2`** - Get a count of webhooks schedule.
- **`addigy-cli oa create-webhooks-3`** - Get a list of webhooks status.
- **`addigy-cli oa list`** - Get a policy home screen layout
- **`addigy-cli oa list-benchmarks`** - Get pre-built benchmarks. <br><b>Permission Required: </b>View Benchmarks.
- **`addigy-cli oa list-compliancerules`** - Get pre-built compliance rules. <br><b>Permission Required: </b>View Benchmarks.
- **`addigy-cli oa list-compliancerules-2`** - Get a compliance rule usage. <br><b>Permission Required: </b>View Benchmarks.
- **`addigy-cli oa list-integrations`** - Get Azure Conditional Access unique enabled tenants metadata
- **`addigy-cli oa list-reports`** - Get a report.
- **`addigy-cli oa list-reports-2`** - Get a list of available reports.
- **`addigy-cli oa list-selfservice`** - Get the self service configurations by OS for a policy

### prebuilt-apps

Manage prebuilt apps

- **`addigy-cli prebuilt-apps create`** - Create a Prebuilt App Version
- **`addigy-cli prebuilt-apps create-prebuiltapps`** - Create a prebuilt app
- **`addigy-cli prebuilt-apps create-prebuiltapps-2`** - Query the prebuilt app library
- **`addigy-cli prebuilt-apps create-prebuiltapps-3`** - Query Prebuilt App Versions
- **`addigy-cli prebuilt-apps delete`** - Delete a prebuilt app
- **`addigy-cli prebuilt-apps delete-prebuiltapps`** - Delete a Prebuilt App Version
- **`addigy-cli prebuilt-apps get`** - Get a prebuilt app
- **`addigy-cli prebuilt-apps get-prebuiltapps`** - Get a Prebuilt App Version
- **`addigy-cli prebuilt-apps update`** - Update a prebuilt app
- **`addigy-cli prebuilt-apps update-prebuiltapps`** - Update a Prebuilt App Version

### self-service-configurations

Manage self service configurations

- **`addigy-cli self-service-configurations create`** - Creates a new self service configuration in the organization. <br><b>Permission Required: </b>Create Instruction.

### static-fields

Manage static fields

- **`addigy-cli static-fields create`** - Creates a new static field in the organization. <br><b>Permission Required: </b>View Devices.
- **`addigy-cli static-fields create-staticfields`** - Assign static field values to device(s) in the organization. <br><b>Permission Required: </b>View Devices.
- **`addigy-cli static-fields delete`** - Removes the static field from the organization. <br><b>Permission Required: </b>View Devices.
- **`addigy-cli static-fields list`** - Gets a list of all static fields available for the organization. <br><b>Permission Required: </b>View Devices.
- **`addigy-cli static-fields list-staticfields`** - Gets a list of all static fields assigned to devices for the organization. <br><b>Permission Required: </b>View Devices.
- **`addigy-cli static-fields update`** - Updates the name of an existing static field in the organization. <br><b>Permission Required: </b>View Devices.

### system-events

Manage system events

- **`addigy-cli system-events create`** - Allow to query for a set of system events. <br><b>Permission Required: </b>View System Events.
- **`addigy-cli system-events create-systemevents`** - Allow to search system events. <br><b>Permission Required: </b>View System Events.

### system-updates

Manage system updates

- **`addigy-cli system-updates create`** - Requests available system updates for a device via MDM command.<br><br><b>Permission Required: </b>View Device List, Execute Predefined Commands.
- **`addigy-cli system-updates create-systemupdates`** - Requests a system updates scan for a device via MDM command.<br><br><b>Permission Required: </b>View Device List, Execute Predefined Commands.
- **`addigy-cli system-updates create-systemupdates-2`** - Requests the schedule of system updates via MDM command.<br><br><b>Permission Required: </b>View Device List, Execute Predefined Commands.
- **`addigy-cli system-updates create-systemupdates-3`** - Requests to create or update system updates settings for a policy.<br><br><b>Permission Required: </b>Create System Updates Settings.<br><br><b>Requirements: </b>The MDM update command only works with macOS 12+, iOS 9+, iPadOS 13+, or tvOS 12+. Devices must be in <b>supervised</b> mode. Unsupervised devices will not receive the update command.<br><br><b>Minor Updates and Patches: </b>Use version values to control which major, minor or patch updates are sent. Addigy will strictly follow your rules. Version values follow the major.minor.patch standard.<br><br>For example:<br>12.0.99 will allow patches, but not the minor update to 12.1<br>12.9.9 will not allow 12.9.91<br><br><b>System Updates Settings: </b>This is what each of the more specialized fields represents within the system updates settings for each os and which are the allowed values for the request.<br><br>install_action - represents the install action when sending the schedule os command to the device: 1.Default (all OS) 2.InstallForceRestart (macOS Only) 3.InstallLater (macOS Only)<br>max_user_deferrals - represents how many times the user can defer the updates, this is an optional parameter and it only works when 'Install Action' is 'InstallLater' and for minor os updates<br>resend_update_command_hour - The time in hours needed to re-send an os update command if the last command status is older than this value. Currently, the default value is 24 hours and the valid values ranges from 1 hour up to 24<br>days_after_release, hours_after_release and minutes_after_release (DDM updates only): The number of days, hours and minutes to force an update installation via DDM, after the update is released.<br><br><b>Schedule (excludes updates via DDM): </b>System Updates commands are scheduled to be sent daily at 2AM UTC, but you can schedule them to run on the device's time and which days of the week. The schedule is optional, if you would like to continue to use the default daily schedule, just set the schedule.enabled field to false. However, if you would like to opt in to use the schedule, just set the schedule.enabled field to true and fill all fields since they are required as part of the schedule request. Please note that your organization must have a monthly paid plan to use this feature.<br><br>This is what each field represents and what are the allowed values for the schedule request:<br>enabled - represents if the schedule is enabled or disabled (true or false)<br>week_days - represents days of the week ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"]<br>starting_time - represents the schedule starting time (hours: 0-23h, min: 0 or 30)<br>cut_off_time - represents last time within the maintenance window to send updates commands to the devices (min: 30, 45, 60)<br>maintenance_window - represents how long do the schedule runs for in 2x hour intervals (hours: 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24)<br><br><br>For more information about System Updates please visit [System Updates via MDM](https://support.addigy.com/hc/en-us/articles/10073419654931)
- **`addigy-cli system-updates create-systemupdates-4`** - Requests system updates statuses for a device via MDM command.<br><br><b>Permission Required: </b>View Device List, Execute Predefined Commands.
- **`addigy-cli system-updates create-systemupdates-5`** - Gets available updates reported for multiple devices
- **`addigy-cli system-updates create-systemupdates-6`** - Gets installed system updates reported for multiple devices
- **`addigy-cli system-updates create-systemupdates-7`** - Requests to schedule system updates (on-demand) for devices via MDM command.<br><br><b>Permission Required: </b>[View System Updates Settings, Create System Updates Settings]<br><br><b>[MDM System Updates Only] </b>System Updates commands are scheduled to be sent daily at 2AM UTC, but you can send them now to the device(s) on this list. Please note that your organization must have a monthly paid plan to use this feature.
- **`addigy-cli system-updates create-systemupdates-8`** - Requests to schedule system updates (on-demand) for policy devices via MDM command.<br><br><b>Permission Required: </b>[View System Updates Settings, Create System Updates Settings]<br><br><b>[MDM System Updates Only] </b>System Updates commands are scheduled to be sent daily at 2AM UTC, but you can send them now to the device(s) in this policy. Please note that your organization must have a monthly paid plan to use this feature.
- **`addigy-cli system-updates create-systemupdates-9`** - Requests to send installed system updates reported for policy devices to user email.<br><br><b>Permission Required: </b>View System Updates Settings.
- **`addigy-cli system-updates list`** - Gets available system updates reported for a device.<br><br><b>Permission Required: </b>View Device List.
- **`addigy-cli system-updates list-systemupdates`** - Get latest system updates available by os type
- **`addigy-cli system-updates list-systemupdates-2`** - Gets system updates settings for a policy.<br><br><b>Permission Required: </b>View System Updates Settings.
- **`addigy-cli system-updates list-systemupdates-3`** - Gets current system updates statuses reported for a device.<br><br><b>Permission Required: </b>View Device List, Execute Predefined Commands.
- **`addigy-cli system-updates list-systemupdates-4`** - Gets available system updates reported for a device, with their current installation statuses.<br><br><b>Permission Required: </b>View Device List, Execute Predefined Commands.
- **`addigy-cli system-updates list-systemupdates-5`** - Gets device system updates statuses via ddm status report.<br><br><b>Permission Required: </b>View Device List.
- **`addigy-cli system-updates list-systemupdates-6`** - Gets installed system updates reported for a device.<br><br><b>Permission Required: </b>View System Updates Settings.

### users

Manage users

- **`addigy-cli users delete`** - Deletes a user from the organization. <br><b>Permission Required: </b>Remove User.
- **`addigy-cli users update`** - Update a user. <br><b>Permission Required: </b>Edit User.


## Output Formats

```bash
# Human-readable table (default in terminal, JSON when piped)
addigy-cli configuration

# JSON for scripting and agents
addigy-cli configuration --json

# Filter to specific fields
addigy-cli configuration --json --select id,name,status

# Dry run — show the request without sending
addigy-cli configuration --dry-run

# Agent mode — JSON + compact + no prompts in one flag
addigy-cli configuration --agent
```

## Agent Usage

This CLI is designed for AI agent consumption:

- **Non-interactive** - never prompts, every input is a flag
- **Pipeable** - `--json` output to stdout, errors to stderr
- **Filterable** - `--select id,name` returns only fields you need
- **Previewable** - `--dry-run` shows the request without sending
- **Explicit retries** - add `--idempotent` to create retries and `--ignore-missing` to delete retries when a no-op success is acceptable
- **Confirmable** - `--yes` for explicit confirmation of destructive actions
- **Piped input** - write commands can accept structured input when their help lists `--stdin`
- **Offline-friendly** - sync/search commands can use the local SQLite store when available
- **Agent-safe by default** - no colors or formatting unless `--human-friendly` is set

Exit codes: `0` success, `2` usage error, `3` not found, `4` auth error, `5` API error, `7` rate limited, `10` config error.

## Use with Claude Code

Install the focused `addigy` skill, then invoke `/addigy <query>` in Claude Code. The skill is the most efficient path — Claude Code drives the CLI directly without an MCP server in the middle.

<details>
<summary>Use as an MCP server in Claude Code (advanced)</summary>

If you'd rather register this CLI as an MCP server in Claude Code, install the MCP binary first:

```bash
go install github.com/joelcodess/addigy-cli/cmd/addigy-mcp@latest
```

Then register it:

```bash
claude mcp add addigy addigy-mcp -e ADDIGY_DOCUMENTATION_API_KEY=<your-key>
```

</details>

## Use with Claude Desktop

Install both binaries and add the MCP server to your Claude Desktop config. Install **both** `addigy-mcp` and `addigy-cli` — several MCP tools (`compliance`, `devices_stale`, `fleet_summary`, `sync`, …) shell out to the `addigy-cli` binary, so it must be on `PATH` alongside `addigy-mcp`.

```bash
go install github.com/joelcodess/addigy-cli/cmd/addigy-cli@latest
go install github.com/joelcodess/addigy-cli/cmd/addigy-mcp@latest
```

Add to your Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "addigy": {
      "command": "addigy-mcp",
      "env": {
        "ADDIGY_DOCUMENTATION_API_KEY": "<your-key>",
        "ADDIGY_CLI_PATH": "/absolute/path/to/addigy-cli"
      }
    }
  }
}
```

`ADDIGY_CLI_PATH` is optional if `addigy-cli` is already on `PATH`; set it to the absolute path otherwise.

## Health Check

```bash
addigy-cli doctor
```

Verifies configuration, credentials, and connectivity to the API.

## Configuration

Config file: `~/.config/addigy-cli/config.toml`

Static request headers can be configured under `headers`; per-command header overrides take precedence.

Environment variables:

| Name | Kind | Required | Description |
| --- | --- | --- | --- |
| `ADDIGY_DOCUMENTATION_API_KEY` | per_call | Yes | Set to your API credential. |
| `ADDIGY_BASE_URL` | session | No | Overrides the default API host `https://api.addigy.com`. |
| `ADDIGY_BASE_PATH` | session | No | Overrides the default API base path `/api/v2`. |
| `ADDIGY_CONFIG` | session | No | Path to the config file (default `~/.config/addigy-cli/config.toml`). |

## Troubleshooting
**Authentication errors (exit code 4)**
- Run `addigy-cli doctor` to check credentials
- Verify the environment variable is set: `echo $ADDIGY_DOCUMENTATION_API_KEY`
**Not found errors (exit code 3)**
- Check the resource ID is correct
- Run the `list` command to see available items

### API-specific

- **HTTP 429 on every call** — Rate-limit lockout. Addigy locks API keys for 24 hours when the 1,000-req/10s budget is exceeded. Any live command will surface the `429` (exit code 7); wait for the lockout to clear before retrying.
- **`auth status` returns 401** — Token missing or revoked. Regenerate at https://app.addigy.com/integrations and `export ADDIGY_DOCUMENTATION_API_KEY=...`.
- **`compliance` returns empty against a known noncompliant fleet** — Mirror is stale. Run `addigy-cli sync --since 24h` to refresh policies, rules, and device facts; `compliance` reads only the local mirror.
- **`rollout` shows zero devices for a Smart Software item that is clearly deployed** — Smart Software has no list endpoint in the Addigy v2 API, so its install state is not mirrored. See Known limitations below.

---

## Known limitations

- **The mirror does not include devices, policies, or Smart Software.** The Addigy v2 API exposes devices through a query endpoint (`POST /devices`) rather than a flat list, policies are organization-scoped, and Smart Software has no list endpoint at all. `sync` therefore hydrates only the resources with bulk list endpoints (facts, MDM profiles/definitions, monitoring alerts, OmniAgent benchmarks and compliance rules, static fields, system updates). Compound commands that read device/policy data (`compliance`, `devices stale`, `fleet-summary`, `policy-coverage`, `rollout`) will return empty against a mirror that has no device rows.
- **A few resources sync with zero stored rows.** `device-script-assignments`, `mdm-configurations-definitions`, and `oa` reports emit a `sync_anomaly` event because their payloads have no extractable primary key; they are surfaced as warnings rather than failing the run.
- **Output-format flags are limited on scalar-array payloads.** Endpoints that return an array of plain strings (e.g. `configuration`) render as JSON under `--csv`/`--plain`, and `--quiet` produces no output. Use `--json` for these. `--select`/`--compact` apply only to object-array responses.
- **Two upstream endpoints are currently broken server-side** (the CLI builds the requests correctly): `o billing get-o` (`GET /o/{org}/billing/cards`) returns HTTP 500, and `oa create-prebuiltapps` (`POST /oa/prebuilt-apps/configurations/query`) returns HTTP 404 (route absent). These are Addigy API issues, not CLI defects.

## License

Licensed under the Apache License, Version 2.0 — see [`LICENSE`](LICENSE) and [`NOTICE`](NOTICE).

## Sources & Inspiration

This CLI was built by studying these projects and resources:

- [**pliancy/addigy-node**](https://github.com/pliancy/addigy-node) — JavaScript
- [**Addigy-Community/addytool**](https://github.com/Addigy-Community/addytool) — Python
- [**Addigy-Community/upload-software**](https://github.com/Addigy-Community/upload-software) — Python
- [**dsrosen6/addigy-cli**](https://github.com/dsrosen6/addigy-cli) — Go
