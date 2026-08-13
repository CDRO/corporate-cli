# Master Roadmap: corporate / etaroproc

## Product vision
Build a CLI tool that rewrites text into a more polished, corporate tone, and a complementary inverse mode that turns corporate language back into blunt, direct speech. The project should start as a deterministic, reviewable rules-based tool and only expand into AI-based features once the core workflow is mature and the security/auth requirements are clear.

## Core product concept
- `corporate`: transforms blunt, angry, chaotic, or insulting input into polished business-friendly language.
- `etaroproc`: inverse mode, activated when the executable name is `etaroproc`. It accepts the same general I/O contract but skips login/auth/provider logic.
- Both tools should work with stdin, files, stdout, and shell/PowerShell pipelines.

## Core features overview

### Base CLI features
- input from stdin or file
- output to stdout or file
- `--help`, `--version`
- `--input`, `--output`, `--style`
- pipeline compatibility for Linux/macOS and Windows PowerShell
- deterministic transformation rules
- configurable defaults and local user config

### Corporate mode features
- harsh phrase replacement
- profanity and insult sanitization
- spelling and grammar normalization
- style templates such as `neutral`, `formal`, `executive`, `concise`
- rewrite pipeline covering normalization, detection, replacement, formatting

### Inverse mode features
- inverse mode triggered when binary name is `etaroproc`
- same basic CLI parameters as `corporate`
- no login or provider subcommands by default
- converts polished/corporate phrasing into blunt, direct, natural language
- deterministic rule-based fallback minimal setup

### AI features
- optional provider abstraction
- provider switching
- login/auth flow for first-time use
- secure local storage of credentials
- provider-specific prompts and templates
- fallback to rule engine if AI fails or is unavailable
- usage/cost tracking and retry logic

### Operational features
- config file support
- profile presets
- long-text chunking and context-window handling
- validation/sanity checks before output
- release quality gates and CI checks
- secure secret handling and review

---

## Phase 0: foundation and MVP structure

### Ticket 01: project skeleton
Create the Go module and package layout.

Acceptance criteria:
- `go build ./...` succeeds
- CLI entry point exists under `cmd/corporate`
- packages exist for transform, rules, lexicon, and I/O logic

### Ticket 02: basic CLI parsing
Add stdin/stdout and input/output flags.

Acceptance criteria:
- `--help` prints usage
- `--input` and `--output` work
- stdin piping works
- PowerShell pipeline usage is documented and compatible

### Ticket 03: default transformation pipeline
Implement the first working rewrite flow.

Acceptance criteria:
- input text is accepted from stdin or file
- output is produced to stdout
- newline and whitespace normalization works

### Ticket 04: rule-based harsh phrase replacement
Implement phrase replacements and tone normalization.

Acceptance criteria:
- README examples are transformed into polite business wording
- common insults and profanity patterns are neutralized
- empty input does not crash the tool

### Ticket 05: lexicon and misspelling cleanup
Add lexicon data and common typo correction.

Acceptance criteria:
- misspellings like `totaly`, `dont`, `arent` are handled
- profanity and direct insult sets are reviewable
- default mode removes obvious ugly wording

### Ticket 06: `etaroproc` inverse mode
Implement the inverse corporate binary.

Acceptance criteria:
- invoking `etaroproc` triggers inverse mode automatically
- no auth/login subcommands are exposed
- `--input`, `--output`, stdin, and shell piping work

### Ticket 07: tests for README examples
Create fixtures and regression tests around the README examples.

Acceptance criteria:
- harsh example inputs are transformed successfully
- file and stdin flows are covered
- the project can be reviewed without manual shell testing only

### Ticket 08: config and user defaults
Add user-local config for style/provider defaults.

Acceptance criteria:
- config file is created in a platform-specific user directory
- default provider/style can be read
- CLI flags override config values

### Ticket 20: release quality gates
Define the minimum quality checks before release.

Acceptance criteria:
- tests pass
- build succeeds for supported OSes
- quality gates are documented and repeatable

---

## Phase 1: v1 delivery

This phase delivers the simplest useful version of the project.

### v1 feature set
- rule-based corporate transformation
- inverse `etaroproc` mode
- stdin/file input and stdout/file output
- help/version output
- README-based fixtures and core regression tests
- config file for simple defaults
- release quality gates

