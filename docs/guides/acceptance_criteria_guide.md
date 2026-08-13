# Acceptance Criteria Guide

## Purpose
Acceptance criteria define the conditions that must be true before a milestone, feature, or issue is considered complete. They should be specific enough that a developer or reviewer can decide with confidence whether the work is done.

## Core rules
- Write each criterion in plain language.
- Make criteria observable and testable.
- Prefer behavior over implementation details.
- Keep the criteria measurable where possible.
- Include both positive and negative cases if they matter.
- Avoid vague terms such as "works well", "nice", or "better" without a testable meaning.

## Recommended format
Use a short, consistent structure:

1. Goal: what problem is being solved?
2. Condition: what must be true?
3. Proof: how the result will be checked.

Example:

- Goal: the CLI accepts piped input from stdin.
- Condition: the tool reads text from stdin and emits rewritten output to stdout.
- Proof: run the binary with a pipe and compare the output against the expected transformation.

## Good acceptance criteria
- The program exits successfully with no panic on empty input.
- The command prints usage text when `--help` is passed.
- The output contains no profanity when the default corporate mode is used.
- A file-based input flow matches the stdin-based flow for the same content.
- The config file is created in the platform user directory on first run.

## Weak acceptance criteria
- The CLI works properly.
- The transformation looks good.
- The feature is implemented.
- The code is clean.

These are weak because they cannot be checked objectively.

## Template
Use this template for each issue or milestone:

```text
### Acceptance criteria
- [ ] The feature does X.
- [ ] The feature handles Y edge case.
- [ ] The CLI or API output matches the expected result.
- [ ] Errors are handled clearly and without crashes.
- [ ] The behavior is covered by tests or a reproducible check.
```

## Review checklist
Before closing an issue or milestone, confirm that:
- each requirement can be observed or tested,
- the success condition is not ambiguous,
- failure modes are covered when necessary,
- the acceptance criteria reflect user-facing behavior rather than internal code structure.

## Tip for this project
For `corporate-cli`, acceptance criteria should usually verify:
- CLI behavior for stdin/file inputs,
- output readability and safety,
- deterministic rule-based rewriting,
- auth/provider fallback behavior when relevant,
- release readiness and test coverage.
