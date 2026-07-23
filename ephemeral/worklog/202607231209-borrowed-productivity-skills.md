decision: Borrowed skill provenance lives in BORROWED_SKILLS.yaml with one source commit and explicit local-to-upstream mappings for every copied skill.
decision: The copied productivity subtree carries Matt Pocock's MIT license notice at skills/productivity/LICENSE because the upstream license requires the notice with substantial copies.
decision: The repo-local updater skill is canonical under .agents/skills and exposed to Claude through a relative symlink under .claude/skills.
decision: The user's explicit request to copy skills/productivity supersedes the existing borrow-map preference for submodules in this task; the pinned manifest makes the vendored exception visible and updateable.
friction: The repository root was already ahead of origin/main with an unrelated modified skill, and an initial nested worktree path appeared as untracked root state -> use an external task worktree based on freshly fetched origin/main.
