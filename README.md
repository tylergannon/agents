# agents

Shared tooling for Codex projects.

This repo is the development home for generalized agent skills and
infrastructure that should not live inside a single product repository,
including reusable prompt and workflow skills, scripts, reference indexes, and
operational notes.

## Current Contents

- `skills/` contains published portable agent skills.
- `.agents/skills/` contains repo-local maintainer skills that are
  not shipped to target repositories.
- `ephemeral/` stores tracked operating material that daily docs jobs mine into
  durable docs and skills.
- `submodules/` mounts upstream reference repositories for agent-tooling research.
- `submodules/index/` tracks the local inventory, lessons, and borrow map for
  those references.

Current skill focus:

- prompt writing and editing,
- Attractor pipeline design/edit/proof,
- Claude Code / Codex consensus review loops,
- proof-of-work and dependency-upgrade PR lifecycle ownership,
- document-maintenance skills for target repositories,
- session worklogs and behavior proof,
- reusable Svelte factoring, E2E coverage triage, Supabase query, Resend
  preview verification, customer-feedback triage, and tech-debt scouting skills,
- skill evaluation and agent-tooling research as repo-local maintainer workflows.

## Install

Install all published skills globally for Codex, Claude Code, and OpenCode:

```sh
vp dlx skills add tylergannon/agents \
  -g \
  -a codex -a claude-code -a opencode \
  -s '*' \
  -y
```

Plugin-aware hosts can consume the same published skill set through
`.codex-plugin/plugin.json` or `.claude-plugin/plugin.json`. Repo-local
maintainer skills live under `.agents/skills/` and are not distributed to
target repositories.

`package.json` is retained for metadata validation and package dry-runs; the
documented distribution path is the public skills CLI.

To inspect the published skill set locally:

```sh
find skills -maxdepth 2 -name SKILL.md -print
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
