# Milestone 11: Release and Distribution

## Goal
Define the release process for a Go-built CLI before the project is shown to external users.

## Release considerations
- keep installation easy
- be explicit about supported OSes
- provide binary names that are easy to discover
- keep the README and build examples in sync

## Release artifacts
- Linux binary
- macOS binary
- Windows executable
- checksums file
- changelog stub

## Release workflow
1. update version number
2. run tests
3. build for target platforms
4. package binaries
5. publish artifacts
6. update README if necessary

## Versioning suggestions
Use a simple semantic versioning pattern:

```text
v0.1.0
v0.2.0
v1.0.0
```

## Acceptance criteria
- release process is simple and repeatable
- binary naming is consistent
- users can install the tool without confusion

## Open questions
- Should the project use GitHub Releases or a different distribution mechanism?
- Is a package manager integration desired later?
- Are release notes required for every milestone or only for stable versions?


