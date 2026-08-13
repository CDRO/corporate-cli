# Milestone 37: Numbered Implementation Tickets

## Goal
Turn the milestone ideas into a concrete backlog with numbered implementation tickets, each with acceptance criteria.

## Ticket 01: project skeleton
**Goal:** create the initial Go project scaffold.
**Acceptance criteria:**
- `go build ./...` succeeds
- CLI entry point exists under `cmd/corporate`
- packages exist for transform, rules, lexicon, and I/O logic

## Ticket 02: basic CLI parsing
**Goal:** support input and output flags and stdin/stdout behavior.
**Acceptance criteria:**
- `--help` prints usage
- `--input` and `--output` work
- stdin piping works
- PowerShell pipeline usage is documented and kept compatible

## Ticket 03: default transformation pipeline
**Goal:** implement a first working rewrite flow from input to output.
**Acceptance criteria:**
- plain text input is accepted
- output is produced on stdout
- whitespace and newline normalization works

## Ticket 04: rule-based harsh phrase replacement
**Goal:** implement common phrase replacements for profanity and insults.
**Acceptance criteria:**
- examples from the README are transformed convincingly
- common phrases are replaced in a deterministic way
- no crash on empty input

## Ticket 05: lexicon and misspelling cleanup
**Goal:** define the lexicon used by the rules engine.
**Acceptance criteria:**
- known misspellings are corrected
- profanity and insult lists are reviewable in code or config
- output no longer contains direct insults in the default mode

## Ticket 06: `etaroproc` inverse mode
**Goal:** implement the inverse corporation mode as a separate binary.
**Acceptance criteria:**
- invoking `etaroproc` triggers inverse rewrite logic automatically
- no login or provider subcommands are exposed
- `--input`, `--output`, and stdin flow work

## Ticket 07: tests for README examples
**Goal:** add tests based on the README example inputs.
**Acceptance criteria:**
- example harsh messages are transformed successfully
- integration tests cover stdin and file usage
- regression tests can be extended with new fixtures

## Ticket 08: config and user defaults
**Goal:** support simple config and default style/provider choices.
**Acceptance criteria:**
- config file is created in a user-local directory
- default style and provider can be read
- CLI flags override config values

## Ticket 09: authentication flow for AI mode
**Goal:** implement login and credential storage groundwork.
**Acceptance criteria:**
- first-time login can happen interactively
- credentials are stored locally in a secure way
- login errors are clear and non-destructive

## Ticket 10: AI provider abstraction
**Goal:** add a provider interface and a first actual backend implementation.
**Acceptance criteria:**
- provider interface exists and is swappable
- an OpenAI-compatible provider can be configured
- failure falls back to the rules engine

## Ticket 11: AI prompt templates
**Goal:** define prompts for different styles.
**Acceptance criteria:**
- prompt templates are reviewable and per-style
- output is concise and professional
- prompts avoid inventing facts

## Ticket 12: fallback and retry logic
**Goal:** graceful degradation when AI fails or is unavailable.
**Acceptance criteria:**
- no crash when provider fails
- warning is emitted to stderr
- rules engine is still used as fallback

## Ticket 13: provider switching and profile presets
**Goal:** allow provider choice and named presets.
**Acceptance criteria:**
- provider list and switch commands work
- profiles save style/provider combinations
- default profile remains simple

## Ticket 14: chunking and long-input support
**Goal:** handle long texts safely.
**Acceptance criteria:**
- large inputs are chunked where needed
- output ordering remains stable
- long requests do not fail silently

## Ticket 15: result validation and output safety
**Goal:** verify generated output before final writing.
**Acceptance criteria:**
- profanity can be checked and filtered again
- empty or suspicious outputs are caught
- output is not written if clearly invalid

## Ticket 16: style templates and user profiles
**Goal:** expose formal/executive/neutral modes.
**Acceptance criteria:**
- at least three styles work reliably
- style selection is documented in help output
- style-specific behaviors are testable

## Ticket 17: local model support
**Goal:** add local offline backend support.
**Acceptance criteria:**
- offline mode is configurable
- local provider fallback works when configured
- system remains usable without internet access

## Ticket 18: usage tracking and cost visibility
**Goal:** improve transparency around AI use.
**Acceptance criteria:**
- token or cost metadata is captured when available
- a user can inspect usage or logs
- usage logging does not leak secrets

## Ticket 19: logout and credential refresh
**Goal:** allow users to revoke and refresh credentials.
**Acceptance criteria:**
- logout removes stored auth state
- refresh works when supported
- user can re-login cleanly

## Ticket 20: release quality gates
**Goal:** define release checks and CI gates.
**Acceptance criteria:**
- tests run automatically before release
- build checks pass on supported OSes
- quality gates are documented and repeatable
