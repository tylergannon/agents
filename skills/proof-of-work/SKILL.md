---
name: proof-of-work
description: >
  Prove implementation work by demonstrating changed behavior in a full running
  application stack, then publish the PR with evidence and auto-merge by
  default. Use for application behavior changes before claiming work is done,
  PR-ready, or merge-ready. Do not use for docs-only, policy-only, or
  scripts/tooling-only changes.
---

# Proof Of Work

Unit tests, typechecks, linters, builds, and format checks are hygiene. They are
closeout checks. They are useful after the behavior works, but they do not prove
the software works and should not drive the implementation.

This skill applies when the application is being changed. For docs-only,
policy-only, or scripts/tooling-only work, use targeted checks for those files;
do not require full-stack app proof unless the change affects application
behavior.

Proof means a concrete demonstration that the full running application stack
actually performs the feature or fix being shipped. Develop that proof path from
the start of the task and use it as the main implementation loop.

## Proof Standard

Default proof is browser proof: run the app, exercise the changed behavior
through the real UI in Chrome.app with computer use or Playwright, and capture
screenshots, traces, video, or equivalent browser evidence.

If there is no UI, prove the behavior through the running app's real
API/job/webhook/runtime path. Attach the concrete artifact the feature produces:
response, persisted state, generated PDF, email, storage object, callback log,
or similar.

In all cases, proof must use the full running application stack. Do not claim
proof from one-off scripts, mocked paths, isolated unit tests, or library calls
outside the running app.

Do not be stingy with proof. This is how you demonstrate that the work is done:
capture and post screenshots, traces, rendered artifacts, responses, or
persisted state for every user-visible or contract-bearing claim that bears
proving.

Local paths are not proof. Evidence files must be uploaded and posted to the PR
or issue so reviewers can inspect them. If artifacts cannot be uploaded, say the
work is not fully proved and report the blocker.

## Workflow

1. Start from current `origin/main` in a task branch or worktree.
2. Create or update the session worklog before material edits.
3. Start the relevant runtime stack and establish the smallest behavior proof
   that exercises the real user/API/job/webhook path. Run it against the current
   code when useful to show the missing or failing behavior.
4. Implement in tight loops against that behavior proof until the changed
   behavior passes on the real path.
5. Run only the hygiene checks that are relevant to the changed files and fix
   regressions after the behavior proof is working.
6. Upload and post proof artifacts: screenshots, traces, rendered files,
   responses, logs, or persisted-state evidence.
7. Commit, push, and open or update the PR.
8. Put the proof in the PR body or a PR comment.
9. Verify the PR head SHA before trusting CI, previews, or browser evidence.
10. Enable squash auto-merge by default once proof passes and required checks
    are green or pending.
11. If the branch goes stale, update it, rerun impacted proof, and keep
    auto-merge configured.
12. Follow through until the PR merges or a real blocker is reported.

## PR Evidence

Lead with what was proved, not a list of commands.

For browser changes, include the URL, exact user path exercised, uploaded
screenshot or trace links, and what the evidence shows.

For non-browser changes, include the runtime path exercised, input, output,
persisted state or produced artifact, and the assertion it proves.

If the target repo has a proof-artifact uploader, use that repo's documented
helper and verify the uploaded PR/issue links instead of citing local file
paths. If no uploader exists, say which evidence is only local and why.

## Blockers

Only stop before merge for a real blocker: current-diff product bug, broken full
stack, incompatible dependency breakage, missing secret, failing required check,
or external service outage.

Inherited failures should be recorded with evidence, but they should not replace
proof of the changed behavior.

## Closeout

Report branch, worktree, PR, head SHA, proof evidence, hygiene checks,
exact-head CI/preview state, auto-merge or merge state, and any blocker.

source: agents
