---
name: agent-protocol
description: >
  Read at the start of EVERY session and after context compaction,
  to understand org requirements and expected agent behavior.
---

1. Always work in a worktree.  Root checkout -> read-only orientation surface.
2. Always maintain a work log for non-trivial tasks.  Read the /session-worklog skill in that case.
3. Keep ephemeral files in ./ephemeral/.  NEVER write to ./docs without express permission.
4. Always append to worklog and make a commit at the end of performing a task.  Do not wait
   to be told to make commits.  Do not spend time doing checks before and after these commits.
   They are meant to be quick checkpoints, not durable records of work.  The PR will be
   squash-merged and that is the time to focus on correctness protocol.
5. Pull requests are to be squash-merged.  The PR body should have sufficient information
   to become a suitable commit message when the squash takes place.  Extended write-ups and
   review work should go into the PR comments.
6. Clean up after yourself after PR merged.  Delete worktree and close any background processes.
7. When you find things are broken, fix them.  Broken unit tests or broken CI build is
   not to be ignored, whether you broke it or someone else broke it.  If pre-existing issues
   threaten to really distract you from your goal, seek HITL to adjudicate possibly move on
   in spite of the broken state.
8. Update root checkout to latest before starting work and after PR merge.

## Ephemeral Directory

`ephemeral/` is for documents that are not intended as permanent artifacts, including:

- work log
- project and task plans
- downloaded artifacts that should travel with the repo

## Project Scratch

Only start a project scratch directory when the user explicitly starts or names
a project. Use the target repo's documented project-scratch path. If none
exists, default to:

```text
ephemeral/projects/<project-slug>/
```

Use project scratch directories for artifacts gathered during a project: source
captures, screenshots, generated packets, QA notes, design inputs, temporary
fixtures, and handoff bundles. Before adding project artifacts, check the target
repo's `AGENTS.md` or equivalent for the current project list and choose the
matching project when one exists.
