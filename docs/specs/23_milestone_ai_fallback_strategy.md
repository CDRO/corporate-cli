# Milestone 23: AI Fallback and Graceful Degradation

## Goal
Define how the CLI behaves when AI is not configured, is unreachable, returns malformed output, or fails authentication.

## Fallback hierarchy
1. deterministic rules-based rewrite
2. direct prompt-based AI rewrite if configured
3. partial recovery or warning message
4. safe default output when all attempts fail

## Example logic

```text
if no provider configured:
    use rules engine
if provider fails auth:
    print warning and use rules engine
if provider returns invalid output:
    sanitize and retry once
if all attempts fail:
    return best-effort corporate-safe rewrite from rules only
```

## Important design choice
The CLI should not fail for a user just because an AI provider is unavailable. The tool should keep working with the built-in deterministic engine.

## Acceptance criteria
- AI is optional and additive
- the tool remains useful even without a configured backend
- users are informed, not blocked, when AI features are unavailable

## Open questions
- Should a failed AI call produce a warning on stderr or be silent by default?
- How much of the AI response should be trusted before it is accepted?
- Is a retry policy necessary or should it be kept intentionally simple?
