# Milestone 27: AI Failover, Retry, and Rate-Limit Handling

## Goal
Define how the CLI behaves under transient AI provider failures, rate limits, and network problems.

## Failure modes to handle
- HTTP 429 rate limit
- authentication failure
- timeouts
- network disruption
- malformed provider responses
- overlong inputs

## Expected behaviors
- retry a request a limited number of times with backoff
- fall back to the rules engine when AI is unavailable
- surface a concise warning to the user
- avoid sending multiple duplicate requests in tight loops

## Example policy
- 3 retries with exponential backoff
- then fall back to offline rules-based mode
- log request details only in debug mode

## Acceptance criteria
- user experience remains stable despite provider issues
- the tool does not crash on transient AI errors
- fallback is automatic and documented

## Open questions
- Should retry/backoff be provider-specific from the start?
- Do we need a global `--no-ai` switch to force offline mode?
- Should large prompts be auto-truncated before calling the provider?
