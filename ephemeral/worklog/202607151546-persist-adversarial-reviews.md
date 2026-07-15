correction: Adversarial review must write its complete findings to an artifact under `ephemeral/`; the reviewer is read-only for implementation files, not forbidden from all writes.
decision: Use one immutable artifact per review round under `ephemeral/reviews/`, and have the reviewer return only the artifact path plus a one-line result to the caller.
decision: The caller assigns the review path, reads the populated artifact, and either summarizes it for a one-pass review or carries its session ID and outcome into consensus.
decision: All three review skills validate; git diff --check and package validation with 29 files pass.
