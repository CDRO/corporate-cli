#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
STATE_FILE="$REPO_ROOT/.copilot/v0_agent_state.md"

write_agent_state() {
  local branch="${1:-$(git branch --show-current 2>/dev/null || echo unknown)}"
  local ticket="${2:-$(awk -F': ' '/^- current ticket:/{print $2; exit}' "$STATE_FILE" 2>/dev/null || echo not_started)}"
  local status="${3:-$(awk -F': ' '/^- status:/{print $2; exit}' "$STATE_FILE" 2>/dev/null || echo started)}"
  local cmd="${4:-bash scripts/agent/bootstrap.sh}"
  local last_success="${5:-repo validated and environment checked}"
  local next_action="${6:-read README.md and the v0 build plan, then confirm repo health and start with project bootstrap}"
  local checklist

  if [[ -f "$STATE_FILE" ]]; then
    checklist=$(awk '
      /^## Resume checklist/ {in_checklist=1; print; next}
      in_checklist && /^## / {exit}
      in_checklist { print }
    ' "$STATE_FILE" 2>/dev/null | sed '/^$/d')
  fi

  if [[ -z "$checklist" ]]; then
    checklist=$'## Resume checklist\n- [ ] read README.md\n- [ ] read docs/guides/project_roadmap_and_execution_guide.md\n- [ ] read docs/guides/acceptance_criteria_guide.md\n- [ ] read this file\n- [ ] confirm git status\n- [ ] confirm branch and repo root\n- [ ] run the last known validation command\n- [ ] continue from the next action only after verification'
  fi

  cat > "$STATE_FILE" <<EOF
# v0 Agent State

- branch: $branch
- current ticket: $ticket
- status: $status
- last verified command: $cmd
- last successful result: $last_success
- next action: $next_action
- blocker: none
- repo root: $REPO_ROOT
- last updated: $(date -u +"%Y-%m-%dT%H:%M:%SZ")

$checklist
EOF
}

printf '\n== corporate-cli agent bootstrap ==\n'
printf 'repo_root=%s\n' "$REPO_ROOT"

if [[ ! -f "$REPO_ROOT/go.mod" ]]; then
  echo "ERROR: go.mod not found at $REPO_ROOT. This script must be run from the repo root or its child scripts folder." >&2
  exit 1
fi

if ! command -v go >/dev/null 2>&1; then
  echo "ERROR: Go is not installed or not on PATH." >&2
  exit 1
fi

if [[ ! -f "$REPO_ROOT/README.md" ]] || [[ ! -f "$REPO_ROOT/docs/guides/v0_ai_agent_build_plan.md" ]]; then
  echo "ERROR: required project docs are missing. Restore README.md and the v0 build plan before continuing." >&2
  exit 1
fi

validate_work_branch() {
  local branch
  branch="$(git branch --show-current 2>/dev/null || true)"

  if [[ -z "$branch" || "$branch" == "main" || "$branch" == "master" ]]; then
    echo "ERROR: work must not start on main/master. Create a milestone branch and a ticket subbranch before implementation. Pattern: version/<phase>/<ticket-slug>." >&2
    exit 1
  fi

  if [[ "$branch" =~ ^version/[0-9]+$ ]]; then
    echo "ERROR: milestone branch alone is not enough. Create a ticket subbranch before implementation: git switch -c version/<phase>/<ticket-slug>" >&2
    exit 1
  fi

  if [[ ! "$branch" =~ ^version/[0-9]+/[A-Za-z0-9._/-]+$ ]]; then
    echo "ERROR: invalid branch '$branch'. Required pattern: version/<phase>/<ticket-slug>." >&2
    exit 1
  fi

  printf 'Validated work branch: %s\n' "$branch"
}

cd "$REPO_ROOT"
validate_work_branch

mkdir -p "$REPO_ROOT/.copilot"
if [[ ! -f "$STATE_FILE" ]]; then
  write_agent_state "unknown" "not started" "not_started" "bash scripts/agent/bootstrap.sh" "none" "start by reading README.md and the v0 build plan, then confirm repo health"
  printf '\n== created empty state file ==\n'
  cat "$STATE_FILE"
  printf '\n'
else
  printf '\n== state file ==\n'
  cat "$STATE_FILE"
  printf '\n'
fi

printf '\n== repo state ==\n'
cd "$REPO_ROOT"

git rev-parse --show-toplevel
printf '\n'
git status --short --branch || true
printf '\n'
go version
printf '\n'

printf '== source of truth ==\n'
for file in README.md docs/guides/project_roadmap_and_execution_guide.md docs/guides/acceptance_criteria_guide.md .copilot/v0_agent_state.md; do
  if [[ -f "$REPO_ROOT/$file" ]]; then
    printf '%s\n' "$file"
  else
    printf 'MISSING: %s\n' "$file"
  fi
done

printf '\n== running bootstrap validation ==\n'
if go test ./...; then
  echo "go test ./... passed"
else
  echo "go test ./... failed; this is expected before implementation is complete. Continue with the v0 ticket backlog and fix the failing code next." >&2
fi

printf '\n\n== next work ==\n'
printf '1. Read README.md\n'
printf '2. Read docs/guides/project_roadmap_and_execution_guide.md\n'
printf '3. Read docs/guides/acceptance_criteria_guide.md\n'
printf '4. Start with the first unimplemented v0 ticket from the backlog\n'
printf '5. Keep .copilot/v0_agent_state.md updated after every checkpoint\n\n'
printf 'The agent should not assume prior memory. It must verify the repo state continuously and resume from the repo files only.\n'
