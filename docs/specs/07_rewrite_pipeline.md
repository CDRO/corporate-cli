# Milestone 07: Rewrite Pipeline Design

## Goal
Define the processing stages between raw text input and polished corporate output.

## Proposed pipeline

```text
Input text
  -> read content
  -> normalize whitespace and line endings
  -> split into sentences/paragraphs
  -> apply lexicon replacements
  -> apply rule-based tone adjustments
  -> clean grammar and punctuation
  -> emit final text
```

## Stage details

### 1. Input normalization
- convert CRLF to LF
- trim trailing whitespace
- preserve meaningful paragraph separation

### 2. Sentence segmentation
- split on `.`, `!`, `?`, and newline boundaries
- keep punctuation context for later output

### 3. Lexicon lookup
- detect harsh or insulting phrases
- detect curse words and profanity
- fix common misspellings

### 4. Tone mapping
- replace emotionally charged statements with neutral ones
- remove direct insults and demeaning language
- keep message content and intent intact

### 5. Grammar cleanup
- handle odd spacing
- avoid weird repeated punctuation
- keep output readable and concise

### 6. Output formatting
- return plain text to stdout
- ensure file output is valid UTF-8
- preserve paragraph separation where useful

## Example transform

Input:

```text
these dumbasses are totaly incompetent and the whole thing is a fucking mess
```

Output candidate:

```text
The team appears to be experiencing significant operational challenges, and the current situation requires immediate attention and improved coordination.
```

## Design notes
- The pipeline should not be a black box.
- Each stage should be independently testable.
- This makes later revisions easier when the team challenges the rewrite quality.

## Acceptance criteria
- Each stage has a clear responsibility
- Intermediate transforms can be observed or tested
- Pipeline output can be reproduced with fixture-based tests

## Open questions
- Should the tool rewrite full paragraphs first or sentence by sentence?
- Should there be a â€œsafeâ€ mode that protects important technical nouns?
- Is it acceptable to lose some emotional intensity to achieve professionalism?


