---
name: consensus
description: >
  Caller-side loop for resolving material findings from independent review of
  a design, implementation, or proof. Use after `/request-adversarial-review`
  when substantial work needs repeated review and adjudication.
---

# Consensus

Run each review through Gimbal's headless prompt runner in the target worktree.
Give the reviewer the `/adversarial-review` skill and the complete current
authoritative context. Each round is a new Gimbal prompt run with a fresh
review artifact; do not invoke the separate `agent` executable.

1. Use `/request-adversarial-review` to start the review and retain the Gimbal
   run ID and review-file path. Read the review artifact.
2. Fix each legitimate material finding. Challenge an incorrect or overstated
   finding with concrete evidence in the next review prompt.
3. After fixes, do not summarize earlier findings, fixes, resolved areas, or
   expected remaining problems. Do not add positive or negative subject-matter
   limits, expected conclusions, or desired verdicts. Start a fresh Gimbal
   prompt run with a minimal prompt and a new output path.

Bad:
```text
gimbal run-prompt --workdir <workdir> "/adversarial-review Re-check the fixes
for X and Y. Do not broaden into Z. Write to <next-review-file>."
```

Good:
```text
gimbal run-prompt --workdir <workdir> "/adversarial-review Re-review the entire
current work against the same authoritative sources. Inspect the current
implementation and proof. Write the complete review to <next-review-file>. Do
not edit any other files."
```

4. Use a new `ephemeral/reviews/...-round-NN.md` path for every round; never
   overwrite an earlier review artifact.
5. Re-review the entire current target; previous findings are not the review
   scope.
6. Read each artifact and stop when its outcome is `only nitpicks remain` or
   `no findings`.

If the reviewer returns `invalid review request`, discard that run and restart
with `/request-adversarial-review`; the rejected round cannot count as
consensus.

If evidence does not resolve a disputed finding after three exchanges, stop and
ask HITL to adjudicate. Consensus means either the independent reviewer has no
material findings or HITL has resolved the remaining dissent.
