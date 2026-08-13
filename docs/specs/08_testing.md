# Milestone 08: Testing Strategy

## Goal
Define a test approach that supports iterative rewriting and rule review without overcomplicating the project.

## Test types

### Unit tests
- lexicon replacements
- regex behavior
- sentence splitting
- whitespace normalization
- error handling paths

### Integration tests
- stdin to stdout flow
- file input and output flow
- CLI help output
- command line errors

### Regression tests
- store examples from the README
- verify that harsh input is converted into professional output
- keep tests stable while the rewrite rules evolve

## Example fixtures

```text
input: these dumbasses are totaly incompetent and this is a fucking mess
expect: the team appears to be experiencing significant operational challenges and the current situation requires immediate attention
```

## Suggested file structure

```text
internal/
  transform/
    transform_test.go
  rules/
    rules_test.go
  lexicon/
    lexicon_test.go
cmd/
  corporate/
    main_test.go
```

## Recommended approach
- use table-driven tests
- store examples in small strings or fixture files
- avoid brittle assertions on exact wording too early
- prefer testing behavior like â€œcontains no profanityâ€ and â€œcontains expected corporate phrasingâ€

## Acceptance criteria
- example README inputs are converted into professional output in tests
- edge cases are covered
- the CLI contract is testable without manual shell steps

## Open questions
- Should tests assert exact output or general tone?
- Do we add snapshot tests now or wait until the rewrite engine stabilizes?
- How much of the pipeline should be public for testing?


