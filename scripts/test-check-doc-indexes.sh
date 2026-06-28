#!/usr/bin/env bash
set -euo pipefail

tmpdir="$(mktemp -d)"
trap 'rm -rf "$tmpdir"' EXIT

checker="${CHECK_DOC_INDEXES:-skills/daily-docs-fold/scripts/check-doc-indexes.py}"

readme_repo="$tmpdir/readme-router"
mkdir -p "$readme_repo/docs/manuals"
cat >"$readme_repo/docs/README.md" <<'MD'
# Docs

- [Manuals](manuals/)
MD
cat >"$readme_repo/docs/manuals/index.md" <<'MD'
# Manuals
MD

python3 "$checker" --repo "$readme_repo" --stats-json >"$tmpdir/readme-router.out"
grep -q '"path": "docs/README.md"' "$tmpdir/readme-router.out"

index_repo="$tmpdir/index-router"
mkdir -p "$index_repo/docs/manuals"
cat >"$index_repo/docs/index.md" <<'MD'
# Docs

- [Manuals](manuals/)
MD
cat >"$index_repo/docs/manuals/index.md" <<'MD'
# Manuals
MD

python3 "$checker" --repo "$index_repo" --stats-json >"$tmpdir/index-router.out"
grep -q '"path": "docs/index.md"' "$tmpdir/index-router.out"

design_readme_repo="$tmpdir/design-readme-router"
mkdir -p "$design_readme_repo/docs/design/leaves"
cat >"$design_readme_repo/docs/README.md" <<'MD'
# Docs

- [Design](design/)
MD
cat >"$design_readme_repo/docs/design/README.md" <<'MD'
# Design

- [Leaf](leaves/one.md)
MD
cat >"$design_readme_repo/docs/design/leaves/one.md" <<'MD'
# One
MD

python3 "$checker" --repo "$design_readme_repo" --design-dir docs/design --stats-json >"$tmpdir/design-readme-router.out"
grep -q '"path": "docs/design/README.md"' "$tmpdir/design-readme-router.out"

missing_repo="$tmpdir/missing-router"
mkdir -p "$missing_repo/docs/manuals"
cat >"$missing_repo/docs/manuals/index.md" <<'MD'
# Manuals
MD

if python3 "$checker" --repo "$missing_repo" >"$tmpdir/missing-router.out" 2>"$tmpdir/missing-router.err"; then
  echo "expected missing root docs index check to fail" >&2
  exit 1
fi

grep -q 'no root docs index' "$tmpdir/missing-router.err"

echo "docs index checker tests ok"
