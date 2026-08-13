# Milestone 24: Style Templates and Rewrite Modes

## Goal
Define structured rewrite styles for the tool so users can choose different output personalities.

## Style modes
- `neutral`
- `formal`
- `executive`
- `concise`
- `polite`
- `balanced`

## Example concept

```bash
corporate --style executive < input.txt
corporate --style concise --input input.txt --output output.txt
```

## Template design
Each style can define:
- tone description
- sentence density
- acceptable vocabulary
- profanity handling strategy
- maximum sentence length or complexity

## Example policy for executive mode
- no slang
- no insults
- no profanity
- prefer structured, objective wording
- keep meaning and urgency intact

## Acceptance criteria
- each style can be mapped to a clear transformation profile
- style selection is discoverable from help output
- a user can request a polished output without editing source code

## Open questions
- Should styles be static or use a config file later?
- Should styles be implemented by a rules engine, AI prompt templates, or both?
- Is a `--style auto` mode useful for MVP?


