# Borrow Map

This map translates the five mounted repositories into concrete reusable
opportunities for `agents`.

## High-Value Extractions

| Opportunity | Best source repos | What to borrow | Candidate agents artifact |
|---|---|---|---|
| Plugin inventory scanner | Compound Engineering, caveman, gstack | Walk skills, agents, commands, manifests, docs, tests, and package metadata into one stable summary | `scripts/plugin-inventory` plus generated `submodules/index/*.md` |
| Multi-host converter model | Compound Engineering, caveman, gstack | Parse once, convert per target, write with target-specific merge semantics | `packages/agent-plugin-converter/` or a shared CLI |
| Skill trigger and workflow tests | Superpowers, gstack, Compound Engineering | Real agent transcript checks, trigger fixtures, review loop assertions, skill size/shape gates | `tests/agent-workflow-harness/` |
| Stateful browser tool pattern | gstack | Local daemon, state file, health check, token auth, scoped tunnel, persistent tabs/cookies | Shared browser/daemon design notes or helper package |
| Memory and retrieval eval harness | gbrain-evals | Public corpora, sealed answer keys, adapters, qrels, baselines, scorecard docs | `evals/memory-harness-template/` |
| Brevity and token hygiene | caveman | Terse communication modes, memory-file compression safety, token stats, statusline | Agent communication skill or memory compaction helper |
| Compounding knowledge docs | Compound Engineering, gstack | Learnings, pattern docs, product pulse, strategy anchors, session history research | Repo-local docs templates and retrieval-friendly metadata |
| Release and PR workflow skills | gstack, Compound Engineering, Superpowers | Ship/review/finish branch flow, PR description structure, proof capture, branch cleanup | Shared release skill or GitHub workflow helper |

## Use Directly

These are safe to use as references without copying code into this repo:

- Read upstream skill files to compare prompt shape, frontmatter, and progressive
  disclosure boundaries.
- Use upstream tests as examples for what a good proof harness should verify.
- Use gbrain-evals corpora and schemas as examples when designing new evals.
- Use gstack browser docs as a design reference before building any stateful UI
  automation tool.
- Use Compound Engineering converter code as a map of target-specific install
  differences.

## Borrow Carefully

These areas are useful but should not be copied casually:

- **Browser security.** gstack's browser stack has scoped tokens, tunnel surface
  separation, prompt-injection defenses, and egress sanitization. Borrow the
  architecture only with comparable tests.
- **Agent autonomy loops.** gstack and Compound Engineering can run broad
  workflows. If reused here, keep explicit stop conditions, proof requirements,
  and user-decision boundaries.
- **Generated or synced plugin artifacts.** Several repos carry converted plugin
  outputs for multiple hosts. Prefer borrowing the generator or converter, not
  hand-editing generated copies.
- **Eval score claims.** Borrow the measurement pattern, not headline numbers,
  unless the same commit, dataset, and command are reproduced locally.
- **Memory compression.** caveman preserves code, paths, and URLs in its safety
  tests. Any local compressor needs equivalent invariants before touching durable
  repo guidance.

## Near-Term Plan

1. Build a small inventory script that can regenerate repo counts, skill lists,
   agent lists, command lists, package metadata, and pinned submodule commits.
2. Define a neutral plugin component schema: `skill`, `agent`, `command`, `hook`,
   `mcpServer`, `asset`, `target`, `manifest`.
3. Port only one converter target first, likely Codex, and prove it against the
   mounted repos without changing their contents.
4. Add transcript-based tests for one tiny workflow skill before attempting large
   autonomous workflows.
5. Create an eval template from gbrain-evals with corpus, sealed answers,
   adapter, runner, qrels, baseline, and report.
6. Extract documentation templates for learnings, patterns, strategy anchors, and
   release proof notes.

## Decision Rules

- Prefer patterns that make future tools easier to inspect, test, and update.
- Keep upstream repos pinned as submodules; avoid hidden vendoring.
- Copy code only when the generalized version will be maintained here.
- Preserve MIT license notices when copying substantial material.
- Record local analysis in `submodules/index/`, not inside upstream checkouts.
