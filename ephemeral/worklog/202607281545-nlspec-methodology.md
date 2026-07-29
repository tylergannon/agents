# nlspec methodology — spec-writing suite

decision: technical specs are ephemeral-as-authority (per-slice, archived in
PR, never maintained); nlspec + code are the only durable authorities. Source:
adversarial Fable debate adjudicated with HITL; full record in
ephemeral/projects/nlspec-methodology/journal.md.

decision: replace economics over detect-and-repair — drift triggers (budget
breach, contract failure, undemonstrable claim, seam spray, friction) condemn
a cell for rebuild; no drift baselines. Budget is the primary tripwire, kept
slightly tight. Source: HITL.

decision: ADR-style decision records rejected (unverifiable prose that slowly
lies); replaced by the catalytic converter — tests elevated to claims,
behavioral evidence only. Source: HITL, explicit slippery-slope ruling on
"permanent language lessons."

correction: "minimally spanning" is a basis-vectors metaphor; minimal ≠ thin.
"Answer all hard questions up front" is a named pitfall — spanning is
maintained over the spec's lifetime via two pumps (spec-defect escalation,
catalytic converter), with a stopping rule (two grill rounds surfacing only
derivable material) and a deferred-decisions tripwire list.

doc_bug: skills/spec-writing referenced `/grill-me`; the exposed skill is
`/grilling` -> fixed in the spec-authoring rewrite.

skill_issue: spec-writing source=repo severity=design -> fused a stable
notation standard with a HITL process in one file; split into spec-writing
(standard + notation.md) plus verb skills spec-authoring, spec-review,
slice-design, cell-lifecycle.

decision: the mechanical consistency pass dropped from the spec-writing draft
was restored as its own skill (spec-review); it is the work models do reliably
and humans do badly.

correction: HITL hard-rejected TaskNodeBackend Checkpoint()/Restore() (attractor
spec 4.5). Checkpointing is a workflow-engine concept only; harness sessions
are natively durable (that is the premise of building on full harnesses), and
the thread binding table should be write-through durable in the run directory
at bind time, not snapshot/restored through an opaque blob in checkpoint.json.
I defended the spec text instead of re-deriving — the exact "confident stale
document as evidence" failure the methodology names. Re-derive before
defending; the spec author is the authority, not the spec draft.
