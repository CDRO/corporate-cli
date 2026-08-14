package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"corporate-cli/internal/auth"
	"corporate-cli/internal/transform"
	"corporate-cli/internal/update"
)

const defaultAppVersion = "v0.1.0"

func main() {
	if err := run(os.Args[1:], os.Stdin, os.Stdout, filepath.Base(os.Args[0])); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func currentAppVersion() string {
	if version := strings.TrimSpace(os.Getenv("CORPORATE_VERSION")); version != "" {
		return version
	}
	if version := strings.TrimSpace(os.Getenv("VERSION")); version != "" {
		return version
	}
	return defaultAppVersion
}

func releaseRepoURL() string {
	if repo := strings.TrimSpace(os.Getenv("CORPORATE_GITHUB_REPO")); repo != "" {
		return normalizeReleaseRepoURL(repo)
	}
	if repo := strings.TrimSpace(os.Getenv("GITHUB_REPOSITORY")); repo != "" {
		return "https://api.github.com/repos/" + strings.Trim(repo, "/")
	}
	return "https://api.github.com/repos/CDRO/corporate-cli"
}

func normalizeReleaseRepoURL(repo string) string {
	repo = strings.TrimSpace(repo)
	repo = strings.TrimSuffix(repo, "/")
	repo = strings.TrimPrefix(repo, "https://api.github.com/repos/")
	repo = strings.TrimPrefix(repo, "http://api.github.com/repos/")
	repo = strings.TrimPrefix(repo, "https://github.com/")
	repo = strings.TrimPrefix(repo, "http://github.com/")
	repo = strings.TrimPrefix(repo, "/")
	if repo == "" {
		return "https://api.github.com/repos/CDRO/corporate-cli"
	}
	return "https://api.github.com/repos/" + repo
}

type cliConfig struct {
	DefaultProvider string `json:"defaultProvider"`
	DefaultStyle    string `json:"defaultStyle"`
	AIEnabled       bool   `json:"aiEnabled"`
	Provider        string `json:"provider,omitempty"`
	APIKey          string `json:"apiKey,omitempty"`
}

func defaultConfigPath() string {
	path, err := auth.DefaultStorePath()
	if err != nil {
		return filepath.Join(os.TempDir(), "corporate-config.json")
	}
	return path
}

func loadCLIConfig(path string) (cliConfig, error) {
	cfg := cliConfig{}
	if path == "" {
		path = defaultConfigPath()
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

func writeCLIConfig(path string, cfg cliConfig) error {
	if path == "" {
		path = defaultConfigPath()
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	payload, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, payload, 0o600)
}

func resolvedStyle(args []string, config cliConfig) string {
	for i := 0; i < len(args)-1; i++ {
		if args[i] == "--style" {
			return strings.ToLower(strings.TrimSpace(args[i+1]))
		}
	}
	if config.DefaultStyle != "" {
		return strings.ToLower(strings.TrimSpace(config.DefaultStyle))
	}
	return "neutral"
}

func run(args []string, stdin io.Reader, stdout io.Writer, appName string) error {
	inverse := isInverseMode(appName, args)
	if len(args) > 0 && args[0] == "--inverse" {
		inverse = true
		args = args[1:]
	}
	args = stripInverseFlag(args)

	configPath := defaultConfigPath()
	cleanArgs := make([]string, 0, len(args))
	for i := 0; i < len(args); i++ {
		if args[i] == "--config" {
			if i+1 < len(args) {
				configPath = args[i+1]
				i++
			}
			continue
		}
		cleanArgs = append(cleanArgs, args[i])
	}
	args = cleanArgs

	cfg, err := loadCLIConfig(configPath)
	if err != nil {
		return err
	}
	style := resolvedStyle(args, cfg)

	if len(args) == 0 {
		data, err := io.ReadAll(stdin)
		if err != nil {
			return err
		}
		text := strings.TrimSpace(string(data))
		if text == "" {
			return printHelp(stdout, inverse)
		}
		_, err = fmt.Fprintln(stdout, transformText(text, inverse, style))
		return err
	}

	switch args[0] {
	case "update":
		if inverse {
			return fmt.Errorf("inverse mode does not support update commands")
		}
		return handleUpdate(args[1:], stdout)
	case "login":
		if inverse {
			return fmt.Errorf("inverse mode does not support login commands")
		}
		if len(args) < 3 {
			return fmt.Errorf("usage: corporate login <provider> <api-key>")
		}
		store := auth.Store{Path: configPath}
		if err := store.Login(args[1], args[2]); err != nil {
			return err
		}
		cfg.DefaultProvider = strings.TrimSpace(strings.ToLower(args[1]))
		cfg.Provider = cfg.DefaultProvider
		cfg.APIKey = args[2]
		cfg.AIEnabled = true
		if err := writeCLIConfig(configPath, cfg); err != nil {
			return err
		}
		_, err = fmt.Fprintf(stdout, "Login saved for %s\n", cfg.DefaultProvider)
		return err
	case "logout":
		if inverse {
			return fmt.Errorf("inverse mode does not support logout commands")
		}
		cfg.DefaultProvider = ""
		cfg.AIEnabled = false
		if err := os.Remove(configPath); err != nil && !os.IsNotExist(err) {
			return err
		}
		_, err = fmt.Fprintln(stdout, "Logged out and cleared local credentials.")
		return err
	case "provider":
		if inverse {
			return fmt.Errorf("inverse mode does not support provider commands")
		}
		return handleProvider(args[1:], configPath, stdout)
	case "config":
		if inverse {
			return fmt.Errorf("inverse mode does not support config commands")
		}
		return handleConfig(args[1:], configPath, stdout)
	case "--help", "-h", "help":
		return printHelp(stdout, inverse)
	case "--input":
		if len(args) < 2 {
			return fmt.Errorf("missing value for --input")
		}
		if args[1] == "-" {
			data, err := io.ReadAll(stdin)
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(stdout, transformText(string(data), inverse, style))
			return err
		}
		data, err := os.ReadFile(args[1])
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(stdout, transformText(string(data), inverse, style))
		return err
	case "--style":
		if len(args) < 2 {
			return fmt.Errorf("missing value for --style")
		}
		style = strings.TrimSpace(strings.ToLower(args[1]))
		cfg.DefaultStyle = style
		if err := writeCLIConfig(configPath, cfg); err != nil {
			return err
		}
		if len(args) == 2 {
			return printHelp(stdout, inverse)
		}
		data, err := io.ReadAll(stdin)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(stdout, transformText(string(data), inverse, cfg.DefaultStyle))
		return err
	default:
		if inverse {
			return fmt.Errorf("unknown inverse command %q", args[0])
		}
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func isInverseMode(appName string, args []string) bool {
	base := strings.TrimSuffix(filepath.Base(appName), filepath.Ext(filepath.Base(appName)))
	if strings.EqualFold(base, "etaroproc") {
		return true
	}
	for _, arg := range args {
		if arg == "--inverse" {
			return true
		}
	}
	return false
}

func stripInverseFlag(args []string) []string {
	filtered := make([]string, 0, len(args))
	for _, arg := range args {
		if arg == "--inverse" {
			continue
		}
		filtered = append(filtered, arg)
	}
	return filtered
}

func transformText(input string, inverse bool, style string) string {
	if inverse {
		return transform.InverseCorporateize(input)
	}
	style = strings.TrimSpace(strings.ToLower(style))
	switch style {
	case "neutral", "":
		return transform.Corporateize(input)
	case "formal", "executive", "concise", "polite", "balanced":
		return transform.Corporateize(input)
	default:
		return transform.Corporateize(input)
	}
}

func handleProvider(args []string, configPath string, stdout io.Writer) error {
	if len(args) == 0 || args[0] == "list" {
		_, err := fmt.Fprintln(stdout, "Supported providers: mock, noop, openai")
		return err
	}
	if len(args) >= 2 && args[0] == "set" {
		cfg, err := loadCLIConfig(configPath)
		if err != nil {
			return err
		}
		cfg.DefaultProvider = strings.TrimSpace(strings.ToLower(args[1]))
		cfg.AIEnabled = true
		if err := writeCLIConfig(configPath, cfg); err != nil {
			return err
		}
		_, err = fmt.Fprintf(stdout, "Active provider set to %s\n", cfg.DefaultProvider)
		return err
	}
	return fmt.Errorf("usage: corporate provider list | corporate provider set <name>")
}

func handleConfig(args []string, configPath string, stdout io.Writer) error {
	cfg, err := loadCLIConfig(configPath)
	if err != nil {
		return err
	}
	if len(args) == 0 || args[0] == "show" {
		payload, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(stdout, string(payload))
		return err
	}
	if len(args) >= 2 && args[0] == "set" {
		key := args[1]
		value := args[2]
		if key == "defaultStyle" {
			cfg.DefaultStyle = strings.TrimSpace(strings.ToLower(value))
		} else if key == "defaultProvider" {
			cfg.DefaultProvider = strings.TrimSpace(strings.ToLower(value))
		} else if key == "aiEnabled" {
			cfg.AIEnabled = strings.EqualFold(value, "true") || strings.EqualFold(value, "1")
		} else {
			return fmt.Errorf("unknown config key %q", key)
		}
		if err := writeCLIConfig(configPath, cfg); err != nil {
			return err
		}
		_, err = fmt.Fprintf(stdout, "Updated config: %s=%s\n", key, value)
		return err
	}
	return fmt.Errorf("usage: corporate config show | corporate config set <key> <value>")
}

func printHelp(w io.Writer, inverse bool) error {
	if inverse {
		_, err := fmt.Fprint(w, strings.Join([]string{
			"etaroproc CLI",
			"",
			"Usage:",
			"  etaroproc --help",
			"  etaroproc --input input.txt",
			"  etaroproc < input.txt > output.txt",
			"",
			"Inverse mode turns polished language back into blunt, direct wording.",
		}, "\n"))
		return err
	}

	_, err := fmt.Fprint(w, strings.Join([]string{
		"corporate CLI",
		"",
		"Usage:",
		"  corporate < input.txt",
		"  corporate --input input.txt",
		"  corporate --style executive < input.txt",
		"  corporate login <provider> <api-key>",
		"  corporate provider list",
		"  corporate provider set openai",
		"  corporate config show",
		"  corporate update --check",
		"  corporate update --install",
		"  corporate --help",
		"",
		"The default behavior stays deterministic and safe when AI is not configured.",
		"Provider and style settings can be stored locally for repeatable use.",
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
		result, err := update.CheckForRelease(releaseRepoURL(), currentAppVersion())
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
