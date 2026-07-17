# Adversarial Review — Issue #1, Round 01

## Review target

Working tree of branch `codex/issue-1-agent-argument-order` (worktree
`.worktrees/agents-issue-1-agent-argument-order`), reviewed against GitHub
issue #1 ("Clarify review skills: agent no longer accepts a model positional")
and repository instructions (`AGENTS.md`, `CLAUDE.md`).

Changed files:

- `skills/consensus/SKILL.md` (modified)
- `skills/request-adversarial-review/SKILL.md` (modified)
- `ephemeral/worklog/202607171049-agent-workdir-positional.md` (new)

## Evidence inspected

- Issue #1 body via `gh issue view 1` (no comments).
- Full `git diff` of the two modified skill files; full current contents of
  both files.
- `skills/adversarial-review/SKILL.md`, `skills/agent-protocol/SKILL.md`,
  `skills/README.md`, `README.md`, `AGENTS.md` for surrounding context and
  neighboring examples.
- Independent repository-wide audit for stale model/agent positional forms:
  `grep -rniE 'agent (fable|sonnet|opus|haiku|claude[^-]|codex|gpt)'` across
  the tracked tree (including `submodules/`, excluding `ephemeral/` history).
  Sole hit: `README.md:40` `--agent codex claude-code`, which is a flag to the
  `skills` installer CLI, not an `agent` CLI positional — not a stale form.
- Worklog claims in `ephemeral/worklog/202607171049-agent-workdir-positional.md`
  cross-checked against the audit above; they hold.

## Requirements check

Issue #1 requested three things:

1. **State explicitly in both skills that agent selection is automatic and
   model/agent positionals are forbidden.** Done. Both
   `skills/request-adversarial-review/SKILL.md:15-16` and
   `skills/consensus/SKILL.md:11-12` add: "The `agent` CLI selects the reviewer
   automatically. Never place a model or agent name before `<workdir>`; the
   first positional argument is always the workdir." The verbatim duplication
   in both files is what the issue asked for, and is appropriate since each
   skill can be loaded independently.
2. **Use only `agent <workdir> <prompt>` and
   `agent <workdir> --session <id> <prompt>` examples.** Done.
   `request-adversarial-review` shows only `agent <workdir> "..."` (line 27);
   `consensus` shows only `agent <workdir> --session <id> "..."` (line 21). No
   other `agent` CLI invocation examples exist in either file or elsewhere in
   the published skills.
3. **Audit neighboring skill/docs examples for obsolete positional forms.**
   Done and independently re-verified (see Evidence). No `agent fable`,
   `agent sonnet`, or similar positional forms exist in the tracked tree.

Repository instructions: the change is concise, portable, and free of
product-specific policy (`AGENTS.md`); a session worklog was maintained per
`agent-protocol`. No unrequested behavior, tooling, or infrastructure was
added.

## Findings

No material findings. One genuine nitpick:

1. **[nitpick] Prohibition wording covers only the "before `<workdir>`" error
   mode.** `skills/request-adversarial-review/SKILL.md:15-16` and
   `skills/consensus/SKILL.md:11-12` say "Never place a model or agent name
   before `<workdir>`". A caller misled by stale external examples could
   instead insert a model name *after* the workdir (e.g.
   `agent <workdir> fable "<prompt>"`), where it would be silently absorbed
   into the prompt text rather than failing loudly. Impact is low — the
   preceding sentence ("selects the reviewer automatically") plus the
   examples-only-two-forms rule already implies no model argument belongs
   anywhere — but a phrasing like "never pass a model or agent name at all"
   would close the gap completely.

## Outcome

only nitpicks remain
