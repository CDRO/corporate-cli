package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Config holds provider auth metadata in a local file.
type Config struct {
	Provider string `json:"provider"`
	APIKey   string `json:"apiKey,omitempty"`
}

// Store writes config to a local path.
type Store struct {
	Path string
}

// DefaultStorePath returns the per-user config path for the CLI.
func DefaultStorePath() (string, error) {
	base := os.Getenv("APPDATA")
	if base == "" {
		base = os.Getenv("HOME")
	}
	if base == "" {
		return "", fmt.Errorf("unable to determine config directory")
	}
	return filepath.Join(base, "corporate", "config.json"), nil
}

// LoadConfig reads config from disk, returning an empty config if missing.
func LoadConfig(path string) (Config, error) {
	cfg := Config{}
	if path == "" {
		return cfg, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return cfg, err
	}
	if len(data) == 0 {
		return cfg, nil
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

// WriteConfig writes the config to disk.
func WriteConfig(path string, cfg Config) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	payload, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, payload, 0o600)
}

// Login persists a provider name and API key for future use.
func (s Store) Login(provider, apiKey string) error {
	provider = strings.TrimSpace(strings.ToLower(provider))
	apiKey = strings.TrimSpace(apiKey)
	if provider == "" {
		return fmt.Errorf("provider name is required")
	}
	if apiKey == "" {
		return fmt.Errorf("api key is required")
	}

	cfg := Config{Provider: provider, APIKey: apiKey}
	if s.Path == "" {
		path, err := DefaultStorePath()
		if err != nil {
			return err
		}
		s.Path = path
	}
	return WriteConfig(s.Path, cfg)
}

// Status reports whether a credential is configured.
func (s Store) Status() (string, error) {
	if s.Path == "" {
		path, err := DefaultStorePath()
		if err != nil {
			return "", err
		}
		s.Path = path
	}
	cfg, err := LoadConfig(s.Path)
	if err != nil {
		return "", err
	}
	if cfg.Provider == "" || cfg.APIKey == "" {
		return "not configured", nil
	}
	return fmt.Sprintf("configured for %s", cfg.Provider), nil
}
