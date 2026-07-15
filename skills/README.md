# agents Skills

This directory is the source of published agent skills.

This directory contains only published portable skills. Repo-local maintainer
skills live in `../.agents/skills/`.

Install all published skills globally for Codex, Claude Code, and OpenCode:

```sh
vp dlx skills add tylergannon/agents \
  -g \
  -a codex -a claude-code -a opencode \
  -s '*' \
  -y
```
