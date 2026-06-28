---
name: e2e-coverage-triage
description: >
  Triage Playwright or BDD E2E scenarios into the right proof layer. Use when
  revising E2E scenarios, shared authenticated sessions, seeded records,
  Storybook/component coverage splits, slow/flaky CI reduction, or
  record-addressed selectors.
---

# E2E Coverage Triage

Use this skill to reduce E2E runtime and flakes without losing real coverage.
The aim is not to hide setup; it is to put each behavior at the right proof
layer.

## Core Rule

Scenario text describes product behavior. Step code handles test mechanics.

Only mention empty state, account creation, fixture IDs, seeded records, or
test isolation in Gherkin when that is the product behavior under test. Do not
mention helper internals, cleanup, Storybook, or `ctx` fields in feature text.

## Coverage Boundary

Keep in app E2E when the scenario proves a real cross-boundary workflow:

- authentication, role, tenant, or permission behavior;
- browser file input, upload, storage, background jobs, or async workflows;
- retry/cancel/dismiss actions that mutate persisted server state;
- routing/navigation persistence, cross-page behavior, or protected redirects;
- download endpoints, generated artifacts, webhooks, email callbacks, or
  API/UI integration.

Convert to seeded app E2E when the browser workflow matters but slow setup is
incidental:

- row actions for a known record;
- acknowledgement or approval actions for a known item;
- retry or state transitions when the initial state can be seeded;
- search, filter, navigation, download, or menu behavior around a known record.

Move to Storybook/component tests when the behavior is a pure UI state or
render contract:

- empty, pending, failed, duplicate, loading, terminal, and hidden states;
- button loading/error/disabled states;
- popover/menu visual states;
- copy, layout, responsive rendering, and component-only state transitions.

Use unit/API tests for pure logic and server contracts. Delete or collapse
redundant E2E only after another gated proof layer covers the behavior. If the
replacement layer is not in CI, mark the intended move but do not claim the
coverage is replaced.

Keep at least one real browser lane for each critical browser-only integration
surface.

## Shared-Session Pattern

The implementation may reuse a user, org, tenant, or session, but the scenario
should simply say the product role:

```gherkin
Given I am signed in as an admin
```

The step implementation should create or reuse the session, keep browser
contexts isolated unless sharing is deliberate, create scenario-specific
records under the shared scope, and store the target record in a named context
field.

## Record-Addressed UI Pattern

Tests should not require a pristine dashboard or list. Use stable selectors and
helper functions that address the specific record under test.

Prefer:

```ts
ctx.recordId = seeded.id;
await expect(rowForRecord(page, ctx.recordId)).toBeVisible();
```

Avoid:

```ts
page.getByTestId('row').first();
page.locator('.row.is-pending').first();
expect(page.getByTestId('row')).toHaveCount(1);
```

Global counts and empty-list assertions belong in E2E only when the product
behavior is truly about the whole list. Otherwise target the scenario's record.

## Rewrite Workflow

1. Pick one feature cluster, not the whole suite.
2. Inventory each scenario and write its proof boundary: E2E, seeded E2E,
   Storybook/component, unit/API, or delete.
3. Rewrite feature text into product language before changing step internals.
4. Convert setup inside steps: sign in in product language, reuse or establish
   session in code, seed only the needed record, and store its identity.
5. Replace broad locators with record-addressed selectors.
6. Add fail-fast waits when there are competing terminal states; use
   `Promise.race`-style helpers, not long generic waits.
7. Prove with the smallest relevant gates: generated BDD steps, targeted helper
   tests, targeted Playwright proof, or hosted preview E2E when behavior is
   environment-sensitive.
8. In the PR body, lead with the testing-strategy intent, then list mechanics.

source: agents
