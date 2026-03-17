# lix-agents

A CLI tool that lets AI agents request temporary [Lix](https://lix-it.com) API tokens with human approval via email.

Agents can't — and shouldn't — hold long-lived API keys. `lix-agents` gives them a way to request short-lived tokens that require explicit human permission before they're issued.

## Add to your agent

Install the skill and it handles everything — including installing the CLI for you. No manual setup required.

### Claude Code

First, add the marketplace:

```
/plugin marketplace add lix-it/lix-agents
```

Then install the plugin:

```
/plugin install lix-agents@lix-agents
```

The `/lix-agents` skill will be available immediately. It tracks the `stable` branch.

### Claude Cowork

1. Open Claude Desktop and switch to the **Cowork** tab.
2. Click **Customize** in the left sidebar.
3. Click **Browse plugins**.
4. If `lix-agents` isn't listed, click **Upload** and provide the GitHub URL: `https://github.com/lix-it/lix-agents`
5. Click **Install** on the `lix-agents` plugin.

Alternatively, type these commands directly in a Cowork task:

```
/plugin marketplace add lix-it/lix-agents
```

```
/plugin install lix-agents@lix-agents
```

### Amp

```
amp skill add lix-it/lix-agents/lix-agents
```

Or open the command palette (`Ctrl+O`), select **skill: add**, and enter `lix-it/lix-agents/lix-agents`.

### Claude Code / Cowork (manual)

If the marketplace commands aren't working, you can install the plugin manually:

1. Clone the repo:
   ```bash
   git clone https://github.com/lix-it/lix-agents.git
   ```
2. In Claude Code, run:
   ```
   /plugin install /path/to/lix-agents
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
