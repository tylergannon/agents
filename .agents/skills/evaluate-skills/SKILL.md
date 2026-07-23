---
name: evaluate-skills
description: >
  Evaluate agents skills for clarity, coherence, usability,
  actionability, compression, routing quality, and alignment with neighboring
  skills. Use when maintaining this repo's SKILL.md files, repo-local skills,
  distributed skills, plugin surfaces, or when the user asks whether skill
  instructions are word salad, overfit, unclear, duplicated, or inconsistent.
metadata:
  internal: true
---

# Evaluate Skills

Use this skill to review whether agent skills will actually help the next agent
do better work with less context.

The goal is not to make skills longer. The goal is to remove confusion, expose
missing contracts, and make related skills behave like one coherent system.

## Workflow

1. Identify the review scope: one `SKILL.md`, one skill directory, a related
   skill set, or a plugin/package surface.
2. Read the routing surfaces first: frontmatter `name` and `description`,
   skill index docs, plugin metadata, and any AGENTS/CLAUDE routing text.
3. Read the skill body and only the references needed to understand the review
   target. Do not bulk-load unrelated references.
4. Compare neighboring skills that might overlap, delegate to, or contradict
   the target skill.
5. Review against the rubric below and propose the smallest edits that would
   materially improve agent behavior.
6. Verify with the target repo's checks. If no deterministic check exists, use a
   realistic dry run, independent review, or explicit unproved-risk note.

## Rubric

- **Routing:** The description says when to use the skill and distinguishes it
  from neighboring skills using words users and agents will actually say.
- **Scope:** The skill owns one repeatable job. It does not absorb adjacent
  workflows that belong in another skill, script, policy file, or test.
- **Coherence:** The title, description, procedure, references, and closeout all
  point at the same job without changing terminology midstream.
- **Actionability:** The workflow tells an agent what to inspect, edit, verify,
  and report. It avoids vague advice that cannot change behavior.
- **Compression:** The skill assumes model competence, deletes obvious coaching,
  avoids defensive case piles, and moves rare detail into one-level references.
- **Layering:** Deterministic checks live in scripts/tests/schemas where
  practical. Judgment and routing guidance stay in prose.
- **Alignment:** Related skills hand off cleanly, share vocabulary where needed,
  and do not issue conflicting instructions.
- **Closeout:** The skill names evidence expected at completion: commands,
  artifacts, review findings, proof gaps, or remaining risks.

## Findings

Report findings first, ordered by severity:

- `critical`: likely to route agents into the wrong workflow or cause harmful
  edits.
- `bug`: unclear, contradictory, or missing instructions that can plausibly make
  normal use fail.
- `design`: skill works but is overbroad, undercompressed, poorly layered, or
  misaligned with neighboring skills.
- `nit`: wording or structure polish that does not affect normal use.

For each finding, cite the file and tight line range, explain the agent behavior
it would cause, and give a concrete repair.

## Closeout

Report:

- skills and metadata reviewed,
- related skills compared,
- edits made or recommended,
- checks, dry runs, or independent reviews used,
- remaining assumptions or unproved behavior.

source: agents
