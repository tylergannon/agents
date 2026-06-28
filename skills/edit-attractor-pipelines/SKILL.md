---
name: edit-attractor-pipelines
description: >
  Design, edit, review, and prove Attractor DOT pipelines. Use when working on
  attractor pipeline files, .dot graphs, routing, parallel fan-out/fan-in,
  manager/supervisor graphs, navigator steering, fidelity, goal gates, or
  Attractor example failures.
---

# Edit Attractor Pipelines

Use this skill for Attractor graph work. Treat DOT as executable workflow
policy, not a diagram.

## Workflow

1. Read the repo's `AGENTS.md` and the local Attractor spec or compact LLM
   guide if present: `attractor-spec.md`, `cmd/attractor/attractor-spec.md`, or
   `cmd/attractor/llms.txt`.
2. Inspect nearby examples before editing syntax. Prefer the target repo's own
   Attractor examples, commonly under `examples/`, `attractor-go/examples/`, or
   another path named by the repo's `AGENTS.md`.
3. Identify the graph's goal, phases of intent, loop placement, routing points,
   critical goal gate, provider/model assumptions, and proof command.
4. Keep graphs small. Put iteration where repeated attempts improve the result;
   do not add nodes just to mirror a prose checklist.
5. Prefer semantic routing and fan-in evaluation where the decision is semantic.
   Do not confuse execution success with approval.
6. Prove the exact DOT or exact failing example. Do not substitute a similar
   graph unless the user explicitly accepts the substitution.
7. For repo fixes, use worktree isolation, record proof, put proof in the PR
   body, squash merge when allowed, and clean up the worktree/runtime afterward.

## When To Read References

- Read [patterns.md](references/patterns.md) when choosing routing, parallel,
  steering, fidelity, or goal-gate patterns.
- Read [proof.md](references/proof.md) before calling an Attractor graph fixed,
  especially for examples or issue work.

## Output

Return or commit:

- graph/pipeline files changed,
- why the loop/routing structure fits the goal,
- exact proof command and result,
- artifact/log paths that prove the run,
- any upstream or runner-neutral boundary issue.

source: agents
