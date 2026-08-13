# Milestone 24: Configuration File Design

## Goal
Define how the CLI stores user preferences, credentials metadata, provider defaults, and style preferences in a local configuration file.

## Why this matters
A config layer makes the CLI easier to use and lets users customize behavior without editing code or environment variables manually.

## Proposed storage locations
Use platform-appropriate user config directories:

- Linux: `~/.config/corporate/`
- macOS: `~/Library/Application Support/corporate/`
- Windows: `%APPDATA%\corporate\`

## Initial config file
`config.json` or `config.yaml` could be used. JSON is simpler for Go tooling and easier to review in early milestones.

```json
{
  "defaultProvider": "openai",
  "defaultStyle": "executive",
  "aiEnabled": false,
  "providerSettings": {
    "openai": {
      "model": "gpt-4o-mini",
      "endpoint": "https://api.openai.com/v1"
    }
  }
}
```

## Responsibilities
- store default style
- store selected provider
- store credential metadata references
- allow future config overrides via flags
- keep secrets out of the file unless explicitly required

## Recommended precedence
1. explicit CLI flags
2. environment variables
3. local config file
4. built-in defaults

## Acceptance criteria
- users can change defaults without editing source code
- config is friendly to both CLI and future GUI integrations
- configuration does not contain raw secrets by default

## Open questions
- Should the config file be human-editable YAML or JSON?
- Should the tool support multiple profiles for different users or teams?
- Should `aiEnabled` be a global switch or per-provider switch?
