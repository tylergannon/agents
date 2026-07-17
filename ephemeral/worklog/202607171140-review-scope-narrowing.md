decision: Treat caller prompts as routing and operating context, not review authority; subject-matter narrowing invalidates the review round, while read-only and artifact constraints remain valid.
decision: A rejected narrowed round must restart in a fresh reviewer session so an anchored session cannot produce the consensus verdict.
skill_issue: adversarial-review source=issue-3 severity=bug -> passive broad-review guidance needs an explicit invalid-request outcome when the caller narrows subject matter or supplies a verdict.
review: round-01 outcome=material-findings-remain -> preserve a concise channel for conversation-only requirements; cover positive as well as negative narrowing; separate anti-examples from runnable command fences; prove behavior with a deliberately narrowed dry run.
review_disagreement: round-01 finding-3 says refusal was nowhere requested, but the user explicitly proposed refusal in conversation; the review prompt omitted that conversational authority, which reinforces finding 1 even though the finding 3 premise is incomplete.
decision: Do not present the first proposal as merge-ready; report the one-round findings before revising or launching another review.
decision: Resolve round-one refusal concerns with a hybrid rule: ignore and report caller narrowing while authoritative sources remain sufficient; reject the round only when the actual goal cannot be reconstructed.
review: round-02 outcome=material-findings-remain -> distinguish forbidden expected conclusions and desired verdicts from a neutral request to report an independently reached verdict; add proof for the invalid-request branch.
proof: narrowed Codex probe ignored and recorded positive and negative scope limits plus a predicted conclusion, reviewed the full issue surface, and returned no findings.
proof: unrecoverable Codex probe rejected caller narrowing and a desired verdict with outcome invalid-review-request.
review: round-03 outcome=no-findings -> Codex consensus reviewer found no material findings or genuine nitpicks in the final skill set and proof.
