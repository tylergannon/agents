# Route review skill guidance through Gimbal

Updated the canonical `consensus` and `request-adversarial-review` skills to launch reviewer prompts with `gimbal run-prompt --workdir`. Each consensus round now starts a fresh Gimbal run and writes to a new review artifact, matching the supported Gimbal run model. The installed copies were updated in place so current sessions pick up the change before the next skill refresh.

The `orchestrate-attractor-loops` installed skill also named the standalone `agent` executable; its installed copy now uses Gimbal prompt runs. Its canonical source is not present in the available Attractor checkout. Provider-specific lanes in `df-sprint-execute` and the separate proof uploader utility are not invocations of the standalone `agent` CLI and were left intact.

Validation: `git diff --check`; searched canonical and installed skill directories for the old `agent <workdir>` examples and reviewer-selection instructions. No runtime review was launched.
