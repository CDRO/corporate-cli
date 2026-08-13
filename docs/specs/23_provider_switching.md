# Milestone 23: Provider Switching and Multi-Provider Setup

## Goal
Define how the CLI handles multiple AI providers and how a user can switch between them without confusing configuration.

## Use cases
- user has an OpenAI key and later wants Azure OpenAI
- same tool is used for different business environments
- multiple model providers are configured for testing or fallback

## Proposed behavior
The CLI stores a current provider and may support multiple credentials in a local profile store.

```bash
corporate provider list
corporate provider set openai
corporate provider set azure
```

## Suggested config shape
```json
{
  "defaultProvider": "openai",
  "providers": {
    "openai": { "type": "openai", "configured": true },
    "azure": { "type": "azure", "configured": false }
  }
}
```

## Acceptance criteria
- user can inspect available providers
- user can switch active provider explicitly
- provider selection is deterministic and visible in help output

## Open questions
- Should provider names be case-insensitive?
- Should each provider have a separate credential profile?
- Is per-project config required or is user-level config enough?


