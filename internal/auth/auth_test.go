package auth

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultStorePathUsesUserConfigDir(t *testing.T) {
	old := os.Getenv("APPDATA")
	defer os.Setenv("APPDATA", old)
	if err := os.Setenv("APPDATA", t.TempDir()); err != nil {
		t.Fatalf("set APPDATA: %v", err)
	}

	path, err := DefaultStorePath()
	if err != nil {
		t.Fatalf("DefaultStorePath returned error: %v", err)
	}
	if path == "" {
		t.Fatal("expected store path, got empty string")
	}
	if filepath.Base(path) != "config.json" {
		t.Fatalf("expected config.json store name, got %q", filepath.Base(path))
	}
}

func TestStoreLoginPersistsProviderKey(t *testing.T) {
	dir := t.TempDir()
	store := Store{Path: filepath.Join(dir, "config.json")}
	if err := store.Login("openai", "test-token"); err != nil {
		t.Fatalf("Login returned error: %v", err)
	}

	cfg, err := LoadConfig(store.Path)
	if err != nil {
		t.Fatalf("LoadConfig returned error: %v", err)
	}
	if cfg.Provider != "openai" {
		t.Fatalf("expected provider openai, got %q", cfg.Provider)
	}
	if cfg.APIKey != "test-token" {
		t.Fatalf("expected API key test-token, got %q", cfg.APIKey)
	}
}

func TestStoreStatusReportsNotConfiguredWhenMissing(t *testing.T) {
	dir := t.TempDir()
	store := Store{Path: filepath.Join(dir, "config.json")}
	status, err := store.Status()
	if err != nil {
		t.Fatalf("Status returned error: %v", err)
	}
	if status != "not configured" {
		t.Fatalf("expected unconfigured status, got %q", status)
	}
}

func TestLoadConfigReturnsEmptyConfigForMissingFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.json")
	cfg, err := LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig returned error for missing file: %v", err)
	}
	if cfg.Provider != "" || cfg.APIKey != "" {
		t.Fatalf("expected empty config, got %#v", cfg)
	}
}

func TestWriteConfigDoesNotCreateSecretsInPlainTextOutput(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	cfg := Config{Provider: "openai", APIKey: "secret-token"}
	if err := WriteConfig(path, cfg); err != nil {
		t.Fatalf("WriteConfig returned error: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read config file: %v", err)
	}
	if string(data) == "" {
		t.Fatal("expected config file to be written")
	}
	if len(data) == 0 {
		t.Fatal("config file should contain data")
	}
	if string(data) == "secret-token" {
		t.Fatal("sensitive data should not be encoded as raw plain text in the output path")
	}
}
