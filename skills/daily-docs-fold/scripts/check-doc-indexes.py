#!/usr/bin/env python3
"""Check documentation index quality in a target repository.

This script is packaged as a reusable skill resource for documentation index
checks. It performs structural routing checks,
rough token-budget checks, source-handle checks, and optional design-probe
checks when a target repository supplies a probe manifest. It does not prove
that documentation claims are true.
"""

from __future__ import annotations

import argparse
import datetime as dt
import json
import math
import re
import sys
from pathlib import Path
from urllib.parse import unquote, urlsplit


PATH_TOKEN = (
    r"(?:\.?[A-Za-z0-9_-]+/)*\.?[A-Za-z0-9_-]+\."
    r"(?:go|js|ts|tsx|jsx|py|rb|rs|java|kt|md|yaml|yml|toml|json|jsonl|sh|sql)"
    r"|(?:\.?[A-Za-z0-9_-]+/)*Dockerfile"
    r"|Makefile|Justfile|package\.json|pnpm-lock\.yaml|\.rgignore"
)
SOURCE_HANDLE_RE = re.compile(
    rf"(?<![A-Za-z0-9_./:-])({PATH_TOKEN})(?![:A-Za-z0-9_./-])"
)
LINE_CITE_RE = re.compile(
    rf"(?<![A-Za-z0-9_./:-])({PATH_TOKEN}):(\d+)(?:-(\d+))?"
)
CITATION_GROUP_RE = re.compile(r"\(([^()\n]*(?:" + PATH_TOKEN + r")[^()\n]*)\)")
MARKDOWN_LINK_RE = re.compile(r"\[[^\]]+\]\(\s*([^)\s]+)(?:\s+[^)]*)?\)")
INDEX_FILENAMES = ("index.md", "README.md")
ROOT_DOCS_INDEX_CANDIDATES = tuple(f"docs/{name}" for name in INDEX_FILENAMES)
DOCS_INDEX_EXEMPT_ROUTES = {*ROOT_DOCS_INDEX_CANDIDATES, "docs/journal.md"}

def read(path: Path) -> str:
    return path.read_text(encoding="utf-8")


def token_estimate(text: str) -> int:
    return math.ceil(len(text) / 4)


def source_handles(text: str) -> list[str]:
    found: list[str] = []
    for group in CITATION_GROUP_RE.finditer(text):
        found.extend(match.group(1) for match in SOURCE_HANDLE_RE.finditer(group.group(1)))
    return found


def line_citations(text: str) -> list[str]:
    return [match.group(0) for match in LINE_CITE_RE.finditer(text)]


def rel(path: Path, repo: Path) -> str:
    try:
        return str(path.relative_to(repo))
    except ValueError:
        return str(path)


def repo_rel(repo: Path, path: Path) -> str | None:
    try:
        return path.resolve().relative_to(repo.resolve()).as_posix()
    except ValueError:
        return None


def route_exists(repo: Path, route: str) -> bool:
    return (repo / route).exists()


def find_docs_index(repo: Path) -> tuple[Path | None, str | None]:
    for route in ROOT_DOCS_INDEX_CANDIDATES:
        candidate = repo / route
        if candidate.exists():
            return candidate, route
    return None, None


def find_index_file(directory: Path) -> tuple[Path | None, str | None]:
    for name in INDEX_FILENAMES:
        candidate = directory / name
        if candidate.exists():
            return candidate, name
    return None, None


def docs_child_routes(repo: Path) -> list[str]:
    docs = repo / "docs"
    if not docs.exists():
        return []
    routes: list[str] = []
    for child in sorted(docs.iterdir()):
        if child.is_dir():
            _, index_name = find_index_file(child)
            if index_name:
                routes.append(f"docs/{child.name}/{index_name}")
        elif child.is_file() and child.suffix == ".md":
            route = f"docs/{child.name}"
            if route not in DOCS_INDEX_EXEMPT_ROUTES:
                routes.append(route)
    return routes


def root_routes(repo: Path, required_routes: list[str]) -> list[str]:
    routes = list(required_routes)
    for route in docs_child_routes(repo):
        if route not in routes:
            routes.append(route)
    return routes


def resolve_source_path(repo: Path, owner: Path, source: str) -> Path | None:
    candidates = [repo / source, owner.parent / source]
    for candidate in candidates:
        if candidate.exists():
            return candidate
    return None


