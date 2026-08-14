# corporate

`corporate` is a small command-line tool that turns blunt, heated, and often chaotic text into polished corporate-sounding language.

It is useful when you need to rephrase angry, messy, or unprofessional writing into a more mature and presentable tone without losing the key message.

## Why this tool exists

Sometimes a message starts out as a rant full of frustration, insult words, bad spelling, and blunt language. `corporate` helps convert that into a professional, executive-ready version.

This is intentionally a playful utility. It does not fix facts; it fixes tone.

## Example input

The following text is intentionally harsh, misspelled, and full of cleartext insults. It is meant to demonstrate the kind of input the tool is built to sanitize.

```txt
these dumbasses in the project team are totaly incompetent, the whole thing is a fucking mess, and the deadline is getting pushed because these clowns keep making stupid mistakes and dont even understand basic requirements.
The PM is a useless sack of shit, the devs are lazy as hell, and half the specs are written by a toddler with a broken keyboard.
This is a joke, this whole rollout is a disaster, and if anyone asks me, the whole operation is badly managed and full of idiots.
```

## Example file input

Save the text above as `fire_rant.txt`:

```txt
these dumbasses in the project team are totaly incompetent, the whole thing is a fucking mess, and the deadline is getting pushed because these clowns keep making stupid mistakes and dont even understand basic requirements.
The PM is a useless sack of shit, the devs are lazy as hell, and half the specs are written by a toddler with a broken keyboard.
This is a joke, this whole rollout is a disaster, and if anyone asks me, the whole operation is badly managed and full of idiots.
```

Then pipe it into the tool and write the transformed result to a new file:

### Linux / macOS / sh / bash

```sh
cat fire_rant.txt | corporate > corporate_version.txt
cat corporate_version.txt
```

Example result:

```txt
The project team appears to be experiencing significant challenges in meeting current expectations, and the overall delivery trajectory is not aligned with the established timeline.
The current planning and execution environment suggests a need for improved coordination, clearer requirement definition, and stronger operational discipline across the project lifecycle.
The current situation reflects a substantial opportunity for process improvement, with a particular emphasis on accountability, communication, and execution quality.
```

### Windows PowerShell

```powershell
Get-Content .\fire_rant.txt | corporate | Set-Content .\corporate_version.txt
Get-Content .\corporate_version.txt
```

Example result:

```txt
The project team appears to be experiencing significant challenges in meeting current expectations, and the overall delivery trajectory is not aligned with the established timeline.
The current planning and execution environment suggests a need for improved coordination, clearer requirement definition, and stronger operational discipline across the project lifecycle.
The current situation reflects a substantial opportunity for process improvement, with a particular emphasis on accountability, communication, and execution quality.
```

## Basic usage

### Read from standard input

```sh
printf '%s\n' 'this is a hot mess and the whole team is incompetent' | corporate
```

### Read from a file

```sh
corporate < input.txt
```

### Write output to a file

```sh
corporate < input.txt > output.txt
```

### Use with a pipe in PowerShell

```powershell
Get-Content .\input.txt | corporate | Set-Content .\output.txt
```

## Help

```sh
corporate --help
```

```powershell
corporate --help
```

## Update checks

The CLI exposes an explicit update flow instead of silently overwriting the installed binary.

```sh
corporate update --check
corporate update --install
```

```powershell
corporate update --check
corporate update --install
```

`--check` compares the installed version to the latest GitHub release tag. `--install` prints the safe next step for a human-approved reinstall instead of forcing an in-place upgrade.

## Installation

### Linux / macOS

From the project root, build the binary:

```sh
go build -o corporate ./cmd/corporate
```

Or, if the source is structured differently in your project:

```sh
gcc -O2 -o corporate main.c
```

Then install it to a directory on your `PATH`:

```sh
sudo install -m 755 corporate /usr/local/bin/corporate
```

### Windows

Build the executable:

```powershell
go build -o .\corporate.exe .\cmd\corporate
```

Or, if it is a single-file C/C++ project:

```powershell
cl /O2 /Fe:corporate.exe main.c
```

You can then run it directly:

```powershell
.\corporate.exe --help
```

## Example workflow

### Linux / macOS shell

```sh
cat > fire_rant.txt <<'EOF'
these dumbasses in the project team are totaly incompetent, the whole thing is a fucking mess, and the deadline is getting pushed because these clowns keep making stupid mistakes and dont even understand basic requirements.
This whole rollout is a disgrace and the lead devs are acting like useless bagheads who dont know what they are doing.
EOF

cat fire_rant.txt | corporate > polished.txt
cat polished.txt
```

### Windows PowerShell

```powershell
@'
these dumbasses in the project team are totaly incompetent, the whole thing is a fucking mess, and the deadline is getting pushed because these clowns keep making stupid mistakes and dont even understand basic requirements.
This whole rollout is a disgrace and the lead devs are acting like useless bagheads who dont know what they are doing.
'@ | Set-Content .\fire_rant.txt

Get-Content .\fire_rant.txt | corporate | Set-Content .\polished.txt
Get-Content .\polished.txt
```

## Notes

- `corporate` is intended for tone conversion, not for factual verification.
- It is most useful for turning emotionally charged drafts into more professional language.
- The tool does not guarantee perfect corporate wording; it is a style helper.

## License

This project is provided as-is for demonstration and experimentation purposes.

## Contributing

Open an issue or send a pull request if you want to improve tone detection, output quality, or platform support.

## Issue-driven working model

The active backlog is defined by GitHub milestones and issues. The repository does not use `docs/specs` for execution planning or day-to-day delivery work.

The source-of-truth order is:

1. Work is grouped by GitHub milestone, ordered by milestone number ascending, starting with `v0`, then `v1`, then `v2`, and so on.
2. Within each milestone, work is ordered by the GitHub issue number ascending.
3. Each ticket is implemented on its own branch, using the pattern `version/xy-release/<ticket-slug>`.
4. `docs/specs` remains historical or reference material only; it is never used to decide what to build next.
5. The issue body, acceptance criteria, and linked review comments are the authoritative requirements for each task.

The required workflow is:

1. Create or update the version release branch for the milestone being worked on: `version/xy-release`.
2. Create a ticket branch from that release branch: `version/xy-release/<ticket-slug>`.
3. Deliver the work only on the ticket branch and push it to GitHub for review.
4. Review the GitHub issue and any PR feedback before merging.
5. Merge the ticket branch into the release branch, close the linked issue, and then pull the release branch locally with `git pull --ff-only origin version/xy-release`.
6. After every accepted merge into the release branch, create or update the automated release tag and package as required for the milestone.
7. Do not merge ticket branches directly into `main`.

The intended flow is:

```text
main -> version/xy-release -> version/xy-release/<ticket-slug> -> push -> GitHub review -> fix-up -> GitHub merge -> close linked issue -> git pull --ff-only origin version/xy-release -> validate -> main
```

This keeps the project aligned with the live GitHub backlog and avoids spec drift between local planning docs and the actual milestone work.

The current v1 backlog starts with the lowest-numbered open v1 issue, which is the `AI Provider Integration` item. That issue is the first execution target for the v1 milestone branch.
