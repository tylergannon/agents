---
name: consensus
description: >
  Caller-side loop for resolving material findings from independent review of
  a design, implementation, or proof. Use after `/request-adversarial-review`
  when substantial work needs repeated review and adjudication.
---

# Consensus

The `agent` CLI selects the reviewer automatically. Never pass a model or agent
name as a positional argument; `<workdir>` must immediately follow `agent`.

1. Use `/request-adversarial-review` to start the review and retain the spawned
   reviewer's session ID and review-file path. Read the review artifact.
2. Fix each legitimate material finding. Challenge an incorrect or overstated
   finding with concrete evidence in the same reviewer session.
3. After fixes, resume the reviewer with a minimal prompt:

```text
agent <workdir> --session <id> "/adversarial-review Re-review the current work
after the fixes. Report the most severe findings that remain, including defects
in the fixes. Write the complete review to <next-review-file>. Do not edit any
other files."
```

4. Use a new `ephemeral/reviews/...-round-NN.md` path for every round; never
   overwrite an earlier review artifact.
5. Do not narrow subsequent review to the previous findings.
6. Read each artifact and stop when its outcome is `only nitpicks remain` or
   `no findings`.

If evidence does not resolve a disputed finding after three exchanges, stop and
ask HITL to adjudicate. Consensus means either the independent reviewer has no
material findings or HITL has resolved the remaining dissent.
