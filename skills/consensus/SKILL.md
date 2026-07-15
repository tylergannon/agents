---
name: consensus
description: >
  Caller-side loop for resolving material findings from independent review of
  a design, implementation, or proof. Use after `/request-adversarial-review`
  when substantial work needs repeated review and adjudication.
---

# Consensus

1. Use `/request-adversarial-review` to start the review and retain the spawned
   reviewer's session ID.
2. Fix each legitimate material finding. Challenge an incorrect or overstated
   finding with concrete evidence in the same reviewer session.
3. After fixes, resume the reviewer with a minimal prompt:

```text
agent <workdir> --session <id> "/adversarial-review Re-review the current work
after the fixes. Report the most severe findings that remain, including defects
in the fixes. Do not edit files."
```

4. Do not narrow subsequent review to the previous findings.
5. Stop when the reviewer reports only nitpicks.

If evidence does not resolve a disputed finding after three exchanges, stop and
ask HITL to adjudicate. Consensus means either the independent reviewer has no
material findings or HITL has resolved the remaining dissent.
