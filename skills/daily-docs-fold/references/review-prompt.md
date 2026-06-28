# Daily Docs Fold Review Prompt

Use this shape for an independent review. Replace bracketed text.

```text
Here is my goal:

[GOAL]

Please review this documentation fold at the level of the repository operating
system, not just prose style.

Focus on demonstrable issues:

- material inputs from commits, PRs, worklogs, reviews, research, or ephemeral
  files that were not represented,
- duplicate input rows or duplicated durable rules,
- mandatory behavior placed outside AGENTS.md or docs/policy,
- manuals or skills that overfit one incident instead of capturing a durable
  repo pattern,
- stale docs, misleading routes, missing routes, or wrong-directory content,
- durable claims without stable source handles,
- prompt or skill wording that is overly narrow, pedantic, or model-limiting.
- material shared-skill failures that should be filed upstream, added to an
  existing upstream issue, or drafted instead of patched only in the target repo.

Write the exact prompt you received and your findings to:

[EPHEMERAL_REVIEW_PATH]

Label findings as critical, bug, design, or nit. Return findings first. For
each material finding include file/path, why it matters, and the smallest
repair. Then say whether another Build Up, Compress, or Rebalance pass is
needed. If only nits remain, say that explicitly.
```
