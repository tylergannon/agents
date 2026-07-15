---
name: adversarial-review
description: >
  Ask an independent agent to find material flaws in a design or implementation.
  Use while planning or after implementing a non-trivial feature or refactor.
---

This skill instructs the calling agent. Load it locally, then use the `agent`
CLI with a broad, goal-level prompt. Do not send `/adversarial-review` or any
other skill invocation to the reviewer.

Give the reviewer the issue, requirements, and relevant changes, but do not
steer it toward your suspected file, pattern, or answer.

Write a prompt file in this shape:

```text
Review the feature work on branch codex/issue-XYZ for real implementation
failures. Read the issue and the changes.

Report at most five findings, keeping only the most severe:

- incomplete requirement: cite the missed instruction;
- incorrect implementation: cite the misunderstood requirement;
- verifiable bug: provide a reproduction;
- over-engineering: identify unrequested behavior, complexity, or infrastructure;
- critical antipattern: explain the concrete failure it creates;
- race or crash condition: reproduce it or give a detailed causal explanation.

Classify each finding as critical, issue, or nitpick.
```

Run it with `agent <workdir> --file <prompt-file>`.
