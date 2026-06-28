---
name: dependency-upgrade-pr
description: >
  Upgrade dependencies through a proof-backed pull request with exact-head CI,
  preview/runtime verification, auto-merge, and merge follow-through. Use when
  changing package versions, lockfiles, package-manager metadata, native or
  platform-sensitive dependencies, framework/compiler/adapters, npm dist-tags,
  or recurring dependency-update automation.
---

# Dependency Upgrade PR

Use this skill for dependency updates where build success is not enough. The
agent owns the branch until it merges or the dependency change proves a real
bug or incompatibility.

## Workflow

1. Use `proof-of-work` for the PR lifecycle: current `origin/main`, `codex/`
   branch/worktree, worklog, PR, exact-head checks, auto-merge, and merge
   verification.
2. Inspect current package surfaces before editing: `package.json`, lockfile,
   package-manager config, workspace files, release notes when needed, npm
   dist-tags, and `pnpm outdated` or the repo's package-manager equivalent.
3. Handle prerelease lines intentionally. For Svelte/SvelteKit/adapters,
   compilers, and framework packages, do not downgrade from an intentional
   prerelease channel just because `latest` points at an older stable line.
4. Update versions and lockfiles with the repo's package manager. Do not hand
   edit lockfiles except for mechanical conflict resolution.
5. Run minimum local proof for dependency changes: frozen install, typecheck,
   unit tests, and build. Add any repo-specific proof documented by the target
   repo.
6. For framework, compiler, adapter, build-tool, native, or
   platform-sensitive packages, prove the deployed preview/runtime path too.
7. Commit, push, update the PR body with decisions and proof, and enable
   auto-merge once local proof passes and required exact-head checks are pending
   or green.
8. Monitor until merged. If checks fail, inspect logs and fix, rebase, rerun, or
   back out. Do not leave the merge path ambiguous.

## Runtime Proof

- Treat build/deploy readiness as readiness, not runtime proof.
- For native or platform-sensitive packages such as image processors, crypto,
  database drivers, or optional native binaries, exercise a deployed route,
  function, worker, or browser flow that imports or uses the package on the
  target runtime.
- For Vercel or similar previews, verify the PR head SHA and confirm the actual
  browser/runtime test step ran. A skipped, stale, or unrelated preview check is
  not evidence.
- If runtime proof needs a temporary probe route, script, or diagnostic, remove
  it before merge unless it is a useful permanent regression test.

## Failure Triage

- Rebase or restart stale branches from current `origin/main`.
- Keep auto-merge enabled after local proof unless the dependency change causes
  a real bug or incompatible dependency breakage.
- Distinguish unrelated flaky CI, dirty branch state, and preview-environment
  failures from dependency-caused failures; work through the former.
- If the dependency itself is bad, document the exact package/version,
  reproduction, logs, and rollback or pinning decision.

## Closeout

Report:

- packages changed and why,
- prerelease/dist-tag decisions,
- local proof commands,
- preview/runtime proof path and exact PR head SHA,
- auto-merge or merge state,
- unresolved dependency-caused bug or external blocker.

source: agents
