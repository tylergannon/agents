---
name: request-adversarial-review
description: >
  Construct and launch a minimal prompt for an independent agent to load
  `/adversarial-review` and review a design, implementation, or proof. Use from
  the calling agent for a one-pass review or to start a `/consensus` loop.
---

# Request Adversarial Review

This skill is for the calling agent. The spawned reviewer owns the review
method by loading `/adversarial-review`; do not reproduce that skill's rubric in
the launch prompt.

1. Identify the worktree and authoritative issue, specification, plan, or proof
   artifact the reviewer should read. Choose a new review path such as
   `ephemeral/reviews/YYYYMMDDHHMM-<task>-round-01.md`.
2. Include only context needed to locate the review target. If a requirement
   exists only in conversation, state it concisely without turning it into a
   checklist or directing attention to suspected defects.
3. Launch the reviewer with a short prompt:

```text
agent <workdir> "/adversarial-review Review the current implementation against
the issue and repository instructions. Inspect the working tree and relevant
surrounding code. Write the complete review to <review-file>. Do not edit any
other files."
```

Retain the reported session ID and review-file path. Read the artifact before
acting on findings. For a one-pass review, summarize and link the artifact. For
repeated review and adjudication, continue with `/consensus`.
