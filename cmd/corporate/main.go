package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"corporate-cli/internal/update"
)

const appVersion = "v0.1.0"

func main() {
	if err := run(os.Args[1:], os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	if len(args) == 0 {
		return printHelp(stdout)
	}

	switch args[0] {
	case "update":
		return handleUpdate(args[1:], stdout)
	case "--help", "-h", "help":
		return printHelp(stdout)
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func printHelp(w io.Writer) error {
	_, err := fmt.Fprint(w, strings.Join([]string{
		"corporate CLI",
		"",
		"Usage:",
		"  corporate update --check",
		"  corporate update --install",
		"  corporate --help",
		"",
		"The update flow is explicit and safe. It never overwrites the current binary without consent.",
	}, "\n"))
	return err
}

func handleUpdate(args []string, stdout io.Writer) error {
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		_, err := fmt.Fprint(stdout, strings.Join([]string{
			"Usage:",
			"  corporate update --check",
			"  corporate update --install",
			"",
			"--check: fetch the latest GitHub release metadata and compare it with the local version.",
			"--install: print the safe follow-up action for a human-approved upgrade.",
		}, "\n"))
		return err
	}

	switch args[0] {
	case "--check", "-c":
		result, err := update.CheckForRelease("https://api.github.com/repos/tizia/corporate-cli", appVersion)
		if err != nil {
			return err
		}
		_, writeErr := fmt.Fprintln(stdout, result.Message)
		return writeErr
	case "--install", "-i":
		_, err := fmt.Fprintln(stdout, "Update install is disabled in this build. Visit the latest GitHub release and reinstall with explicit consent.")
		return err
	default:
		return fmt.Errorf("unknown update command %q", args[0])
	}
}
