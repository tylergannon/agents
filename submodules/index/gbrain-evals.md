# gbrain-evals

- Path: `../gbrain-evals`
- Upstream: `garrytan/gbrain-evals`
- Pinned commit: `565b80754ffa6abb9afb041026f2fab048aa7553`
- License: MIT
- Package: `gbrain-evals` `0.2.0`

## What Is In It

- A reproducible benchmark suite for long-term agent memory systems.
- Public corpora under `eval/data/`, including fictional lives, messy weeks,
  synthetic content, public dataset adapters, sealed gold answers, qrels, and
  baseline files.
- Benchmark runners under `eval/runner/`, including retrieval, identity, time,
  provenance, skill compliance, workflow, adversarial, multimodal, and
  precision-memory tests.
- Schemas under `eval/schemas/` for corpus manifests, evidence contracts, public
  probes, scorecards, tool schemas, and transcripts.
- Published benchmark reports and charts under `docs/benchmarks/`.
- Unit tests under `test/eval/`.
- Skill optimization fixtures under `eval/data/skillopt-v1/`.

Local scan: 676 files, 102 TypeScript files, 294 JSON files, 253 Markdown files,
27 docs files, 22 test files.

## What We Can Learn

- Public evals are strongest when the corpus, runner, baseline, qrels, and report
  are all committed and reproducible from a commit hash.
- Sealed answer keys prevent the system under test from seeing expected answers.
- Retrieval systems should report recall and precision together; optimizing one
  metric can make the other worse.
- Honest scorecards include weak numbers and explain when a metric is a mismatch
  for the default product behavior.
- Eval harnesses should test trust boundaries and adversarial cases, not just
  happy-path retrieval.

## What We Can Use

- `eval/runner/multi-adapter.ts` and related adapters as a model for scoring
  multiple systems through one interface.
- `eval/schemas/` for schema shape around scorecards, evidence, and transcripts.
- `qrels/` and `baselines/` for baseline storage layout.
- `docs/benchmarks/` for report shape and benchmark provenance.
- `eval/data/skillopt-v1/` for skill-improvement eval examples.

## What We Can Borrow

- A small agents eval template with these parts: corpus, sealed answers,
  adapter interface, runner, qrels, baseline, and report.
- A scorecard convention that records command, commit, dataset, thresholds, and
  known weaknesses.
- A "gold answers never enter context" rule for agent-tool evals.
- A precision/recall reporting pattern for search, memory, docs, and corpus
  tooling.

## Cautions

- Do not cite upstream benchmark numbers as local proof unless the same command
  is reproduced at the pinned commit.
- Some runs require API keys or public dataset downloads. Keep local smoke tests
  offline where possible.
