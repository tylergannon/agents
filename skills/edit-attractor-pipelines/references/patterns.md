# Attractor Pipeline Patterns

Use current repo specs over this reference when they differ.

## Core Shape

Every graph needs one start and one exit:

```dot
start [shape=Mdiamond]
exit  [shape=Msquare]
```

Set a concrete graph goal. Each work node should represent a phase of intent,
not a low-level implementation step.

## Routing

Attractor routing is ordered roughly as:

1. explicit edge `condition=`;
2. `preferred_label` matching an edge label;
3. suggested next IDs;
4. highest edge weight;
5. lexical tie-break.

Use diamond nodes as gates that preserve the upstream outcome and label.

## Parallel And Consensus

`shape=component` fans out to unconditional outgoing branches. `shape=tripleoctagon`
fans results back in.

Join policies count execution success, not semantic agreement. If reviewers say
`APPROVED` or `REJECTED`, use a fan-in LLM node to read those results and make
the semantic decision.

Useful examples:

- `examples/parallel/01-fan-out-fan-in.dot`
- `examples/parallel/03-k-of-n.dot`
- `examples/code_review.dot`

## Goal Gates

Use `goal_gate=true` on the node that must succeed before exit. Give it a
`retry_target` when a failed gate should loop back to a specific stage.

Useful example: `examples/goal-gate.dot`.

## Steering

Manager/supervisor graphs use a `house` node with `stack.child_dotfile` and
`manager.actions`. They supervise another pipeline between observation cycles.

Navigator steering watches an in-session agent loop for repeated tool calls,
contradictions, or circular progress. In official mode, require a backend that
supports active steering.

Useful examples:

- `examples/steering/01-manager-supervisor.dot`
- `examples/steering/03-navigator.dot`

## Fidelity

Choose fidelity deliberately. Use compact or summaries for long pipelines; use
full thread reuse only when the next node truly needs the prior interaction.

Useful examples: `examples/fidelity/`.
