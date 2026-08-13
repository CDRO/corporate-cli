# Milestone 14: Edge Cases and Defensive Handling

## Goal
Capture the likely failure modes and weird input patterns the tool must tolerate before production use.

## Important edge cases

### Empty input
- empty string from stdin
- file exists but is zero bytes
- whitespace-only text

Expected behavior: output should be empty or a minimal valid text response, but never crash.

### Mixed-quality input
Examples:
- profanity mixed with technical language
- short one-line fragments
- heavily abbreviated insults
- message with punctuation but no periods

Expected behavior: transform the message while preserving actions or technical nouns where relevant.

### Very long inputs
- huge file or long stdin stream
- repeated paragraphs
- multi-sentence rant with inconsistent formatting

Expected behavior: process without excessive memory use; stream if possible.

### Input with names, acronyms, and file paths
Examples:
- `PM`, `CTO`, `ADR`, `SLA`, `QA`, `devs`
- file names like `release-v2-final.txt`

Expected behavior: preserve technical names and domain-specific terminology where they are not abusive.

### Broken spellings and noise
Examples:
- `teh`, `wrok`, `mangment`, `proejct`
- repeated keywords like `bad bad bad`

Expected behavior: correct obvious typos but do not rewrite technical identifiers incorrectly.

### Non-UTF-8 input
Potential behavior: fail with a clear message or decode with a fallback strategy.

## Defensive requirements
- avoid panics on malformed input
- print human-readable errors to stderr
- keep stdout clean for successful pipeline output
- fail gracefully if a file is missing

## Acceptance criteria
- the tool remains stable on empty or malformed input
- error output is not mixed into transformed text
- unusual but realistic input does not break the pipeline

## Open questions
- Should the tool attempt spell correction at all, or only fix common known phrases?
- Should it treat uppercase shouting differently from normal text?
- Do we need a `--strict` mode that is more aggressive on profanity removal?


