# nlspec Methodology — Project Journal

Session journal of observations and decisions. Started 2026-07-28.
Companion document: [methodology.md](methodology.md) coalesces these into a
single reasoned plan.

## Sources examined

- `skills/spec-writing/SKILL.md` (working draft, committed as 2a7fa0d)
- humanlayer `wsff.md` ("Why Software Factories Fail", Dex/HumanLayer)
- Two adversarial Fable subagent debates (ephemeral vs. durable technical
  specs), adjudicated in-session

## Observations

### On the existing spec-writing skill

- The notation half (RECORD / ENUM / INTERFACE / state machine / FUNCTION,
  language-neutral, "required semantics not prescribed implementation") is the
  strong, finished part. Protect it.
- Two documents are fused in one file: a stable notation **standard** and a
  HITL **process**. Different lifetimes, different readers.
- No stated test for whether a proposed seam decomposition is *good* — only
  requirements on how seams are written down.
- No mapping between definition-of-done claims and seams; the relation is what
  makes a spec agent-executable and tells a subagent what it may prove alone.
- The mechanical consistency pass (re-read whole spec, trace every claim to
  specified behavior) was dropped in the working-draft revision; it is exactly
  the work models do reliably and humans do badly.
- Stale reference: skill says `/grill-me`; the exposed skill is `/grilling`.

### On wsff

- Core argument accepted: RL scores pass/fail; maintainability has no fast
  oracle; therefore it is absent from the reward signal; harness engineering
  raises the floor but cannot move the ceiling. Sharpest line: "if a model
  could reliably tell good code from bad, it might have written the good
  version to begin with."
- Stolen outright: **vertical slices** (models default to horizontal
  stack-order plans; nothing is touchable until the end), **call-stack trees**
  and **file-tree diffs** as planning artifacts.
- Rejected: wsff's terminal move is "read the dang code" with human review
  bandwidth (100–200 lines) as the unit of control. That imports a bottleneck
  this methodology deliberately does not have — the steerer here is a frontier
  model.
- Observed: wsff's "program design" phase is partly a workaround for the
  absence of a real spec; much of it falls out of authoritative seam
  definitions.
- Observed: for a piece whose thesis is that maintainability lacks a
  mechanical oracle, wsff proposes no mechanical check at all. That gap is
  what this methodology targets.

## Decisions

1. **The substitution.** Replace the uncheckable objective (maintainability)
   with a checkable proxy: **seam integrity**. Four mechanical checks — (a)
   cell's actual public surface matches its stated contract; (b) nothing
   outside reaches past the contract; (c) cell is under its code-volume
   budget; (d) claims routed through the cell are demonstrated. Shotgun
   surgery — wsff's named failure — *is* a seam failure, so quality is not
   conceded, it is attacked through a channel that has a verifier.

2. **Budget is a tripwire, not an implementer rule.** Enforced splitting by
   the implementer produces splits at arbitrary places (more seams than
   concepts). On breach: stop, report; the spec (HITL + frontier model)
   decides where the split goes. Decomposition authority stays in the spec.

3. **Seam-quality criterion: the Parnas test.** A seam belongs where the
   *reason to change* differs on either side. What changes together lives
   together. This is the one piece of quality theory that is load-bearing and
   agent-legible.

4. **Claims are classified** cell-local vs. integration-level. The
   classification doubles as a diagnostic on the decomposition (if most claims
   are integration-only, the decomposition bought nothing).

5. **No-amendment rule.** Implementers never amend the nlspec. On hitting an
   underivable decision they stop and file a spec defect; the frontier
   model/HITL adjudicates; the answer lands in the nlspec; work resumes.

6. **The technical spec is ephemeral — ephemeral-as-authority, not
   ephemeral-as-deleted.** Program-design detail (types, signatures,
   call-stack trees, file layouts) is regenerated per work slice from
   nlspec + current code, authoritative for that slice only, then archived in
   the PR as a historical record — never maintained. Grounds:
   - Maintenance economics of docs sized at a large fraction of the code are
     disqualifying; the shadow copy has no compiler and rots.
   - For an LLM reader a confident stale document *is evidence*; a wrong
     signpost is strictly worse than none.
   - In a repo, durability is authority: durably documenting cell internals
     hoists them above the authority line and defeats the fungibility premise.
   - Regeneration per slice continuously tests the nlspec's spanning.

