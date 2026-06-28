---
name: debt-finder
description: >
  Scout a repository for one material maintainability or design-debt candidate
  and produce a concise proposal artifact. Use for read-only antipattern
  discovery, factoring scouting, ownership review, complexity scouting, missing
  harness proposals, and tech-debt bounty candidate selection.
---

# Debt Finder

Use this skill as a scout. The output is a compact proposal, not an
implementation plan and not a PRD. Do not edit product code. Do not start a
refactor. Do not create a branch or PR unless explicitly asked outside this
skill.

Do not write files by default. Return the proposal in the final response. Write
exactly one proposal artifact only when explicitly asked, when handing the
proposal to consensus or implementation, or when preserving a no-fix result
would materially help a later run.

## Read First

Read only what is needed to scout cheaply:

1. the repo's `AGENTS.md` or equivalent;
2. any framework-specific factoring skill when component boundaries are in
   scope;
3. architecture/design docs only when ownership terms are unclear.

Use current framework docs only when needed to judge a candidate. Do not use
docs lookup as ceremony for non-framework scouting.

## Search Budget

Favor fast structural scans over deep reading:

- token/size hotspots;
- broad imports and route/controller-shaped files;
- components with many state buckets, effects, remotes, or modal/action
  workflows;
- reactive/model files that look like hidden page controllers;
- repeated domain logic across product paths;
- tests that reveal painful setup or missing isolated harnesses;
- recent worklogs or design docs that mention confusion, repeated corrections,
  or ownership drift.

Read deeply only for the strongest one to three candidates. Stop once one
candidate has enough evidence for a concrete proposal.

## Materiality Bar

A candidate is material when it plausibly makes future changes harder in a
specific way. Prefer candidates where at least two are true:

- a routine bug fix or feature would require reading unrelated concerns;
- one file owns multiple independent state machines, workflows, or proof
  surfaces;
- behavior is coordinated by a broad controller, command bundle, generic data
  bag, or peer component internals;
- effects synchronize state that should be derived, or side effects lack clear
  ownership;
- duplicate domain logic creates drift risk;
- a missing harness makes a user-visible component risky or expensive to
  change;
- the next agent/developer is likely to put a change in the wrong place.

Do not select a candidate for size alone. Token and line counts are smoke, not
proof.

## Proposal Output

Keep the output short and contract-shaped:

```md
# Debt Finder Proposal

## Candidate
<file/module/component and one-sentence design failure>

## Materiality
- <specific future change made harder>
- <specific mixed ownership/state/proof problem>

## Evidence
- `<path>:<line>` - <what this shows>
- `<path>:<line>` - <what this shows>

## Boundary Contract
- `<new or existing owner>`: <responsibility>
- State ownership: <short contract>
- Remote/mutation ownership: <short contract or "none">
- Events/context, if any: <short contract or "none">

## Non-Goals
- <what must not be changed>

## Proof Surface
- <tests, component harness, E2E, typecheck, reviewer attestation, or no-fix proof>

## Rejected Candidates
- <candidate>: <why weaker>
```

If no candidate clears the bar, report inspected surfaces, best rejected
candidates, and why none cleared materiality.

source: agents
