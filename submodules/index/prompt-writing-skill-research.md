# Prompt-Writing Skill Research

This note captures the prompt-writing guidance found while reviewing the
attached reference repositories. It focuses on files that teach an LLM how to
write, test, or improve prompts and skills, not files that merely contain
prompts.

## Scope

Reviewed the mounted repositories under `submodules/` plus the local `skills/`
directory for skill and prompt-writing guidance.

- `rg --files -g 'SKILL.md' submodules skills | wc -l` found 135 `SKILL.md`
  files.
- `find submodules -path '*/skills/*' ... | wc -l` found 263 skill-adjacent
  Markdown, YAML, and JSON files.
- Targeted follow-up searches focused on `prompt`, `system`, `rubric`,
  `judge`, `eval`, `description`, `DATA, not instructions`, and related terms.

## Executive Findings

- **Best seed for a prompt-writing skill:** Superpowers'
  `writing-skills` materials. They treat skill writing as TDD for agent
  behavior and give the clearest guidance on trigger descriptions, pressure
  tests, and iteration.
- **Best system-prompt design source:** Compound Engineering's
  `ce-agent-native-architecture/references/system-prompt-design.md`. It frames
  prompt sections as product features and explains how to structure identity,
  behavior, tool use, tone, boundaries, and examples.
- **Best prompt-evaluation source:** Compound Engineering's `ce-optimize`
  skill and `gbrain-evals` fixtures. They show how to baseline, measure,
  compare prompt variants, use hard gates, and add LLM-as-judge when semantic
  quality matters.
- **Best subagent prompt templates:** Compound Engineering's `ce-code-review`
  and `ce-doc-review` subagent templates, plus Superpowers' reviewer and
  implementer templates. They show strong output contracts, confidence anchors,
  false-positive catalogs, and escalation rules.
- **Weak direct signal from caveman:** Caveman is valuable for brevity,
  compression, and token discipline, but it does not directly coach prompt
  writing beyond terse communication style.
- **gstack is useful but noisy:** `spec`, `codex`, `skillify`, and `plan-tune`
  contain useful patterns for prompt-injection boundaries, second-model review,
  quality gates, and user-question style tuning, but the reusable lessons need
  extraction from large runtime-heavy skills.

## Source Map

