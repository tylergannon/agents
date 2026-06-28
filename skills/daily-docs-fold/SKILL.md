---
name: daily-docs-fold
description: >
  Fold recent commits, worklogs, reviews, research notes, user corrections, and
  source discoveries into durable documentation in a target repository. Use for
  daily or periodic documentation maintenance, docs-index repair, policy/manual
  updates, or when raw repo work needs to become durable docs.
---

# Daily Docs Fold

Use this skill to turn raw operating material in a target repository into
durable, searchable documentation.

The goal is not to write more docs. The goal is to preserve the smallest useful
set of policy, manuals, design routes, source indexes, skill updates, and
journal history that future agents can actually use.

First read the target repo's `AGENTS.md` or equivalent policy. Use its local
names for raw material directories, such as `ephemera/` or `ephemeral/`, and its
current project list when it has one.

## Inputs

Build an input ledger from:

- landed commits and PR metadata since the target repo's last ingestion marker,
- every file under the target repo's raw-material directory,
- worklogs, review artifacts, research captures, and source notes,
- user corrections, decisions, and rule discoveries,
- skill telemetry lines such as `skill_use`, `skill_issue`, and
  `skill_fix_request`,
- git changes to installed, vendored, or copied shared skill files,
- reference-repo, submodule, or external source discoveries,
- stale docs found during ordinary work.

Prefer the landed commit as the canonical row when its message contains the
useful PR body. Use PR metadata as fallback or supplemental metadata, not as a
duplicate row.

## Classification

Classify each row into one or more targets:

- `AGENTS.md`, `CLAUDE.md`, or the repo's policy surface for mandatory agent
  behavior.
- `docs/policy/` or equivalent for repo policy.
- `docs/manuals/` or equivalent for preferred techniques and routes.
- `docs/design/` or equivalent only for concept routing and architecture
  claims, not process logs.
- repo-local skills for reusable executable workflows.
- upstream skill issue requests when telemetry names a shared skill source such
  as `source=agents`.
- source indexes for upstream inventory, lessons, and borrow candidates.
- journal, changelog, or release notes for chronology and ingestion markers.
- raw-material directories only when raw evidence still has active value.
- project-scoped scratch material such as `ephemeral/projects/<project>/`: list
  active projects in the ledger but defer folding their contents until done.
- `AGENTS.md`, `CLAUDE.md`, or the repo's policy surface for the current project
  list.
- discard when the row is transient, duplicate, or already represented.

## Loop

Run the three-phase loop from [phases.md](references/phases.md):

1. Build Up: represent every material input in the right durable target.
2. Compress: remove redundancy, stale wording, and oversized leaves.
3. Rebalance: move content to the right layer and fix routing.

For substantial folds, get independent review for each phase. The review prompt
should ask for wrong-directory, missing-source, stale-doc, duplicate-input,
overfit-skill, and routing failures. Continue until only nits remain or a
blocker needs human judgment.

## Source Handles

Durable docs should cite stable handles: paths, headings, script names, test
names, workflow names, skill names, commit SHAs, PR numbers, artifact paths, or
explicit drift notes. Avoid durable `file:line` citations.

## Procedure

1. Create or update a session worklog in the target repo's worklog path.
2. Identify the ingestion range from the target repo's journal, changelog, or
   prior fold marker.
3. Write the input ledger and disposition for each row. Include every current
   project scratch directory from `ephemeral/projects/` or the repo's equivalent
   path. Treat these directories as scratch space for artifacts gathered during a
   user-started project, not as general worklog storage.
4. Run Build Up, Compress, and Rebalance.
5. For `skill_issue`, `skill_fix_request`, or local edits to installed, copied,
   or vendored shared skills, identify the source repo from a source marker,
   package metadata, install manifest, or nearby docs. For material
   `source=agents` issues, search the agents issue tracker
   for an existing open report first. If one exists, add the new evidence there
   or record that it was already covered. If none exists, open a GitHub issue
   with the worklog evidence and smallest requested repair. Do not file `nit`
   severity telemetry upstream unless it is recurring, blocking, or bundled into
   a material repair. If issue creation is unavailable, write a ready-to-file
   issue draft into the target repo's raw-material directory.
6. Run deterministic index checks when the repo has documentation indexes:
   `python3 <installed daily-docs-fold skill dir>/scripts/check-doc-indexes.py --repo <repo>`.
   Resolve the skill dir from the installed skill path; if unavailable, skip
   with the reason in the worklog. The checker also validates source handles in
   `docs/policy`, `docs/manuals`, any `--design-dir`, and extra `--source-dir`
   values.
7. Verify changed durable docs are routed from the root docs index or target
   index.
8. Delete, truncate, or retain ingested raw files only with ledger disposition.
   Leave `ephemeral/projects/<project>/` alone while the project is active.
   Treat a project directory as done when it has not changed for at least seven
   days, using git history when available and filesystem mtimes otherwise, or
   when the project root contains a clear done sentinel such as `done.md`. Once
   a project directory is done, fold its material into durable docs, record the
   ledger disposition, and delete that project directory. Never delete an active
   worklog.
9. Re-render the current project list in `AGENTS.md` or the repo's equivalent
   agent policy file after project disposition. Keep the list tight: one row per
   active project with the project path, project name, and a five-to-eight-word
   description. Remove projects that were folded and deleted.
10. Run proof gates from the target repo's policy. If none are documented, use
   the smallest relevant deterministic checks for the changed docs, scripts, and
   metadata.
11. For docs-only fold PRs, keep the PR on the merge path when targeted proof
   for the changed docs passes. Inherited repo-wide checker debt should be
   recorded in the worklog and PR body, not used as a reason to stop at PR
   creation or leave auto-merge disabled. If the PR becomes dirty or behind,
   rebase/update it and rerun impacted proof unless the fold itself caused the
   failure.
12. Update the journal, changelog, or fold marker with summary, proof, and
   remaining debt.

## Definition Of Done

- Every input row has a disposition.
- Durable knowledge is in the correct target, not left only in raw material.
- The root docs index or equivalent routes to changed documentation areas.
- The target repo's agent policy has a current project list when project
  scratch directories exist.
- Stale or misleading docs were edited, clipped, or marked as drift.
- Build Up, Compress, and Rebalance reached "only nits" or accepted remaining
  debt is recorded.
- Consumed raw material is deleted, truncated, or retained with a reason.
- Done project directories are folded and deleted; active project directories
  are left alone.
- Required proof gates pass or the exact blocker is recorded.
- Docs-only fold PRs are updated, auto-merge-ready, or merged unless the fold
  itself caused the blocking failure.
- Material shared-skill problems were filed upstream, drafted for filing, or
  explicitly deferred with a reason.

## References

- [phases.md](references/phases.md) for the phase loop and review questions.
- [review-prompt.md](references/review-prompt.md) for a reusable reviewer
  prompt.
- [scripts/check-doc-indexes.py](scripts/check-doc-indexes.py) for generic
  docs-index routing lint plus optional design-index/probe checks when invoked
  with `--design-dir` and target-repo probe or second-layer flags.

source: agents
