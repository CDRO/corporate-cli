# Milestone 19: Logout, Refresh, and Credential Rotation

## Goal
Define how users can sign out, refresh tokens, and recover from expired credentials without breaking the CLI workflow.

## Use cases
- user wants to log out and remove stored credentials
- access token expired
- provider changed or account switched
- credential was stored in the wrong place or on the wrong machine

## Commands

```bash
corporate logout
corporate auth logout
corporate login --provider openai
corporate login --refresh
```

## Expected behavior
- `logout` removes local credentials and invalidates CLI state
- `login` can trigger a fresh credential acquisition
- refresh flow should use provider refresh tokens if available
- stale credentials should show a clear error rather than silently failing

## Security notes
- invalidate local storage securely
- do not leave cached secrets in plain text in project folders
- if the OS keychain is used, respect its native storage and expiry behavior

## Acceptance criteria
- the user can switch accounts or re-login without editing config files manually
- stale or expired credentials are handled clearly
- logout is non-destructive to the rest of the tool

## Open questions
- Should `logout` also remove provider configuration or just credentials?
- Should refresh be automatic or manual?
- Do we need a `status` command to show current auth state?
