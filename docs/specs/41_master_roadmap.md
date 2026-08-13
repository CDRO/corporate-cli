# Master Roadmap: corporate / etaroproc

## Product vision
Create a deterministic, reviewable CLI that rewrites text into a polished corporate tone and exposes an inverse `etaroproc` mode that converts corporate language back into blunt, direct wording. The project should keep the rules-based core simple and testable before layer-by-layer AI and auth features are added.

## Core product concept
- `corporate`: turns blunt, angry, chaotic, or insulting input into cleaner business-friendly language.
- `etaroproc`: inverse mode activated by binary name; accepts the same broad I/O contract but does not expose login or provider flows.
- Both modes should work with stdin, files, stdout, and shell/PowerShell pipelines.

## Roadmap structure
The milestone docs in this repository are the source of truth. The sequence below matches the actual numbered files and avoids duplicate tickets and stale references.

---

## Phase 0: foundation and MVP

### Ticket 01: project setup and skeleton
Create the Go module and package layout.

### Ticket 02: CLI contract and UX
Add stdin/stdout flow and core flags such as `--help`, `--input`, and `--output`.

### Ticket 03: input/output handling
Handle file-based and pipeline-based text flows reliably.

### Ticket 04: rewrite rule engine
Implement the first deterministic harsh-phrase replacement logic.

### Ticket 05: lexicon design
Add shared word lists, typo normalization, and profanity/insult replacements.

### Ticket 06: inverse corporate language mode
Implement the `etaroproc` binary and its reverse transformation behavior.

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

### Ticket 17: implementation order
Record the recommended ordering for the first meaningful build path.

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

---

## Priority breakdown

### v1 must-have
- 01 project setup and skeleton
- 02 CLI contract and UX
- 03 input/output handling
- 04 rewrite rule engine
- 05 lexicon design
- 06 inverse corporate language mode
- 07 rewrite pipeline design
- 08 testing strategy
- 09 release quality gates
- 10 cross-platform build strategy
- 11 release and distribution

### v2 important
- 18 AI provider integration
- 19 AI login and authentication flow
- 20 first-use credential fetch and local storage
- 21 AI rewrite pipeline and prompt design
- 22 logout, refresh, and credential rotation
- 23 provider switching and multi-provider setup
- 24 style templates and rewrite modes
- 25 security and secret handling review
- 26 AI fallback and graceful degradation
- 27 configuration file design
- 28 CLI command catalog
- 29 AI cost, usage, and token awareness
- 30 AI failover, retry, and rate-limit handling
- 31 prompt template design for AI rewrite

### Later / optional
- 32 user feedback and improvement loop
- 33 model selection strategy
- 34 local model support and offline mode
- 35 context window and chunking strategy
- 36 presets and user profiles
- 37 result validation and safety checks

---

## Recommended order for the first 10 work items
1. Ticket 01: project setup and skeleton
2. Ticket 02: CLI contract and UX
3. Ticket 03: input/output handling
4. Ticket 04: rewrite rule engine
5. Ticket 05: lexicon design
6. Ticket 06: inverse corporate language mode
7. Ticket 07: rewrite pipeline design
8. Ticket 08: testing strategy
9. Ticket 09: release quality gates
10. Ticket 10: cross-platform build strategy

## Implementation philosophy
- Start with deterministic, explainable behavior.
- Keep the core CLI usable without AI or auth.
- build the inverse mode early because it is easy to test and a strong product differentiator.
- add AI and provider complexity only after the rules engine is stable.
- guard security, config, and output validation before broader release.

## Risk areas to watch
- over-sanitizing text and losing meaning
- brittle regex replacements that are hard to review
- AI output becoming generic or inaccurate
- auth/config handling exposing secrets
- long-input processing failing due to model limits
- release quality dropping without a repeatable gate

## Final summary
The project should first become a reliable, deterministic CLI with two modes: `corporate` and `etaroproc`. Once that base is stable, the repository can safely add AI support, provider abstraction, login flow, prompt design, and operational polish. This keeps the project reviewable and avoids mixing feature work before the core rewrite engine is proven.

