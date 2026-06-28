# agents Skills

This directory is the source of published agent skills.

This directory contains only published portable skills. Repo-local maintainer
skills live in `../.agents/skills/`.

Install these skills with the public `skills` CLI:

```sh
vp dlx skills add tylergannon/agents -a codex --yes --all
```

## Skills

- `claude-codex-consensus` - run and reconcile independent Claude/Codex review.
- `customer-feedback-triage` - classify customer feedback into parsing, output, invariant, or HITL paths.
- `daily-docs-fold` - fold raw operating material into durable docs and skills.
- `debt-finder` - scout one material maintainability candidate.
- `dependency-upgrade-pr` - prove dependency upgrades through preview/runtime PRs.
- `e2e-coverage-triage` - route E2E scenarios to the right proof layer.
- `edit-attractor-pipelines` - design, edit, and prove Attractor DOT workflows.
- `preview-resend-verification` - verify preview auth emails through Resend.
- `proof-of-work` - prove application behavior through the running stack and own PR closeout.
- `session-worklog` - maintain tracked worklogs under `ephemeral/worklog/`.
- `supabase-query` - run narrow linked Supabase CLI reads from the current worktree.
- `svelte-component-factoring` - apply ownership-first Svelte component boundaries.
- `tech-debt-bounty` - orchestrate bounded material tech-debt scouting, consensus, fix, and proof.
- `write-prompts` - write, edit, and trim prompts or prompt-like skill prose.
