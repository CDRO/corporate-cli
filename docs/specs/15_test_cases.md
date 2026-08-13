# Milestone 15: Sample Test Matrix

## Goal
Define a review-friendly set of fixture-based test cases that reflect the README examples and realistic CLI usage.

## Core test groups

### 1. Basic transformation tests
- `these dumbasses` -> neutral corporate wording
- `fucking mess` -> `significant operational challenges`
- `lazy as hell` -> `demonstrating limited execution efficiency`
- `incompetent` -> `not meeting expected performance standards`

### 2. Misspelling normalization
- `totaly` -> `totally`
- `dont` -> `don't`
- `arent` -> `aren't`
- `mangment` -> `management`

### 3. Punctuation and stress handling
- repeated `!!!` is reduced
- repeated profanity is normalized
- all-caps shouting is softened but not removed entirely

### 4. Pipeline tests
- stdin -> stdout
- file -> stdout
- stdin -> file
- file -> file

### 5. Error handling tests
- nonexistent input file
- unreadable file
- invalid flag value
- empty input file

## Example test fixture

```text
Input:
these dumbasses in the project team are totaly incompetent, the whole thing is a fucking mess, and the deadline is getting pushed because these clowns keep making stupid mistakes and dont even understand basic requirements.

Expected contains:
- incompetent
- significant operational challenges
- requirements
- improved coordination
```

## Suggested Go testing pattern

```go
func TestCorporateize(t *testing.T) {
    tests := []struct {
        name  string
        input string
        want  []string
    }{
        {
            name:  "harsh collective criticism",
            input: "these dumbasses are totaly incompetent and this is a fucking mess",
            want: []string{"significant operational challenges", "current situation"},
        },
    }

    for _, tt := range tests {
        got := transform.Corporateize(tt.input)
        for _, s := range tt.want {
            if !strings.Contains(got, s) {
                t.Fatalf("Corporateize(%q) = %q, want substring %q", tt.input, got, s)
            }
        }
    }
}
```

## Acceptance criteria
- the README scenarios have corresponding tests
- style transformation is testable without expensive or fragile infrastructure
- the test matrix is easy to extend as the rules evolve

## Open questions
- Should the tests assert exact phrasing or approximate content?
- Do we want to keep fixtures in Go code or in external `.txt` files?
- Are there any terms we must preserve for technical accuracy?


