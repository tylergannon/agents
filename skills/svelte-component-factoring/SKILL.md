---
name: svelte-component-factoring
description: >
  Apply ownership-first Svelte component factoring rules. Use when creating,
  splitting, composing, or reviewing Svelte components; moving remote
  functions/state between route, page, and component layers; or judging whether
  a Svelte file is overstuffed.
---

# Svelte Component Factoring

Use this skill when the real question is not "can this markup be extracted?"
but "who owns this behavior, state, remote result, and proof?"

## Required Sources

- Read the current component, route, tests, and nearby patterns before proposing
  a boundary.
- Use the official Svelte/SvelteKit docs for APIs touched by the change,
  especially context, `$derived`, `$effect`, remote functions, async
  boundaries, and testing.
- Read any accepted design artifact or ADR for the affected route before
  moving ownership.

## Factoring Standard

Prefer components that own one product job end to end:

- they import the remotes whose results, pending state, errors, and retry
  behavior they own;
- they keep local UI state local;
- they expose factual events or callback props for facts peers may react to;
- they do not import peer component internals, call peer refresh functions, or
  carry controller-style method bundles through unrelated layers.

Split a component when it has multiple independent reasons to change, such as
uploading files, observing async status, rendering a catalog, owning modal
state, or styling a repeated row. Do not split only to move markup while
leaving state and remotes centralized in a route controller.

Keep route pages thin. A route page may compose siblings, pass route data, set
route-scoped context, and own route-only feature flags. It should not own
business workflows, pagers, modal action state, or remote refresh logic unless
those concerns are truly route-wide.

## State And Remote Placement

Place state where the user-visible responsibility lives:

- If one leaf component is the only consumer of a remote result, import the
  remote in that component and pass route params or ids down to it.
- If several siblings render the same result, load it at their nearest common
  owner and pass plain data down.
- If a mutation changes a specific list or row, keep the mutation refresh next
  to the component that owns that list or row.
- If state is derived from other state, use `$derived`, not `$effect`.
- Use `$effect` for browser side effects: polling, subscriptions, timers,
  query refresh in response to a consumed event, analytics, DOM work, or
  third-party integration. Effects that start work need explicit termination or
  cleanup.

## Context And Events

Use Svelte context sparingly and only for route-scoped cross-cutting services
that would otherwise be prop-drilled through passive layers. Context values
must stay narrow:

- allowed: typed event channels, current route-scoped capability objects,
  small immutable config, or bounded reactive state owned by the provider;
- forbidden: broad route data bags, remote query results, row arrays, modal
  state, or generic command bundles.

Events must be factual, not imperative. Publish facts like
`record-saved(id)` or `job-terminal(id)`, not commands like
`refreshList()` or `clearCard()`. Consumers decide whether the fact matters to
their own state.

Prefer Svelte 5 callback props for direct parent-child component events. Use
route-scoped context for known Svelte descendants that publish/consume through
nested leaves. Avoid document-level `CustomEvent` buses unless the event must
cross out of the Svelte route subtree.

## Review Checklist

Before implementing or approving a factoring:

- Name each new component's product job in one sentence.
- Name every remote function and which component owns it.
- Name every state bucket and why it lives there.
- Identify any cross-component event or context value and prove it is a fact,
  not a command or broad shared controller.
- Verify no sibling imports another sibling's internals.
- Verify async siblings sit under their own `<svelte:boundary>` when one slow
  remote should not blank the page.
- State which existing behaviors are invariants and which tests cover them.

source: agents
