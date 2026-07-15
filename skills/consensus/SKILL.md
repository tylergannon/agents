---
name: consensus
description: >
  Resolve material findings from independent review of a design,
  implementation, or proof strategy. Use for substantial planning or completed
  work that benefits from adversarial review.
---

## Orchestration Boundary

`/adversarial-review` is a skill for you, the calling agent. Load and follow it
locally before invoking the independent reviewer. The reviewer receives only
the review prompt; never put `/adversarial-review`, `/consensus`, or another
skill invocation in that prompt.

Write the review prompt to a file and pass it with `--file`:

```text
wrong: agent <workdir> "/adversarial-review Review this implementation..."
right: agent <workdir> --file <prompt-file>
```

## Review Loop

1. Load `/adversarial-review` locally and write the initial review prompt using
   its broad, goal-level contract.
2. Run `agent <workdir> --file <prompt-file>` and retain the reported session ID.
3. Fix each legitimate material finding. Challenge an incorrect or overstated
   finding with concrete evidence in the same agent session.
4. Write a follow-up prompt asking for the most severe findings that remain,
   including defects in the fixes. Do not narrow the review to previous issues.
5. Resume with `agent <workdir> --session <id> --file <follow-up-prompt>`.
6. Stop when the reviewer reports only nitpicks.

If evidence does not resolve a disputed finding after three exchanges, stop and
ask HITL to adjudicate. Consensus means either the independent reviewer has no
material findings or HITL has resolved the remaining dissent.
