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

```bash
# Install if missing
which lix-agents || (brew tap lix-it/lix-agents && brew install lix-agents)

# Start login
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
