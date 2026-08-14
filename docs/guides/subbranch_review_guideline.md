# Subbranch Review Guideline

## Purpose
This review is mandatory before a completed ticket branch is merged into its parent version branch. The goal is to detect implementation mistakes, design drift, and hidden quality problems before the code is accepted upstream.

This review must be adversarial, explicit, and actionable. It is not a formality. It exists to catch issues that a hasty implementation or an AI agent may miss.

The review is designed to be repeated up to 10 times for the same subbranch. Each pass must produce concrete action items, and the branch must not be merged until all blocking issues are fixed.

---

## Scope
This review applies to any branch of the form:

- `version/xy-release/<ticket-slug>`

and is performed before the ticket branch is squash-merged into:

- `version/xy-release`

The reviewer must treat the code as if it were being audited for a real release candidate.

---

## Required review model
The review must be conducted in up to 10 passes. Each pass is a separate review cycle. The reviewer should not stop after the first pass unless the branch is already perfect.

### Pass 1: project integrity
Check whether the branch still matches the repo contract.

Questions to answer:
- Does the branch still compile?
- Does the repo still have the expected project structure?
- Are files missing, duplicated, or incorrectly placed?
- Has the agent changed unrelated files?

Fix rule:
- If a file or package does not fit the project structure, restore it to the expected layout before continuing.

### Pass 2: requirement alignment
Compare the implementation to the milestone and ticket descriptions.

Questions to answer:
- Did the implementation actually solve the ticket goal?
- Did the agent add unrelated scope?
- Did the agent skip an explicit requirement?
- Did the agent interpret the issue incorrectly?

Fix rule:
- If the code does not satisfy the ticket requirements, remove or rewrite the incorrect logic until the requirement is truly met.

### Pass 3: correctness of behavior
Check the actual runtime behavior, not only the code shape.

Questions to answer:
- Does the CLI do what it is supposed to do?
- Are CLI flags correct?
- Do stdin, file input, file output, and stdout behave as expected?
- Are empty-input, malformed-input, and edge-case scenarios handled safely?

Fix rule:
- If the behavior is broken, fix the logic and add or update tests that prove the correct behavior.

### Pass 4: test and regression review
Check whether the branch includes proof that the behavior is correct.

Questions to answer:
- Are tests present for the changed behavior?
- Are they user-facing and realistic?
- Do they cover both success and failure cases?
- Does the branch avoid test-only shortcuts or irrelevant mocks?

Fix rule:
- If there is no test coverage for the changed behavior, add a test or evidence-based verification command.

### Pass 5: code quality and maintainability
Evaluate if the code is readable, reviewable, and not overly clever.

Questions to answer:
- Is the implementation understandable?
- Is there dead code, confusing naming, duplication, or unnecessary complexity?
- Is the logic consistent with the rest of the repository?
- Is the code too brittle or too magical?

Fix rule:
- Simplify the code. Remove accidental complexity, duplicate logic, or unreadable abstractions.

### Pass 6: security and correctness review
Evaluate the code for hidden risk.

Questions to answer:
- Does the implementation handle secrets or config safely?
- Does it leak credentials, tokens, or local state unexpectedly?
- Does it silently ignore failures?
- Does it crash on invalid input or partial data?

Fix rule:
- Remove any unsafe behavior, hard-coded secrets, or silent failure paths.

### Pass 7: release readiness review
Check if the result is ready to be merged into the version branch and then later into `main`.

Questions to answer:
- Does the branch build successfully?
- Does the branch pass the relevant tests?
- Is the output aligned with the README and project docs?
- Is the scope still narrow enough for the ticket?

Fix rule:
- If the branch does not build or test cleanly, fix the failing issue before merging.

### Pass 8: API and contract review
Check boundary conditions and compatibility.

Questions to answer:
- Do the interfaces remain stable?
- Do existing public behaviors remain compatible?
- Did the agent modify common logic in a way that could break neighboring tickets?

Fix rule:
- Restore compatibility and avoid broad-breaking changes unless the issue explicitly requires them.

### Pass 9: documentation and handoff review
Check if the branch leaves enough evidence for the next agent or reviewer.

Questions to answer:
- Is the implementation understandable from the code alone?
- Is the next handoff clear?
- Did the agent leave a trace of the decisions and blockers?
- Are README or docs updates required or missing?

Fix rule:
- Add concise documentation or comments only where they materially help the next reader.

### Pass 10: final gate
This is the final blocking review before approval.

The reviewer must answer only one question:
- Is this branch safe to merge into the version branch?

The answer must be yes only if:
- all required tests pass
- all blocking findings from passes 1-9 are fixed
- no unknown or unverified risks remain
- the ticket goal is fully satisfied
- the implementation is consistent with the project architecture

If the answer is no, the branch is rejected and must be fixed before merge.

---

## Mandatory review output format
The reviewer must produce an explicit review result. It must not be vague.

Use this format:

```md
# Subbranch Review Result

## Branch
- <branch-name>

## Verdict
- PASS | FAIL | REJECTED

## Blockers
- <issue 1>
- <issue 2>

## Required fixes
1. <fix>
2. <fix>

## Verification commands
- <command 1>
- <command 2>

## Passed checks
- <check 1>
- <check 2>

## Remaining risk
- <brief description>
```

A review is not complete unless it includes all of the above.

---

## Hard rules
The following rules are mandatory:

1. The reviewer must be critical. Assume the code may be wrong until proven otherwise.
2. The review must check real behavior, not only naming or structure.
3. The review must call out missing tests or missing verification.
4. The review must prescribe explicit fix actions, not vague statements like "improve quality".
5. The review must stop the merge if any blocker remains.
6. The reviewer should not accept a branch just because it looks plausible.
7. The agent must fix all blockers before requesting approval.
8. A branch may be reviewed up to 10 times, but after the final failed pass, it must be corrected and re-reviewed before any merge can happen.

---

## Approval rule
A ticket branch is allowed to merge into the version branch only if:

- the branch has passed all required reviews without unresolved blockers
- the final verification commands have been run successfully
- the branch matches the issue goal and milestone scope
- the code is stable enough for the version branch to absorb it safely

If any item is missing, the answer is no and the branch remains open.

---

## Default reviewer stance
The default reviewer stance is adversarial but fair:

- look for what is broken
- look for what is hidden
- look for what is untested
- look for what is suspiciously broad or loosely reasoned
- reject on uncertainty, not just on style

This is a quality gate, not a ceremony.
