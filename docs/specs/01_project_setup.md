# Milestone 01: Project Setup and Skeleton

## Goal
Create the initial project scaffold so Go developers can start implementing the CLI without needing to re-establish the repository layout.

## Scope
- initialize Go module
- create CLI entry point under `cmd/corporate`
- create placeholder packages for transform, rules, lexicon, and I/O helpers
- create minimal `.gitignore`
- leave obvious TODO markers for later implementation

## Proposed structure

```text
corporate-cli/
├── cmd/
│   └── corporate/
│       └── main.go
├── internal/
│   ├── io/
│   │   └── io.go
│   ├── lexicon/
│   │   └── lexicon.go
│   ├── rules/
│   │   └── rules.go
│   └── transform/
│       └── transform.go
├── docs/
│   └── specs/
│       ├── 01_project_setup.md
│       ├── 02_cli_contract.md
│       ├── 03_input_output.md
│       ├── 04_rules_engine.md
│       ├── 05_lexicon.md
│       ├── 06_rewrite_pipeline.md
│       ├── 07_testing.md
│       ├── 08_builds.md
│       ├── 09_release.md
│       ├── 10_roadmap_and_future_work.md
│       ├── 11_regex_strategy.md
│       ├── 12_edge_cases.md
│       ├── 13_test_cases.md
│       ├── 14_review_checklist.md
│       ├── 15_implementation_order.md
│       ├── 16_ai_provider_integration.md
│       ├── 17_auth_login.md
│       ├── 18_first_run_credentials.md
│       ├── 19_ai_prompt_pipeline.md
│       ├── 20_logout_refresh.md
│       ├── 21_provider_switching.md
│       ├── 22_style_templates.md
│       ├── 23_security_review.md
│       ├── 24_ai_fallback_strategy.md
│       ├── 25_config_file_design.md
│       ├── 26_cli_command_catalog.md
│       ├── 27_ai_cost_and_tokens.md
│       ├── 28_ai_failover_and_retry.md
│       ├── 29_prompt_templates.md
│       ├── 30_user_feedback_loop.md
│       ├── 31_model_selection_strategy.md
│       ├── 32_local_model_support.md
│       ├── 33_context_window_and_chunking.md
│       ├── 34_presets_and_profiles.md
│       ├── 35_result_validation.md
│       ├── 36_release_quality_gates.md
│       ├── 37_inverse_corporate_mode.md
│       ├── 38_numbered_tickets.md
│       ├── 39_v1_v2_later_priorities.md
│       ├── 40_recommended_first_10.md
│       └── 41_master_roadmap.md
├── .gitignore
├── go.mod
├── README.md
└── LICENSE
```

## Notes
- This is intentionally minimal and easy to revise.
- The package boundaries are illustrative and should be challenged during review.
- The implementation should not lock into one rewrite mechanism before evaluation.

## Acceptance criteria
- Repository builds with `go build ./...` even if the logic is still a placeholder.
- Program has a clear entry point.
- All major conceptual areas are represented as packages or files.

## Open questions
- Should the core transformation logic live in `internal/transform` or at the package root?
- Do we want a `pkg` layer for public APIs later?
- Should we keep `internal/io` or merge it into `main` if this stays tiny?


