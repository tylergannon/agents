# Attractor Proof

Use this reference before claiming an Attractor graph or pipeline fix works.

## Proof Rules

- Run the exact DOT the user cares about.
- For a failure fix, reproduce or identify the failing behavior before editing
  when feasible.
- Inspect logs and artifacts, not just the final status row.
- Treat fallback execution, missing required tools, `partial_success`, and
  review blocker metadata as suspicious until explained.
- If the problem is runner-neutral or provider-specific, state that boundary
  explicitly and avoid burying it in graph syntax.

## Example Sweep Discipline

For broad example sweeps:

1. Use a disposable proof root under `/tmp`.
2. Keep the real checkout clean.
3. Track each example, status, log path, and artifact path.
4. File or record a repro as soon as a failure is confirmed.
5. Validate green rows by spot-checking response artifacts for exact-output and
   invalid-proof classes.

## Repo Fix Closeout

For Attractor issue work:

1. Work in a branch or worktree.
2. Prove the exact failing DOT or exact feature path.
3. Rebase onto current `main`.
4. Publish a PR with proof in the body.
5. Squash merge when checks and proof are complete.
6. Remove worktrees, temp proof roots, and stale runner processes.

## Artifact Upload Path

Some repos use PR helpers that upload proof artifacts and rewrite local
filenames in a draft body to hosted proof URLs. When borrowing that pattern,
verify the published PR body, not only the local body file.
