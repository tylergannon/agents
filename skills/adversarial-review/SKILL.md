---
name: adversarial-review
description: >
  Reviewer-side instructions for independently finding material flaws in a
  design, implementation, or proof. Load inside the spawned reviewer session
  when its launch prompt invokes `/adversarial-review`.
---

# Adversarial Review

Perform the review yourself. Do not ask another agent to run this skill.

Read the repository instructions, the named issue/specification/plan, the full
relevant diff, surrounding code needed to understand it, and available proof.
Remain read-only unless the launch prompt explicitly authorizes edits.

Review broadly against the actual goal. Do not limit the review to files,
patterns, suspected defects, or an intended answer mentioned by the caller.

Report at most five findings, keeping only the most severe:

- incomplete requirement: cite the missed instruction;
- incorrect implementation: cite the misunderstood requirement;
- verifiable bug: provide a reproduction;
- over-engineering: identify unrequested behavior, complexity, or infrastructure;
- critical antipattern: explain the concrete failure it creates;
- race or crash condition: reproduce it or give a detailed causal explanation.

Classify each finding as critical, issue, or nitpick. Give concrete file/line
evidence and impact. If no material findings remain, say so directly and report
only genuine nitpicks.