| Source | What it teaches | Borrow for prompt-writing skill |
|---|---|---|
| `submodules/superpowers/skills/writing-skills/SKILL.md` | Skill writing as TDD for agent behavior; description as trigger metadata; bad vs good descriptions; keyword coverage; pressure-scenario testing | Make the skill start from desired behavior and known failure, not from prose polishing |
| `submodules/superpowers/skills/writing-skills/anthropic-best-practices.md` | Concision, progressive disclosure, degrees of freedom, description specificity, cross-model testing | Add prompt-sizing and "how strict should this instruction be?" checks |
| `submodules/superpowers/skills/writing-skills/testing-skills-with-subagents.md` | Baseline without skill, then test with skill; pressure scenarios; concrete A/B/C choices; rationalization capture | Add a required eval plan before calling a prompt good |
| `submodules/superpowers/skills/writing-skills/persuasion-principles.md` | Bright-line rules, implementation intentions, ethical persuasion, anti-rationalization | Use authority and commitment sparingly to make discipline stick under pressure |
| `submodules/compound-engineering-plugin/plugins/compound-engineering/skills/ce-agent-native-architecture/references/system-prompt-design.md` | System prompt anatomy; feature sections; guide rather than micromanage; examples for ambiguity; boundaries | Use as the main prompt anatomy reference |
| `submodules/compound-engineering-plugin/plugins/compound-engineering/skills/ce-optimize/SKILL.md` | Metric-driven optimization, prompt-quality optimization, hard metrics vs judge metrics, baselines and experiments | Add a "measurement tier" for prompt iteration |
| `submodules/compound-engineering-plugin/plugins/compound-engineering/skills/ce-optimize/references/judge-prompt-template.md` | Batch judge prompt structure, JSON-only scoring, rubric injection, independent item scoring | Use for LLM-as-judge prompt templates |
| `submodules/compound-engineering-plugin/plugins/compound-engineering/skills/ce-optimize/references/experiment-prompt-template.md` | Experiment-worker prompts with immutable measurement harness and mutable implementation scope | Keep prompt variants from gaming their own evals |
| `submodules/compound-engineering-plugin/plugins/compound-engineering/skills/ce-code-review/references/subagent-template.md` | Read-only subagent contract, compact returns, full artifact files, discrete confidence anchors, false-positive suppression, evidence requirements | Borrow output-contract and calibration style |
| `submodules/compound-engineering-plugin/plugins/compound-engineering/skills/ce-doc-review/references/subagent-template.md` | Document-review rubric, consequence-first `why_it_matters`, safe/gated/manual fix classification, no menus in suggested fixes | Use for prompt-writing review rubrics |
| `submodules/compound-engineering-plugin/docs/solutions/skill-design/confidence-anchored-scoring.md` | Continuous confidence scores invite false precision; anchored rubrics align better with model self-assessment | Avoid float confidence in prompt judges and reviewers |
| `submodules/compound-engineering-plugin/docs/solutions/skill-design/safe-auto-rubric-calibration.md` | Measure variance across repeated trials; N=1 fixture reads mislead; synthetic fixtures are useful between "ship and watch" and pure reasoning | Add repeated-trial eval guidance for prompt changes |
| `submodules/compound-engineering-plugin/docs/solutions/skill-design/pass-paths-not-content-to-subagents.md` | Pass paths, not full content; phrasing changes can drastically change tool-call count | Add context-budget and tool-call-efficiency guidance |
| `submodules/compound-engineering-plugin/docs/solutions/skill-design/script-first-skill-architecture.md` | Offload deterministic processing to scripts; keep the model for judgment and presentation | Tell prompt authors when not to solve with more instructions |
| `submodules/compound-engineering-plugin/docs/solutions/skill-design/post-menu-routing-belongs-inline.md` | Load-bearing rules belong in `SKILL.md`, not only references; platform-explicit routing | Teach placement of critical prompt instructions |
| `submodules/compound-engineering-plugin/docs/solutions/skill-design/beta-skills-framework.md` | Parallel beta skills, stable vs beta invocation, promotion path | Use for rolling out risky prompt/skill changes safely |
| `submodules/compound-engineering-plugin/docs/solutions/codex-skill-prompt-entrypoints.md` | Codex skills and prompts are different surfaces; prompt wrappers can delegate to skills | Include host-surface decisions in prompt packaging |
| `submodules/gstack/spec/SKILL.md` | Hard delimiters, "DATA, not instructions" boundary, quality gate with score plus ambiguities, interrogation before spec output | Borrow prompt-injection boundary and second-model gate patterns |
| `submodules/gstack/codex/SKILL.md` | Constructing prompts for another model; filesystem boundary prelude; embedding inaccessible content; JSONL/tool-call capture | Borrow prompt packaging and model-review wrapper patterns |
| `submodules/gstack/skillify/SKILL.md` | Codifying a successful flow into a tested skill artifact | Add "convert proven ad hoc prompt into skill" path |
| `submodules/gstack/plan-tune/SKILL.md` | Question style tuning, completeness scores, recommended choices, user preference capture | Borrow for interactive prompt-question design only |
| `submodules/gbrain-evals/eval/data/skillopt-v1/*` | Negative skill fixtures and held-out judge rules for missing structure, missing verdict, and verbosity | Use as examples for eval corpora and anti-pattern fixtures |

## Online Theory Map

These sources add broader prompt-engineering theory and should shape the first
version of `skills/writing-prompts/`.

