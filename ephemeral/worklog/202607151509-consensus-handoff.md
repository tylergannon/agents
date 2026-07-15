correction: The compressed consensus skill let the caller interpret `Use /adversarial-review` as prompt text for the subprocess; the parent agent must load the skill locally and send only the resulting review prompt to `agent`.
friction: Ambiguous skill handoff produced `agent <workdir> "/adversarial-review ..."` instead of a prompt file and direct reviewer instructions -> make ownership and forbidden invocation shape explicit.
decision: Consensus now states that the caller loads adversarial-review locally, writes the reviewer prompt to a file, and never sends slash-command skill invocations to the reviewer.
decision: Targeted proof passed: agent CLI exposes --file and --session, git diff --check is clean, and package validation includes 27 files.
