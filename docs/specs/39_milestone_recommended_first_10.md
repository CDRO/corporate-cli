# Milestone 39: Recommended Order for the First 10 Work Items

## Goal
Provide a clear implementation order for the first ten tickets so the team starts with the minimal viable and reviewable build path.

## Recommended order

### 1. Ticket 01: project skeleton
Create the Go module, package structure, and empty entry points.

### 2. Ticket 02: basic CLI parsing
Add stdin/stdout and input/output flags.

### 3. Ticket 03: default transformation pipeline
Implement the first working text-to-text rewrite flow.

### 4. Ticket 04: rule-based harsh phrase replacement
Add the most obvious phrase transformations from the README examples.

### 5. Ticket 05: lexicon and misspelling cleanup
Add the replacement lists and simple misspelling handling.

### 6. Ticket 07: tests for README examples
Write fixtures around the README scenarios to stabilize the behavior.

### 7. Ticket 06: `etaroproc` inverse mode
Add the inverse binary after the main corporate mode is working.

### 8. Ticket 08: config and user defaults
Add local config for style and provider defaults.

### 9. Ticket 20: release quality gates
Document and implement the quality checks needed for a minimal release.

### 10. Ticket 09: authentication flow for AI mode
Add the first secure login path once the deterministic core is stable.

## Why this order
This sequence builds the project from simple and reliable behavior to AI-enabled features. It keeps the tool useful, reviewable, and testable before expanding into auth, multi-provider logic, or advanced prompt systems.

## Acceptance criteria for the first 10
- the CLI works without AI
- the inverse mode works as a separate binary
- README examples are represented in tests
- the project is buildable and reviewable
- the next steps are clearly scoped for AI features without forcing them into v1
