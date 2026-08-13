# Milestone 34: Result Validation and Safety Checks

## Goal
Define how the CLI validates rewritten output to ensure it is still safe, coherent, and useful before writing it to stdout or a file.

## Why this matters
When using AI or aggressive rewriting, a result may contain invented wording, quality problems, or unwanted content. Validation should catch obvious issues.

## Validation checks
- ensure no secrets were leaked into output
- check that profanity was actually removed when requested
- ensure output still contains the original message meaning roughly
- ensure no nonsense or broken formatting is produced
- detect empty or suspiciously short output

## Example checks
- if input had profanity and output still contains it, warn or retry
- if AI output looks empty or nonsense, fallback to rules-based output
- if rewriting removed all content, route to a more conservative default

## Proposed validation flow

```text
rewrite result -> validate length and quality -> if bad -> fallback -> emit final
```

## Acceptance criteria
- outputs are sanity-checked before release
- bad or malformed output does not silently pass through
- users can still get useful results in degraded mode

## Open questions
- Should validation be a simple rules pass or more advanced scoring?
- Is output validation required in v1 or only for AI mode?
- How much of this logic should be public vs internal?
