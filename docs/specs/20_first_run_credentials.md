# Milestone 20: First-Use Credential Fetch and Local Storage

## Goal
Define the behavior for loading credentials on the very first run and storing them in a safe local location for future use.

## Requirement
When a user runs the tool for the first time with AI features enabled, it should either:
- fetch credentials automatically if configured, or
- prompt the user to authenticate, or
- load credentials from a standard environment variable or local secret store

## Example behavior

```bash
$ corporate --style executive
No AI credentials found. Would you like to log in now? [Y/n]
```

Then after login:

```bash
$ corporate --style executive < input.txt
```

## Credential sources
- environment variables
- OS keychain / credential manager
- config file in user home directory
- provider-specific login flow

## Local storage design
Store a minimal auth object such as:
- provider name
- access token
- refresh token
- expiry timestamp
- optional metadata

Keep this under a user-scoped directory such as:

```text
~/.config/corporate/
~/.corporate/
```

or a platform-appropriate equivalent for Windows/macOS/Linux.

## Security requirements
- never write secrets to the repo
- never print tokens to terminal output
- avoid plaintext credentials in obvious config files unless explicitly requested
- prefer OS keychain when available

## Acceptance criteria
- first run triggers a login or credential resolution flow
- credentials become available for later runs without manual re-entry
- a user can clear the stored credentials easily

## Open questions
- Should a user be able to force a re-login command?
- Should the tool support explicit `--login` and `--logout` commands?
- Is OS keychain support a v1 requirement or a later enhancement?


