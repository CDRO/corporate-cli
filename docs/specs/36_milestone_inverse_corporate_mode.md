# Milestone 36: Inverse Corporate Language Mode

## Goal
Define a second CLI mode that does the opposite of the main `corporate` tool: it converts polished or corporate-sounding text back into blunt, direct, unfiltered language.

## Product idea
If the executable is called `etaroproc` (the inverse of `corporate`), the CLI should activate the inverse mode automatically. This mode intentionally does not include user account, login, provider setup, or sub-command functionality.

## Core rule
- `corporate` = corporate-speak generator
- `etaroproc` = inverse corporate-speak generator

This means the same core logic may be reused, but with a different transformation profile and a much simpler CLI contract.

## Intended behavior
When invoked as `etaroproc`, the tool should:
- accept the same basic input sources as `corporate`
- accept CLI parameters like `--input`, `--output`, `--help`, `--version`
- skip login/auth/provider commands entirely
- transform text from corporate or polished phrasing into blunt, awkward, direct, or intentionally unprofessional wording
- remain simple and predictable, without AI or account features unless explicitly added later

## Examples

Input:

```text
The team appears to be experiencing significant challenges in meeting current expectations, and the overall delivery trajectory is not aligned with the established timeline.
```

Inverse corporate output candidate:

```text
The team is kinda failing hard, the timeline is slipping, and nobody seems to have a grip on what is actually going on.
```

## CLI contract for `etaroproc`

```bash
etaroproc --help
etaroproc input.txt
etaroproc < input.txt > output.txt
etaroproc --input input.txt --output output.txt
```

```powershell
etaroproc --help
Get-Content .\input.txt | etaroproc | Set-Content .\output.txt
```

## Design principles
- keep the inverse mode simple and explicit
- do not inherit login, AI, or config subcommands unless a later product decision requires it
- prefer deterministic rewrite rules over AI-generated inversion
- keep both binaries aligned around the same core transformation model

## Acceptance criteria
- invoking the binary as `etaroproc` triggers inverse mode automatically
- the inverse mode accepts the same general input/output parameters as `corporate`
- no login or provider subcommands are exposed by default
- the output is intentionally less polished and more direct than the corporate version
- the CLI works in shell and PowerShell pipelines

## Open questions
- Should `etaroproc` share the exact same options as `corporate` or only the basic subset?
- Should inverse mode be available as a flag on `corporate` too, for example `corporate --inverse`?
- Should the inverse mode be rules-driven only, or can it use the same provider abstraction later?
