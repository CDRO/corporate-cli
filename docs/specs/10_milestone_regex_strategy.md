# Milestone 10: Regex Strategy and Text Normalization

## Goal
Define the regular expression and text-normalization strategy for identifying harsh phrases, profanity, and malformed text before rewriting.

## Why this matters
A deterministic rewrite engine needs clear and repeatable detection steps. Regex is a good starting point because it is transparent, fast, and easily reviewable by a Go developer.

## Text normalization steps
- unify all line endings to `\n`
- replace multiple spaces with single spaces unless paragraph breaks are meaningful
- normalize common apostrophe variants
- lowercase for dictionary matching, but keep the original case for final output if desired
- collapse repeated punctuation like `!!!` or `???` into a single terminal symbol

## Proposed normalization rules

```text
CRLF -> LF
multiple spaces -> single spaces
"totaly" -> "totally"
"dont" -> "don't"
"arent" -> "aren't"
"!!!" -> "!"
"???" -> "?"
```

## Regex categories

### 1. Profanity detection
Detect phrases like:
- `fucking`
- `shit`
- `damn`
- `hell`
- `bastard`
- `idiot`
- `moron`

### 2. Harsh phrase detection
Detect patterns such as:
- `these dumbasses`
- `useless sack of shit`
- `lazy as hell`
- `totaly incompetent`
- `whole thing is a mess`

### 3. Repetition detection
- multiple exclamation marks
- repeated punctuation
- repeated words: `very very very`
- all-caps shouting: `THIS IS A DISASTER`

### 4. Structural patterns
Patterns like:
- `X is a [very] bad Y`
- `this is a [fucking] mess`
- `nobody knows what they are doing`

## Notes
Regex should be used thoughtfully. Overly broad patterns can rewrite harmless content incorrectly. The first version should favor known phrases and short windows rather than broad linguistic guessing.

## Acceptance criteria
- common user inputs from the README are detected reliably
- normalization does not damage the output structure
- pattern matching remains readable and maintainable

## Open questions
- Should matching be case-insensitive by default?
- Should profanity detection be separated from tone detection?
- Are we okay with a small amount of false positives in early milestones?
