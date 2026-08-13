# Milestone 00: Project Setup and Skeleton

## Goal
Create the initial project scaffold so Go developers can start implementing the CLI without needing to re-establish the repository layout.

## Scope
- initialize Go module
- create CLI entry point under `cmd/corporate`
- create placeholder packages for transform, rules, lexicon, and I/O helpers
- create minimal `.gitignore`
- leave obvious TODO markers for later implementation

## Proposed structure

```text
corporate-cli/
├── cmd/
│   └── corporate/
│       └── main.go
├── internal/
│   ├── io/
│   │   └── io.go
│   ├── lexicon/
│   │   └── lexicon.go
│   ├── rules/
│   │   └── rules.go
│   └── transform/
│       └── transform.go
├── docs/
│   └── specs/
│       ├── 00_milestone_project_setup.md
│       ├── 01_milestone_cli_contract.md
│       ├── 02_milestone_input_output.md
│       ├── 03_milestone_rules_engine.md
│       ├── 04_milestone_lexicon.md
│       ├── 05_milestone_rewrite_pipeline.md
│       ├── 06_milestone_testing.md
│       ├── 07_milestone_builds.md
│       ├── 08_milestone_release.md
│       └── 09_milestone_roadmap.md
├── .gitignore
├── go.mod
├── README.md
└── LICENSE
```

## Notes
- This is intentionally minimal and easy to revise.
- The package boundaries are illustrative and should be challenged during review.
- The implementation should not lock into one rewrite mechanism before evaluation.

## Acceptance criteria
- Repository builds with `go build ./...` even if the logic is still a placeholder.
- Program has a clear entry point.
- All major conceptual areas are represented as packages or files.

## Open questions
- Should the core transformation logic live in `internal/transform` or at the package root?
- Do we want a `pkg` layer for public APIs later?
- Should we keep `internal/io` or merge it into `main` if this stays tiny?
