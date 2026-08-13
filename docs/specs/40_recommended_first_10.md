# Milestone 40: Recommended Order for the First 10 Work Items

## Goal
Provide a clear implementation order for the first ten tickets so the team starts with the minimal viable and reviewable build path.

## Recommended order

### 1. Ticket 01: project setup and skeleton
Create the Go module, package structure, and initial entry points.

### 2. Ticket 02: CLI contract and UX
Add the basic command-line flags and stdin/stdout behavior.

### 3. Ticket 03: input/output handling
Handle file input, file output, and pipeline-friendly text flow.

### 4. Ticket 04: rewrite rule engine
Add the first rule-based transformation logic for harsh phrases and tone normalization.

### 5. Ticket 05: lexicon design
Add the replacement lists and common typo corrections that drive rewriting quality.

### 6. Ticket 06: inverse corporate language mode
Build the `etaroproc` mode once the base corporate flow is stable.

### 7. Ticket 07: rewrite pipeline design
Connect normalization, sentence handling, and replacement steps into one reviewable pipeline.

### 8. Ticket 08: testing strategy
Add fixture-based tests so README example behavior is locked in.

### 9. Ticket 09: release quality gates
Define the minimal verification checks required before a release is considered usable.

### 10. Ticket 10: cross-platform build strategy
Make the build path reproducible for Linux, macOS, and Windows.

## Why this order
This sequence follows the actual project backlog and keeps the product reviewable before AI or auth complexity is added. It prioritizes predictable behavior, testability, and a believable release path before optional expansion.

## Acceptance criteria for the first 10
- the CLI works without AI or accounts
- the inverse mode works as a separate binary
- README examples are covered by tests
- the project is buildable and reviewable on common platforms
- the next steps are clearly scoped for later AI and auth work without forcing them into v1


