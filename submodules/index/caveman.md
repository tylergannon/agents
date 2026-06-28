# caveman

- Path: `../caveman`
- Upstream: `JuliusBrussee/caveman`
- Pinned commit: `655b7d9c5431f822264b7732e9901c5578ac84cf`
- License: MIT
- Package: `caveman-installer` `0.1.0`

## What Is In It

- A terse communication skill/plugin for Claude Code, Codex, Gemini, Cursor,
  OpenCode, OpenClaw, and many other agent hosts.
- Installers: `install.sh`, `install.ps1`, `bin/install.js`.
- Skill families: `skills/caveman`, `skills/caveman-compress`,
  `skills/caveman-stats`, `skills/caveman-commit`, `skills/caveman-review`,
  `skills/cavecrew`, and host-specific plugin copies.
- Agent files: `agents/cavecrew-builder.md`,
  `agents/cavecrew-investigator.md`, `agents/cavecrew-reviewer.md`.
- Commands: TOML commands under `commands/` and OpenCode command markdown under
  `src/plugins/opencode/commands/`.
- Runtime helpers: Claude hooks under `src/hooks/` and an MCP tool wrapper under
  `src/mcp-servers/caveman-shrink/`.
- Proof surfaces: benchmark prompts, eval scripts, compression safety tests,
  installer tests, hook tests, and repository verification.

Local scan: 145 files, 15 `SKILL.md` files, 7 agent files, 9 command files, 23
test files, 3 installer/bin files.

## What We Can Learn

- Brevity is more durable when packaged as a skill, commands, hooks, tests, and
  stats rather than as a one-line instruction.
- Memory compression needs invariants. caveman explicitly preserves code, URLs,
  and paths and carries tests for compressed memory files.
- Cross-agent support benefits from a single installer that detects host layouts
  and writes the right artifact shape for each host.
- A token-savings feature is easier to trust when it ships benchmark scripts and
  raw result snapshots.

## What We Can Use

- `skills/caveman-compress/scripts/` as a reference for safe memory-file
  compression.
- `bin/install.js` and `bin/lib/settings.js` as a reference for host detection
  and idempotent installer behavior.
- `src/hooks/` as a reference for session-state activation and statusline
  integration.
- `src/mcp-servers/caveman-shrink/` as a reference for wrapping tool metadata.
- `tests/` and `benchmarks/` as a reference for proving a communication skill is
  doing what it claims.

## What We Can Borrow

- A small `agents` terse-mode skill for agents that need compact status
  updates.
- A safe memory compression helper, but only with byte-preservation tests for
  code fences, paths, URLs, commands, and structured data.
- An installer pattern that writes host-specific artifacts from one source tree.
- A stats command pattern that turns invisible context/token savings into
  visible proof.

## Cautions

- Do not compress canonical repo policy files unless the repo owner asks for it.
- Do not copy the style wholesale into docs that need normal prose.
- Preserve MIT license notice if copying substantial scripts or prompt text.
