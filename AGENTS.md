# agents

- This repository publishes reusable agent skills and Go tools that support them.
- `skills/` is the canonical source for published skills.
- `.agents/skills/` contains maintainer-only skills and is not distributed.
- Keep skills concise, project-portable, and free of product-specific policy.
- Put a skill's scripts and references beside that skill.
- Put shared Go commands in `cmd/` and reusable Go packages in `internal/` or `pkg/`.
- Add tools only when deterministic code is more reliable than prompt instructions.
- Use focused checks that prove the changed behavior; do not treat a generic test command as sufficient proof.
- Keep package and plugin metadata synchronized when published contents change.
- Preserve licenses and attribution when adapting external work.
- Keep optional sources for skill research in `SKILL_REPOSITORIES.md`; consult them only when useful.
