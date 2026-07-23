# agents

Shared tooling for Codex projects.

This repo is the development home for generalized agent skills and
infrastructure that should not live inside a single product repository,
including reusable prompt and workflow skills, scripts, reference indexes, and
operational notes.

## Current Contents

- `skills/` contains published portable agent skills.
- `borrowed-skills.json` pins the provenance of copied upstream skills.
- `.agents/skills/` contains repo-local maintainer skills that are
  not shipped to target repositories.
- `ephemeral/` stores tracked operating material that daily docs jobs mine into
  durable docs and skills.
- `submodules/` mounts upstream reference repositories for agent-tooling research.
- `submodules/index/` tracks the local inventory, lessons, and borrow map for
  those references.

Current skill focus:

- prompt writing and editing,
- productivity interviews, handoffs, teaching, and skill-writing guidance,
- Attractor pipeline design/edit/proof,
- adversarial review and consensus loops,
- proof-of-work and proof-artifact publishing,
- document-maintenance skills for target repositories,
- session worklogs and behavior proof,
- reusable Svelte factoring, E2E coverage triage, Supabase query, Resend
  preview verification, and customer-feedback triage skills,
- skill evaluation and agent-tooling research as repo-local maintainer workflows.

## Install

Install all published skills globally for Codex and Claude Code:

```sh
npx skills add tylergannon/agents --global --skill '*' --agent codex claude-code --yes
```

The skills CLI uses symlinks by default. Do not pass `--copy`.

Plugin-aware hosts can consume the same published skill set through
`.codex-plugin/plugin.json` or `.claude-plugin/plugin.json`. Repo-local
maintainer skills live under `.agents/skills/` and are not distributed to
target repositories.

`package.json` is retained for metadata validation and package dry-runs; the
documented distribution path is the public skills CLI.

To inspect the published skill set locally:

```sh
find skills -name SKILL.md -print
```

## Proof artifacts

The Go proof CLI uploads temporary artifacts to an S3-compatible bucket,
rewrites relative Markdown links, and removes old objects:

```sh
go run ./cmd/proof upload-file FILE...
go run ./cmd/proof prepare-proof DOCUMENT OUTPUT_OR_-
go run ./cmd/proof vacuum 2w
```

Configuration is read from `~/.proof-uploader/config.yaml`; `.env` and process
environment variables override it. Follow
[`docs/cloudflare-r2.md`](docs/cloudflare-r2.md) for setup and credential
rotation.

Public object URLs are returned by default. Set `download_mode: signed` to
return native R2 presigned URLs valid for one week instead.
