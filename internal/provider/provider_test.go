package provider

import (
	"context"
	"os"
	"testing"
)

func TestBuildProviderDefaultsToNoop(t *testing.T) {
	provider, err := BuildProvider(Config{})
	if err != nil {
		t.Fatalf("BuildProvider returned error for default config: %v", err)
	}
	if provider.Name() != "noop" {
		t.Fatalf("expected noop provider by default, got %q", provider.Name())
	}

	out, err := provider.GenerateText(context.Background(), PromptRequest{Text: "hello"})
	if err != nil {
		t.Fatalf("noop provider returned error: %v", err)
	}
	if out != "hello" {
		t.Fatalf("noop provider should return the original text, got %q", out)
	}
}

func TestBuildProviderRejectsUnknownProvider(t *testing.T) {
	_, err := BuildProvider(Config{Enabled: true, Name: "not-real"})
	if err == nil {
		t.Fatal("expected error for unknown provider")
	}
}

func TestNewConfigFromEnvUsesEnabledProvider(t *testing.T) {
	old := os.Getenv("CORPORATE_PROVIDER")
	defer os.Setenv("CORPORATE_PROVIDER", old)

	if err := os.Setenv("CORPORATE_PROVIDER", "mock"); err != nil {
		t.Fatalf("set env: %v", err)
	}

	cfg := NewConfigFromEnv()
	if !cfg.Enabled {
		t.Fatal("expected provider config to be enabled when provider name is set")
	}
	if cfg.Name != "mock" {
		t.Fatalf("expected provider name to be mock, got %q", cfg.Name)
	}
}
