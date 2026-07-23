#!/usr/bin/env node

import { existsSync, readdirSync, readFileSync, statSync } from "node:fs";
import path from "node:path";

const manifestPath = path.resolve(process.argv[2] ?? "borrowed-skills.json");
const repoRoot = path.dirname(manifestPath);
const errors = [];

function fail(message) {
  errors.push(message);
}

function localPath(relativePath) {
  return path.join(repoRoot, relativePath);
}

function collectSkills(relativeRoot) {
  const root = localPath(relativeRoot);
  if (!existsSync(root)) {
    fail(`missing local root: ${relativeRoot}`);
    return [];
  }

  const skills = [];
  const visit = (directory) => {
    for (const name of readdirSync(directory)) {
      const entry = path.join(directory, name);
      if (statSync(entry).isDirectory()) {
        visit(entry);
      } else if (name === "SKILL.md") {
        skills.push(
          path.relative(repoRoot, path.dirname(entry)).split(path.sep).join("/"),
        );
      }
    }
  };

  visit(root);
  return skills.sort();
}

let manifest;
try {
  manifest = JSON.parse(readFileSync(manifestPath, "utf8"));
} catch (error) {
  console.error(`invalid manifest: ${error.message}`);
  process.exit(1);
}

if (manifest.version !== 1) {
  fail("version must be 1");
}
if (!Array.isArray(manifest.sources) || manifest.sources.length === 0) {
  fail("sources must be a non-empty array");
}

const allLocalSkills = new Set();
let skillCount = 0;

for (const [sourceIndex, source] of (manifest.sources ?? []).entries()) {
  const label = `sources[${sourceIndex}]`;
  if (typeof source.repository !== "string" || source.repository.length === 0) {
    fail(`${label}.repository must be a non-empty string`);
  }
  if (
    typeof source.commit !== "string" ||
    !/^[0-9a-f]{40}$/.test(source.commit)
  ) {
    fail(`${label}.commit must be a full lowercase Git commit SHA`);
  }
  if (typeof source.localRoot !== "string" || source.localRoot.length === 0) {
    fail(`${label}.localRoot must be a non-empty string`);
  }
  if (
    typeof source.upstreamRoot !== "string" ||
    source.upstreamRoot.length === 0
  ) {
    fail(`${label}.upstreamRoot must be a non-empty string`);
  }
  if (!Array.isArray(source.skills) || source.skills.length === 0) {
    fail(`${label}.skills must be a non-empty array`);
    continue;
  }

  const declaredSkills = [];
  for (const [skillIndex, skill] of source.skills.entries()) {
    const skillLabel = `${label}.skills[${skillIndex}]`;
    if (typeof skill.local !== "string" || skill.local.length === 0) {
      fail(`${skillLabel}.local must be a non-empty string`);
      continue;
    }
    if (typeof skill.upstream !== "string" || skill.upstream.length === 0) {
      fail(`${skillLabel}.upstream must be a non-empty string`);
    }
    if (!skill.local.startsWith(`${source.localRoot}/`)) {
      fail(`${skillLabel}.local must be under ${source.localRoot}`);
    }
    if (!skill.upstream.startsWith(`${source.upstreamRoot}/`)) {
      fail(`${skillLabel}.upstream must be under ${source.upstreamRoot}`);
    }
    if (!existsSync(localPath(path.join(skill.local, "SKILL.md")))) {
      fail(`${skillLabel}.local has no SKILL.md`);
    }
    if (allLocalSkills.has(skill.local)) {
      fail(`duplicate local skill mapping: ${skill.local}`);
    }
    allLocalSkills.add(skill.local);
    declaredSkills.push(skill.local);
    skillCount += 1;
  }

  const actualSkills = collectSkills(source.localRoot);
  const declared = [...new Set(declaredSkills)].sort();
  for (const skill of actualSkills.filter((item) => !declared.includes(item))) {
    fail(`unmapped borrowed skill: ${skill}`);
  }
  for (const skill of declared.filter((item) => !actualSkills.includes(item))) {
    fail(`mapped skill missing from local root: ${skill}`);
  }

  if (source.license !== undefined) {
    if (
      typeof source.license.spdx !== "string" ||
      source.license.spdx.length === 0
    ) {
      fail(`${label}.license.spdx must be a non-empty string`);
    }
    if (
      typeof source.license.upstream !== "string" ||
      source.license.upstream.length === 0
    ) {
      fail(`${label}.license.upstream must be a non-empty string`);
    }
    if (
      typeof source.license.local !== "string" ||
      source.license.local.length === 0
    ) {
      fail(`${label}.license.local must be a non-empty string`);
    } else if (!existsSync(localPath(source.license.local))) {
      fail(`missing local license: ${source.license.local}`);
    }
  }
}

if (errors.length > 0) {
  for (const error of errors) {
    console.error(error);
  }
  process.exit(1);
}

console.log(
  `borrowed-skills manifest ok (${manifest.sources.length} sources, ${skillCount} skills)`,
);
