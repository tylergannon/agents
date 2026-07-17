# Adversarial Review — Issue #3 review-scope changes, round 01

## Review target

Working-tree changes on `codex/issue-3-review-scope-proposal` (base `origin/main`,
7ca3360) implementing GitHub issue #3 "Make review skills explicitly forbid
narrowing adversarial-review scope":

- `skills/adversarial-review/SKILL.md`
- `skills/consensus/SKILL.md`
- `skills/request-adversarial-review/SKILL.md`
- `ephemeral/worklog/202607171140-review-scope-narrowing.md` (untracked worklog)

## Evidence inspected

- GitHub issue #3 full body (incident, suggested rules 1–8, acceptance criteria).
- `git diff` of all three skill files plus their full current contents.
- Pre-change contents via diff context (old step 2/3 of
  `request-adversarial-review`, old step 3/5 of `consensus`, old broad-review
  paragraph of `adversarial-review`).
- Repository instructions (`AGENTS.md`), `README.md`, `skills/agent-protocol/SKILL.md`.
- Cross-reference grep for `adversarial`, `no findings`, `only nitpicks` across
  `skills/`, `.agents/skills/`, `docs/` — no other file references the review
  outcomes or the removed prompt text, so no stale cross-references.
- Validation evidence: the worklog (3 decision/skill_issue lines). No other proof
  artifacts exist for this change.

## Findings

### 1. [issue] Regression: conversation-only requirements lost their sanctioned channel

`skills/request-adversarial-review/SKILL.md` step 2 previously read: "If a
requirement exists only in conversation, state it concisely without turning it
into a checklist or directing attention to suspected defects." The new step 2
(lines 25–28) drops this provision entirely, and new step 3 (lines 29–30)
instructs the caller to "Remove any sentence that tells the reviewer what not to
examine or what it **should find**" — which a caller can reasonably read as
requiring removal of a concisely stated conversational requirement, since a
requirement statement is exactly a statement of what the reviewer should verify.

Meanwhile `skills/adversarial-review/SKILL.md` lines 18–20 tell the reviewer to
"Derive that scope from the authoritative user request" — but a spawned reviewer
has no access to the caller's conversation. When the authoritative request exists
only in conversation (a common case for this skill set), the reviewer now has no
way to obtain the normative source, and the caller has no sanctioned way to
supply it. Issue #3's own suggested rule 4 says review scope is set "by
**identifying** the authoritative user request" — for conversation-only
requirements, identification requires restating. Impact: reviews of
conversation-specified work silently lose their success criteria, producing
exactly the unanchored/incomplete reviews the issue is trying to prevent.

### 2. [issue] Refusal trigger misses the issue's primary narrowing phrasings

Issue #3 rule 1 names the prohibited clauses: "`do not broaden into`, `ignore`,
`focus only on`, or `only review`." The new reviewer-side trigger
(`skills/adversarial-review/SKILL.md` lines 24–26) fires only when the launch
prompt "tells you what **not** to examine, what to find, or which verdict to
reach." Positive narrowing — "focus only on the migration files", "only review
the three artifacts below" — does not literally say what not to examine, states
no expected finding, and names no verdict, so the two most common narrowing
phrasings from the issue neither trip the invalid-request rule nor are otherwise
addressed on the reviewer side (the old blanket sentence "Do not limit the
review to files, patterns, suspected defects, or an intended answer mentioned by
the caller" was deleted in this same diff). Relatedly, acceptance criterion 3
("Both skills prohibit negative reviewer-scope clauses and expected-conclusion
framing") is only partially met in `skills/consensus/SKILL.md`: its prose (lines
18–19) prohibits summarizing findings/fixes/expected remaining problems, but
negative scope clauses and desired-verdict framing are prohibited only by
implication of the `Bad:` example, not stated. Impact: the narrowing pattern
that caused the original incident can recur through positive-only phrasing and
pass both skills' explicit rules.

### 3. [issue] Over-engineering: refuse-and-restart protocol replaces the requested review-broadly rule

Issue #3 requests explicit caller-side prohibitions, reviewer-side scope
clarifications, and examples (rules 1–8). The diff additionally introduces an
enforcement protocol nowhere requested: a fourth outcome `invalid review
request` (`skills/adversarial-review/SKILL.md` lines 24–28, 49–52), a mandate to
stop without a verdict, and discard/restart handling in both caller-side skills
(`skills/consensus/SKILL.md` lines 39–41, `skills/request-adversarial-review/SKILL.md`
lines 45–47). This replaces the prior, simpler behavior — review broadly and
ignore the narrowing — with a hard refusal that burns a full reviewer round and
session on any prompt the reviewer judges narrowed, including borderline cases
where target identification shades into scoping (e.g. "review the three changed
skills"). Combined with finding 2's under-specified trigger, the mechanism is
simultaneously easy to evade and prone to false-positive dead rounds. Concrete
failure: a mildly imperfect but honest prompt now yields no review at all, where
the old text yielded a full-surface review regardless.

### 4. [nitpick] Bad examples embedded inside the runnable prompt fence

In both `skills/consensus/SKILL.md` (lines 21–30) and
`skills/request-adversarial-review/SKILL.md` (lines 32–41), the `Bad:` example
was placed inside the same ```text fence as the good launch command. Previously
the fenced block contained exactly the template to copy; now a mechanically
copying agent can pick up the anti-example, and the "Bad:/Good:" labels are not
part of any runnable prompt. Separate fences (or prose for the bad example)
would preserve the copy-the-block invariant.

### 5. [nitpick] Thin validation evidence

The only proof artifact is a three-line worklog. `AGENTS.md` requires "focused
checks that prove the changed behavior"; for prompt-skill changes a minimal
check exists — e.g. a dry-run showing a previously narrowed prompt is now
rewritten/rejected, or at least a mapping of each issue acceptance criterion to
the satisfying lines. No such mapping or check is present.

## Outcome

`material findings remain`