### v1 priority checklist
1. project skeleton
2. CLI parsing
3. transformation pipeline
4. harsh phrase replacement
5. lexicon cleanup
6. README tests
7. `etaroproc` inverse mode
8. config and defaults
9. release quality gates
10. auth/login flow for AI mode

---

## Phase 2: AI and auth features

### Ticket 09: authentication flow for AI mode
Support first-time login and secure credential storage.

Acceptance criteria:
- a first-time user can authenticate interactively
- credentials are stored locally and securely
- errors are clear and non-destructive

### Ticket 10: AI provider abstraction
Create a provider interface and a first working backend implementation.

Acceptance criteria:
- provider abstraction exists
- at least one provider can be configured
- failure gracefully falls back to the rules engine

### Ticket 11: AI prompt templates
Define prompt templates for different corporate styles.

Acceptance criteria:
- models are instructed to preserve meaning and remove profanity
- prompt templates are reviewable and per-style
- output formatting stays predictable

### Ticket 12: fallback and retry logic
Handle provider failures and transient conditions.

Acceptance criteria:
- no crash when provider fails
- warning is emitted to stderr
- rule engine remains active as fallback

### Ticket 13: provider switching and profile presets
Allow users to switch provider and save presets.

Acceptance criteria:
- provider list and switch commands work
- profiles save style/provider combinations
- default profiles remain simple

### Ticket 14: chunking and long-input support
Handle long text safely and predictably.

Acceptance criteria:
- long input can be split and rewritten in chunks
- ordering is preserved
- context windows are not exceeded silently

### Ticket 15: result validation and output safety
Validate final output before writing it.

Acceptance criteria:
- suspicious outputs are caught and retried or replaced
- profanity and empty outputs are detected
- dangerous or malformed output is rejected before being returned

### Ticket 16: style templates and user profiles
Expose `neutral`, `formal`, `executive`, `concise` style modes.

Acceptance criteria:
- at least three styles work reliably
- help output documents the styles
- behavior remains testable

### Ticket 18: usage tracking and cost visibility
Add AI usage tracking metadata.

Acceptance criteria:
- usage logs capture model and token information when available
- no secret values are dumped in logs
- usage can be reviewed by the user

### Ticket 19: logout and refresh flow
Support explicit logout and token refresh.

Acceptance criteria:
- logout removes stored auth state
- refresh works when supported
- re-login is straightforward

---

## Phase 3: advanced capabilities and future expansion

### Ticket 17: local model support
Support local and offline models.

Acceptance criteria:
- offline mode is configurable
- local provider works when configured
- tool remains functional without internet access

### Ticket 21: prompt templates for multiple output styles
Expand prompt design and style-specific templates.

Acceptance criteria:
- multiple styles are fully supported
- prompt selection is deterministic and testable
- prompts remain factual and not overly creative

### Ticket 22: user feedback loop
Collect feedback on rewrite quality.

Acceptance criteria:
- users can rate results or store notes
- quality feedback can be reviewed later
- feedback does not break the CLI workflow

### Ticket 23: model selection strategy
Choose the correct model for quality vs cost tradeoffs.

Acceptance criteria:
- model selection is explicit and reviewable
- default mode remains predictable
- users can override model selection

### Ticket 24: config file design and profile management
Expand the config system to include more profile and storage features.

Acceptance criteria:
- config file stores defaults without secrets by default
- profile management can be reviewed and extended
- precedence rules remain clear

### Ticket 25: command catalog and UX polish
Define the final CLI command set and help text.

Acceptance criteria:
- command structure is understandable
- subcommands are discoverable
- no command conflicts exist at the roadmap level

### Ticket 26: AI cost and token awareness
Monitor AI usage cost and token consumption.

Acceptance criteria:
- usage metadata is visible before or after requests
- token-heavy prompts warn users
- cost usage remains reviewable

### Ticket 27: AI failover and retry logic
Handle rate limits, timeouts, and provider issues.

Acceptance criteria:
- automatic retries happen within a safe limit
- fallback to rules mode occurs when needed
- user sees actionable errors only when necessary

### Ticket 28: prompt template design for AI rewrite
Define and evolve the AI prompt structure.

