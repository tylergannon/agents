# Adversarial Review — Issue #3 review-scope changes, round 03

## Review target

The complete current proposal in
`/tmp/agents-issue-3-proposal.W0zZzW`, based on `origin/main` at
`7ca336082cd6031e0f85028ac9812cca317bfddf`, implementing GitHub issue #3,
"Make review skills explicitly forbid narrowing adversarial-review scope."
The target includes the combined committed and uncommitted changes to the three
review skills, the authoritative issue-scope proposal, the worklog and earlier
review rounds, and both focused behavior probes.

## Evidence inspected

- GitHub issue #3 in full: incident, requested rules 1–8, examples, and all four
  acceptance criteria.
- Repository instructions in `AGENTS.md`, including concise portable skill
  guidance, focused proof requirements, and metadata synchronization.
- `ephemeral/issue-3-review-scope-proposal.md`, including the additional
  reviewer-side refusal requirement and the hybrid continue/invalid behavior.
- The complete current contents, frontmatter, pre-change contents, and full
  `origin/main` diff for `skills/adversarial-review/SKILL.md`,
  `skills/request-adversarial-review/SKILL.md`, and
  `skills/consensus/SKILL.md`.
- Routing and package surfaces: `skills/request-adversarial-review/agents/openai.yaml`,
  `README.md`, `skills/README.md`, `package.json`, `.codex-plugin/plugin.json`,
  `.claude-plugin/plugin.json`, and `.claude-plugin/marketplace.json`.
- `ephemeral/worklog/202607171140-review-scope-narrowing.md`, the round-one and
  round-two reviews, the narrowed-but-recoverable behavior probe, and the
  unrecoverable `invalid review request` behavior probe.
- Current-state checks: remote `main` and local `origin/main` both resolve to
  `7ca336082cd6031e0f85028ac9812cca317bfddf`; `git diff --check origin/main`
  reports no errors.

## Findings

No material findings or genuine nitpicks remain.

The request skill explicitly separates implementation scope from review scope,
preserves concise conversation-only authoritative context, prohibits positive
and negative subject-matter limits plus expected-answer framing, and gives
separate concise bad/good examples
(`skills/request-adversarial-review/SKILL.md:15-17`, `:22-46`). The consensus
skill prohibits prior-finding and fix summaries, subject-matter limits, expected
conclusions, and desired verdicts in later rounds; its example and whole-target
rule require a fresh broad inspection without erasing reviewer session context
(`skills/consensus/SKILL.md:14-46`).

The reviewer skill independently rejects caller narrowing while retaining valid
target and operating constraints, makes gaps caused by authoritative
implementation exclusions reviewable, treats severity as an output filter
rather than a subject-matter limit, and reserves `invalid review request` for
cases where authoritative sources cannot reconstruct the goal
(`skills/adversarial-review/SKILL.md:18-34`). The round-two neutral-verdict
mismatch is resolved consistently: initial and repeat callers prohibit desired
verdicts, while the reviewer rejects requests for a particular verdict
(`skills/request-adversarial-review/SKILL.md:30-32`,
`skills/consensus/SKILL.md:18-20`, `skills/adversarial-review/SKILL.md:25-30`).

Focused proof now covers both new reviewer behaviors. The first probe records
and ignores positive file limits, negative limits, and a predicted conclusion,
then continues broadly from sufficient authoritative sources
(`ephemeral/reviews/202607171205-issue-3-narrowing-behavior-probe.md:9-33`). The
second probe rejects a requested verdict and unrecoverable narrowing, produces
no review findings, and returns exactly `invalid review request`
(`ephemeral/reviews/202607171220-issue-3-invalid-request-behavior-probe.md:3-18`).

## Outcome

`no findings`
