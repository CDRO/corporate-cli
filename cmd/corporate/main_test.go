package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunWithoutArgsReadsStdinWhenInputIsPiped(t *testing.T) {
	var stdout bytes.Buffer

	err := run(nil, strings.NewReader("This is a hot mess and the whole team is incompetent."), &stdout, "corporate")
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	got := stdout.String()
	if got == "" {
		t.Fatal("expected transformed output, got empty string")
	}
	if strings.Contains(got, "Usage:") {
		t.Fatalf("run unexpectedly printed help instead of transforming stdin: %q", got)
	}
}

func TestRunWithoutArgsShowsHelpWhenNoInputExists(t *testing.T) {
	var stdout bytes.Buffer

	err := run(nil, strings.NewReader(""), &stdout, "corporate")
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	got := stdout.String()
	if !strings.Contains(got, "Usage:") {
		t.Fatalf("expected help output, got %q", got)
	}
}

func TestRunInverseModeUsesInverseCorporateize(t *testing.T) {
	var stdout bytes.Buffer
	input := "The project team appears to be experiencing significant challenges in meeting current expectations."

	err := run(nil, strings.NewReader(input), &stdout, "etaroproc")
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	got := stdout.String()
	if strings.Contains(got, "Usage:") {
		t.Fatalf("run unexpectedly printed help in inverse mode: %q", got)
	}
	lower := strings.ToLower(got)
	for _, want := range []string{"mess", "requirements"} {
		if !strings.Contains(lower, want) {
			t.Fatalf("expected inverse output to include %q, got %q", want, got)
		}
	}
}

func TestRunInverseFlagUsesInverseMode(t *testing.T) {
	var stdout bytes.Buffer
	input := "The project team appears to be experiencing significant challenges in meeting current expectations."

	err := run([]string{"--inverse"}, strings.NewReader(input), &stdout, "corporate")
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	got := stdout.String()
	lower := strings.ToLower(got)
	for _, want := range []string{"mess", "requirements"} {
		if !strings.Contains(lower, want) {
			t.Fatalf("expected inverse mode output to include %q, got %q", want, got)
		}
	}
}

func TestRunInputFileFlagReadsFileAndTransformsIt(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "input.txt")
	content := "these dumbasses are totaly incompetent and this is a fucking mess"
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write input file: %v", err)
	}

	var stdout bytes.Buffer
	err := run([]string{"--input", path}, strings.NewReader("ignored"), &stdout, "corporate")
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	got := stdout.String()
	if !strings.Contains(strings.ToLower(got), "significant") {
		t.Fatalf("expected output to include corporate rewrite content, got %q", got)
	}
	for _, banned := range []string{"dumbasses", "fucking"} {
		if strings.Contains(strings.ToLower(got), banned) {
			t.Fatalf("expected harsh content to be removed, got %q", got)
		}
	}
}

func TestRunMissingInputFileReturnsError(t *testing.T) {
	var stdout bytes.Buffer
	err := run([]string{"--input", filepath.Join(t.TempDir(), "missing.txt")}, strings.NewReader("ignored"), &stdout, "corporate")
	if err == nil {
		t.Fatal("expected missing input file error")
	}
	if !strings.Contains(err.Error(), "no such file") && !strings.Contains(err.Error(), "missing") {
		t.Fatalf("expected file-not-found error, got %v", err)
	}
}

func TestRunEmptyInputFileProducesEmptyOutput(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "empty.txt")
	if err := os.WriteFile(path, []byte(""), 0o600); err != nil {
		t.Fatalf("write empty input file: %v", err)
	}

	var stdout bytes.Buffer
	err := run([]string{"--input", path}, strings.NewReader("ignored"), &stdout, "corporate")
	if err != nil {
		t.Fatalf("run returned error for empty input file: %v", err)
	}
	if got := stdout.String(); strings.TrimSpace(got) != "" {
		t.Fatalf("expected empty output for empty input file, got %q", got)
	}
}

func TestRunInvalidFlagValueReturnsError(t *testing.T) {
	var stdout bytes.Buffer
	err := run([]string{"--input"}, strings.NewReader("ignored"), &stdout, "corporate")
	if err == nil {
		t.Fatal("expected error for invalid --input flag value")
	}
	if !strings.Contains(err.Error(), "missing value") {
		t.Fatalf("expected missing value error, got %v", err)
	}
}

func TestRunUnknownCommandReturnsError(t *testing.T) {
	var stdout bytes.Buffer
	err := run([]string{"not-a-command"}, strings.NewReader(""), &stdout, "corporate")
	if err == nil {
		t.Fatal("expected error for unknown command")
	}
	if !strings.Contains(err.Error(), "unknown command") {
		t.Fatalf("expected unknown-command error, got %v", err)
	}
}