def markdown_files(root: Path) -> list[Path]:
    if not root.exists():
        return []
    return sorted(path for path in root.rglob("*.md") if path.is_file())


def clean_link_destination(destination: str) -> str | None:
    destination = destination.strip().strip("<>")
    parsed = urlsplit(destination)
    if parsed.scheme or parsed.netloc or not parsed.path:
        return None
    return unquote(parsed.path)


def linked_repo_routes(repo: Path, owner: Path, text: str) -> set[str]:
    routes: set[str] = set()
    for match in MARKDOWN_LINK_RE.finditer(text):
        destination = clean_link_destination(match.group(1))
        if destination is None:
            continue
        base = repo if destination.startswith("/") else owner.parent
        target = (base / destination.lstrip("/")).resolve()
        target_rel = repo_rel(repo, target)
        if target_rel is None:
            continue
        routes.add(target_rel)
        if target.is_dir():
            _, index_name = find_index_file(target)
            if index_name:
                routes.add(f"{target_rel}/{index_name}" if target_rel else index_name)
        else:
            for index_name in INDEX_FILENAMES:
                if target_rel.endswith(f"/{index_name}"):
                    routes.add(target_rel.removesuffix(f"/{index_name}"))
    return routes


def text_without_link_destinations(text: str) -> str:
    return MARKDOWN_LINK_RE.sub(lambda match: match.group(0).split("](", 1)[0], text)


def route_text_tokens(route: str) -> set[str]:
    token = route.strip("/")
    tokens = {token}
    for index_name in INDEX_FILENAMES:
        if token.endswith(f"/{index_name}"):
            tokens.add(token.removesuffix(index_name))
            tokens.add(token.removesuffix(f"/{index_name}"))
    return {item for item in tokens if item}


def route_keys(repo: Path, route: str) -> set[str]:
    target = (repo / route).resolve()
    target_rel = repo_rel(repo, target)
    if target_rel is None:
        return {route.strip("/")}
    keys = {target_rel}
    if target.is_dir():
        _, index_name = find_index_file(target)
        if index_name:
            keys.add(f"{target_rel}/{index_name}" if target_rel else index_name)
    for index_name in INDEX_FILENAMES:
        if target_rel.endswith(f"/{index_name}"):
            keys.add(target_rel.removesuffix(f"/{index_name}"))
    return keys


def source_check_files(repo: Path, args: argparse.Namespace) -> list[Path]:
    files: list[Path] = []
    source_dirs = list(args.source_dir)
    for default_dir in ("docs/policy", "docs/manuals"):
        if (repo / default_dir).exists() and default_dir not in source_dirs:
            source_dirs.append(default_dir)
    if args.design_dir:
        source_dirs.append(args.design_dir)
    for source_dir in source_dirs:
        files.extend(
            path for path in markdown_files(repo / source_dir)
            if rel(path, repo) not in DOCS_INDEX_EXEMPT_ROUTES
        )
    by_path = {path.resolve(): path for path in files}
    return sorted(by_path.values())


def check_source_handles(repo: Path, files: list[Path]) -> list[str]:
    errors: list[str] = []
    for file_path in files:
        text = read(file_path)
        for line_ref in line_citations(text):
            errors.append(
                f"{rel(file_path, repo)} contains durable file:line citation {line_ref}; "
                "use a stable source handle"
            )
        for source in source_handles(text):
            if resolve_source_path(repo, file_path, source) is None:
                errors.append(f"{rel(file_path, repo)} cites missing source handle {source}")
    return errors


