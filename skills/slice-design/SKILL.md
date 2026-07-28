---
name: slice-design
description: >
  Turn an accepted nlspec plus the current code into vertical work slices,
  each with an ephemeral technical spec and an implementer handoff contract.
  Use when planning implementation work governed by an nlspec.
---

# Slice Design

Load the spec-writing standard first. The inputs are the nlspec and the
current code; the outputs are vertical slices, each carrying an ephemeral
slice spec. The slice spec is authoritative for its slice only, then archived
in the PR — never maintained. It may contain nothing with cross-slice
lifetime (the lifetime rule).

## Slicing

Slice vertically: each slice crosses the stack thinly and yields something
exercisable early — curl-able, browsable, claim-demonstrable. Reject
horizontal stack-order plans (migrations → services → API → frontend); that
is the default model tic, and under it nothing is touchable until the end.
Sequence slices so each demonstrates at least one claim or de-risks a seam,
and so no slice cuts across a seam mid-contract.

## The slice spec

Derive it fresh from the nlspec and the code as it is today, conditioned on
what already exists. Contents, using the artifacts from the notation standard:

- call-stack tree for the control-flow change (diff syntax);
- file-tree diff, respecting cell boundaries — the countable preview of "is
  this change about to spray across seams";
- types and signatures for key new functions;
- the existing functions and types to modify;
- which claims this slice demonstrates, and how.

Content that is derivable by the implementer in context is omitted; content a
future slice must agree with is promoted to the nlspec before this slice
ships, via HITL.

## Implementer handoff

Every slice ships with the handoff contract:

- Contracts are inviolable. If the work forces a contract re-evaluation, stop
  and report; do not amend the nlspec.
- An underivable decision, or any flagged deferred decision, escalates
  immediately as a spec defect. Improvisation at a seam is a defect in itself.
- A budget breach stops work; the spec decides the split.
- Claims are demonstrated per /proof-of-work, not asserted.

## Exit valves

Before a slice's PR merges, verify:

1. **Promotion** — nothing with cross-slice lifetime remains in the slice
   spec; promoted decisions are in the nlspec.
2. **Conversion** — new tests worth keeping as behavior are run through the
   catalytic converter (see cell-lifecycle) as claim candidates.
3. **Archive and budgets** — the slice spec is archived in the PR, and every
   touched cell's budget is re-checked mechanically.
