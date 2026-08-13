# Milestone 36: Presets and User Profiles

## Goal
Define how users can save and reuse preferred styles, providers, and CLI behaviors through presets or profiles.

## Why this matters
Users may want different defaults depending on context, team, or role. A profile system reduces repeated flag input.

## Example profiles
- `default`
- `executive`
- `customer-facing`
- `team-internal`
- `strict-corporate`

## Example config

```json
{
  "profiles": {
    "default": {
      "style": "executive",
      "provider": "openai",
      "aiEnabled": true
    },
    "customer-facing": {
      "style": "formal",
      "provider": "azure",
      "aiEnabled": true
    }
  }
}
```

## CLI examples

```bash
corporate profile list
corporate profile use executive
corporate --profile customer-facing < input.txt
```

## Acceptance criteria
- users can switch behaviors without changing code or environment all the time
- defaults remain simple for first-time users
- profiles can be added later without major redesign

## Open questions
- Should profiles be file-based or stored in a single config file?
- Should profiles be editable by the CLI or only by hand?
- Is a single default profile sufficient for v1?


