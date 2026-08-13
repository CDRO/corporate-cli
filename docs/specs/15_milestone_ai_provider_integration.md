# Milestone 15: AI Provider Integration

## Goal
Define how the CLI can optionally call an AI service to improve rewrite quality beyond the deterministic rule engine.

## Why this matters
The base tool can work without AI, but a later milestone can add a smarter backend that improves tone conversion and handles edge cases more naturally.

## Proposed architecture
Introduce a provider abstraction with a small interface so different backends can be swapped in later.

```go
// conceptual only
type Provider interface {
    Name() string
    GenerateText(ctx context.Context, req PromptRequest) (string, error)
}
```

## Supported provider options
- OpenAI-compatible API
- Anthropic-compatible API
- Azure OpenAI
- local/self-hosted model endpoint
- future optional offline provider

## Responsibilities
- send the user text to the selected provider
- pass a clear prompt instructing the model to rewrite into a corporate tone
- handle system prompt and user prompt separately
- return only the transformed text or a structured result

## Example flow

```text
user input -> prompt builder -> provider API -> polished corporate output
```

## Design principles
- keep the provider layer behind an interface
- allow rules-based output to remain the default path
- prefer explicit opt-in configuration over implicit remote calls
- log provider errors without exposing secrets

## Acceptance criteria
- the project can plug in a provider implementation without rewriting the CLI
- default behavior remains deterministic when no AI backend is configured
- provider configuration is readable and reviewable

## Open questions
- Should the AI path be enabled only when a flag or config is set?
- Should the tool support multiple providers in one build?
- Is a local model path required for offline usage later?
