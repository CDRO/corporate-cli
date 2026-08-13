# Milestone 05: Lexicon Design

## Goal
Create the vocabulary layer that powers the toolâ€™s transformations.

## Lexicon responsibilities
The lexicon should track:
- insulting phrases
- broken spelling variants
- profanity or curse words
- corporate-safe replacements
- weak/ambiguous wording patterns

## Suggested lexicon layout

```text
harsh_terms/
  insult_phrases.txt
  profanity.txt
  weak_phrases.txt
  misspellings.txt
  corporate_replacements.txt
```

## Example entries

```text
# insult_phrases.txt
these dumbasses
useless sack of shit
clowns
idiots
lazy as hell

# profanity.txt
fucking
shit
damn
hell

# misspellings.txt
totaly => totally
arent => aren't
dont => don't

# corporate_replacements.txt
these dumbasses => the team appears to be experiencing challenges
fucking mess => significant operational challenges
lazy as hell => demonstrating limited execution efficiency
```

## Implementation concept
The lexicon could be represented as:
- a set of maps for direct replacements
- arrays of regex patterns
- a sentence-level phrase table for longer replacements

## Notes
- A lexicon-first solution keeps the tool explainable.
- The README examples should be reflected in the initial test corpus.
- It may be helpful to separate â€œharshâ€ and â€œreplacementâ€ sets, so they are easier to review.

## Acceptance criteria
- The lexicon is easy to inspect and edit
- Replacements are grouped by type
- Example phrases from the README are covered by new tests

## Open questions
- Should the lexicon be statically compiled into code or loaded from data files?
- Should the project support multiple languages later?
- Is spelling normalization a core feature or an optional enhancement?


