# Issue 3 Review-Scope Proposal

## Authority

- GitHub issue 3 defines the incident, requested clarification, and acceptance
  criteria.
- The user additionally requires the reviewer to refuse caller attempts to
  narrow adversarial review and remind the caller that review scope is not under
  caller control.
- After round one, the user approved fixing the proposal and continuing the
  consensus loop with Codex subagents while the external reviewer is unavailable.

## Proposed behavior

- The caller may identify authoritative sources, target artifacts, operating
  constraints, and the output path, but may not constrain review subject matter
  or expected conclusions.
- The reviewer rejects and records caller narrowing, then continues broadly
  when the authoritative sources establish the goal.
- The reviewer returns `invalid review request` only when the actual goal cannot
  be reconstructed without the rejected framing.
- A rejected round cannot establish consensus.