7. **The lifetime rule.** The ephemeral slice spec may only contain decisions
   whose lifetime ≤ the slice. Any decision a future slice must agree with is
   by that property seam-level and must be promoted into the nlspec before the
   slice ships. Discarding becomes safe by construction.

8. **Two cell modes; no per-edit regeneration.** (User correction to the
   "churn engine" and "nondeterminism" objections, which assumed
   regeneration-per-edit.)
   - *Maintenance mode* (common): ordinary contract-guarded agent work —
     "fix this; don't touch the contract; if the fix forces contract
     re-evaluation, stop and report."
   - *Rebuild mode* (rare, deliberate): the cell is **condemned** and rebuilt
     from its contract, conditioned on surviving elements — never a fresh
     stochastic sample.

9. **Replace economics, not detect-and-repair economics.** (User correction to
   the "rationalization" objection.) No drift *baseline* is needed, only drift
   *triggers*: budget breach, contract-check failure, claim no longer
   demonstrable, supposedly cell-local change spraying across seams, plus a
   *friction* trigger (repeated failed local fixes). On trigger: condemn and
   rebuild. Archived slice specs are forensics, not load-bearing.

10. **Budget is the primary drift tripwire.** Degraded internals show up as
    volume first (duplication, defensive scaffolding, flag soup). Check it
    mechanically on every slice; keep it slightly tight on purpose so the wire
    trips while the rebuild is still cheap.

11. **The bullet bitten openly.** Internal degradation that breaches no
    contract, fails no claim, and stays under budget is defined as
    not-a-problem until it trips a wire. If it matters (e.g. complexity), it
    should have been a measurable claim. The budget cap bounds the eventual
    rebuild cost.

12. **Decision records (ADRs) dropped.** Replaced by the **catalytic
    converter**: test cases discovered in code are evaluated for elevation
    into claims, up to inclusion in the nlspec. Promoted tests are *checkable*
    memory; rationale prose is unverifiable and slowly lies. Converter intake
    is behavioral evidence only — no opinions, no "permanent language
    lessons" (slippery slope, rejected by user).

13. **Minimal ≠ thin: the basis-vector metaphor is canonical.** "Minimally
    spanning" is borrowed from basis vectors: the whole space is
    constructible from the materials given plus the rules of engagement.
    Membership test: a decision belongs in the nlspec iff it is not derivable
    from what is already there plus the rules of engagement. Detector:
    two capable readers would derive different answers *and* the difference
    matters at a seam, claim, budget, or exclusion.

14. **Spanning is maintained over the document's lifetime, not achieved at
    authoring.** Missing basis vectors are discovered by attempting
    construction — the cheap way. "Answer all hard questions up front" is a
    named pitfall.

15. **Stopping rule for authoring.** Keep answering questions only while the
    answers land in the nlspec's own sections (seams, claims, budgets,
    exclusions). Operationally: stop after two consecutive grill rounds that
    surface only derivable material. Deferred hard questions are listed in a
    short **deferred-decisions** section — a tripwire list converting unknown
    spanning failures (silent improvisation) into known ones (immediate
    escalation).

16. **The two pumps.** The nlspec has exactly two intake channels:
    spec-defect escalation (decisions, from blocked slices) and the catalytic
    converter (behavioral evidence, from code/tests). Authority flows only
    downward; content flows only upward through these pumps. Nothing else
    writes to the nlspec.

17. **Artifact hierarchy** (final): nlspec (permanent, maintained, sole
    authority) / per-slice tech spec (one slice, then archived in PR) /
    code + tests (permanent fact; tests double as candidate claims). Routing
    question for any content: *how long must someone agree with this?*

18. **Rebuild protocol order is fixed: harvest first.** Run the converter over
    the dying cell's tests before condemning — a condemned cell's tests are
    the only record of everything its ephemeral specs never wrote down. Then
    condemn → rebuild from contract → demonstrate claims.

19. **Skill suite axis: noun vs. verbs**, not wsff's four phases. One standard
    (reference, loaded by all) plus verb skills for authoring, review, slice
    design, and cell lifecycle. Verb skills stay short because the noun skill
    carries the vocabulary.

## Open questions

- Whether slice-design and cell-lifecycle eventually merge or split further
  (converter and condemnation could be argued as separate verbs).
- Concrete budget metric (raw LOC vs. language-weighted) and default numbers.
- How the archived slice spec is attached (PR body vs. artifact upload via
  proof-uploader conventions).
- Whether `spec-review` should always invoke `/request-adversarial-review`
  for material specs or only on request.
