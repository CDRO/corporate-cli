# Milestone 04: Rewrite Rule Engine

## Goal
Define how the text transformation logic will work before choosing a strategy such as dictionary replacements, regex rules, or LLM integration.

## Design direction
The first implementation should be deterministic and explainable. For a tool like this, a rule engine is safer than a hidden opaque model.

## Core idea
The tool reads text, splits it into sentences or clauses, applies a rule set, and emits a more polished statement.

## Rule categories
- phrase replacements
- phrase level tone reductions
- profanity masking
- sentence restructuring
- emphasis neutralization
- inconsistency cleanup

## Example replacement table

```text
these dumbasses -> the team appears to be experiencing challenges
useless sack of shit -> underperforming in key responsibilities
fucking mess -> significant operational challenges
lazy as hell -> showing limited execution efficiency
stupid mistakes -> recurring avoidable errors
```

## Rule structure
A rule may have:
- trigger pattern
- replacement text
- optional priority
- optional scope (sentence, paragraph, whole text)
- optional style restrictions

## Possible rule format

```go
// conceptual only
type Rule struct {
    Name      string
    Pattern   string
    Replace   string
    Priority  int
    AppliesTo string
}
```

## Implementation options
1. regex-based replacements
2. dictionary-based mapping
3. sentence normalizer
4. hybrid rule engine (recommended)

## Recommended baseline
Hybrid approach:
- dictionary for common insulting and destructive phrases
- regex for patterns like repeated punctuation, bad spelling clusters, and profanity markers
- sentence transformation for repeated structure patterns

## Acceptance criteria
- A set of example insults is transformed into neutral corporate wording
- Rule ordering does not cause incorrect phrases to cascade
- Output remains readable and coherent

## Open questions
- Should we support user-defined custom rule files later?
- Should the tool only sanitize profanity, or should it also improve grammar?
- How aggressive should the rewrite be by default?
- Are there safety or policy concerns for transforming abusive content?


