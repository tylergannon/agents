# gstack

- Path: `../gstack`
- Upstream: `garrytan/gstack`
- Pinned commit: `1626d4857bfe30da2690dd6a3217961934aa3192`
- License: MIT
- Package: `gstack` `1.57.7.0`

## What Is In It

- A large AI development-team toolkit centered on Claude Code, with support for
  other hosts.
- Fifty-eight `SKILL.md` files covering product office hours, plan review,
  engineering/design/DX review, QA, browser testing, security audits, shipping,
  deployment, retros, iOS QA, gbrain sync, benchmarking, and more.
- A persistent browser stack under `browse/`, implemented in Bun/TypeScript with
  a local server, browser manager, CDP bridge, security filters, tests, and
  compiled CLI targets.
- Seventy-two helper binaries/scripts under `bin/`, including config, browser,
  gbrain, telemetry, learning, decision, benchmark, security, and release helpers.
- Design tooling under `design/` and `design-*` skills.
- Supabase functions/migrations for telemetry and update-check flows.
- A large test suite: 510 test files in the local scan.

Local scan: 1121 files, 58 `SKILL.md` files, 72 `bin/` files, 42 docs files, 510
test files.

## What We Can Learn

- Stateful agent tools should not cold-start expensive resources. gstack's
  browser uses a persistent localhost daemon, a state file, health checks,
  version restart, and short command round trips.
- Browser automation needs security boundaries: local-only listeners, scoped
  tokens for tunnels, command allowlists, prompt-injection scanning, and egress
  sanitization.
- Skills can be a product taxonomy. gstack names specialists by job to be done:
  QA lead, security officer, release engineer, design partner, performance
  engineer.
- CLI helpers are often better than stuffing logic into prompts. The skills call
  small tools for config, redaction, learning, benchmarks, timelines, and browser
  commands.
- Large agent systems need their own tests for prompt shape, template generation,
  browser behavior, security, e2e skill routing, and regressions.

## What We Can Use

- `browse/ARCHITECTURE` material via `ARCHITECTURE.md` and `browse/src/` as the
  strongest reference for local browser/daemon tooling.
- `browse/test/` for security and browser behavior test coverage ideas.
- `bin/gstack-*` for command taxonomy and standalone helper ergonomics.
- `docs/designs/` for design rationale examples.
- `scripts/gen-skill-docs.ts` and skill docs tests for generated skill
  documentation patterns.
- `slop-scan.config.json` and review-related tests as examples of automated
  prompt/style quality gates.

## What We Can Borrow

- A persistent local-service template for future stateful tools in `agents`.
- A browser QA/security checklist before implementing any shared browser runner.
- A skill taxonomy around plan, review, QA, ship, deploy, learn, and retro.
- A generated-doc pipeline that keeps skill docs aligned with source skills.
- A local learning/decision-log CLI shape, adapted to this repo's simpler docs
  model.

## Cautions

- gstack is broad. Borrow one subsystem at a time with tests.
- Do not reuse browser tunnel or cookie flows without equivalent security tests.
- Treat README performance/productivity claims as upstream context, not local
  proof for this repo.
