# Milestone 30: Model Selection Strategy

## Goal
Define how the CLI chooses which AI model or provider to use for a given task and user configuration.

## Why this matters
Different models trade speed, quality, cost, and context size. A CLI tool that may run in a business context needs clear logic for choosing the right model.

## Potential model modes
- `fast` for short and simple rephrasing
- `balanced` for default usage
- `high-quality` for more complex or sensitive texts
- `offline` for local or rules-only operation

## Suggested configuration logic

```text
if user explicitly set model: use it
else if user selected style = executive: use balanced model
else if text is short: use fast model
else if text is long or complex: use higher-quality model
else: use default provider model
```

## Data points to consider
- prompt length
- quality needs
- cost thresholds
- provider availability
- user configuration

## Acceptance criteria
- model choice is explicit and reviewable
- the default path is predictable
- users can override the model without editing source code

## Open questions
- Should model choice be a config file property or CLI flag?
- How much should the CLI expose about provider-specific models?
- Is a “smartest available model” mode useful or too opaque?
