# Reconciliation

Use this when independent agents disagree.

## Finding Labels

Use these labels for review artifacts:

- **critical:** must fix before proceeding.
- **bug:** demonstrable incorrect behavior, broken contract, race, or regression.
- **design:** architecture, boundary, scope, maintainability, or proof issue
  that is materially likely to cause problems.
- **nit:** small cleanup that should not block progress.

When rendering a review prompt, include those definitions so the called agent
does not need this reference file.

## Resolution Order

1. Fix `critical` and `bug` findings or reject them with specific evidence.
2. Fix `design` findings when they threaten correctness, maintainability, scope
   control, or proof; otherwise record why they are non-blocking.
3. Treat `nit` findings as optional cleanup.
4. Reject unsupported findings explicitly and say what evidence is missing.
5. Repeat the independent review after material revisions until the remaining
   findings are only nits.
6. Escalate to the user instead of looping if three rounds have run, the same
   finding repeats without new evidence, or the last revision produced no
   material change.

## Anti-Patterns

- Treating "both models agree" as proof.
- Passing one agent the other's conclusions before it has done an independent
  pass.
- Asking for broad approval instead of concrete findings.
- Asking a narrow prompt such as "check formatting and tests" when the real
  risk is product, design, architecture, or boundary correctness.
- Letting review expand scope without tying it back to the user's goal.

## Final Decision Record

Use this compact format:

```md
Decision: [what we will do]
Evidence: [files, commands, artifacts]
Accepted findings: [short list]
Rejected/deferred findings: [short list with reason]
Review artifact: [path containing prompt and findings]
Proof still required: [commands or artifacts]
```
