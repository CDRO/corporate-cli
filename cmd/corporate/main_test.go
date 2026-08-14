package main

import (
	"bytes"
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
