---
name: claude-codex-consensus
description: >
  Get Claude Code and Codex to independently review designs, implementation
  plans, diffs, or proof strategy, then reconcile the disagreement. Use when
  the user asks for Claude/Codex consensus, a second model opinion, adversarial
  review, design convergence, cross-agent implementation review, or when the
  work is difficult, risky, architectural, broad, or otherwise non-trivial.
---

# Claude / Codex Consensus

Use this skill when a decision or implementation benefits from independent model
judgment. The point is not ceremony. The point is to surface disagreement before
implementation, during implementation, and before merging risky work.

Default to this skill when the human mentions consensus, Claude Code, second
opinion, design review, adversarial review, or when the agent recognizes the
work as difficult or non-trivial.

## Workflow

1. State the goal under review in the user's terms.
2. Gather the minimum evidence: relevant files, constraints, work order,
   proposed plan, current diff, failing proof, and repo conventions.
3. Choose a per-round review artifact path such as
   `ephemeral/reviews/<YYYYMMDDNN>-<task>.md` or the repo's review/worklog path.
4. Ask the other agent for a broad, goal-level review. Do not narrow the prompt
   to formatting, tests, or your preferred theory of the work.
5. Require the other agent to write the exact prompt it received and its
   findings to the review artifact.
6. Require findings labeled `critical`, `bug`, `design`, or `nit`.
7. Reconcile findings with the label rules in
   [reconciliation.md](references/reconciliation.md).
8. Revise the design, plan, or implementation, then prove the result where
   there is runnable behavior to prove.
9. Repeat review while the other agent finds `critical`, `bug`, or `design`
   issues and material progress is still happening. Escalate to the user after
   three rounds, no material diff, or the same unresolved finding repeating.
10. Report the final consensus and any unresolved dissent.

## Prompt Quality Rules

- Ask at the level of the actual goal, not at the level of a narrow local check.
- Give the called agent enough context to discover the right review frame.
- Do not prohibit useful exploration. Narrow only destructive actions.
- Do not ask for approval. Ask for real findings.
- Do not leak your intended answer unless the task is to critique that answer.

## Calling Another Agent

Use the `agent` CLI to contact the adversarial reviewer. Prefer the latest-model
aliases provided by the CLI rather than pinning stale model IDs. For Claude,
call the latest Claude Opus reviewer with:

```sh
agent opus <workdir> -f <prompt-file>
```

For Codex, call the latest flagship GPT reviewer with:

```sh
agent gpt <workdir> -f <prompt-file>
```

If the available CLI differs, run `agent --help` and choose the alias whose help
text says it targets the latest appropriate model family. Do not fall back to
older pinned model names when a latest alias exists.

Keep prompts compact and artifact-based, but leave the review open-ended enough
for the called agent to find architectural and boundary problems.

The request should not be read-only: the called agent must write the review
artifact. If you need to protect the worktree, constrain product-code edits, not
artifact creation.

For the full prompt shape, use
[review-prompts.md](references/review-prompts.md). If the callee has a code
review skill, ask it to use that skill; otherwise ask it to review directly.

When the orchestrator is Claude Code rather than Codex, use the same CLI pattern
in reverse: call `agent gpt` for the Codex pass. If the CLI is unavailable,
write a Codex-ready review prompt artifact under `ephemeral/reviews/` and ask
the user to run it; do not claim consensus until the independent result exists.

## When To Read References

- Read [review-prompts.md](references/review-prompts.md) for reusable prompts.
- Read [reconciliation.md](references/reconciliation.md) when two agents
  disagree or when the result needs to become an implementation plan.

## Closeout

Report:

- what each agent reviewed,
- top findings from each side,
- accepted changes,
- rejected or unresolved disagreements,
- proof that followed the consensus.

source: agents