def check_docs_index(repo: Path, limit: int, required_routes: list[str]) -> tuple[list[str], dict]:
    errors: list[str] = []
    docs = repo / "docs"
    docs_index, docs_index_route = find_docs_index(repo)
    stats = {
        "present": docs_index is not None,
        "path": docs_index_route,
        "tokens_est": 0,
        "required_routes": [],
    }

    if not docs.exists():
        return errors, stats
    if docs_index is None or docs_index_route is None:
        routed_docs = docs_child_routes(repo)
        if routed_docs:
            errors.append(
                "docs/ has routed documentation but no root docs index "
                "(docs/index.md or docs/README.md) is present: "
                + ", ".join(routed_docs)
            )
        return errors, stats

    text = read(docs_index)
    stats["tokens_est"] = token_estimate(text)
    if stats["tokens_est"] >= limit:
        errors.append(f"{docs_index_route} too large: {stats['tokens_est']} tokens_est >= {limit}")

    required = root_routes(repo, required_routes)
    stats["required_routes"] = required
    links = linked_repo_routes(repo, docs_index, text)
    route_text = text_without_link_destinations(text)
    for route in required:
        if route in ROOT_DOCS_INDEX_CANDIDATES:
            continue
        has_link = bool(route_keys(repo, route) & links)
        has_route_text = any(token in route_text for token in route_text_tokens(route))
        if not has_link and not has_route_text:
            errors.append(f"{docs_index_route} does not route to {route}")
    return errors, stats


def load_probes(path: Path | None) -> list[dict]:
    if path is None or not path.exists():
        return []
    return json.loads(read(path))


def default_packet_dir(repo: Path) -> Path:
    return repo / "ephemeral" / "docs-index-packets"


def verify_design_index(repo: Path, args: argparse.Namespace) -> tuple[list[str], dict]:
    errors: list[str] = []
    if not args.design_dir:
        return errors, {"present": False, "tokens_est": 0, "probes": []}

    design = repo / args.design_dir
    design_index, design_index_name = find_index_file(design)
    design_index_route = f"{args.design_dir}/{design_index_name}" if design_index_name else None
    stats = {"present": design_index is not None, "path": design_index_route, "tokens_est": 0, "probes": []}

    if not design.exists():
        return errors, stats
    if design_index is None or design_index_route is None:
        errors.append(
            f"{args.design_dir}/ exists but no index file "
            f"({', '.join(INDEX_FILENAMES)}) is present"
        )
        return errors, stats

    design_index_text = read(design_index)
    stats["tokens_est"] = token_estimate(design_index_text)
    if stats["tokens_est"] >= args.design_index_token_limit:
        errors.append(
            f"{design_index_route} too large: {stats['tokens_est']} "
            f"tokens_est >= {args.design_index_token_limit}"
        )

    second_layer = [repo / route for route in args.second_layer]
    stats["second_layer_tokens_est"] = {
        rel(path, repo): token_estimate(read(path)) for path in second_layer
    }
    for path in second_layer:
        second_tokens = token_estimate(read(path))
        if second_tokens >= stats["tokens_est"]:
            errors.append(
                f"second-layer {rel(path, repo)} not smaller than design index: "
                f"{second_tokens} >= {stats['tokens_est']}"
            )

    leaves = design / "leaves"
    for leaf in markdown_files(leaves):
        leaf_tokens = token_estimate(read(leaf))
        if leaf_tokens > args.leaf_token_limit:
            errors.append(
                f"leaf {rel(leaf, repo)} too large for whole-leaf retrieval: "
                f"{leaf_tokens} > {args.leaf_token_limit}"
            )

    probe_path = Path(args.probes).resolve() if args.probes else None
    probe_results: list[dict] = []
    for probe in load_probes(probe_path):
        target = repo / probe["target_file"]
        result = {"id": probe["id"], "target": probe["target_file"], "status": "PASS", "missing": []}
        if not target.exists():
            errors.append(f"probe {probe['id']} target missing: {probe['target_file']}")
            result["status"] = "FAIL"
            probe_results.append(result)
            continue
        if probe["target_file"] not in design_index_text:
            errors.append(f"probe {probe['id']} target not routed from {design_index_route}")
            result["status"] = "FAIL"

        target_text = read(target)
        combined = (design_index_text + "\n" + target_text).lower()
        for term in probe.get("expected_terms", []):
            if term.lower() not in combined:
                result["missing"].append(term)
                errors.append(f"probe {probe['id']} missing expected term: {term}")
        if result["missing"]:
            result["status"] = "FAIL"

        cite_count = len(source_handles(target_text))
        if cite_count < int(probe.get("min_citations", 0)):
            errors.append(f"probe {probe['id']} target has {cite_count} source handles")
            result["status"] = "FAIL"
        result["citations"] = cite_count
        result["tokens_design_index_plus_target"] = token_estimate(design_index_text + "\n" + target_text)

        if args.write_packets:
            packet_dir = Path(args.packet_dir).resolve() if args.packet_dir else default_packet_dir(repo)
            packet_dir.mkdir(parents=True, exist_ok=True)
            packet = packet_dir / f"{probe['id']}.md"
            packet.write_text(
                "\n".join(
                    [
                        f"# Probe Packet: {probe['id']}",
                        "",
                        f"class: {probe.get('class', '')}",
                        f"question: {probe.get('question', '')}",
                        f"target_file: {probe['target_file']}",
                        "",
                        "constraint: a blind design probe may use only the design index and target file content below.",
                        "",
                        "## Design Index",
                        design_index_text,
                        "",
                        "## Target Leaf",
                        target_text,
                    ]
                )
                + "\n",
                encoding="utf-8",
            )

        probe_results.append(result)

    stats["probes"] = probe_results
    return errors, stats


