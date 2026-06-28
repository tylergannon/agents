# Submodule Index

This directory is the local index for the reference repositories mounted under
`submodules/`.

## Layout

- `submodules/<name>/` contains upstream Git submodules pinned by this repo.
- `submodules/index/` contains local analysis and recommendations. It is not a
  submodule.
- Keep observations here instead of editing upstream submodule contents.

## Attached Repositories

| Submodule | Upstream | Pinned commit | Local shape | Index |
|---|---|---:|---|---|
| `caveman` | `JuliusBrussee/caveman` | `655b7d9` | 145 files, 15 `SKILL.md` files, 7 agent files, 9 command files, Node installer | [caveman.md](caveman.md) |
| `superpowers` | `obra/superpowers` | `6fd4507` | 147 files, 14 workflow skills, hook/plugin manifests, 52 tests | [superpowers.md](superpowers.md) |
| `gstack` | `garrytan/gstack` | `1626d48` | 1121 files, 58 skills, 72 helper binaries/scripts, persistent browser stack, 510 tests | [gstack.md](gstack.md) |
| `gbrain-evals` | `garrytan/gbrain-evals` | `565b807` | 676 files, memory eval corpora, schemas, runners, scorecards, 22 tests | [gbrain-evals.md](gbrain-evals.md) |
| `compound-engineering-plugin` | `EveryInc/compound-engineering-plugin` | `e74e298` | 576 files, converter CLI, Compound Engineering plugin, 45 local `SKILL.md` files including fixtures, 46 visible agent files, 88 tests | [compound-engineering-plugin.md](compound-engineering-plugin.md) |

All five upstream repositories currently carry MIT licenses. If code, docs, or
substantial prompt text is copied out of a submodule, preserve the required
license notice. Prefer adapting patterns into new generalized agents
artifacts over vendoring large upstream chunks.

## What We Can Learn

The strongest shared lesson is that agent tooling is no longer just prompts. The
best systems here combine compact skill entrypoints, deeper reference files,
tests or evals that exercise real agent behavior, and installers that understand
multiple hosts.

- **Progressive disclosure works.** Superpowers, gstack, and Compound
  Engineering keep user-facing entrypoints small while linking reference docs,
  scripts, and persona prompts that are loaded only when needed.
- **Workflow skills need proof harnesses.** Superpowers tests full Claude Code
  sessions. gstack and gbrain-evals invest heavily in Bun tests, evals, sealed
  answers, transcript checks, and scorecards.
- **Multi-host packaging is becoming a product surface.** caveman and Compound
  Engineering both treat Claude, Codex, Gemini, OpenCode, Cursor, and other
  hosts as install targets with explicit file layouts.
- **Agent-native tools need local services when latency or state matters.**
  gstack's browser stack uses a persistent local daemon rather than cold-starting
  a browser per action.
- **Memory and learning should be measured.** gbrain-evals is the clearest model
  for public corpora, sealed answers, baselines, and honest reporting of weak
  numbers next to strong ones.
- **Brevity can be a first-class mode.** caveman packages terse communication,
  memory compression, stats, tests, and safety checks as a reusable tool rather
  than an ad hoc style request.

## What We Can Use

Use the submodules as reference implementations and design source material:

- `../caveman/skills/` and `../caveman/plugins/caveman/` for terse-mode skill
  packaging, memory compression, and cross-agent installer ideas.
- `../superpowers/skills/` and `../superpowers/tests/` for disciplined workflow
  skills, TDD enforcement, subagent review loops, and skill-trigger testing.
- `../gstack/browse/`, `../gstack/bin/`, and `../gstack/docs/designs/` for
  persistent browser tooling, local daemon security, QA workflows, and helper CLI
  design.
- `../gbrain-evals/eval/`, `../gbrain-evals/qrels/`, and
  `../gbrain-evals/docs/benchmarks/` for memory/retrieval eval harnesses,
  committed corpora, sealed answer keys, and scorecard shape.
- `../compound-engineering-plugin/src/`, `../compound-engineering-plugin/plugins/`,
  and `../compound-engineering-plugin/docs/solutions/` for converter/writer
  architecture, plugin release validation, persona reviews, and compounding
  knowledge docs.

For prompt-writing-specific skill research, see
[prompt-writing-skill-research.md](prompt-writing-skill-research.md). It maps
the strongest prompt design, skill-writing, rubric, and eval guidance across the
mounted repositories.

For recent AI-first engineering sources from OpenAI, Anthropic, Cursor,
Cognition, Replit, Vercel, and high-signal operators, see
[recent-ai-engineering-loop-resources.md](recent-ai-engineering-loop-resources.md).
It tracks loop engineering, dynamic workflows, harness design, and the shift
from one-shot prompts to agent feedback systems.

## What We Can Borrow

See [borrow-map.md](borrow-map.md) for a cross-repo extraction plan. The short
version:

1. A reusable plugin inventory scanner and index generator.
2. A host-target abstraction for converting skills/plugins across Codex, Claude,
   Gemini, OpenCode, Cursor, and related agents.
3. A workflow-skill test harness that can prove trigger behavior, transcript
   shape, and real task completion.
4. A small eval corpus pattern with sealed answers and scorecard docs.
5. A persistent local-service pattern for stateful agent tools, especially
   browser or UI tools.
6. A repo-local learnings/solutions format that compounds solved problems
   without turning `AGENTS.md` into a catch-all.
7. A prompt-writing skill built from Superpowers' skill-writing TDD discipline,
   Compound Engineering's prompt architecture and optimization loops, gstack's
   prompt-boundary patterns, and gbrain-evals-style held-out fixtures.

## Maintenance

When updating submodules, run:

```sh
git submodule update --init --recursive
git submodule status
```

Then refresh this index with the new commit ids and any material changes to
skills, agents, tests, or reusable patterns.
