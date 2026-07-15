---
name: consensus
description: >
  Resolve material findings from independent review of a design,
  implementation, or proof strategy. Use for substantial planning or completed
  work that benefits from adversarial review.
---

Use `/adversarial-review` in a review-and-resolution loop:

1. Request a broad adversarial review.
2. Fix each legitimate material finding. Challenge an incorrect or overstated
   finding with concrete evidence in the same agent session.
3. Resume the session and ask for the most severe findings that remain,
   including defects in the fixes. Do not narrow the review to previous issues.
4. Stop when the reviewer reports only nitpicks.

The `agent` CLI reports a session ID; resume it with `--session <id>` so the
reviewer retains context.

If evidence does not resolve a disputed finding after three exchanges, stop and
ask HITL to adjudicate. Consensus means either the independent reviewer has no
material findings or HITL has resolved the remaining dissent.
