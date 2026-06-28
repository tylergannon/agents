---
name: tech-debt-bounty
description: >
  Run a bounded tech-debt hunt that turns one material scout proposal into a
  crisp refactoring contract through consensus, delegates or implements the
  bounded fix, proves it, and avoids ceremonial cleanup. Use for scheduled or
  explicit architecture/refactoring sweeps, antipattern hunts, overstuffed
  component/module reviews, and agent-context-budget work.
---

# Tech-Debt Bounty

Use this skill for bounded, high-signal maintainability work. The job is not to
find any possible cleanup. The job is to turn one material scout proposal into a
clear responsibility contract, fix that contract safely, and prove the fix. If
no proposal clears the materiality bar, stop without making a cosmetic PR.

## Operating Rule

Do not optimize for file count, line count, or visible churn. Optimize for:

- obvious ownership: a future bug fix has a clear home;
- local change surface: routine work usually touches one to three primary
  files;
- context fit: the relevant implementation, tests, and proof fit in one agent
  context window most of the time;
- coherent metaphors: every component, module, and reactive model has one named
  product job.

## Setup

Start from a fresh task worktree on current `origin/main`. Read the repo's
agent instructions, coding guide, and design docs only as needed. Use
framework-specific factoring skills or docs when the candidate touches a
framework boundary.

## Roles

This skill is an orchestrator. It should use smaller skills and agents instead
of doing every phase in one context:

- `debt-finder`: read-only scout that returns one compact candidate proposal or
  a no-fix outcome.
- `claude-codex-consensus`: independent design review that sharpens the scout
  proposal into a refactoring contract.
- worker agent: bounded implementation against the accepted contract and
  invariants.
- reviewer or consensus pass: attestation that implementation satisfies the
  contract while preserving behavior.

Do not skip straight from scout proposal to implementation when the change is
architectural, state-heavy, or crosses component/module boundaries.

## Materiality Bar

A candidate is material when at least two are true:

- a routine feature or bug fix requires reading or editing unrelated concerns;
- one file owns multiple state machines or independent workflows;
- behavior is coordinated by a broad controller, command bundle, generic data
  bag, or peer internals;
- effects synchronize state that should be derived, or side effects lack clear
  cleanup and ownership;
- tests or proof require excessive setup because ownership is unclear;
- duplicate domain logic causes drift risk across product paths;
- existing agents or developers are likely to misplace the next change.

Do not count formatting-only churn, dependency bumps, renames without ownership
improvement, splitting solely because a file is long, or extracting helpers that
leave the same controller in charge.

## Orchestration Loop

1. Run or accept a `debt-finder` scout result.
2. Judge materiality. If the candidate is weak, stale, or cosmetic, stop with a
   no-fix report.
3. Run `claude-codex-consensus` on the selected proposal. The consensus target
   is a responsibility contract, not a PRD and not a task recipe.
4. Stop if consensus shows the proposal is not material, the boundary is
   incoherent, or the proof surface cannot protect the change in budget.
5. Hand a worker only the accepted contract and invariants.
6. Review the diff against the contract.
7. Prove behavior with `proof-of-work` for application changes, or the target
   repo's documented checks for skills/docs/tooling-only changes.

## Consensus Contract Shape

```md
Decision: <one-sentence target boundary>
Materiality: <why this makes future work materially easier>
Component/module responsibilities:
- <owner>: <one product job>
State ownership: <where state lives and why>
Remote/mutation ownership: <where remote functions/forms/mutations live>
Events/context: <factual events, narrow context service, or none>
Invariants: <behavior, authorization, semantics, and tests that must not change>
Non-goals: <what would be gratuitous or out of scope>
Proof surface: <specific component, E2E, typecheck, or reviewer evidence>
```

## Proof And Closeout

Use `proof-of-work` for application behavior changes. For docs, skills, scripts,
or workflow-only changes, use the target repo's documented checks.

For every completed bounty run, produce either a merged or merge-ready PR with
proof and reviewer attestation where needed, or a no-fix report that names
inspected surfaces and explains why no candidate cleared materiality.

source: agents
