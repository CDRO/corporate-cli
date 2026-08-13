# Milestone 25: Security and Secret Handling Review

## Goal
Capture the expected security rules for AI integration and stored credentials before a real release is considered.

## Security concerns
- API keys in terminal history
- secrets accidentally written to stdout logs
- plaintext credentials in config files
- accidental exposure in debug output
- use of untrusted or local model endpoints

## Security requirements
- never print secrets to stdout or logs
- keep secrets out of source control
- favor OS keychain or secure local storage
- mask credentials in any debug information
- use environment variables or safe config paths when appropriate
- require explicit consent before requesting login or store credentials

## Example safe behaviors
- `corporate login` should prompt in terminal, not echo the token
- when a credential is invalid, print a clear error without dumping the credential itself
- debug mode should redact anything that looks like a token

## Acceptance criteria
- storage and retrieval of secrets is reviewed before release
- the README includes clear security warnings for AI auth
- no obvious security anti-patterns are present in the design

## Open questions
- Should the tool support a `--debug` flag that redacts secrets automatically?
- Is OS keychain integration required or optional in early milestones?
- Should configuration files allow secure password storage or only token-based storage?


