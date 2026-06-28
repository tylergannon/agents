# agents

This repo is a skills library and shared agent-tooling package for projects that
want reusable agent capabilities.

The deliverable is reusable agent capability: skills, skill references, bundled
scripts, evals, harnesses, workflow templates, and reference indexes that can be
installed into other repositories. Do not treat this repo as the target software
project when importing a workflow from another repo; translate the workflow into
portable skills and resources that other repos can use.

## Notes

- Keep repo-local notes in this file when they affect how future agents should
  work in this repo.
- Keep durable implementation notes small and close to the maintained surface.
  Do not create a repo-local documentation operating system here when the
  requested deliverable is a portable skill for other repositories.
- Use `ephemeral/` for operating worklogs, review artifacts, decision/correction
  capture, and other material that daily jobs will mine into durable docs or
  skills. Session worklogs belong under `ephemeral/worklog/`; Claude Code /
  Codex consensus reviews belong under `ephemeral/reviews/`. Project scratch
  artifacts belong under `ephemeral/projects/<project>/` only after the user
  starts or names a project.
- `ephemeral/` is tracked but ignored by routine ripgrep. Search it
  deliberately during documentation maintenance.
- Do not assume a tool belongs here just because it is useful once. Prefer this
  repo for generalized workflows, reusable scripts, agent skills, and shared CI
  infrastructure.
- Treat this repo as an active lab for reusable agent practice. Core skills
  should improve over time as new techniques, models, tools, and workflows become
  available.
- Always look for inspiration from the strongest current sources in AI-first
  software engineering: frontier model teams, serious agent-tool builders,
  excellent public repos, production writeups, eval harnesses, and high-signal
  operator notes. Capture useful sources in `submodules/index/` so they can be
  mined later.
- Favor skills that encode durable judgment, compact prompts, good loops,
  executable checks, and project-portable workflows. Avoid baking narrow
  project-specific instructions into core skills unless the pattern clearly
  generalizes.
- When borrowing from a project repo, preserve the useful skill behavior,
  scripts, review loops, and proof gates as skill resources. Do not copy that
  repo's local docs taxonomy here unless this library needs it as a
  library-level convention.
- `submodules/` is the mount point for upstream reference repositories used to
  study reusable agent tooling. `submodules/index/` is the local tracked index of
  what is in those submodules, what can be learned from them, and what patterns
  are candidates to borrow.
- Keep local analysis in `submodules/index/`; do not edit upstream submodule
  contents for repo-local notes.
- When updating, adding, or removing submodules, refresh `submodules/index/` with
  the pinned commits and material changes to skills, agents, tests, docs, and
  reusable patterns.
- If copying substantial code, docs, or prompt text out of a submodule, preserve
  upstream license notices. Prefer extracting generalized tooling over vendoring
  large upstream chunks.
- `skills/` is the canonical source for published, portable agent skills.
- `.agents/skills/` holds repo-local maintainer skills that should not be
  distributed to target repositories.
- Use the public skills CLI to install published skills into target agent
  systems once this repo has a remote. Do not maintain a parallel custom linker
  unless a missing capability is explicitly identified.
- When adding or changing skills, keep `.codex-plugin/plugin.json`,
  `.claude-plugin/plugin.json`, `README.md`, `skills/README.md`, package
  metadata, and validation checks in sync.
- End published `SKILL.md` files with `source: agents` so installed copies can
  route fixes back to this repo.
- Use `skills/session-worklog`, `skills/proof-of-work`, and
  `skills/daily-docs-fold` to preserve operating evidence, own behavior proof,
  and fold raw material into durable docs. Use `skills/dependency-upgrade-pr` for
  dependency-update PRs that need exact-head CI and preview/runtime proof.
- Worklogs in target repos should record shared-skill telemetry with
  `skill_use`, `skill_issue`, and `skill_fix_request`; daily docs folds should
  turn material `source=agents` failures into upstream issues or ready-to-file
  issue drafts.

## Skill Roadmap

The first skills to develop here are:

- **Prompt writing / editing:** help agents write concise, high-leverage prompts
  that assume model competence, separate intent from deterministic enforcement,
  and move repeated work into tools, evals, skills, or loops.
- **Attractor pipeline editing:** help agents safely design, edit, prove, and
  document Attractor pipeline changes across local harnesses and project repos.
- **Claude Code / Codex consensus:** help Claude Code and Codex review each
  other's design and implementation plans, surface disagreements, and converge on
  stronger implementation processes.

## Current Tooling

- `npm test` validates package metadata, package contents, and documentation
  index checker behavior.
- Public install path for supported systems:
  `vp dlx skills add tylergannon/agents -a <agent> --yes --all`.
- Current reference submodules:
  - `submodules/caveman`
  - `submodules/superpowers`
  - `submodules/gstack`
  - `submodules/gbrain-evals`
  - `submodules/compound-engineering-plugin`
