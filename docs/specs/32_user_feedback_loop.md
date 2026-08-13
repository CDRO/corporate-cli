# Milestone 32: User Feedback and Improvement Loop

## Goal
Define how the tool collects user feedback about rewrite quality so the rules engine and AI prompts can improve over time.

## Why this matters
A tone-conversion tool needs iterative improvement. The product will not be perfect at first, and the review process should include feedback loops.

## Possible feedback mechanisms
- `--rating` or `--feedback` flags for a manual review flow
- optional local feedback log
- prompt for feedback after transformation in interactive mode
- simple quality report after command execution

## Example workflow

```bash
corporate --style executive < input.txt
Would you like to rate this rewrite? [good/bad/neutral]
```

## Feedback data to store
- original input
- transformed output
- selected style
- user rating
- timestamp
- provider used (if any)

## Storage options
- local JSON log file
- CSV for easy review
- optional future remote analytics backend

## Acceptance criteria
- feedback is easy to capture without extra infrastructure
- the tool remains useful even when feedback is not enabled
- feedback data is stored in a review-friendly way

## Open questions
- Is user feedback local-only for v1 or a later cloud feature?
- Should feedback be anonymous or tied to the user profile?
- Should the tool allow â€œwhy was this bad?â€ prompts later?


