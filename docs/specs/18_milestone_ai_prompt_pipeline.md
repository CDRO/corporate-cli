# Milestone 18: AI Rewrite Pipeline and Prompt Design

## Goal
Define how the tool uses AI once credentials are available, while preserving the deterministic rules engine as a fallback.

## Prompt pipeline

```text
normalize input -> inspect for profanity/harshness -> decide strategy ->
if rules are sufficient -> use rules engine
else -> call AI provider with prompt -> return polished output
```

## Prompt design goals
- ask the AI to rewrite text into a professional corporate tone
- maintain the original message meaning
- remove abuse and vulgarity without fabricating facts
- be explicit about the expected output format

## Example prompt

```text
Rewrite the following text into a concise, professional, corporate tone.
Keep the facts and intent intact.
Remove insults, profanity, and aggressive wording.
Do not invent new facts.
Return only the rewritten text.

Input:
these dumbasses are totaly incompetent and this is a fucking mess
```

## Fallback strategy
- if the provider is unavailable or auth fails, fall back to the rules engine
- if the provider output is unacceptable, keep the deterministic version
- if the provider returns malformed output, sanitize and retry if possible

## Safety considerations
- never make up facts
- never leak original secrets from the user input to logs
- keep provider errors separate from user-visible transformed text

## Acceptance criteria
- AI output can be compared against rule-based output in tests
- failures degrade gracefully
- the pipeline remains reviewable and understandable

## Open questions
- Should the model be asked to return only plain text or JSON with metadata?
- Should the tool allow different prompt templates per style like `formal`, `executive`, `concise`?
- Is AI output optional in v1 or always available once configured?
