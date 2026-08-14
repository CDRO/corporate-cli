package update

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNeedsUpdate(t *testing.T) {
	if !needsUpdate("v0.1.0", "v0.2.0") {
		t.Fatal("expected update when latest is newer")
	}

	if needsUpdate("v0.2.0", "v0.2.0") {
		t.Fatal("expected no update when versions match")
	}

	if needsUpdate("v0.3.0", "v0.2.0") {
		t.Fatal("expected no update when local version is newer")
	}
}

func TestCheckForUpdate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/repos/example/project/releases/latest" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"tag_name":"v0.2.0"}`))
	}))
	defer server.Close()

	result, err := checkForRelease(server.URL+"/repos/example/project", "v0.1.0")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !result.Available {
		t.Fatal("expected update to be available")
	}

	if result.Latest != "v0.2.0" {
		t.Fatalf("got latest %q, want v0.2.0", result.Latest)
	}
}
