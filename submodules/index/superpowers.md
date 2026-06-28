# Superpowers

- Path: `../superpowers`
- Upstream: `obra/superpowers`
- Pinned commit: `6fd4507659784c351abbd2bc264c7162cfd386dc`
- License: MIT
- Package: `superpowers` `5.1.0`

## What Is In It

- A composable software-development methodology packaged as agent skills.
- Fourteen workflow skills under `skills/`, including brainstorming,
  test-driven development, systematic debugging, worktrees, writing plans,
  requesting/receiving review, subagent-driven development, and branch finishing.
- Host/plugin surfaces for Claude, Codex, Cursor, OpenCode, Gemini, and related
  agents.
- Session hooks under `hooks/`.
- Planning/spec docs under `docs/`.
- Integration and trigger tests under `tests/`, including headless Claude Code
  transcript verification for subagent-driven development.

Local scan: 147 files, 14 `SKILL.md` files, 17 docs files, 52 tests.

## What We Can Learn

- The system treats workflow skills as mandatory operating procedures, not
  optional tips. That makes the agent check for relevant skills before work.
- TDD and debugging rules are effective when they are explicit about sequence:
  reproduce, write the failing test, make it pass, refactor, verify.
- Subagent-driven work can be tested by running real headless agent sessions and
  inspecting transcripts for tool use, task dispatch, review order, commits, and
  passing tests.
- Plans are written for execution by a less-contextual worker. That forces exact
  paths, small tasks, and verification steps.

## What We Can Use

- `skills/test-driven-development/SKILL.md` for red-green-refactor sequencing.
- `skills/systematic-debugging/SKILL.md` for root-cause discipline.
- `skills/writing-plans/SKILL.md` for plan granularity and review prompts.
- `skills/subagent-driven-development/SKILL.md` and
  `tests/claude-code/` for transcript-based workflow proof.
- `tests/skill-triggering/` and `tests/explicit-skill-requests/` for skill
  routing tests.

## What We Can Borrow

- A shared skill-testing harness that checks whether a skill actually triggers
  and produces the expected workflow evidence.
- A branch-finishing checklist for verification, merge/PR decision, and cleanup.
- A worktree lifecycle skill for projects that require isolated issue branches.
- A minimal "write plans for another agent" format for cross-agent handoff.

## Cautions

- Superpowers is intentionally prescriptive. Borrow the testable workflow shape,
  but adapt any process language to this repo's existing preference for proof and
  concise execution.
- Long-running integration tests can be expensive. Keep local variants small and
  reserve full agent-session tests for high-risk workflow changes.
