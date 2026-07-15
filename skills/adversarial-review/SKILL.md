---
name: adversarial-review
description: >
  Perform an review on a design or implementation.  Use when planning or after
  implementation on non-trivial feature or refactoring work.
---

Use the `agent` CLI to request adversarial review.  Understand the sense of
"adversarial" --> you intentially invoke an agent to try to find

* incompleteness
* incorrect implementation
* verifiable bugs
* over-engineering:
  agent added complexity, "beld and suspenders", or any other extra logic, layering, or unnecessary abstractions not requested
* bad, smelly anti-patterns in the code
* clear, real race conditions or crash conditions

When you seek adversarial review, DO NOT condition your prompt.  Give the agent
a broad, short, open-ended prompt that briefly explains what is being reviewed
and let the agent find the flaws and vulnerabilities wherever they may be.

You'll be tempted to tell the agent to review something very specific.  Resist
the temptation.

Write the prompt into a file and pass it to the agent.  It should read like the following:

-------

Please review the feature work done on branch codex/issue-XYZ to find real signs of
implementation failure.

Read the issue description as well as the changes.

Report up to five issues.  If more than five issues are found, report only the most severe.

* incompleteness -- a real requirement has not been met.  Be specific and point to the missed instruction.
* incorrect implementation -- the implementation misunderstands the instruction or some higher-order project requirement.  Be specific.
* verifiable bugs -- must be accompanied by a reproducible test case
* over-engineering:
  agent added complexity, "beld and suspenders", or any other extra logic, layering, or unnecessary abstractions not requested.
  Be clear about how the implementation overlays unwanted and unrequested behavior, compexity, or infrastructure.
* bad, smelly anti-patterns in the code (critical only)
* clear, real race conditions or crash conditions (always critical) -- if not reproducible, must be accompanied by a detailed and believable explanation.

Write up your findings and classify them as critical, issue, or nit-pick.

-------

Note how the prompt does NOT tell the agent to focus its activity on any particular
file, idiom, pattern, issue, etc, but gives it full license to find any flaws or vulnerabilities.

That is your most important task here, is getting out of your own way in order
to ensure that the prompt you write gives plenty of room for the agent to use
its own capabilities.  Your own attempts to guide the agent to where you think
it should go can be quite destructive and should be avoided.

```
Usage:
  agent <workdir> [prompt text] [flags]
  agent [command]

Examples:
  agent /path/to/repo "Fix the failing tests"
  agent . --file PROMPT.md

Flags:
  -f, --file string            Read the prompt from a file
  -h, --help                   help for agent
      --session string         Resume an existing session ID
```
