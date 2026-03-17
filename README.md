# lix-agents

A CLI tool that lets AI agents request temporary [Lix](https://lix-it.com) API tokens with human approval via email.

Agents can't — and shouldn't — hold long-lived API keys. `lix-agents` gives them a way to request short-lived tokens that require explicit human permission before they're issued.

## How it works

1. **A human logs in once** — `lix-agents auth login` prints a URL. The user visits it on any device, signs in, and the CLI picks up the credentials automatically.
2. **The agent requests a token** — `lix-agents auth token` sends an approval request to the user's email.
3. **The human approves** — The user clicks the link in the email, chooses how long the token should last (1 hour to 30 days), and approves.
4. **The agent gets the token** — The temporary API token is printed to stdout, ready to use.

No passwords or long-lived secrets ever pass through the agent.

## Installation

### Homebrew

```bash
brew tap lix-it/lix-agents
brew install lix-agents
```

### Go install

```bash
go install github.com/lix-it/lix-agents@latest
```

### Binary download

Grab the latest release for your platform from [GitHub Releases](https://github.com/lix-it/lix-agents/releases). Binaries are available for macOS (Intel & Apple Silicon), Linux, and Windows.

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

## For AI agents

This repo ships with a skill (`skills/lix-agents/SKILL.md`) that teaches any compatible agent how to install and use `lix-agents` to authenticate with the Lix API.

### Claude Code (Plugin)

Install as a Claude Code plugin from the marketplace:

```
/plugin marketplace add lix-it/lix-agents
/plugin install lix-agents@lix-agents
```

The `/lix-agents` skill will then be available in Claude Code. The plugin tracks the `stable` branch for reliable releases.

### Amp

Install the skill directly from GitHub:

```
amp skill add lix-it/lix-agents/lix-agents
```

Or from the Amp command palette (`Ctrl+O`), select **skill: add** and enter `lix-it/lix-agents/lix-agents`.

### Manual (any agent that supports SKILL.md)

Copy the skill file into your project or user-level skill directory:

```bash
# Project-level (committed to your repo)
mkdir -p .claude/skills/lix-agents
curl -fsSL https://raw.githubusercontent.com/lix-it/lix-agents/stable/skills/lix-agents/SKILL.md \
  -o .claude/skills/lix-agents/SKILL.md

# Or for Amp's project-level convention
mkdir -p .agents/skills/lix-agents
curl -fsSL https://raw.githubusercontent.com/lix-it/lix-agents/stable/skills/lix-agents/SKILL.md \
  -o .agents/skills/lix-agents/SKILL.md
```

Any agent that reads `SKILL.md` files from `.claude/skills/`, `.agents/skills/`, or `~/.config/agents/skills/` will pick it up automatically.

## Release channels

The plugin uses a `stable` branch for production-ready releases. The `main` branch tracks the latest development state.

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

## Releasing

This project uses [goreleaser](https://goreleaser.com) to build cross-platform binaries and publish to Homebrew. See `.goreleaser.yaml` for configuration.

## License

Apache 2.0
