# Milestone 26: AI Cost, Usage, and Token Awareness

## Goal
Capture the operational concerns of using an AI backend, especially cost and usage patterns.

## Why this matters
The tool can start as a deterministic utility, but AI support introduces token usage, rate limits, and billable operations.

## Desired behaviors
- show usage estimate before long AI transformations if possible
- warn on large input size
- support chunking or truncation for very large texts
- keep a simple local usage log for review

## Suggested fields for tracking
- provider name
- model name
- prompt tokens
- completion tokens
- total tokens used
- request timestamp
- cost estimate if available

## Example usage log
```json
{
  "provider": "openai",
  "model": "gpt-4o-mini",
  "timestamp": "2026-08-13T12:00:00Z",
  "promptTokens": 200,
  "completionTokens": 160,
  "totalTokens": 360
}
```

## Acceptance criteria
- AI usage is visible enough to review later
- token-heavy commands can be estimated before execution
- default behavior remains usable when the user chooses not to use AI

## Open questions
- Should the CLI log usage by default or only when a flag is set?
- Do we need a `usage` or `billing` command for review?
- Is AI cost awareness required for v1 or only for later releases?
