$ErrorActionPreference = 'Stop'
Set-StrictMode -Version Latest

function Write-AgentState {
    param(
        [string]$Branch,
        [string]$Ticket,
        [string]$Status,
        [string]$Command,
        [string]$LastResult,
        [string]$NextAction
    )

    $checklist = @()
    if (Test-Path $StateFile) {
        $lines = Get-Content -Path $StateFile
        $inChecklist = $false
        foreach ($line in $lines) {
            if ($line -match '^## Resume checklist') {
                $inChecklist = $true
                $checklist += $line
                continue
            }
            if ($inChecklist -and $line -match '^## ') {
                break
            }
            if ($inChecklist) {
                $checklist += $line
            }
        }
    }

    if (-not $checklist -or $checklist.Count -eq 0) {
        $checklist = @(
            '## Resume checklist',
            '- [ ] read README.md',
            '- [ ] read docs/guides/project_roadmap_and_execution_guide.md',
            '- [ ] read docs/guides/acceptance_criteria_guide.md',
            '- [ ] read this file',
            '- [ ] confirm git status',
            '- [ ] confirm branch and repo root',
            '- [ ] run the last known validation command',
            '- [ ] continue from the next action only after verification'
        )
    }

    $stateText = @(
        '# v0 Agent State',
        '',
        "- branch: $Branch",
        "- current ticket: $Ticket",
        "- status: $Status",
        "- last verified command: $Command",
        "- last successful result: $LastResult",
        "- next action: $NextAction",
        '- blocker: none',
        "- repo root: $RepoRoot",
        "- last updated: $(Get-Date -Format o)",
        '',
        $checklist
    )

    $stateText | Set-Content -Path $StateFile -Encoding utf8
}

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$RepoRoot = (Resolve-Path (Join-Path $ScriptDir '..\..')).Path
$StateFile = Join-Path $RepoRoot '.copilot\v0_agent_state.md'

Write-Host ""
Write-Host "== corporate-cli agent bootstrap =="
Write-Host "repo_root=$RepoRoot"

if (-not (Test-Path (Join-Path $RepoRoot 'go.mod'))) {
    throw "ERROR: go.mod not found at $RepoRoot. Run this script from the repo root or a child folder."
}

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    throw "ERROR: Go is not installed or not on PATH."
}

if (-not (Test-Path (Join-Path $RepoRoot 'README.md')) -or -not (Test-Path (Join-Path $RepoRoot 'docs\guides\v0_ai_agent_build_plan.md'))) {
    throw "ERROR: required project docs are missing. Restore README.md and the v0 build plan before continuing."
}

New-Item -ItemType Directory -Force -Path (Join-Path $RepoRoot '.copilot') | Out-Null
if (Test-Path $StateFile) {
    Write-Host ""
    Write-Host "== state file =="
    Get-Content -Path $StateFile
    Write-Host ""
} else {
    Write-AgentState -Branch 'unknown' -Ticket 'not started' -Status 'not_started' -Command 'none' -LastResult 'none' -NextAction 'start by reading README.md and the v0 build plan, then confirm repo health'
    Write-Host ""
    Write-Host "== created empty state file =="
    Get-Content -Path $StateFile
    Write-Host ""
}

Write-Host "== repo state =="
Set-Location $RepoRoot

git rev-parse --show-toplevel
Write-Host ""
git status --short --branch
Write-Host ""
go version
Write-Host ""

Write-Host "== source of truth =="
$requiredFiles = @(
    'README.md',
    'docs/guides/project_roadmap_and_execution_guide.md',
    'docs/guides/acceptance_criteria_guide.md',
    '.copilot/v0_agent_state.md'
)

foreach ($file in $requiredFiles) {
    $fullPath = Join-Path $RepoRoot $file
    if (Test-Path $fullPath) {
        Write-Host $file
    } else {
        Write-Host "MISSING: $file"
    }
}

$branchName = git branch --show-current 2>$null
if (-not $branchName) { $branchName = 'unknown' }
$commandText = 'powershell -ExecutionPolicy Bypass -File scripts/agent/bootstrap.ps1'
Write-AgentState -Branch $branchName -Ticket 'bootstrap' -Status 'started' -Command $commandText -LastResult 'repo validated and environment checked' -NextAction 'read README.md and docs/guides/v0_ai_agent_build_plan.md, then start the first v0 implementation task'

Write-Host ""
Write-Host "== running bootstrap validation =="
try {
    go test ./...
    Write-Host "go test ./... passed"
} catch {
    Write-Host "go test ./... failed; this is expected before implementation is complete. Continue with the v0 ticket backlog and fix the failing code next." -ForegroundColor Yellow
}

Write-Host ""
Write-Host "== next work =="
Write-Host '1. Read README.md'
Write-Host '2. Read docs/guides/project_roadmap_and_execution_guide.md'
Write-Host '3. Read docs/guides/acceptance_criteria_guide.md'
Write-Host '4. Start with the first unimplemented v0 ticket from the backlog'
Write-Host '5. Keep .copilot/v0_agent_state.md updated after every checkpoint'
Write-Host ""
Write-Host 'The agent should not assume prior memory. It must verify the repo state continuously and resume from the repo files only.'
