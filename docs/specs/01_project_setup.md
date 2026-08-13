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
â”œâ”€â”€ cmd/
â”‚   â””â”€â”€ corporate/
â”‚       â””â”€â”€ main.go
â”œâ”€â”€ internal/
â”‚   â”œâ”€â”€ io/
â”‚   â”‚   â””â”€â”€ io.go
â”‚   â”œâ”€â”€ lexicon/
â”‚   â”‚   â””â”€â”€ lexicon.go
â”‚   â”œâ”€â”€ rules/
â”‚   â”‚   â””â”€â”€ rules.go
â”‚   â””â”€â”€ transform/
â”‚       â””â”€â”€ transform.go
â”œâ”€â”€ docs/
â”‚   â””â”€â”€ specs/
â”‚       â”œâ”€â”€ 01_project_setup.md
â”‚       â”œâ”€â”€ 02_cli_contract.md
â”‚       â”œâ”€â”€ 03_input_output.md
â”‚       â”œâ”€â”€ 04_rules_engine.md
â”‚       â”œâ”€â”€ 05_lexicon.md
â”‚       â”œâ”€â”€ 06_rewrite_pipeline.md
â”‚       â”œâ”€â”€ 07_testing.md
â”‚       â”œâ”€â”€ 08_builds.md
â”‚       â”œâ”€â”€ 09_release.md
â”‚       â”œâ”€â”€ 10_roadmap_and_future_work.md
â”‚       â”œâ”€â”€ 11_regex_strategy.md
â”‚       â”œâ”€â”€ 12_edge_cases.md
â”‚       â”œâ”€â”€ 13_test_cases.md
â”‚       â”œâ”€â”€ 14_review_checklist.md
â”‚       â”œâ”€â”€ 15_implementation_order.md
â”‚       â”œâ”€â”€ 16_ai_provider_integration.md
â”‚       â”œâ”€â”€ 17_auth_login.md
â”‚       â”œâ”€â”€ 18_first_run_credentials.md
â”‚       â”œâ”€â”€ 19_ai_prompt_pipeline.md
â”‚       â”œâ”€â”€ 20_logout_refresh.md
â”‚       â”œâ”€â”€ 21_provider_switching.md
â”‚       â”œâ”€â”€ 22_style_templates.md
â”‚       â”œâ”€â”€ 23_security_review.md
â”‚       â”œâ”€â”€ 24_ai_fallback_strategy.md
â”‚       â”œâ”€â”€ 25_config_file_design.md
â”‚       â”œâ”€â”€ 26_cli_command_catalog.md
â”‚       â”œâ”€â”€ 27_ai_cost_and_tokens.md
â”‚       â”œâ”€â”€ 28_ai_failover_and_retry.md
â”‚       â”œâ”€â”€ 29_prompt_templates.md
â”‚       â”œâ”€â”€ 30_user_feedback_loop.md
â”‚       â”œâ”€â”€ 31_model_selection_strategy.md
â”‚       â”œâ”€â”€ 32_local_model_support.md
â”‚       â”œâ”€â”€ 33_context_window_and_chunking.md
â”‚       â”œâ”€â”€ 34_presets_and_profiles.md
â”‚       â”œâ”€â”€ 35_result_validation.md
â”‚       â”œâ”€â”€ 36_release_quality_gates.md
â”‚       â”œâ”€â”€ 37_inverse_corporate_mode.md
â”‚       â”œâ”€â”€ 38_numbered_tickets.md
â”‚       â”œâ”€â”€ 39_v1_v2_later_priorities.md
â”‚       â”œâ”€â”€ 40_recommended_first_10.md
â”‚       â””â”€â”€ 41_master_roadmap.md
â”œâ”€â”€ .gitignore
â”œâ”€â”€ go.mod
â”œâ”€â”€ README.md
â””â”€â”€ LICENSE
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


