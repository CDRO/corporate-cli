# Milestone 34: Local Model Support and Offline Mode

## Goal
Define how the tool could support local AI models or offline-only execution for environments with no cloud access.

## Why this matters
Some users will want no external API dependency. A rules engine plus optional local-model support makes the tool more flexible.

## Potential local model modes
- llama.cpp-based local model
- Ollama integration
- local HTTP API endpoint
- future GPU-assisted local inference

## CLI behavior

```bash
corporate --offline
corporate --provider local
corporate --model llama3.1
```

## Responsibilities
- detect whether a local model endpoint is available
- choose the local mode automatically when configured
- keep fallback to rules-based output if local inference fails

## Acceptance criteria
- the tool can operate offline when configured
- the CLI still works in no-AI environments
- local-provider setup is documented and easy to test

## Open questions
- Should local model support be part of v1 or a later milestone?
- Is offline mode required for corporate or secure environments?
- What is the minimum local model setup that is worth supporting?


