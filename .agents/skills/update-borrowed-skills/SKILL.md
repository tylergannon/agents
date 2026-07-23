---
name: update-borrowed-skills
description: >
  Update copied skills from upstream repositories and maintain
  borrowed-skills.json. Use when importing, refreshing, auditing, or changing
  the provenance of borrowed skills in this repository.
metadata:
  internal: true
---

# Update Borrowed Skills

## Workflow

1. Read `AGENTS.md` and `borrowed-skills.json`. Resolve every requested local
   path to one source entry. Complete this step when the scope and pinned commit
   are explicit.
2. Clone the source under `ephemeral/vendor/`, then fetch both the pinned commit
   and the target ref. Compare the local root with the pinned upstream
   root and compare the two upstream commits. Complete this step when every
   local divergence and upstream change is classified.
3. Stage the target subtree away from its destination. Replace the borrowed
   root as a whole, copy the source license to the manifest's local license
   path, and reapply only intentional downstream changes. Complete this step
   when every resulting file traces to the target commit or an approved local
   change.
4. Update `borrowed-skills.json` with the source repository, full target commit
   SHA, roots, license, and one local-to-upstream mapping for every borrowed
   `SKILL.md`. Synchronize package, plugin, and inventory text when the published
   surface changes. Complete this step when the manifest checker passes.
5. Compare the refreshed tree with the target checkout, inspect the diff, and
   run:

   ```sh
   node .agents/skills/update-borrowed-skills/scripts/check-manifest.mjs
   node scripts/validate-package.mjs
   git diff --check
   ```

   Remove the temporary clone. Complete the update when the copied tree,
   license, manifest, published package, and reported target SHA agree.
