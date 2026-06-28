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
