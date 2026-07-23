---
name: agent-tooling-research
description: >
  Study mounted agent-tooling reference repositories and turn them into concrete
  reusable agents patterns. Use when comparing Caveman, Superpowers, gstack,
  gbrain-evals, Compound Engineering, or when borrowing skill/plugin/eval ideas.
metadata:
  internal: true
---

# Agent Tooling Research

Use this skill when the task is to inspect, compare, or borrow from reference
agent-tooling repositories mounted under `submodules/`.

## Workflow

1. Read `submodules/index/README.md` and `submodules/index/borrow-map.md`.
2. Open the relevant per-repo index file before reading upstream source.
3. Inspect upstream files directly when making a concrete claim.
4. Keep new local analysis in `submodules/index/` or durable repo docs.
5. Do not edit upstream submodule contents for agents notes.

## Output

Return concrete findings:

- what exists upstream,
- what is reusable here,
- what should only be referenced,
- what needs tests before borrowing,
- what license notice or attribution is required.

Prefer exact paths and command evidence over broad summaries.

source: agents
