package update

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ReleaseInfo models the GitHub release metadata required for a safe update check.
type ReleaseInfo struct {
	Available bool
	Latest    string
	Message   string
}

func needsUpdate(currentVersion, latestVersion string) bool {
	if strings.TrimSpace(currentVersion) == "" || strings.TrimSpace(latestVersion) == "" {
		return false
	}

	current := parseVersion(currentVersion)
	latest := parseVersion(latestVersion)
	if len(current) != 3 || len(latest) != 3 {
		return false
	}

	for i := 0; i < 3; i++ {
		if latest[i] > current[i] {
			return true
		}
		if latest[i] < current[i] {
			return false
		}
	}

	return false
}

func parseVersion(version string) []int {
	version = strings.TrimPrefix(version, "v")
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return []int{-1, -1, -1}
	}

	parsed := make([]int, 3)
	for i, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			return []int{-1, -1, -1}
		}
		parsed[i] = value
	}
	return parsed
}

func CheckForRelease(repoURL, currentVersion string) (ReleaseInfo, error) {
	return checkForRelease(repoURL, currentVersion)
}

func checkForRelease(repoURL, currentVersion string) (ReleaseInfo, error) {
	endpoint := strings.TrimRight(repoURL, "/") + "/releases/latest"
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return ReleaseInfo{Available: false, Message: "update check failed: invalid release URL"}, nil
	}
	request.Header.Set("Accept", "application/vnd.github+json")
	request.Header.Set("User-Agent", "corporate-cli")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(request)
	if err != nil {
		return ReleaseInfo{Available: false, Message: "update check failed: unable to reach release API"}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return ReleaseInfo{Available: false, Message: "update check failed: no published release found yet"}, nil
	}
	if resp.StatusCode != http.StatusOK {
		return ReleaseInfo{Available: false, Message: fmt.Sprintf("update check failed: release API returned %s", resp.Status)}, nil
	}

	var payload struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return ReleaseInfo{Available: false, Message: "update check failed: malformed release metadata"}, nil
	}

	if strings.TrimSpace(payload.TagName) == "" {
		return ReleaseInfo{Available: false, Message: "update check failed: missing release tag"}, nil
	}

	available := needsUpdate(currentVersion, payload.TagName)
	return ReleaseInfo{
		Available: available,
		Latest:    payload.TagName,
		Message: func() string {
			if available {
				return fmt.Sprintf("update available: %s", payload.TagName)
			}
			return "no update available"
		}(),
	}, nil
}
