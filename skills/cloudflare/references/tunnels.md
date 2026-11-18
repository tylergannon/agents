# Publish a service through Cloudflare Tunnel

This workflow uses a remotely managed tunnel. `cf` or the dashboard manages
Cloudflare-side resources. `cloudflared` runs on the origin host using a
connector token scoped to that tunnel.
[Cloudflare's Tunnel API walkthrough](https://developers.cloudflare.com/cloudflare-one/networks/connectors/cloudflare-tunnel/get-started/create-remote-tunnel-api/)

## Prepare the origin and the hostname

Confirm the service URL from the connector host, the intended Cloudflare
account and DNS zone, the hostname and its existing records, and the access
policy. Keep the origin on a private listener where possible. The connector
initiates outbound connections to Cloudflare on port 7844. Use the
[connectivity pre-checks](https://developers.cloudflare.com/cloudflare-one/networks/connectors/cloudflare-tunnel/troubleshoot-tunnels/connectivity-prechecks/)
when the host has outbound filtering.

For a protected hostname, create or verify the
[self-hosted Access application](https://developers.cloudflare.com/cloudflare-one/access-controls/applications/http-apps/self-hosted-public-app/)
and its allow policy before publishing the route. Enable **Protect with Access**
at the connector or validate the Access JWT at the origin as appropriate.
A DNS route to a healthy tunnel can otherwise expose the application to the
internet. Confirm the intended audience before activation.

## Create or reuse the tunnel and route

For API setup, use [scoped credentials](credentials.md) with Tunnel edit access
for the account and DNS edit access for the zone. Access application changes
have their own permissions. Inspect existing tunnels, ingress, and DNS first.
Set `CLOUDFLARE_ACCOUNT_ID` and use `--zone` for DNS commands when selecting the
target explicitly.

1. Create the tunnel with `vp dlx -- cf tunnels create --name "$tunnel_name"
   --config-source cloudflare --dry-run`. Inspect the request and then run the
   authorized create without `--dry-run`. The create response can include
   credentials. Pipe its output through a filter that emits only nonsecret
   fields, such as `jq '{id, name, status}'`, using `vp dlx --silent -- cf …`.
2. Add an ingress entry mapping the hostname to the origin URL before the last
   catch-all rule, typically `{"service":"http_status:404"}`. Keep existing
   rules and settings when updating. The CLI's `tunnels config update` replaces
   the configuration via PUT. Preview the full `{"config":{…}}` request with
   `vp dlx -- cf tunnels config update "$tunnel_id" --body "@$config_json"
   --dry-run`. Keep credentials out of the request body and its preview.
3. Point a proxied DNS CNAME at `<tunnel-id>.cfargotunnel.com`. In the dashboard,
   adding a published application route can create this record. For the CLI,
   preview with `vp dlx -- cf dns records create --zone "$zone_id" --body
   "@$dns_json" --dry-run`. The JSON body contains `type: CNAME`, the hostname
   as `name`, the tunnel target as `content`, and `proxied: true`. Reuse a
   matching record. Preserve any existing route until its replacement is
   authorized.

## Run the connector

Save the connector token directly into a private file on the origin host using
its secret store or a protected local pipe. With authorized CLI access on that
host, the capture helper can consume the JSON string returned by
`tunnels token get`. Set `capture_script` to the absolute path of the
[capture helper](../scripts/capture-credentials.py).

```bash
(
  set -o pipefail
  vp dlx --silent -- cf tunnels token get "$tunnel_id" |
    python3 "$capture_script" --tunnel-token-file "$private_tunnel_token_file"
) && cloudflared tunnel run --token-file "$private_tunnel_token_file"
```

The helper creates a new mode-0600 file. `--token-file` requires
`cloudflared` 2025.4.0 or later. For a persistent connector, configure the
host's service manager to run the same command under the account that owns the
token file. Use a token-file path in the service definition.
[Run parameters](https://developers.cloudflare.com/cloudflare-one/networks/connectors/cloudflare-tunnel/configure-tunnels/run-parameters/#token-file) ·
[Tunnel token permissions](https://developers.cloudflare.com/cloudflare-one/networks/connectors/cloudflare-tunnel/configure-tunnels/remote-tunnel-permissions/)

## Verify the route

Check that Cloudflare shows an active connector and that the hostname reaches
the expected origin from an external client. For a protected route, prove that
an authorized request reaches the app and an unauthenticated request is
challenged or denied. Verify after a service restart if persistent startup was
part of the setup. Report the hostname, origin, tunnel identity, service state,
and observed responses.
