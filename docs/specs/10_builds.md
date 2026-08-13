# Milestone 10: Cross-Platform Build Strategy

## Goal
Define how to build the CLI for Linux, macOS, and Windows so the `README` examples are credible and reproducible.

## Build requirements
- Go should be the default implementation language
- builds should support the main supported operating systems
- output binaries should be easy to name and install

## Linux / macOS build

```bash
go build -o corporate ./cmd/corporate
sudo install -m 755 corporate /usr/local/bin/corporate
```

## Windows build

```powershell
go build -o .\corporate.exe .\cmd\corporate
```

## Cross-compilation examples

```bash
GOOS=linux GOARCH=amd64 go build -o corporate-linux-amd64 ./cmd/corporate
GOOS=darwin GOARCH=amd64 go build -o corporate-darwin-amd64 ./cmd/corporate
GOOS=windows GOARCH=amd64 go build -o corporate-windows-amd64.exe ./cmd/corporate
```

## Release automation requirements
- every merge into a version branch must trigger GitHub release creation automatically
- release tags must use the pattern `<version>.<minor>.0`, where:
  - `<version>` is the name of the GitHub milestone that the ticket is attached to, for example `v0`, `v1`, or `v2`
  - `<minor>` is the milestone number recorded in the ticket description, for example `1`, `16`, or `18`
- examples: `v0.1.0`, `v0.16.0`, `v1.18.0`
- the version branch must automatically publish a release bundle containing all supported OS builds
- the release bundle should include Linux, macOS, and Windows binaries, plus checksums and a short note of the merged work

## Packaging ideas
- zip or tarball each OS binary
- add expected SHA256 checksums for release artifacts
- document install paths in the README
- store release artifacts in GitHub Releases generated from the version branch

## Acceptance criteria
- Linux build command works
- Windows build command works
- cross-compilation is demonstrable
- a user can run the binary with no extra framework dependencies
- a GitHub release is created automatically from each version branch merge
- the release bundle contains builds for all supported operating systems

## Open questions
- Should the project ship artifacts on release or only provide build instructions?
- Do we want a Makefile or a small build script later?
- Is a simple `go build` enough for v1, or do we need packaging automation?



