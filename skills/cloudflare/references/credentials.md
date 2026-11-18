# Cloudflare credentials

| Credential | Consumer |
| --- | --- |
| Cloudflare OAuth grant or scoped API token | `cf` and the Cloudflare REST API |
| R2 Access Key ID and Secret Access Key | S3-compatible SDKs and tools |
| Tunnel connector token | `cloudflared` for one remotely managed tunnel |

Prefer a scoped account-owned API token for supported service endpoints. Use a
scoped user token when the operation needs user ownership or the endpoint
requires one. Confirm the account, zone or buckets, permissions, and any
expiration or IP restrictions for the consumer.
[API token creation](https://developers.cloudflare.com/fundamentals/api/get-started/create-token/) ·
[Account-owned tokens](https://developers.cloudflare.com/fundamentals/api/get-started/account-owned-tokens/)

## Check CLI access

Reuse an appropriate API token or OAuth grant. If interactive login is needed,
request the necessary scopes and check the returned grant with `auth whoami`.
For a Safari handoff, use `vp dlx -- cf auth login --no-browser` and open the
displayed verification URL while the device code is fresh.

The CLI can expose an API operation that the active grant cannot authorize.
Mission Control's R2 setup created buckets using OAuth, but the grant omitted
`account_api_tokens:create` and the account-token create operation returned
403 / 9109. Check the current grant and the actual operation if authorized.
A denied permission-group listing alone does not establish whether token
creation works. Use the dashboard when the required create permission is
unavailable through the CLI. [Cloudflare CLI issue 607](https://github.com/cloudflare/cf/issues/607)

## Prepare the dashboard in Safari

1. Open the intended account's **Manage Account → API Tokens** page, or
   [My Profile → API Tokens](https://dash.cloudflare.com/profile/api-tokens/)
   for a user token.
   For R2, open `https://dash.cloudflare.com/<account-id>/r2/api-tokens`.
2. Prepare the token name, permissions, and resource restrictions. For R2
   object access choose the required object permission and **Apply to specific
   buckets only**. Use **Object Read & Write** only for writers.
3. Review the summary, present the prepared form to the human for the final
   **Create Token** click, and ask them to leave the one-time result page open.
   Handle local capture and configuration after the click.

## Save the one-time result

Run [capture-credentials.py](../scripts/capture-credentials.py) with Python 3 on
the workstation. Set `capture_script` to its absolute path in the installed
skill. `--source safari` reads the active tab via Apple Events when Safari
already allows that access. The helper checks the dashboard URL and prints
only status and destination metadata.

```sh
# Account-owned bearer token. Omit --account for a user token.
python3 "$capture_script" --source safari --account "$account_id" \
  --api-token-file "$private_token_file"

# R2 keys for an AWS SDK profile.
python3 "$capture_script" --source safari --account "$account_id" \
  --profile "$profile"
```

If Safari Apple Events are unavailable, use a browser tool that can transfer
JSON directly into local stdin without displaying or logging the payload. Use
`--source stdin` with the same destination arguments. Pass `url` and either
`text` from the result page or its exact credential fields. The API field is
`api_token`. The R2 fields are `access_key_id` and `secret_access_key`.
If a tool cannot transfer secrets privately, have the human save the result
locally in the target secret store. Keep the result out of chat.

The bearer token file is mode 0600 and must have a new path outside the work
repository. R2 capture saves a mode-0600 AWS profile, preserving other profiles.
Use `--credentials-file` for a custom store and `--replace` for an intended
profile rotation. `--dry-run` validates a capture without persisting it.

The R2 dashboard shows the S3 key pair explicitly. If issuing the R2 token via
the API, Cloudflare defines the S3 Access Key ID as the returned `id` and the
S3 Secret Access Key as the SHA-256 hex digest of the token `value`.
[R2 authentication](https://developers.cloudflare.com/r2/api/tokens/)

Verify an API token with a scoped operation on the intended resource. For R2,
use the consumer's AWS profile and the correct jurisdiction's S3 endpoint
with region `auto`. Verify the required object operations and use a disposable
object if verifying writes. During rotation, verify the replacement through
the consumer before revoking the old credential.
