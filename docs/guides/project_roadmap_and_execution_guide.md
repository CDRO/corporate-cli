# Project Roadmap and Execution Guide

## Product vision
Create a deterministic, reviewable CLI that rewrites text into polished corporate tone and exposes an inverse etaroproc mode that converts corporate language back into blunt, direct wording. The project should keep the rules-based core simple and testable before layer-by-layer AI and auth features are added.

## Core product concept
- corporate: turns blunt, angry, chaotic, or insulting input into cleaner business-friendly language.
- etaroproc: inverse mode activated by binary name; accepts the same broad I/O contract but does not expose login or provider flows.
- Both modes should work with stdin, files, stdout, and shell/PowerShell pipelines.

## Roadmap structure
The roadmap is deliberately staged so the team can ship a reliable core first and only add AI and auth complexity after the base behavior is stable.

## Branching policy (mandatory)
The repository must follow a version-based branch workflow so that both developers and AI agents stay aligned with the GitHub version milestones used for delivery.

1. When work begins on a GitHub version milestone, create a branch from `main` named `version/xy`, where `xy` is the version number. Examples: `version/00`, `version/01`, `version/02`.
2. A version branch must be created only once active work for that version begins; it must not exist before implementation starts.
3. For each ticket or subtask worked on within that version, create a new branch from the version branch, not from `main`.
4. Ticket branch names must follow the pattern `version/xy/<ticket-slug>`, for example `version/01/rewrite-pipeline` or `version/02/provider-integration`.
5. All ticket work must be completed on the ticket branch and then squash-merged back into the version branch.
6. Ticket branches must never be merged directly into `main`.
7. The version branch acts as the integration branch for that release phase and remains separate from `main` until the version is complete and approved.
8. Only after the version is fully reviewed and validated should the version branch be merged into `main`.
9. In short: `main` -> `version/xy` -> `version/xy/<ticket-slug>` -> squash merge back to `version/xy` -> review -> `main`.

This prevents long-lived feature branches from drifting away from `main`, keeps each version isolated, and gives every issue a clear, reviewable delivery path.

---

## Phase 0: foundation and MVP

### Ticket 01: project setup and skeleton
Create the Go module and package layout.

### Ticket 02: CLI contract and UX
Add stdin/stdout flow and core flags such as --help, --input, and --output.

### Ticket 03: input/output handling
Handle file-based and pipeline-based text flows reliably.

### Ticket 04: rewrite rule engine
Implement the first deterministic harsh-phrase replacement logic.

### Ticket 05: lexicon design
Add shared word lists, typo normalization, and profanity/insult replacements.

### Ticket 06: inverse corporate language mode
Implement the etaroproc binary and its reverse transformation behavior.

### Ticket 07: rewrite pipeline design
Combine normalization, sentence handling, rule application, and formatting into a single pipeline.

### Ticket 08: testing strategy
Create fixture-based tests for README examples and core CLI behavior.

### Ticket 09: release quality gates
Define minimal checks for build health, test coverage, and release readiness.

### Ticket 10: cross-platform build strategy
Make the project build reproducibly for Linux, macOS, and Windows.

### Ticket 11: release and distribution
Package the CLI and document install and release steps.

### Ticket 12: roadmap and future work
Keep the long-term architecture visible without forcing a premature design.

### Ticket 13: regex strategy and text normalization
Define a repeatable normalization and regex approach for detection and rewriting.

### Ticket 14: edge cases and defensive handling
Cover empty input, repeated punctuation, malformed text, and noisy output paths.

### Ticket 15: sample test matrix
Document and organize concrete regression cases for common use patterns.

### Ticket 16: review checklist
Define a lightweight review process for future iterations.

---

## Phase 1: AI and auth features

### Ticket 18: AI provider integration
Define the provider abstraction and platform extension point.

### Ticket 19: AI login and authentication flow
Support secure first-time login and local credential handling.

### Ticket 20: first-use credential fetch and local storage
Persist credentials in a simple, safe local configuration flow.

### Ticket 21: AI rewrite pipeline and prompt design
Connect provider input to prompt construction and output handling.

### Ticket 22: logout, refresh, and credential rotation
Allow users to re-authenticate and recover cleanly from expired sessions.

### Ticket 23: provider switching and multi-provider setup
Support multiple backends and provider selection in a testable way.

### Ticket 24: style templates and rewrite modes
Expose `neutral`, `formal`, `executive`, and other style variants.

### Ticket 25: security and secret handling review
Audit auth flow, config storage, and secret exposure risk.

### Ticket 26: AI fallback and graceful degradation
Keep the rules engine active when AI is unavailable or fails.

### Ticket 27: configuration file design
Create a local config model for defaults and profile settings.

### Ticket 28: CLI command catalog
Document the final command layout and help output clearly.

### Ticket 29: AI cost, usage, and token awareness
Track usage metadata and warn on token-heavy calls.

### Ticket 30: AI failover, retry, and rate-limit handling
Handle transient provider issues without crashing the CLI.

### Ticket 31: prompt template design for AI rewrite
Refine the text-generation prompts and response constraints.

### Ticket 32: user feedback and improvement loop
Capture quality feedback in a local, reviewable format.

### Ticket 33: model selection strategy
Choose the best model by quality, cost, and availability trade-offs.

### Ticket 34: local model support and offline mode
Allow users to run local-offline or self-hosted backends when configured.

---

## Phase 2: advanced operational features

### Ticket 35: context window and chunking strategy
Support long text inputs by splitting and reassembling output safely.

### Ticket 36: presets and user profiles
Allow reuse of common provider/style combinations through profiles.

### Ticket 37: result validation and safety checks
Validate output before writing, and fall back when generation fails or looks suspicious.

## Acceptance criteria
- the implementation sequence is easy to follow
- milestones are not skipped without justification
- review feedback can be incorporated without rewriting the whole project
- the project has a clear path from prototype to release

## Recommended working pattern
- Keep the rules engine deterministic and reviewable.
- Grow complexity only after the core output is stable.
- Treat AI, provider, and auth layers as optional enhancements rather than prerequisites.
- Use this guide as the source of truth for both scope and sequencing.
