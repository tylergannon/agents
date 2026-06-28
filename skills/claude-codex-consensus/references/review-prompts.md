# Consensus Review Prompts

Use these as starting points. Trim them to the task, but do not shrink them into
local checks that prevent the reviewing agent from finding the real problem.
Before sending a prompt, expand `<LABEL_INSTRUCTIONS>` from the canonical labels
in [reconciliation.md](reconciliation.md). If `<LABEL_INSTRUCTIONS>` is still
present, the prompt is not ready to send.

## Complete Code Review

```md
Here is my goal:

<GOAL>

If you have a code review skill, use it. Otherwise review directly. Give a
complete review of the current design and implementation against that goal.

Focus on:

- Demonstrable bugs.
- Over-engineered components.
- Features or layers that were not expressly requested or directly implied by
  the work order.
- Demonstrable race conditions.
- Inelegant architecture decisions.
- Smelly architecture.
- Logic placed on the wrong side of component boundaries.
- Leaky components.
- Places where the project's metaphors or APIs are being misunderstood in ways
  that will clearly create maintainability or correctness problems.

This is not an invitation to push for gold plating. Prefer real findings over
style preferences.

Write the exact prompt you were given and your findings to:
<REVIEW_ARTIFACT_PATH>

<LABEL_INSTRUCTIONS>
```

## Design Review

```md
Here is my goal:

<GOAL>

Review this design before implementation. Look for correctness problems,
missing constraints, unnecessary layers, wrong ownership boundaries, leaky
abstractions, data-flow mistakes, failure modes, and proof gaps.

Write the exact prompt you were given and your findings to:
<REVIEW_ARTIFACT_PATH>

<LABEL_INSTRUCTIONS>
```

## Implementation Plan Review

```md
Here is my goal:

<GOAL>

Review this implementation plan before coding. Look for hidden dependencies,
wrong ownership boundaries, over-engineered components, unrequested features or
layers, missing tests, migration risk, proof gaps, and simpler implementation
paths that still satisfy the work order.

Write the exact prompt you were given and your findings to:
<REVIEW_ARTIFACT_PATH>

<LABEL_INSTRUCTIONS>
```

## Diff Review

```md
Here is my goal:

<GOAL>

Review the current diff against that goal. Prioritize demonstrable bugs,
regressions, missing proof, over-engineering, unrequested features or layers,
wrong component boundaries, leaky abstractions, and architecture decisions that
will likely create maintainability or correctness problems.

Write the exact prompt you were given and your findings to:
<REVIEW_ARTIFACT_PATH>

<LABEL_INSTRUCTIONS>

Use file/line references when possible.
```

## Proof Review

```md
Here is my goal:

<GOAL>

Review whether this proof actually demonstrates the behavior. Check that the
command ran against the right checkout, the right artifact, and the exact
failing or acceptance path. Flag static checks being presented as behavioral
proof, missing rendered artifacts, and proof that does not match the work order.

Write the exact prompt you were given and your findings to:
<REVIEW_ARTIFACT_PATH>

<LABEL_INSTRUCTIONS>
```
