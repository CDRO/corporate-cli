# v0 AI Agent Build Plan

## Goal
Build the v0 milestone for `corporate-cli` as a deterministic, reviewable, and testable CLI that supports the core product flow:

- rewrite blunt or angry text into polished corporate language
- support the inverse `etaroproc` mode from the binary name
- handle stdin, files, stdout, and pipeline usage reliably
- pass a basic release-quality gate for the first shippable version

This plan is intended for a Copilot agent or another autonomous coding agent. It is written so the work can be resumed from any machine, after a crash, or after a partial session without needing memory from the prior run.

---

## Source of truth
Use these files as the authoritative references before starting or resuming work:

- [README.md](README.md)
- [docs/guides/project_roadmap_and_execution_guide.md](docs/guides/project_roadmap_and_execution_guide.md)
- [docs/guides/acceptance_criteria_guide.md](docs/guides/acceptance_criteria_guide.md)

If a file is missing, stop and restore it from the repo before continuing. Do not assume the implementation state from memory.

---

## Required branch workflow
The project uses version milestones as delivery units. Every v0 task must follow the version workflow:

- create or use a working branch from `main`: `version/00-release` or `version/00-release/<ticket-slug>`
- all implementation work must happen on a ticket branch derived from the release branch
- ticket branches are squash-merged back into the release branch
- do not merge ticket work directly into `main`

Example flow:

```text
main
  -> version/00-release
      -> version/00-release/cli-contract
      -> version/00-release/input-output
      -> version/00-release/rule-engine
  -> squash merge back to version/00-release
  -> final review
  -> merge version/00-release into main
```

---

## v0 scope
The v0 milestone covers the foundational product, not the AI features.

The implementation must include:

1. project setup and Go module structure
2. CLI interface and command behavior
3. stdin/stdout and file-based I/O
4. deterministic rule-based rewrite logic
5. lexicon and typo normalization
6. inverse `etaroproc` mode
7. full rewrite pipeline integration
8. test fixtures and regression protection
9. release quality gates and build checks
10. cross-platform build validation

The AI should not begin provider integration, auth flows, or model logic yet. Those belong to later versions.

---

## Deliverables for v0
The following artifacts are the minimum acceptably complete output for the v0 milestone:

- Go module initialized and buildable
- `cmd/corporate/main.go` CLI entry point
- a stable transform pipeline under `internal/`
- rule engine for harsh-to-corporate rewrite behavior
- lexicon and normalization tables
- `etaroproc` inverse mode behavior
- tests covering core examples and edge cases
- README examples aligned with actual program behavior
- successful `go build ./...` and `go test ./...`

---

## Task order for the agent
Work in the following order. Do not jump ahead unless the current task is fully verified.

### 1. Bootstrap project
- verify the repo state and current branch
- confirm repository root and module name
- initialize or validate Go project structure
- create the required directories if missing
- verify `go build ./...` succeeds before introducing logic

Required files:
- `cmd/corporate/main.go`
- `internal/transform/transform.go`
- `internal/rules/rules.go`
- `internal/lexicon/lexicon.go`
- `internal/io/io.go`

### 2. CLI contract
Implement the command-line behavior:

- `--help`
- `--input`
- `--output`
- standard input and stdout flow
- PowerShell pipeline compatibility

Acceptance focus:
- usage prints correctly
- piped input works
- output is written to stdout or file as expected
- no panic on empty input

### 3. Input and output handling
Implement:

- reading from stdin
- reading from files
- writing to stdout
- writing to output files
- newline and whitespace consistency

Acceptance focus:
- same content produces matching output for stdin and file intake
- output remains readable and stable

### 4. Rewrite rule engine
Implement the core deterministic rewrite behavior:

- harsh phrase replacements
- direct insult removal or neutralization
- shouting and repeated punctuation normalization
- typo cleaning and phrase-level rewriting

Acceptance focus:
- README examples transform as expected
- no crash on malformed text
- replacement logic is deterministic and reviewable

### 5. Lexicon and normalization
Create the shareable replacement vocabulary:

- common rude or insulting phrases
- corporate alternatives
- typo map for common mistakes
- punctuation normalization rules

