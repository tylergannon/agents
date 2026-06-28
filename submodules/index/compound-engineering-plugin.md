# Compound Engineering Plugin

- Path: `../compound-engineering-plugin`
- Upstream: `EveryInc/compound-engineering-plugin`
- Pinned commit: `e74e29864fbfa2f800fc3e08509e2966e4947f1e`
- License: MIT
- Package: `@every-env/compound-plugin` `3.11.2`

## What Is In It

- A multi-host plugin distribution for Compound Engineering skills and agents.
- A Bun/TypeScript CLI under `src/` with commands for conversion, installation,
  listing, cleanup, and plugin-path discovery.
- Target converters/writers for Codex, Gemini, Kiro, OpenCode, Pi, and related
  hosts.
- The main plugin under `plugins/compound-engineering/`, with workflow,
  planning, review, documentation, product, setup, release, browser, Xcode, Slack,
  and worktree skills.
- Specialized agents under `plugins/compound-engineering/agents/`, including
  correctness, security, performance, data integrity, design, product, scope,
  session history, Slack, GitHub issue, and repo research personas.
- A secondary `plugins/coding-tutor/` plugin.
- Large docs areas: skill docs, plans, brainstorms, solution notes, specs, and
  best-practice writeups.
- Release validation scripts and extensive tests.

Local scan: 576 files, 45 `SKILL.md` files including fixtures, 46 visible
`agents/` files, 157 docs files, 88 tests. Upstream docs describe the shipped
plugin as 37 skills and 50+ agents; the local count includes fixtures and support
files differently.

## What We Can Learn

- A plugin is cleaner when parsed once into a neutral model and then converted by
  target-specific converters/writers.
- Domain vocabulary matters. `CONCEPTS.md` distinguishes plugin, skill, agent,
  converter, writer, bundle, marketplace, learning, pattern doc, confidence
  anchor, and autofix class.
- Compounding engineering treats solved problems as reusable assets. Learnings
  and pattern docs are explicit artifacts, not buried in transcripts.
- Review quality improves when persona agents have narrow lenses and findings
  carry confidence anchors and autofix classes.
- Installation docs need to state host limitations directly. Codex, for example,
  currently needs both native plugin installation and a separate Bun agent
  install step for custom agents.

## What We Can Use

- `src/converters/`, `src/targets/`, and `src/types/` as the strongest reference
  for multi-host plugin conversion.
- `src/release/` and `scripts/release/validate.ts` for release metadata and
  validation ideas.
- `docs/solutions/` for reusable learning and best-practice doc structure.
- `docs/skills/` for user-facing skill documentation shape.
- `plugins/compound-engineering/skills/ce-code-review/` and
  `plugins/compound-engineering/agents/` for persona review design.
- `tests/` for path safety, manifest invariants, target writer tests, and skill
  contract tests.

## What We Can Borrow

- A neutral plugin component schema and target writer pattern for shared
  `agents` plugin utilities.
- A release validation command that checks marketplace metadata, manifests,
  target outputs, and generated docs before publishing.
- A confidence-gated review aggregation model.
- A learnings/pattern-doc taxonomy for durable repo-local lessons.
- A setup/health-check skill pattern for validating user environments before a
  workflow starts.

## Cautions

- Some generated or target-specific artifacts exist alongside source files.
  Borrow the converter pipeline rather than editing generated outputs by hand.
- Review persona prompts are powerful but can produce process bloat if copied
  wholesale. Start with the narrow reviewer lenses this repo actually needs.
- Preserve MIT license notice for substantial copied code or prompt text.
