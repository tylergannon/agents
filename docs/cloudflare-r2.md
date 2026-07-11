# Cloudflare R2 configuration

- In **R2 Object Storage**, create or select the proof bucket.
- Under **Public Development URL**, enable public access and copy the `r2.dev` URL.
- Public URLs allow reads by exact object key but do not expose bucket listing. Upload, list, vacuum, and delete still require API credentials.
- Under **Manage R2 API Tokens**, create an account token with **Object Read & Write** restricted to the proof bucket.
- Copy the generated Access Key ID and Secret Access Key immediately; Cloudflare shows the secret only once.
- Create `~/.proof-uploader/config.yaml` with directory mode `0700` and file mode `0600`:

```yaml
account_id: "CLOUDFLARE_ACCOUNT_ID"
access_key_id: "R2_ACCESS_KEY_ID"
secret_access_key: "R2_SECRET_ACCESS_KEY"
bucket: "proof-artifacts"
endpoint: "https://CLOUDFLARE_ACCOUNT_ID.r2.cloudflarestorage.com"
prefix: "proof"
download_mode: public
public_base_url: "https://PUBLIC_BUCKET_ID.r2.dev"
```

- `download_mode: public` is the default and returns non-expiring, unguessable public object URLs.
- Set `download_mode: signed` to return private R2 presigned URLs valid for one week; `public_base_url` is then optional.
- `.env` and process environment variables override YAML; see `.env.example` for their names.
- To rotate credentials, create a replacement bucket-scoped token, update the YAML or environment, verify an upload, then delete the old token.
- Run `go run ./cmd/proof vacuum 2w` periodically; deleting an object revokes both public and signed URLs.
