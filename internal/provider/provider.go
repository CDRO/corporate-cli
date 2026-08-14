package provider

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
)

// PromptRequest carries the text the provider should rewrite.
type PromptRequest struct {
	Text string
}

// Provider is a small abstraction for optional AI-backed rewrite providers.
type Provider interface {
	Name() string
	GenerateText(ctx context.Context, req PromptRequest) (string, error)
}

// Config describes provider selection and opt-in behavior.
type Config struct {
	Enabled bool
	Name    string
	APIKey  string
	BaseURL string
}

// NewConfigFromEnv reads provider configuration from environment variables.
func NewConfigFromEnv() Config {
	name := strings.TrimSpace(os.Getenv("CORPORATE_PROVIDER"))
	if name == "" {
		return Config{}
	}
	return Config{
		Enabled: true,
		Name:    strings.ToLower(name),
		APIKey:  strings.TrimSpace(os.Getenv("CORPORATE_PROVIDER_KEY")),
		BaseURL: strings.TrimSpace(os.Getenv("CORPORATE_PROVIDER_BASE_URL")),
	}
}

// noopProvider is the default safe provider that preserves the original text.
type noopProvider struct{}

func (noopProvider) Name() string { return "noop" }

func (noopProvider) GenerateText(_ context.Context, req PromptRequest) (string, error) {
	return req.Text, nil
}

// BuildProvider creates the configured provider implementation.
// The default behavior is a noop provider so the CLI remains deterministic when AI is not configured.
func BuildProvider(cfg Config) (Provider, error) {
	if cfg.Name == "" && !cfg.Enabled {
		return noopProvider{}, nil
	}

	name := strings.ToLower(strings.TrimSpace(cfg.Name))
	if name == "" {
		name = "noop"
	}

	switch name {
	case "noop":
		return noopProvider{}, nil
	case "mock":
		return mockProvider{}, nil
	default:
		return nil, fmt.Errorf("unsupported provider %q", cfg.Name)
	}
}

// mockProvider is a test-friendly provider used for config validation and opt-in wiring.
type mockProvider struct{}

func (mockProvider) Name() string { return "mock" }

func (mockProvider) GenerateText(_ context.Context, req PromptRequest) (string, error) {
	if strings.TrimSpace(req.Text) == "" {
		return "", nil
	}
	return req.Text, nil
}

// ErrNoProviderConfigured is returned when a remote provider is requested without configuration.
var ErrNoProviderConfigured = errors.New("no AI provider configured")