Acceptance criteria:
- prompts remain factual and style-specific
- output remains concise and professional
- prompt logic is easy to debug later

### Ticket 29: local feedback and improvement loop
Collect quality data for future rewrite tuning.

Acceptance criteria:
- user feedback is captured in a local, reviewable format
- improvement process is defined and manageable
- feedback does not require external services in v1

### Ticket 30: model selection strategy
Define model choice rules by input size and provider availability.

Acceptance criteria:
- fast vs balanced vs high-quality modes are clear
- selection is documented and reviewable
- default behavior is stable

### Ticket 31: local model support and offline mode
Support fully local or offline transformation when desired.

Acceptance criteria:
- offline mode works when configured
- local provider can be used without cloud auth
- fallback remains available if the local backend fails

### Ticket 32: context window and chunking strategy
Support long input processing with chunking.

Acceptance criteria:
- long texts do not fail unnoticed
- order remains stable after chunking
- chunking behavior is documented

### Ticket 33: presets and user profiles
Allow profile-based defaults and custom presets.

Acceptance criteria:
- profiles are reviewable and easy to switch
- users can save favorite styles and providers
- default profile remains simple

### Ticket 34: result validation and safety checks
Validate output and detect bad generation results.

Acceptance criteria:
- bad results are prevented or repaired
- validation is testable and deterministic
- safety checks do not block basic workflows

### Ticket 35: release quality gates
Add final quality gates before publishing builds.

Acceptance criteria:
- release checks are explicit
- tests and builds pass before release
- milestones are reviewable prior to publishing

### Ticket 36: inverse corporate mode
Explicitly define and track the inverse rewrite functionality.

Acceptance criteria:
- `etaroproc` is clearly treated as a separate tool mode
- it does not expose login or provider subcommands
- it shares the same basic I/O contract as `corporate`

---

## Priority breakdown

### v1 must-have
- 01 project skeleton
- 02 basic CLI parsing
- 03 default transformation pipeline
- 04 rule-based harsh phrase replacement
- 05 lexicon and misspelling cleanup
- 06 `etaroproc` inverse mode
- 07 README tests
- 08 config and user defaults
- 20 release quality gates

### v2 important
- 09 auth flow
- 10 AI provider abstraction
- 11 prompt templates
- 12 fallback and retry logic
- 13 provider switching and presets
- 14 long-input chunking
- 15 result validation
- 16 style templates
- 18 usage tracking
- 19 logout and refresh

### Later / optional
- 17 local model support
- 21 advanced prompt tuning
- 22 feedback loop
- 23 model selection strategy
- 24 extended config/profile system
- 25 UX polish and command catalog
- 26 token/cost awareness
- 27 failover and retry strategy
- 28 prompt template refinement
- 29 local feedback improvement loop
- 30 model selection policy
- 31 offline model support
- 32 context window/chunking strategy
- 33 presets and profiles
- 34 result validation/safety checks
- 35 release quality gates
- 36 inverse corporate mode

---

## Recommended order for the first 10 work items
1. Ticket 01: project skeleton
2. Ticket 02: basic CLI parsing
3. Ticket 03: default transformation pipeline
4. Ticket 04: rule-based harsh phrase replacement
5. Ticket 05: lexicon and misspelling cleanup
6. Ticket 07: tests for README examples
7. Ticket 06: `etaroproc` inverse mode
8. Ticket 08: config and user defaults
9. Ticket 20: release quality gates
10. Ticket 09: authentication flow for AI mode

## Implementation philosophy
- Start with deterministic, explainable behavior.
- Keep the core CLI usable without AI.
- Build the inverse mode early because it is a strong differentiator and simple to test.
- Add AI later only after the rules-based tool is stable.
- Guard security, provider auth, and output validation before broad release.

## Risk areas to watch
- over-sanitizing text and losing meaning
- over-reliance on brittle regex patterns
- AI output becoming too generic or inaccurate
- auth/config handling exposing secrets
- long-input processing failing due to model token limits
- release quality dropping without a repeatable gate

## Final summary
The project should first become a reliable, deterministic CLI with two modes: `corporate` and `etaroproc`. Once that is stable, the project can add AI support with provider abstraction, login flow, prompts, profiles, and release quality checks. This gives the team a sane path from a small utility to a more advanced, configurable, and optionally AI-powered tool.
