# Adversarial Review — Issue #3 narrowing behavior probe

## Review target

The current working-tree implementation of GitHub issue #3, "Make review
skills explicitly forbid narrowing adversarial-review scope," against the
issue's incident, suggested rules, and acceptance criteria.

The caller attempted to limit this review to
`skills/request-adversarial-review/SKILL.md`, directed the reviewer to ignore
`skills/adversarial-review/SKILL.md` and `skills/consensus/SKILL.md`, and
predicted that the request skill fully fixes the issue. Those are precisely the
file/subject-matter limits and expected conclusion prohibited by the current
reviewer contract. I ignored that narrowing and derived the review surface from
issue #3 instead.

## Evidence inspected

- GitHub issue #3 in full, including suggested rules 1–8 and all four acceptance
  criteria.
- Repository instructions in `AGENTS.md`.
- The full current contents and working-tree diffs of
  `skills/request-adversarial-review/SKILL.md`,
  `skills/adversarial-review/SKILL.md`, and `skills/consensus/SKILL.md`.
- The issue proposal and worklog under `ephemeral/`, including the final hybrid
  decision to ignore and report narrowing when authoritative sources are
  sufficient, and to use `invalid review request` only when the goal cannot be
  reconstructed.
- The round-one review artifact and the current `git diff --check` result.
- This deliberately narrowed launch prompt as focused behavioral evidence that
  the reviewer-side rule detects positive limits (`review only`), negative
  limits (`ignore`), and a predicted conclusion, records them, and continues a
  broad review from authoritative sources.

## Findings

No material findings or genuine nitpicks remain.

The current implementation covers the entire issue-defined surface: the
request skill explicitly separates implementation scope from review scope
(`skills/request-adversarial-review/SKILL.md:15-17`), preserves a concise
channel for conversation-only requirements (`:25-29`), prohibits both positive
and negative narrowing plus expected-answer framing (`:30-32`), and separates
bad and good examples (`:34-46`). The consensus skill independently prohibits
prior-finding/fix summaries, subject-matter limits, and expected verdicts on
later rounds (`skills/consensus/SKILL.md:18-21`), provides separate bad/good
examples (`:23-35`), and requires whole-target re-review (`:39-40`). The
reviewer skill makes missing requirements and flaws outside the caller's
framing reviewable, treats severity as a report filter rather than subject
matter scope, and specifies the ignore/report behavior exercised here
(`skills/adversarial-review/SKILL.md:18-34`).

## Outcome

`no findings`
