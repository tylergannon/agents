---
name: consensus
description: >
  Achieve consensus on design, implementation, or proof strategy.
  Use to refine architecture and planning documents as well as for
  review and approval after substantial work has been done.
---

Use the `/adversarial-review` skill in a loop to achieve consensus on design, implementation, or proof strategy.

The gist is to loop between review and resolutions until the other agent has only
nitpick review complaints.

1. Ask for adversarial review
2. Resolve all legitimate issues.  If you think a specific issue report is incorrect,
   chat with the agent until you agree on what the issue is and how to resolve it.  See below.
3. Repeat until the adversarial review has only nitpick review complaints.

It is best to continue the existing session on subsequent rounds rather than starting a new one,
to avoid the agent needing to re-read EVERYTHING.  On subsequent rounds after having addressed
the issues, keep the prompt wide open and let the agent explore further.

The `agent` CLI will report the session ID.  Subsequent calls can use `--session [Session ID]`
to continue the same session.

Basically use the same prompt except for telling the agent this time,
"the following issues are believed to be resolved.  please resume your review process and return
the top five issues you believe are still unresolved, including but not limited to defects in
the fixes I have applied."

The MOST IMPORTANT point to consensus is to avoid narrowing, limiting, or conditioning
your review prompts.

## Prompt Quality Rules

- Ask at the level of the actual goal, not at the level of a narrow local check.
- Give the called agent enough context to discover the right review frame.
- Do not prohibit useful exploration. Narrow only destructive actions.
- Do not ask for approval. Ask for real findings.
- Do not leak your intended answer unless the task is to critique that answer.

## Chat With the Agent

If you believe a specific issue report is inappropriately reported or lower priority,
chat with the agent to confirm and resolve it.

Simply phrase your question clearly and send it back to the same session.

e.g.

```
Issue 3 regarding a race condition was implemented correctly.  Note the lock at line 123 that
synchronizes access to the contended object.
```

In this way you can maintain a dialog with the agent until you agree.  If three turns are
made without agreement, stop and seek HITL adjudication.
