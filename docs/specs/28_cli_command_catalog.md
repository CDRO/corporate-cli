# Milestone 28: CLI Command Catalog

## Goal
Define the command set for the CLI in a way that remains reviewable and extensible before implementation.

## Command group ideas

### Core commands
- `corporate` or `corporate --help`
- `corporate < input.txt`
- `corporate --input file.txt --output output.txt`
- `corporate --style executive`

### Auth commands
- `corporate login`
- `corporate logout`
- `corporate auth status`
- `corporate auth refresh`

### Provider commands
- `corporate provider list`
- `corporate provider set openai`
- `corporate provider set azure`

### Config commands
- `corporate config show`
- `corporate config set defaultStyle executive`
- `corporate config reset`

### Future commands
- `corporate version`
- `corporate diagnose`
- `corporate benchmark`
- `corporate dry-run`

## Design principles
- keep the main command easy to use
- keep subcommands explicit and discoverable
- avoid too many flags in v1
- support help text that explains both default and AI-enhanced usage

## Acceptance criteria
- the command set is easy to review and extend
- the CLI remains intuitive for shell users and PowerShell users
- no command name conflicts exist in the early roadmap

## Open questions
- Should auth commands be top-level or namespaced under `auth`?
- Is a `status` command necessary or is `login` enough for v1?
- How much of the command catalog should be implemented before the project is considered usable?


