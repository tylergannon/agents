---
name: proof-uploader
description: Install, configure, and use the proof CLI to upload artifacts, rewrite Markdown links, or vacuum old objects in an S3-compatible bucket.
---

# Proof Uploader

Install the CLI:

```sh
go install github.com/tylergannon/agents/cmd/proof@latest
```

Configure one of these examples, keeping credentials out of tracked files:

- YAML: [`references/config.yaml`](references/config.yaml) at `~/.proof-uploader/config.yaml`
- dotenv: [`references/.env.example`](references/.env.example) as an untracked `.env`

Run `proof upload-file FILE...`, `proof prepare-proof INPUT.md OUTPUT.md`, or `proof vacuum 2w`.

Public URLs are the default. Set `download_mode: signed` or `PROOF_DOWNLOAD_MODE=signed` for one-week presigned URLs.

source: agents
