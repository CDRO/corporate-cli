# Milestone 03: Input/Output Handling

## Goal
Define how the CLI accepts text and emits transformed output across different platforms.

## Design principles
- prefer plain text processing
- work with streams for UNIX-style command composition
- support file-based usage for Windows and shell users
- keep API simple for later expansion

## Supported input sources
- standard input (`stdin`)
- file path from argument parser
- maybe direct string argument in the future

## Supported output destinations
- standard output (`stdout`)
- file via `--output`
- maybe stderr for warnings and errors

## Main responsibilities
- read text as UTF-8
- preserve newline handling and trailing newline behavior
- trim or normalize unwanted whitespace consistently
- avoid writing binary or accidental encoding issues

## Proposed flow

```text
read input -> normalize text -> corporateize -> write output
```

## Example flows

```bash
cat rant.txt | corporate > polished.txt
corporate < rant.txt > polished.txt
corporate --input rant.txt --output polished.txt
```

```powershell
Get-Content .\rant.txt | corporate | Set-Content .\polished.txt
corporate --input .\rant.txt --output .\polished.txt
```

## Failure modes
- unreadable file
- permission denied
- encoding mismatch
- output path creation failure
- partial writes during interrupted runs

## Acceptance criteria
- stdin piping works reliably
- file output behaves as expected
- errors are sent to stderr without corrupting stdout data
- successful runs produce clean, valid text

## Open questions
- Should the CLI accept multiple input files in one invocation?
- Should it detect BOMs and normalize them?
- Should it support `-` as stdin shorthand?
- Should output writing be atomic for safer file writes?


