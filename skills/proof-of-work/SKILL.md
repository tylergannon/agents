---
name: proof-of-work
description: >
  Read while authoring coding tasks and during PR review or when asked to
  reassess proof.  Explains this team's expected proof standard and how to
  demonstrate changes completely.  Use for coding tasks, not docs updates.
---

# Proof Of Work

Prove a feature by actually running the software through the feature being
worked on. You must produce an artifact demonstrating that the proof was made.
The specifics on the protocol and artifact will differ by project and stack. If
the repo provides a `/repo-proof` skill, load it now.

Proof means an unbiased, unconditional demonstration of the software actually
accomplishing the new or fixed behavior.

Unit tests and linters are required checks but they are not proof except when
they explicitly exercise the whole application.

## Scope / Applicability

This skill applies when the application is being changed. For docs-only,
policy-only, or scripts/tooling-only work, use targeted checks for those files;
do not require full-stack app proof unless the change affects application
behavior.

## Presenting Proof

The proof should be concrete, legible, and obvious.  It should be presented on
the pull request body, completely.  Do not use paths to local files but instead
upload the proof artifact to the pull request.

## Proof Standard

When creating a task writeup or github issue, include a "Proof" section.
DO NOT include the obvious checks (unit tests, compile check, etc).
Also DO NOT provide a complete todo list of proof steps.
Provide a clear set of claims that the agent should demonstrate in the proof.

## PR Evidence

Lead with what was proved, not a list of commands or basic linters.
Describe the proof and also show material evidence where possible.

If the target repo has a proof-artifact uploader, use that repo's documented
helper and verify the uploaded PR/issue links instead of citing local file
paths. If no uploader exists, say which evidence is only local and why.

## Closeout

Report branch, worktree, PR, head SHA, proof evidence, hygiene checks,
exact-head CI/preview state, auto-merge or merge state, and any blocker.
