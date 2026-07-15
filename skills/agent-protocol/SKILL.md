---
name: agent-protocol
description: >
  Read at the start of every session and after context compaction to understand
  organization requirements and expected agent behavior.
---

1. Do task work in a worktree. Use the root checkout only for orientation and
   fast-forward synchronization.
2. For non-trivial tasks, maintain a worklog using `/session-worklog`. Record
   actionable intelligence, not routine activity.
3. Keep temporary operating material under `ephemeral/`. Never write to `docs/`
   without express permission.
4. Append to the worklog and make a checkpoint commit after each task. Do not
   rerun checks solely because of the commit. Complete the relevant proof and
   correctness protocol before declaring the PR merge-ready.
5. Squash-merge PRs. Write the PR body so it can serve as the squash commit
   message; put extended write-ups and review discussion in PR comments.
6. After merge, delete the task worktree and stop its background processes.
7. Fix broken tests or CI regardless of origin. If a pre-existing failure would
   materially distract from the task, ask HITL whether to proceed despite it.
8. Fast-forward the root checkout before starting work and after the PR merges.

`ephemeral/` is tracked working material such as worklogs, plans, and downloaded
artifacts. Create `ephemeral/projects/<project-slug>/` only when the user starts
or names a project, and follow any project list documented by the target repo.
