# Milestone 35: Release Quality Gates

## Goal
Define a set of release checks that must be met before the project is considered stable enough for broader use.

## Quality gate ideas
- all core CLI tests pass
- README examples still match actual behavior
- AI mode works with a mocked or test provider
- rules-based mode works without internet access
- auth and config flows are tested in a safe environment
- cross-platform build succeeds for at least Linux and Windows

## Suggested checklist
1. `go test ./...` passes
2. sample transformer cases pass
3. CLI pipeline examples pass
4. invalid input scenarios are handled
5. AI fallback works when provider is unavailable
6. release binaries build for target OSes

## Acceptance criteria
- every release has a minimum verification pass
- the README examples remain realistic
- major regressions are caught before release

## Open questions
- Do we require AI tests to use mocks or real provider sandboxes?
- Should the release gate include security review before publication?
- Is `go vet` or `golangci-lint` expected in the early phase?
