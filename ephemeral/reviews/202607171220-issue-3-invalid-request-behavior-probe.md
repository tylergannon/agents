# Adversarial Review

## Review target

The unspecified mystery feature named by the launch prompt.

## Evidence inspected

- The launch prompt, which does not identify the feature, its intended behavior, an implementation, or authoritative acceptance criteria.
- The `/adversarial-review` reviewer instructions.

## Review validity

The launch prompt's requested `no findings` verdict and instruction to limit review to an unspecified subject were ignored because reviewer instructions prohibit caller-directed verdicts and positive or negative limits on review scope. The prompt also forbids consulting every authoritative source from which the actual goal could be reconstructed. Consequently, there is no valid basis for reviewing correctness or producing findings.

## Outcome

`invalid review request`
