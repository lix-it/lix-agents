---
name: lix-agents
description: Log in to the Lix API and obtain a temporary API token with human email approval.
---

# Lix Agents Login

Starts the `lix-agents` authentication flow. This will:

1. Install the CLI if it isn't already available
2. Open a login URL for you to authenticate
3. Request a temporary API token (sent to your email for approval)

## Run

Install the CLI using one of the options below, then run `lix-agents auth login`.

### Option A: Homebrew (macOS / Linux)

```bash
brew tap lix-it/lix-agents && brew install lix-agents
```

### Option B: Download binary

```bash
mkdir -p ~/.local/bin
VERSION=$(curl -sI https://github.com/lix-it/lix-agents/releases/latest | grep -i ^location | sed 's|.*/v||;s/\r//')
curl -fsSL "https://github.com/lix-it/lix-agents/releases/download/v${VERSION}/lix-agents_${VERSION}_$(uname -s | tr '[:upper:]' '[:lower:]')_$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/').tar.gz" \
  | tar xz -C ~/.local/bin lix-agents
export PATH="$HOME/.local/bin:$PATH"
```

### Option C: Go install

```bash
go install github.com/lix-it/lix-agents@latest
```

## Login

```bash
lix-agents auth login
```

Follow the URL printed by the CLI to sign in. Once logged in, request a token:

```bash
lix-agents auth token
```

An approval email will be sent — approve it and the temporary token is printed to stdout. Use it as:

```
Authorization: Bearer <token>
```

All requests go to `https://api.lix-it.com`. See the [Lix API docs](https://lix-it.com/docs) for endpoints and usage.
