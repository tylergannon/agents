import { execFileSync } from "node:child_process";
import { existsSync, readdirSync, readFileSync, statSync } from "node:fs";
import path from "node:path";

const jsonFiles = [
  ".codex-plugin/plugin.json",
  ".claude-plugin/plugin.json",
  ".claude-plugin/marketplace.json",
  "package.json",
];

for (const file of jsonFiles) {
  JSON.parse(readFileSync(file, "utf8"));
}

function walk(dir) {
  if (!existsSync(dir)) return [];
  const files = [];
  for (const name of readdirSync(dir)) {
    const fullPath = path.join(dir, name);
    const relPath = fullPath.split(path.sep).join("/");
    if (statSync(fullPath).isDirectory()) {
      files.push(...walk(fullPath));
    } else {
      files.push(relPath);
    }
  }
  return files;
}

function isGenerated(file) {
  return file.split("/").includes("__pycache__") || /\.py[cod]$/.test(file);
}

const expectedSkillFiles = walk("skills").filter((file) => !isGenerated(file));
const packOutput = execFileSync("npm", ["pack", "--dry-run", "--json"], {
  encoding: "utf8",
  stdio: ["ignore", "pipe", "pipe"],
});
const [pack] = JSON.parse(packOutput);
const packed = new Set(pack.files.map((file) => file.path));
const errors = [];

for (const file of expectedSkillFiles) {
  if (!packed.has(file)) {
    errors.push(`missing from package: ${file}`);
  }
}

for (const file of packed) {
  if (isGenerated(file)) {
    errors.push(`generated file in package: ${file}`);
  }
  if (/^(\.agents|bin|ephemeral|submodules|tests)\//.test(file)) {
    errors.push(`private path in package: ${file}`);
  }
}

if (errors.length > 0) {
  for (const error of errors) {
    console.error(error);
  }
  process.exit(1);
}

console.log(`metadata and package contents ok (${pack.files.length} files)`);
