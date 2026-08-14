$ErrorActionPreference = 'Stop'
$root = Split-Path -Parent $PSScriptRoot
$dist = Join-Path $root 'dist'
New-Item -ItemType Directory -Force -Path $dist | Out-Null

$builds = @(
    @{ Name = 'corporate'; Output = (Join-Path $dist 'corporate.exe') },
    @{ Name = 'etaroproc'; Output = (Join-Path $dist 'etaroproc.exe') }
)

foreach ($build in $builds) {
    $source = Join-Path $root 'cmd\corporate'
    $binary = $build.Output
    if ($build.Name -eq 'etaroproc') {
        $tmp = Join-Path $dist 'corporate-inverse.exe'
        go build -o $tmp $source
        Copy-Item $tmp $binary -Force
        Remove-Item $tmp -Force
    } else {
        go build -o $binary $source
    }
}

Write-Host "Built:"
Get-ChildItem $dist | Select-Object Name, FullName
