---
name: lix-agents
description: Obtain temporary Lix API tokens via CLI with human email approval.
---

# Lix Agents

Use `lix-agents` to get temporary API tokens for the Lix API. Tokens require human approval via email, so agents never hold unsupervised credentials.

Run `lix-agents --help` for full command reference.

## Installation

```bash
brew tap lix-it/lix-agents
brew install lix-agents
```

Or download a binary from [GitHub Releases](https://github.com/lix-it/lix-agents/releases), or `go install github.com/lix-it/lix-agents@latest`.

## When to use this

- You need authenticated access to the Lix API
- You need to enrich LinkedIn profiles, companies, or other data via Lix
- You need API credentials for any Lix service
- You don't already have a valid Lix API token in your environment

## How it works

1. **`lix-agents auth login`** — Prints a URL for the human to visit and sign in on any device. Only needed once; credentials are saved locally at `~/.lix/credentials.json`.
2. **`lix-agents auth token`** — Requests a temporary token. An approval email is sent to the user. Tell the user to check their email and approve. The token is printed to stdout once approved.
3. Use the token in API requests to `https://api.lix-it.com`.

## Lix API documentation
See the [Lix API docs](https://lix-it.com/docs) for details on available endpoints, request formats, and response formats.
