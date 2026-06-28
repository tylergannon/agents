# Daily Docs Fold Phases

## Build Up

Goal: no material input remains only in raw logs.

Actions:

- inventory landed commits, PRs, worklogs, reviews, research notes, and
  `ephemeral/` files,
- assign one canonical row per merged work item,
- route every durable decision, correction, rule discovery, source lesson, and
  proof rule to the right target,
- extract `skill_use`, `skill_issue`, and `skill_fix_request` telemetry and
  identify the upstream skill source when present,
- add missing routes in the target repo's docs root, policy index, manuals
  index, skills index, source index, or equivalent routing surface.

Review asks:

- What material input is missing?
- Which source row is duplicated?
- Which durable claim lacks a stable source handle?
- Which user correction or rule discovery was not preserved?
- Which shared-skill failure should become an upstream issue?

## Compress

Goal: docs stay compact enough for future agents to use.

Actions:

- delete repeated rules and duplicated summaries,
- replace raw chronology with stable policy, skill, or manual guidance,
- remove stale caveats after they have become policy,
- split swollen leaves only when it improves retrieval.

Review asks:

- Which section says the same thing twice?
- Which prompt or skill has become pedantic or overfit?
- Which details belong in a source artifact instead of durable prose?
- Which old wording will mislead the next agent?

## Rebalance

Goal: every fact lives at the right layer.

Actions:

- move mandatory behavior to the repo's agent policy or policy docs,
- move preferred techniques to manuals or skill references,
- move executable workflows to skills,
- move material shared-skill correction requests to the source repo as issues,
  existing-issue comments, or issue drafts,
- move upstream inventory and borrow candidates to source indexes,
- keep chronological notes and ingestion markers in the repo journal or
  changelog,
- keep raw evidence in the raw-material directory only with a reason.

Review asks:

- Which content is in the wrong directory?
- Which policy is hiding in a manual, journal, or review artifact?
- Which manual reads like mandatory policy?
- Which route from the root docs index fails the two-read lookup path?