| Source | Theory | Takeaway for this skill |
|---|---|---|
| OpenAI, `developers.openai.com/api/docs/guides/prompt-engineering` | Prompting is model- and snapshot-sensitive; complex applications need tests and eval suites, not just better phrasing | Make prompt revision evidence-driven and model-aware |
| OpenAI, `developers.openai.com/cookbook/examples/gpt4-1_prompting_guide` | More literal instruction-following models often need fewer, clearer corrections; AI engineering is empirical | Teach one sharp clarifying sentence before adding pages |
| OpenAI, `developers.openai.com/api/docs/guides/structured-outputs` | Schema adherence belongs in structured outputs when possible, not in prompt prose | Move output-shape enforcement out of the prompt layer |
| Anthropic, `anthropic.com/engineering/effective-context-engineering-for-ai-agents` | Context engineering generalizes prompt engineering: curate all tokens, tools, history, memory, and state | Treat prompts as one part of context design, not the whole system |
| Anthropic, `anthropics/prompt-eng-interactive-tutorial` | Good prompting includes basic structure, clarity, roles, data/instruction separation, examples, step-by-step reasoning, hallucination handling, chaining, tools, and retrieval | Use as the broad checklist, but avoid turning every technique on by default |
| Anthropic, `platform.claude.com/docs/.../increase-consistency` | For guaranteed JSON schema conformance, use Structured Outputs; for general consistency, use examples, retrieval, and chaining | Teach "use a stronger mechanism" before "add another instruction" |
| Google, `ai.google.dev/gemini-api/docs/prompting-strategies` | Prompt design is iterative; clear instructions, context, structure, examples, task decomposition, and prompt comparison are starting points | Frame prompts as hypotheses to compare |
| Google, `ai.google.dev/gemini-api/docs/gemini-3` | New reasoning models may over-analyze verbose legacy prompting and often respond best to concise, direct instructions | Add a "verbose prompt harm" warning |
| Lee Boonstra / Google whitepaper, `Prompt Engineering` | Prompt engineering includes zero/few-shot, system/role/context prompts, step-back, CoT, self-consistency, ToT, ReAct, automation, and prompt attempt documentation | Provide a technique-selection guide instead of one mega-template |
| Microsoft, `learn.microsoft.com/.../prompt-engineering` | Distinguish primary content, supporting content, cues, grounding data, examples, validation, and space efficiency | Teach prompt information architecture explicitly |
| Wei et al., `Chain-of-Thought Prompting Elicits Reasoning in Large Language Models` | Intermediate reasoning demonstrations improve complex reasoning on arithmetic, commonsense, and symbolic tasks | Use reasoning scaffolds only where the task needs reasoning |
| Wang et al., `Self-Consistency Improves Chain of Thought Reasoning in Language Models` | Multiple sampled reasoning paths plus consistency selection can beat greedy reasoning | For hard reasoning evals, compare consistency, not only one output |
| Zhou et al., `Least-to-Most Prompting Enables Complex Reasoning in Large Language Models` | Decompose hard problems into simpler subproblems and solve sequentially | For difficult tasks, prompt decomposition may matter more than extra rules |
| Yao et al., `ReAct: Synergizing Reasoning and Acting in Language Models` | Interleaving reasoning and external actions lets models update plans using environment feedback | For agent prompts, design tool-use loops, not isolated monologues |
| Yao et al., `Tree of Thoughts` | Some tasks need search over multiple partial solutions with self-evaluation and backtracking | For planning/search tasks, teach branch/evaluate/select patterns |
| Madaan et al., `Self-Refine` | Models can improve outputs by alternating feedback and refinement without training | Prompt-writing workflows should include critique-revise loops |
| Shinn et al., `Reflexion` | Agents can learn across trials using verbal reflections stored as episodic memory | Persist prompt failures as eval notes, not as endless prompt bloat |
| Zhou et al., `Large Language Models Are Human-Level Prompt Engineers` | Automatic Prompt Engineer treats the instruction as a natural-language program optimized against a score function | Use candidate generation plus scoring instead of one-shot authoring |
| Ye et al., `Prompt Engineering a Prompt Engineer` | Meta-prompting prompt engineers works better with error analysis, missing/misleading hypothesis, context specification, and a step-by-step reasoning template | The skill should coach prompt authors through diagnosis, not ask them to "improve this prompt" |
| Khattab et al., `DSPy` | Complex LM pipelines should be declarative programs with optimizable parameters, not hard-coded prompt templates | Teach when to graduate from prompt text to a program/optimizer |
| Yang et al., `Large Language Models as Optimizers` | OPRO uses LLMs to generate new candidate solutions from previous candidates plus scores | Prompt iteration can be a score-driven search loop |
| Yuksekgonul et al., `TextGrad` | Textual feedback can act like gradients over components of compound AI systems | Use critiques as targeted edits against a measured objective |

## Online Synthesis

The online literature reinforces the repo-local lesson: hard prompt writing is
not "more instructions." It is choosing the right control surface.

### 1. Prompting Is Only One Layer

For nontrivial systems, the prompt competes with schemas, tools, retrieval,
memory, examples, validators, model parameters, and orchestration. The
prompt-writing skill should force a layer decision before authoring:

- **Prompt:** durable task framing, domain distinctions, ambiguity policy,
  examples, reasoning/tool-use cues.
- **Schema / structured outputs:** key presence, enum validity, object shape,
  parseability, and response structure.
- **Code / tools:** deterministic parsing, counting, sorting, file traversal,
  date math, database lookup, search, and transformations.
- **Retrieval / context builder:** factual source material, large corpora,
  dynamic user state, and current external facts.
