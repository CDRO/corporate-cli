# Milestone 09: Cross-Platform Build Strategy

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

## Packaging ideas
- zip or tarball each OS binary
- add expected SHA256 checksums for release artifacts
- document install paths in the README

## Acceptance criteria
- Linux build command works
- Windows build command works
- cross-compilation is demonstrable
- a user can run the binary with no extra framework dependencies

## Open questions
- Should the project ship artifacts on release or only provide build instructions?
- Do we want a Makefile or a small build script later?
- Is a simple `go build` enough for v1, or do we need packaging automation?