Acceptance focus:
- data is centralized and easy to review
- regex and maps are minimal and understandable

### 6. Inverse mode: `etaroproc`
Add the inverse command behavior:

- binary name determines mode
- if executable is `etaroproc`, reverse the corporate rewrite flow
- do not expose AI or auth flows in this mode
- keep the same stdin/file/stdout contract

Acceptance focus:
- `etaroproc` behaves as an inverse transform mode
- no auth or provider flow is involved

### 7. Pipeline integration
Connect all logic into one pipeline:

- read input
- normalize text
- apply rule engine
- rewrite lexicon-driven phrases
- emit result

Acceptance focus:
- consistent output for the same input
- maintain readability and basic tone quality

### 8. Tests and review
Add tests for:

- README examples
- empty input
- common profanity phrases
- file input and stdin input parity
- `etaroproc` inverse mode
- regression cases

Acceptance focus:
- `go test ./...` passes
- sample cases are covered by fixtures
- reviews are possible without reading the whole algorithm from memory

### 9. Release quality gate
Run and document the release checks:

- `go test ./...`
- `go build ./...`
- verify examples in README still match output
- document the known constraints of v0

### 10. Cross-platform validation
Check that core commands work on the supported target OSes:

- Linux
- macOS
- Windows

Acceptance focus:
- the build path is reproducible
- each OS can build the CLI with a standard `go build` command

---

## Recovery and crash-resume protocol
This is the key requirement for AI continuity.

If the agent crashes, is interrupted, loses context, or gets stuck, it must recover using the repository state alone. The agent must not assume previous session memory.

### Always re-read these before continuing
1. the project README
2. this build plan
3. the roadmap/execution guide
4. the current Git state
5. the current files under `cmd/` and `internal/`

### Required resume state file
Create and maintain a file at:

- `.copilot/v0_agent_state.md`

This file is the handoff checkpoint. It must be updated after each major step or when the agent stops for any reason.

Required contents:

```md
# v0 Agent State

- branch: version/00
- current ticket: <ticket-name>
- status: in_progress | blocked | complete
- last verified command: <command>
- last successful result: <short description>
- next action: <next step>
- blocker: <none or description>
- repo root: <path>
- last updated: <timestamp>
```

When the agent resumes, it must:

1. read `.copilot/v0_agent_state.md`
2. read `git status`
3. read the relevant task file or implementation file
4. verify the last successful command before doing anything new
5. continue from the next action only after confirming the repo is still in the expected state

### Crash recovery rules
- Never continue from an assumption such as "I think this was almost done."
- Always verify by running the last known successful command or a minimal validation command.
- If the repo is dirty, inspect the diff before making a new change.
- If a task is blocked, record the blocker in the state file and stop instead of guessing.
- If a task cannot be completed, write the exact reason and the next verification command required.

### Recovery command sequence
When resuming, run these in order:

```bash
git status

git branch --show-current

ls -la

# if the state file exists
cat .copilot/v0_agent_state.md

# then verify the project health

go test ./...
# or, if tests are not yet implemented:
go build ./...
```

If any of these fail, stop and fix the root issue before continuing to new work.

---

## Completion definition for v0
The v0 milestone is complete only when all of the following are true:

- the project builds without errors
- all tests pass
- the core CLI behavior matches the README examples
- the inverse `etaroproc` mode works
- the repo state is intentional and reviewable
- the final branch is merged back into the version branch and ready for review

---

## Important guardrails
- Do not add AI provider logic before v0 is complete.
- Do not start auth or secret-handling flows in v0.
- Keep the implementation deterministic and easy to review.
- Prefer explicit, minimal code over clever abstractions.
- If the agent becomes stuck, resume from the state file and the last verified command instead of continuing blindly.

---

## Minimal handoff summary for any agent
If a new agent starts from any environment, it should be able to reconstruct work by reading:

1. README.md
2. docs/guides/project_roadmap_and_execution_guide.md
3. docs/guides/acceptance_criteria_guide.md
4. .copilot/v0_agent_state.md
5. the current repo state and failing or passing validation commands

This ensures the work can be resumed from anywhere with no hidden context dependency.