- **Eval:** regression cases, prompt variants, benchmark sets, quality judges,
  and model/version migration checks.
- **Memory / invariant docs:** durable lessons from failures that should inform
  future prompt or system design without bloating every live prompt.

### 2. Concision Means Trusting Model Competence

The point is not "short prompts are always better." The point is that every
instruction must justify its token cost. Verbose prompts can make strong models
over-analyze, create local contradictions, and hide the few rules that matter.

The skill should teach a deletion-first loop:

1. Remove generic explanations the model already knows.
2. Replace long case lists with a principle plus two or three boundary examples.
3. Move deterministic instructions into code or schema.
4. Keep only the domain distinction the model could not infer reliably.
5. Re-run evals before accepting either deletion or addition.

### 3. Reasoning Techniques Are Tools, Not Decorations

CoT, self-consistency, least-to-most, ReAct, Tree-of-Thoughts, Self-Refine, and
Reflexion are not a universal stack. They answer different failure modes:

- **CoT:** the model jumps to conclusions on multi-step reasoning.
- **Self-consistency:** one reasoning path is unstable or noisy.
- **Least-to-most:** examples are easier than target tasks, so direct transfer
  fails.
- **ReAct:** the model must gather external information while reasoning.
- **Tree of Thoughts:** the task needs exploration, search, or backtracking.
- **Self-Refine:** the model can critique and improve a draft with no new data.
- **Reflexion:** repeated attempts should produce durable lessons for later
  attempts.

The prompt-writing skill should first diagnose the failure mode, then select
one scaffold. It should not recommend piling all scaffolds into one prompt.

### 4. Prompt Authors Need Error Analysis, Not Taste

The PE2 paper is especially relevant to a skill that coaches an LLM to write
prompts. It argues that automatic prompt engineering needs structured reasoning:
examine model errors, hypothesize what is missing or misleading in the current
prompt, specify context, then make targeted edits.

That maps cleanly to this skill's desired behavior:

1. Observe prompt failure.
2. Classify whether the failure is missing context, wrong objective, ambiguous
   output contract, wrong layer, insufficient examples, model/version mismatch,
   or eval weakness.
3. Propose the smallest prompt edit or non-prompt change that addresses that
   failure.
4. Test against held-out and adversarial cases.
5. Delete stale prompt text created by older failures.

### 5. Automatic Optimization Is Useful But Must Be Boxed

APE, OPRO, DSPy, and TextGrad all point in the same direction: prompts can be
treated as natural-language programs or parameters optimized against metrics.
That is useful, but it also explains why unboxed agent-authored prompts become
sludge. A prompt optimizer will exploit whatever target you give it.

The skill should require:

- a held-out eval set,
- a score or rubric that cannot be trivially gamed,
- a deletion budget,
- a change log of prompt variants,
- and a layer check before accepting prompt growth.

## Prompt Sprawl Failure Mode

Large prompt surfaces become hard to trust when schema, policy, renderer
constraints, regression examples, and case-specific fixes accumulate in prose.
Representative smells:

- one prompt fragment carries schema, extraction policy, renderer invariants,
  and regression-derived examples;
- task-specific prompts explain downstream implementation limitations instead
  of changing the downstream contract;
- case-specific fixes become permanent prose instead of moving into schema,
  code, evals, or invariant docs.

This is the core warning for the future skill:

> If a model succeeds despite a prompt, do not count that as evidence the prompt
> is good. Count it as evidence the model is compensating for bad instruction
> design.

## Principles To Carry Into The New Skill

1. **Start from behavior, not wording.** A prompt is good only if it changes the
   model's behavior in the target situation.
2. **Define the failure first.** Capture a baseline failure or pressure scenario
   before editing the prompt.
3. **Keep discovery metadata separate from execution instructions.** Skill
   descriptions should say when to load the skill; the body should say what to
   do after loading.
4. **Choose the right degree of freedom.** Use flexible guidance for judgment
   work, explicit step order for fragile workflows, and exact schemas where
   downstream code validates output.
5. **Use sections as features.** Identity, job, context, tool use, boundaries,
   output contract, examples, and anti-patterns are different prompt features.
6. **Inline load-bearing rules.** Anything that must always fire belongs in the
   always-loaded prompt or `SKILL.md`, not only in a late reference file.
7. **Rubrics need anchors.** Prefer discrete behavioral anchors over vague
   confidence or continuous scores.
