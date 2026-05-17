# lix-agents

A CLI tool that lets AI agents request temporary [Lix](https://lix-it.com) API tokens with human approval via email.

Agents can't — and shouldn't — hold long-lived API keys. `lix-agents` gives them a way to request short-lived tokens that require explicit human permission before they're issued.

## Add to your agent

Install the skill and it handles everything — including installing the CLI for you. No manual setup required.

### Claude Code

Add the marketplace:

```
/plugin marketplace add lix-it/lix-agents
```

Then install the plugin:

```
/plugin install lix-agents@lix-agents
```

The `/lix-agents:lix-agents` skill will be available immediately. Run `/reload-plugins` if you don't see it. It tracks the `stable` branch.

### Claude Cowork

1. Open the Claude Desktop app and switch to the **Cowork** tab.
2. Click the **Customize** menu in the left sidebar.
3. Click **Browse plugins**.
4. Click **Upload** and provide the GitHub URL: `https://github.com/lix-it/lix-agents`
5. Click **Install** on the `lix-agents` plugin.

Once installed, type `/` or click the **+** button to see the `lix-agents` skill.

### Amp

```
amp skill add lix-it/lix-agents/lix-agents
```

Or open the command palette (`Ctrl+O`), select **skill: add**, and enter `lix-it/lix-agents/lix-agents`.

### Claude Code / Cowork (manual)

If the marketplace or upload isn't working, you can load the plugin from a local clone:

1. Clone the repo:
   ```bash
   git clone https://github.com/lix-it/lix-agents.git
   ```
2. In Claude Code, load it with:
   ```bash
   claude --plugin-dir /path/to/lix-agents
   ```
   In Cowork, click **Customize** → **Browse plugins** → **Upload**, then select the `lix-agents` folder you cloned.

### Other agents

Any agent that supports [Agent Skills](https://agentskills.io) (Cursor, Gemini CLI, Windsurf, Roo Code, OpenHands, GitHub Copilot, and [many more](https://agentskills.io)) can use this skill. Copy the `SKILL.md` into the skills directory your agent expects:

```bash
# Claude-compatible agents (.claude/skills/)
mkdir -p .claude/skills/lix-agents
curl -fsSL https://raw.githubusercontent.com/lix-it/lix-agents/stable/skills/lix-agents/SKILL.md \
  -o .claude/skills/lix-agents/SKILL.md

# Amp-compatible agents (.agents/skills/)
mkdir -p .agents/skills/lix-agents
curl -fsSL https://raw.githubusercontent.com/lix-it/lix-agents/stable/skills/lix-agents/SKILL.md \
  -o .agents/skills/lix-agents/SKILL.md
```

### Manual CLI install (optional)

Only needed if your agent can't install it automatically, or if you want to use `lix-agents` outside of an agent.

**Homebrew:**
```bash
brew install lix-it/lix-agents/lix-agents
```

**Go install:**
```bash
go install github.com/lix-it/lix-agents@latest
```

**Binary download:** grab the latest from [GitHub Releases](https://github.com/lix-it/lix-agents/releases) (macOS, Linux, Windows).

## How it works

1. **A human logs in once** — `lix-agents auth login` prints a URL. The user visits it on any device, signs in, and the CLI picks up the credentials automatically.
2. **The agent requests a token** — `lix-agents auth token` sends an approval request to the user's email.
3. **The human approves** — The user clicks the link in the email, chooses how long the token should last (1 hour to 30 days), and approves.
4. **The agent gets the token** — The temporary API token is printed to stdout, ready to use.

No passwords or long-lived secrets ever pass through the agent.

## Quick start

```bash
# Step 1: Log in (one-time, prints a URL to visit)
lix-agents auth login

# Step 2: Request a temporary token
lix-agents auth token
# → An approval email is sent to your inbox.
# → Approve it, and the token is printed here.

# Step 3: Use the token
curl -H "X-Api-Key: <token>" https://api.lix-it.com/v1/person
```

## Look up API endpoints (`lix.sh`)

The Lix API reference is also served as a read-only SSH filesystem at `lix.sh` (powered by [OpenLore](https://github.com/aakarim/go-openlore)). No SSH key required — connections are anonymous. Agents can explore the docs with the same `ls`/`cat`/`grep`/`find` they already know, without burning context-window tokens on the whole reference:

```bash
# Discover the API surface
ssh lix.sh "tree -L 2 /"

# Find the right endpoint
ssh lix.sh "grep -rli 'email' /api"

# Read just one endpoint's reference
ssh lix.sh "cat /api/contact.md"

# Pull a single section (saves context)
ssh lix.sh "sed -n '/## Email from LinkedIn profile/,/^## /p' /api/contact.md"
```

Available files under `/api/`: `lix_account.md`, `account.md`, `disambiguation.md`, `enrichment.md`, `activity.md`, `linkedin.md`, `lookc.md`, `ai.md`, `contact.md`, `errors.md`, plus `agents.md` (this guide) and `/index.md` (intro + auth).

Same content is also available on the web at [https://lix.sh](https://lix.sh) and the rendered reference at [https://lix-it.com/api](https://lix-it.com/api).

The bundled [`/lix-agents` skill](skills/lix-agents/SKILL.md) tells agents to consult `lix.sh` after they have a token, so they pick the right endpoint without you having to paste the docs into context.

## Usage

Run `lix-agents --help` for the full command reference.

```
lix-agents auth login    Log in via URL (saves credentials locally)
lix-agents auth token    Request a temporary API token (requires email approval)
```

## Release channels

| Branch   | Purpose                              |
|----------|--------------------------------------|
| `stable` | Tested releases — plugin default     |
| `main`   | Latest development                   |

## Building from source

```bash
git clone https://github.com/lix-it/lix-agents.git
cd lix-agents
go build -o lix-agents .
```

## License

Apache 2.0
