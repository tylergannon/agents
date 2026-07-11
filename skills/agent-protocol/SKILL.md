---
name: agent-protocol
description: >
  Read at the start of EVERY session and after context compaction,
  to understand org requirements and expected agent behavior.
---

# Session Worklog

Use this skill when work creates decisions, corrections, proof, source
discoveries, review results, or operational rules that should not disappear into
chat history.

## Session Start

For non-trivial repository work, start in a task worktree before creating the
worklog or writing operating artifacts. If the target repo documents a worktree
path convention, use it. Otherwise default to:

```text
.worktrees/<short-task>
```

The root checkout is an orientation surface for status, search, and branch
selection. If the session is already in a task worktree, continue there. If a
worktree cannot be created, record the blocker before making repo changes.

## Path

Use the target repo's documented worklog path. If none exists, default to:

```text
ephemeral/worklog/YYYYMMDDHHMM-<short-task>.md
```

Use the user's timezone when known. Keep one worklog per coherent session or
task, not one file per command.

## Project Scratch

Only start a project scratch directory when the user explicitly starts or names
a project. Use the target repo's documented project-scratch path. If none
exists, default to:

```text
ephemeral/projects/<project-slug>/
```

Use project scratch directories for artifacts gathered during a project: source
captures, screenshots, generated packets, QA notes, design inputs, temporary
fixtures, and handoff bundles. Before adding project artifacts, check the target
repo's `AGENTS.md` or equivalent for the current project list and choose the
matching project when one exists.

## What To Record

- user goals, constraints, corrections, and changed instructions,
- worktree path, branch name, and project scratch path when applicable,
- relevant commands and results,
- files changed or intentionally left alone,
- decisions, tradeoffs, and rejected alternatives,
- proof commands, failures, blockers, and final status,
- PR, branch, review, or consensus state,
- material source discoveries,
- skill telemetry when a skill is useful, confusing, wrong, blocked, contradicted
  by higher-priority instructions, missing a needed resource, or corrected by
  the user.

Use structured lines when a later fold should mine the entry:

```text
decision: <durable decision and source>
correction: <user correction or rule change>
rule_discovery: <new operating rule>
doc_lookup: <doc/source consulted> -> <why it mattered>
doc_bug: <stale or wrong doc> -> <needed repair>
external_resource: <source> -> <lesson or candidate borrow>
skill_use: <skill> source=<repo-or-local> -> <why it was used>
skill_issue: <skill> source=<repo-or-local> severity=<critical|bug|design|nit> -> <what failed or confused the work>
skill_fix_request: <skill> source=<repo-or-local> -> <smallest requested repair>
```

## Skill Telemetry

Record `skill_issue` or `skill_fix_request` when:

- the user says the agent is applying a skill incorrectly,
- the skill is hard to read, over-specific, missing context, or model-limiting,
- the skill asks for behavior that conflicts with higher-priority instructions,
- the skill assumes files, tools, paths, permissions, or repo structure that are
  not present,
- the skill omits a proof step, bundled resource, review loop, or escalation
  condition needed for the task,
- the agent has to work around or reinterpret the skill to finish correctly,
- a local installed copy of a shared skill is edited or appears stale.

For shared skills, include their source marker when known. For skills from this
repo,
use `source=agents` so daily folds can route the request upstream.

## Boundaries

- Keep the target repo's raw-material directory tracked unless its policy says
  otherwise.
- Do not delete an active worklog.
- Do not turn the worklog into polished documentation.
- During daily docs fold, delete, truncate, or retain ingested material only
  when the fold ledger records the disposition.

## Closeout

Before final response on substantial work, update the worklog with proof run,
remaining debt, branch/PR state, auto-merge or merge state when applicable, and
any material decisions that were made.

source: agents
