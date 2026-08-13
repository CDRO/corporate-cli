# Milestone 17: Recommended Implementation Order

## Goal
Provide a practical sequence for the later implementation team to work through the project without getting stuck on architecture too early.

## Recommended order

### Phase 1: Skeleton fix
1. confirm package layout
2. create `main.go` with basic CLI flags parsing
3. verify `go build ./...` succeeds

### Phase 2: IO contract
1. implement stdin reading
2. implement file reading
3. implement stdout writing
4. implement file output writing
5. validate pipeline behavior

### Phase 3: Rule engine
1. create phrase replacement table
2. create regex detection helpers
3. add normalization logic
4. test a few canonical examples

### Phase 4: Lexicon and content transformation
1. define harsh phrases list
2. define corporate replacements
3. add common misspelling map
4. add detector for repeated punctuation and shouting

### Phase 5: Full pipeline integration
1. wire `main` to `transform.Corporateize`
2. confirm pipeline output is clean and stable
3. benchmark with small sample corpus

### Phase 6: Testing and review
1. add unit tests for replacements
2. add integration tests for file/stdin usage
3. finalize review checklist
4. rework rules based on test output

### Phase 7: Build and release
1. build for Linux/macOS/Windows
2. write installation examples
3. prepare release notes and version tagging

## Acceptance criteria
- the implementation sequence is easy to follow
- milestones are not skipped without justification
- review feedback can be incorporated without rewriting the whole project

## Open questions
- Should the team start from a single package or from the layered package structure immediately?
- Do we want to add a `Makefile` before or after the first working CLI?
- Is single-binary distribution enough for the first public version?


