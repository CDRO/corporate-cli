# Milestone 19: AI Login and Authentication Flow

## Goal
Define how CLI users authenticate to the AI backend so the tool can call a service securely.

## Authentication goals
- allow first-time login from terminal
- avoid storing secrets in plaintext where possible
- support simple local machine usage without extra setup friction
- keep the design compatible with provider-specific auth systems

## Potential auth modes
- API key input via terminal on first run
- OAuth device-code login flow
- external credential provider integration
- config file based credentials for advanced users

## Recommended first version
Use a simple login flow that asks for the provider credentials on demand and stores them securely in a local config or OS keychain if available.

## CLI experience

```bash
corporate login
corporate auth login
corporate --provider openai login
```

## Expected behavior
- if no valid credential is present, the CLI offers a login or setup path
- the user sees a clear prompt to authenticate
- the tool stores credentials in a safe place for future runs
- the command exits with a helpful message if login fails

## Important considerations
- never print tokens to stdout or into logs
- do not store raw credentials in source-controlled files
- support refresh or re-auth flow later
- if auth fails, show a clear error and keep the app usable

## Acceptance criteria
- a first-time user can authenticate without reading source code
- login status is easy to check
- authentication errors are understandable and actionable

## Open questions
- Should the project support multiple providers with one login command or one per provider?
- Do we want OS keychain support immediately or later?
- Should login only work for interactive terminals, or also in CI or script mode?