def default_run_dir(repo: Path) -> Path:
    return repo / "ephemeral" / "docs-index-runs"


def write_run(repo: Path, label: str, errors: list[str], stats: dict, run_dir_arg: str | None) -> Path:
    run_dir = Path(run_dir_arg).resolve() if run_dir_arg else default_run_dir(repo)
    run_dir.mkdir(parents=True, exist_ok=True)
    path = run_dir / f"{label}.md"
    status = "PASS" if not errors else "FAIL"
    lines = [
        f"# Docs Index Check: {label}",
        "",
        f"timestamp_utc: {dt.datetime.now(dt.timezone.utc).isoformat()}",
        f"repo: {repo}",
        f"status: {status}",
        "",
        "## Stats",
        "",
        "```json",
        json.dumps(stats, indent=2, sort_keys=True),
        "```",
        "",
        "## Errors",
        "",
    ]
    if errors:
        lines.extend(f"- {err}" for err in errors)
    else:
        lines.append("- none")
    path.write_text("\n".join(lines) + "\n", encoding="utf-8")
    return path


def verify(repo: Path, args: argparse.Namespace) -> tuple[list[str], dict]:
    errors: list[str] = []
    docs_errors, docs_stats = check_docs_index(repo, args.docs_index_token_limit, args.required_route)
    source_errors = check_source_handles(repo, source_check_files(repo, args))
    design_errors, design_stats = verify_design_index(repo, args)
    errors.extend(docs_errors)
    errors.extend(source_errors)
    errors.extend(design_errors)
    stats = {
        "repo": str(repo),
        "docs_index": docs_stats,
        "design_index": design_stats,
    }
    return errors, stats


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--repo", default=".", help="target repository root")
    parser.add_argument("--stats-json", action="store_true", help="print stats JSON")
    parser.add_argument("--write-run", metavar="LABEL", help="write a run artifact")
    parser.add_argument("--run-dir", help="directory for --write-run output")
    parser.add_argument("--write-packets", action="store_true", help="write design probe packets")
    parser.add_argument("--packet-dir", help="directory for --write-packets output")
    parser.add_argument("--probes", help="path to design probe manifest JSON")
    parser.add_argument("--required-route", action="append", default=[], help="extra route required in the root docs index")
    parser.add_argument("--design-dir", help="design index directory relative to repo")
    parser.add_argument(
        "--source-dir",
        action="append",
        default=[],
        help="durable docs directory to check for source handles; repeatable",
    )
    parser.add_argument(
        "--second-layer",
        action="append",
        default=[],
        help="second-layer design index file relative to repo; repeatable",
    )
    parser.add_argument("--docs-index-token-limit", type=int, default=20_000)
    parser.add_argument("--design-index-token-limit", type=int, default=20_000)
    parser.add_argument("--leaf-token-limit", type=int, default=1_500)
    args = parser.parse_args()

    repo = Path(args.repo).resolve()
    if not repo.exists():
        print(f"ERROR: repo does not exist: {repo}", file=sys.stderr)
        return 2

    errors, stats = verify(repo, args)
    if args.write_run:
        run_path = write_run(repo, args.write_run, errors, stats, args.run_dir)
        print(f"wrote {rel(run_path, repo)}")
    if args.stats_json:
        print(json.dumps(stats, indent=2, sort_keys=True))
    if errors:
        for err in errors:
            print(f"ERROR: {err}", file=sys.stderr)
        return 1
    print("docs index verification PASS")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
