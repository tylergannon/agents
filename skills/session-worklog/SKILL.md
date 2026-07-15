---
name: session-worklog
description: >
  Capture actionable intelligence during non-trivial repository work:
  corrections, friction, failures, ambiguous instructions, repeated waste, and
  durable lessons that should change future agent behavior, skills, or docs.
---

The worklog is not an activity transcript. Do not record routine commands,
expected results, or uneventful success. Record information that should change
how future work is done.

## Path

Start the worklog in the task worktree. Use the repo's documented path or:

```text
ephemeral/worklog/YYYYMMDDHHMM-<task-name>.md
```

Use local time and one worklog per session. Before creating a PR, rename it to a
short, accurate description of the work.

## What To Record

- a user correction or important misconception;
- a failed, slow, or repetitive approach and the lesson from it;
- ambiguous, stale, missing, or contradictory instructions;
- tooling or infrastructure friction worth repairing;
- a material decision, discovery, or agent disagreement future work needs;
- a requested memory, skill change, or documentation change.

Use structured lines when a later fold should mine the entry:

```text
decision: <durable decision and source>
correction: <user correction or rule change>
friction: <costly failure or repetition> -> <needed repair>
doc_bug: <stale or wrong doc> -> <needed repair>
skill_issue: <skill> source=<source> severity=<critical|bug|design|nit> -> <problem>
```

## Boundaries

- Keep `ephemeral/` tracked; do not add it to `.gitignore`.
- Append facts as they emerge. Do not polish the worklog into documentation.
- Do not add entries after the PR merges.
