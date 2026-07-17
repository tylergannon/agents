# Adversarial Review — Issue #3 review-scope changes, round 02

## Review target

The complete current proposal in
`/tmp/agents-issue-3-proposal.W0zZzW`, based on `origin/main` at
`7ca336082cd6031e0f85028ac9812cca317bfddf`, for GitHub issue #3,
"Make review skills explicitly forbid narrowing adversarial-review scope."
The target includes the combined committed and uncommitted changes to all three
review skills, the issue-scope proposal and worklog, the round-one review, and
the narrowing behavior probe that appeared during this review.

## Evidence inspected

- GitHub issue #3 in full: incident, existing guidance, requested rules 1–8,
  examples, and all four acceptance criteria.
- Repository instructions in `AGENTS.md`, especially the requirements for
  concise portable skills, focused behavioral proof, and synchronized metadata.
- `ephemeral/issue-3-review-scope-proposal.md`, including the user's additional
  reviewer-side refusal requirement and the hybrid continue/invalid behavior.
- The full `origin/main` diff and complete current contents of
  `skills/adversarial-review/SKILL.md`,
  `skills/request-adversarial-review/SKILL.md`, and
  `skills/consensus/SKILL.md`, plus their pre-change contents.
- Routing and package surfaces: the three skills' frontmatter,
  `skills/request-adversarial-review/agents/openai.yaml`, `README.md`,
  `skills/README.md`, `package.json`, `.codex-plugin/plugin.json`,
  `.claude-plugin/plugin.json`, and `.claude-plugin/marketplace.json`.
- `ephemeral/worklog/202607171140-review-scope-narrowing.md`, the round-one
  review, and
  `ephemeral/reviews/202607171205-issue-3-narrowing-behavior-probe.md`.
- Current-state checks: remote `main` and local `origin/main` both resolve to
  `7ca336082cd6031e0f85028ac9812cca317bfddf`; `git diff --check origin/main`
  reports no errors.

## Findings

### 1. [issue] Neutral verdict requests are incorrectly treated as narrowing

Issue #3 prohibits an expected conclusion or a **desired** verdict, not a
neutral instruction to return an independently reached result. The current
reviewer rule instead refuses any caller instruction that "requests a verdict"
(`skills/adversarial-review/SKILL.md:25-30`), and the initial-request skill tells
the caller to remove all "conclusions, or verdicts" without qualifying them as
expected or desired (`skills/request-adversarial-review/SKILL.md:30-32`). This
also conflicts with the more precise consensus wording, which prohibits
"expected conclusions, or desired verdicts"
(`skills/consensus/SKILL.md:18-20`).

Concrete reproduction: a caller can supply only authoritative sources and say,
"Review the implementation independently and report findings and a verdict."
That prompt predicts no finding and requests no particular outcome, but the
reviewer contract requires treating it as rejected narrowing and recording it
in the artifact. This creates a false-positive accusation during normal use and
makes the three related skills disagree about what language is forbidden. The
behavior probe exercises a predicted conclusion, not this neutral-verdict case
(`ephemeral/reviews/202607171205-issue-3-narrowing-behavior-probe.md:9-15`), so
it does not resolve the mismatch.

The smallest repair is to use the consensus skill's distinction everywhere:
remove **expected conclusions or desired verdicts** in the caller skill, and
refuse prompts that predict a conclusion or request a **particular** verdict in
the reviewer skill.

### 2. [nitpick] The focused probe covers only the continue branch

The new probe demonstrates detection, recording, and broad continuation when a
narrowed prompt still names sufficient authoritative sources
(`ephemeral/reviews/202607171205-issue-3-narrowing-behavior-probe.md:19-33`). It
does not exercise the other newly introduced path: returning
`invalid review request` when the goal cannot be reconstructed, followed by
discard/relaunch handling in the caller skills
(`skills/adversarial-review/SKILL.md:28-31`,
`skills/request-adversarial-review/SKILL.md:48-52`,
`skills/consensus/SKILL.md:44-46`). The prose is internally consistent, so this
is an unproved-risk note rather than evidence of a defect; a second deliberately
narrowed probe with no recoverable authoritative goal would close it.

## Outcome

`material findings remain`
