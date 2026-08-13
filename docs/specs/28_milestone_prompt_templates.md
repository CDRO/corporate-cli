# Milestone 28: Prompt Template Design for AI Rewrite

## Goal
Define the prompt structure used when the CLI calls an AI provider to rewrite text into a more corporate tone.

## Why this matters
Prompt quality directly affects output quality. Even if the tool starts with a rules engine, the AI path should have clear, consistent prompts.

## Prompt structure
A prompt should usually include:
- role or system instruction
- task description
- constraints
- input text
- expected output format

## Example system prompt

```text
You are a professional communication assistant.
Rewrite the user text into a polished corporate tone.
Keep the original meaning and factual intent.
Remove profanity, insults, and inflammatory language.
Do not invent facts or add new claims.
Return only the rewritten text.
```

## Example user prompt

```text
Rewrite the following message into a concise and professional corporate style.

Input:
these dumbasses are totaly incompetent and this is a fucking mess
```

## Style-specific variants
- `neutral`: calm and straightforward
- `formal`: more structured and polished
- `executive`: concise and leadership-oriented
- `concise`: shorter, tighter delivery

## Acceptance criteria
- prompts are explicit enough to reduce hallucination risk
- output formatting is predictable
- styles can be mapped to different prompt variants later

## Open questions
- Should the model be forced to return plain text only?
- Should prompt templates live in code or config files?
- Should there be a prompt preview command for debugging?
