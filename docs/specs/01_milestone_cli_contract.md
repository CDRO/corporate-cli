# Milestone 01: CLI Contract and UX

## Goal
Define what the CLI does from a user perspective and what flags and behaviors it should expose.

## Primary user story
A user can pass text into the tool via standard input, a file, or command flags, and receive a more corporate version of the message on stdout.

## Command behavior
The tool should support at least these modes:

1. stdin to stdout
2. stdin to file
3. file to stdout
4. file to file
5. built-in help output

## Proposed interface

```bash
corporate --help
corporate < input.txt
corporate input.txt
corporate input.txt > output.txt
corporate --input input.txt --output output.txt
corporate --style formal
corporate --style executive
```

## Potential flags
- `--input` path to a source file
- `--output` path to write output
- `--style` one of: `neutral`, `formal`, `executive`, `polite`
- `--strict` stronger sanitization
- `--quiet` suppress informational logs
- `--help` display usage text
- `--version` show version information

## Expected behavior
- read from stdin if no input file is provided
- write to stdout by default
- if `--output` is supplied, write a file instead
- preserve roughly the same content length and concept, but remove aggression and profanity
- keep messages understandable and professional

## Edge cases
- empty input should produce empty output
- whitespace-only input should not crash
- large text input should be handled efficiently
- invalid file path should emit a clear error
- unsupported style should fail with a meaningful message

## Acceptance criteria
- Command line help is readable and accurate
- Standard Unix pipeline usage works
- PowerShell usage is viable with `Get-Content` and `Set-Content`
- Error handling is explicit and non-destructive

## Open questions
- Should the first version support only stdin/stdout and no file flags?
- Should output always end with a trailing newline?
- Should there be a `--no-emoji` or `--preserve-tone` mode?
- Is `--style` necessary in v1 or too much scope?
