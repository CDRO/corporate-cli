# Milestone 32: Context Window and Chunking Strategy

## Goal
Define how the CLI handles long input text when using AI providers with limited context windows or token budgets.

## Why this matters
Large messages may exceed provider limits. The tool must decide whether to chunk the text, summarize it, or process it in segments.

## Input size considerations
- short text: process directly
- medium text: keep as single request
- long text: split into chunks or detect the need for summarization

## Proposed strategy
1. check input size and estimated token count
2. if within provider limits, process as one request
3. if over limit, split into paragraphs or sentence groups
4. rewrite each chunk with same style constraint
5. combine outputs in order

## Chunking rules
- split on paragraphs first
- then on sentence boundaries if needed
- preserve ordering and context between chunks
- avoid cutting mid-sentence when possible

## Example behavior

```text
Long rant text
 -> paragraph 1 chunk
 -> paragraph 2 chunk
 -> paragraph 3 chunk
 -> merge rewritten chunks back together
```

## Acceptance criteria
- long text inputs do not fail silently because of context limits
- output order remains stable
- large requests can still produce a useful result

## Open questions
- Should the tool chunk only for AI mode or also for local rules mode?
- How much overlap should chunks have between neighboring sections?
- Should the CLI warn the user when chunking is used?
