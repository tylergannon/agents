---
name: session-worklog
description: >
  Maintain a tracked session worklog for later use repairing skills, developing
  new ones, and improving documentation.
---

The work log exists in order to learn from the agent's experience so that the
same lessons need not be learned repeatedly.

## Session Start

start in a task worktree before creating the worklog or writing operating artifacts.

## Path

Use the target repo's documented worklog path. If none exists, default to:

```text
ephemeral/worklog/YYYYMMDDHHMM-<task-name>.md
```

Use local time. Keep one worklog per session.  Before creating a PR, rename the
worklog to a short, correct, description of what was done.

## What To Record

The main thing we want to find later is _what went wrongly_ and what lessons
were learned, as well as highly repetitive actions that take up lots of time.

- human has to correct the agent on misconception, wrong implementation, or wrong implementation shape, etc
- human is frustrated by the agent needing too much guidance
- human is frustrated by time taken
- discovery that the instructions were unclear or ambiguous, and need to be corrected
- issues with skills being written in confusing, incoherent, ambiguous, contradictory instructions
- infrastructure problems that slow development or proof
- notes when the user says to remember something or that some specific thing should be part of a skill etc.
- summarize interactions with other agents e.g. consensus achieved, conflicts resolved
- skill telemetry when a skill is useful, confusing, wrong, blocked, contradicted
  by higher-priority instructions, missing a needed resource, or corrected by
  the user.
- difficulty finding needed information about this project

Use structured lines when a later fold should mine the entry:

```text
decision: <durable decision and source>
correction: <user correction or rule change>
rule_discovery: <new operating rule>
doc_bug: <stale or wrong doc> -> <needed repair>
external_resource: <source> -> <lesson or candidate borrow>
skill_issue: <skill> source=<repo-or-local> severity=<critical|bug|design|nit> -> <what failed or confused the work>
```

## Skill Telemetry

Record `skill_issue` when it becomes clear (usually via HITL feedback) that
either the agent didn't know to load the skill, or did load the skill but was
unable to finish the task correctly.

## Boundaries

- Do not add ephemeral to .gitignore.  All of these materials should be tracked and follow the repo.
- Append worklog only.  Do not edit or polish the contents.

## Closeout

Do not add final entries after the PR is merged, or there will be orphaned changes.
