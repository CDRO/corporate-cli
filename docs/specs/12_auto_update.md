# Milestone 12: Auto Update

## Goal
Define how the CLI can check for a newer release and either notify the user or install the latest binary automatically without making the workflow fragile or surprising.

## Core requirements
- the project should support release-based updates after new GitHub releases are published
- update checks must be explicit and safe
- the tool should avoid forced installs or silent overwrites without user consent
- update behavior should be documented for CLI users and maintainers

## Update model
- the application checks the current installed version against the latest GitHub release tag
- the release tag follows the pattern `<version>.<minor>.0`, where:
  - `<version>` is the GitHub milestone name the ticket is attached to, such as `v0`, `v1`, or `v2`
  - `<minor>` is the ticket milestone number in the ticket description, such as `1`, `16`, or `18`
- if a newer release is available, the CLI may print an update notice or guide the user to install the newest build
- the update check must fail gracefully when the network is unavailable or the release metadata is malformed

## Safety requirements
- no silent replacement of user binaries without confirmation
- no automatic install from an untrusted or unsigned artifact
- prefer a clear prompt or an explicit command such as `corporate update --check` or `corporate update --install`
- log or surface the reason when the update cannot be validated

## Platform behavior
- support upgrade checks for Linux, macOS, and Windows builds published in the GitHub release bundle
- the release assets must be easy to match against the current operating system and architecture
- update instructions should remain simple enough for end users to follow without extra tooling

## Acceptance criteria
- the project defines a safe, explicit update flow
- the CLI can compare its local version to the latest GitHub release
- the workflow is documented for users and maintainers
- update checks fail gracefully instead of crashing or overwriting data
- the update mechanism is compatible with the release strategy defined in milestone 11

## Open questions
- Should update checks be enabled by default or behind a flag?
- Should the project support direct in-place upgrades or only guidance to re-install?
- Do we want to include a self-update command later, or keep the first iteration as a version check and notification only?
