correction: Adversarial-review should be reviewer-side so the launched agent loads the stable broad-review rubric; a separate caller-side skill should help the top-level agent construct the minimal launch prompt.
decision: Split responsibilities across request-adversarial-review for one-pass launch, adversarial-review for reviewer behavior, and consensus for repeated review plus adjudication.
decision: The caller prompt contains only locator context and the explicit `/adversarial-review` invocation; the reviewer-side skill owns breadth, finding categories, evidence standards, severity, and output limits.
decision: Skill validation passed for all three review skills; package validation includes 29 files and git diff --check is clean.
