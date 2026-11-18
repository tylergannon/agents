---
name: cloudflare
description: >
  Configure Cloudflare Tunnels and scoped Cloudflare credentials. Use for
  cloudflared connectors, hostname routes and Zero Trust Access, Cloudflare API
  tokens, or R2 S3 keys, including Safari handoff and local credential capture.
---

# Cloudflare

Identify the intended Cloudflare account, zone, and resource before applying a
change. Reuse existing resources and access when they fit the task.

## Use the Cloudflare CLI

Run `cf` through VitePlus and pnpm's download cache.

```sh
vp dlx -- cf --version
vp dlx -- cf cli search 'create a remotely managed Cloudflare Tunnel'
vp dlx -- cf tunnels create --help
vp dlx -- cf schema tunnels create
```

Use a version pin when reproducing a workflow and `vp dlx --silent -- cf …`
when piping JSON. Search by action and resource type and keep queries free of
account identifiers, hostnames, and secrets. Discover current syntax with
`cli search`, command help, and `schema`. Preview unfamiliar API writes with
`--dry-run` and inspect the account, zone, and request body. Keep the project's
existing tooling for Workers development and deployment.

Use `auth whoami` to check the active Cloudflare identity and set
`CLOUDFLARE_ACCOUNT_ID` explicitly when ambiguous. For unattended API access,
load a scoped `CLOUDFLARE_API_TOKEN` from the consumer's secret store. Use
`cf docs search` for current product documentation. Cloudflare's
[official cf-cli skill](https://github.com/cloudflare/skills/tree/main/skills/cf-cli)
is an optional reference.

## Workflows

- For API tokens and R2 S3 keys, follow [Credentials](references/credentials.md).
  It covers grant verification, the Safari handoff, local capture, and how to
  verify the credential through its intended consumer.
- To publish a service through a remotely managed tunnel, follow
  [Tunnels](references/tunnels.md). It covers the connector token, DNS, Access,
  and an end-to-end request.

Report the selected resource, any local credential path or profile name, and
the operation actually verified. Keep secret values and credential-bearing
responses out of chat, command arguments, logs, and Git.
