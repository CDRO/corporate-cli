package transform

import "testing"

func TestNormalizeTextCleansCommonNoise(t *testing.T) {
	input := "  totaly  dont!!!  this   is a fucking mess???  "
	got := normalizeText(input)

	if got == input {
		t.Fatal("normalizeText should change the input")
	}
	if !containsFold(got, "totally") {
		t.Fatalf("expected normalized text to contain totally, got %q", got)
	}
	if !containsFold(got, "don't") {
		t.Fatalf("expected normalized text to contain don't, got %q", got)
	}
	if !containsFold(got, "!") {
		t.Fatalf("expected punctuation to be normalized, got %q", got)
	}
}

func TestNormalizeTextHandlesRegexStrategyCases(t *testing.T) {
	input := "THIS IS A DISASTER!!!\r\nvery very very bad???"
	got := normalizeText(input)

	if containsFold(got, "!!!") {
		t.Fatalf("expected repeated punctuation to be collapsed, got %q", got)
	}
	if containsFold(got, "very very") {
		t.Fatalf("expected repeated words to be condensed, got %q", got)
	}
	if !containsFold(got, "this is a disaster") {
		t.Fatalf("expected shout normalization to preserve the message, got %q", got)
	}
}

func containsFold(s, substr string) bool {
	if substr == "" {
		return true
	}
	if len(s) < len(substr) {
		return false
	}
	for i := 0; i+len(substr) <= len(s); i++ {
		if equalFoldASCII(s[i:i+len(substr)], substr) {
			return true
		}
	}
	return false
}

func equalFoldASCII(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if toLowerASCII(a[i]) != toLowerASCII(b[i]) {
			return false
		}
	}
	return true
}

func toLowerASCII(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		return ch + ('a' - 'A')
	}
	return ch
}