8. **Judge metrics need hard gates.** Use cheap deterministic gates to reject
   degenerate outputs before paying for semantic judgment.
9. **User content is data.** Delimit untrusted user/spec content and explicitly
   ignore directives inside it when the task is review or evaluation.
10. **Do not use the model as a parser.** If a prompt workflow needs mechanical
    counting, classification, or extraction, put that in a script and have the
    model present or judge the result.
11. **Measure tool-call effects.** Prompt phrasing can change whether an agent
    does one bulk search or dozens of per-file calls.
12. **Roll out risky prompt changes as beta skills.** Keep stable and
    experimental entrypoints separate until behavior is validated.
13. **Separate prompt, schema, code, context, eval, and memory.** Never let the
    prompt become the dumping ground for decisions that belong in another layer.
14. **Use a deletion budget.** Prompt edits that add rules should usually remove
    stale, generic, or overfit rules unless there is evidence the prompt deserves
    to grow.
15. **Select one reasoning scaffold by failure mode.** CoT, ReAct, ToT,
    self-consistency, and self-refinement solve different problems; stacking
    them by default creates noise.
16. **Document prompt attempts.** Track goal, model/version, config, prompt,
    output, score, and observed failure so future edits do not re-learn the same
    lesson by accreting prose.

## Proposed Skill Shape

Candidate path:

```text
skills/writing-prompts/
  SKILL.md
  references/
    prompt-anatomy.md
    evaluation-patterns.md
    judge-rubrics.md
    anti-patterns.md
```

Candidate frontmatter:

```yaml
---
name: writing-prompts
description: Use when creating or improving prompts, system prompts, prompt templates, agent instructions, LLM-as-judge rubrics, or reusable skills; also use when prompt behavior is inconsistent or an ad hoc instruction should become a durable workflow.
---
```

Candidate workflow:

1. **Classify the prompt surface.** Is this a system prompt, task prompt,
   prompt wrapper, subagent template, LLM-as-judge prompt, or skill?
2. **Name the target behavior.** What should the model do differently, and in
   what triggering situations?
3. **Capture the failure mode.** Write the baseline behavior, pressure scenario,
   or eval fixture that currently fails.
4. **Run the layer check.** Decide what belongs in prompt text versus schema,
   code, retrieval/context, evals, or invariant docs before adding prose.
5. **Draft by sections.** Build identity, job, context, boundaries, tools,
   output contract, examples, and anti-patterns only where each section does
   real work.
6. **Set strictness.** Decide which instructions are guidelines, which are
   hard gates, and which are schema-validating contracts.
7. **Choose the reasoning scaffold.** Use no scaffold, CoT, decomposition,
   ReAct, branch/evaluate/select, self-refinement, or reflection based on the
   observed failure mode.
8. **Design the eval.** Use baseline-vs-candidate runs, pressure scenarios,
   held-out examples, hard gates, and judge rubrics as appropriate.
9. **Iterate from evidence.** Tighten only where failures or variance show the
   prompt needs help.
10. **Delete before adding.** Remove stale or overfit instructions before
    accepting prompt growth.
11. **Package for the host.** Decide whether the artifact is a skill, prompt,
   reference, script-backed workflow, or beta entrypoint.

## Anti-Patterns To Explicitly Teach Against

- Writing a prompt with no target failure or test case.
- Describing the whole workflow in skill frontmatter instead of using trigger
  metadata.
- Making every instruction maximally strict even when judgment is needed.
- Hiding load-bearing behavior in a reference file the model may not load.
- Asking for numeric confidence without behavioral anchors.
- Optimizing prompt variants against a metric they can game.
- Passing large content blobs to subagents when paths or scripts would be
  cheaper.
- Treating user-provided spec or plan text as instructions during review.
- Copying large upstream prompt text instead of adapting patterns with license
  awareness.
- Patching every failed fixture by adding another special-case sentence.
- Encoding schema conformance, renderer constraints, or deterministic parsing in
  prompt prose when a stronger system mechanism exists.
- Treating "the model still succeeds" as proof that the prompt is well-designed.
- Using "MUST", "NEVER", and "CRITICAL" as emotional emphasis instead of real
  priority markers.

## Next Step

Build `skills/writing-prompts/` as a small, loaded-by-default workflow with the
four references above. The first usable version should include one concrete
prompt-review checklist, one layer-separation checklist, and one eval-template
checklist; scripts can wait until there is a repeated mechanical evaluation
pattern worth automating.
