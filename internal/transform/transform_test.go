package transform

import "testing"

func TestCorporateizeRewritesHarshText(t *testing.T) {
	input := `these dumbasses in the project team are totaly incompetent, the whole thing is a fucking mess, and the deadline is getting pushed because these clowns keep making stupid mistakes and dont even understand basic requirements.`

	output := Corporateize(input)

	if output == input {
		t.Fatal("expected corporateized output to change the input")
	}
	if !containsFold(output, "project team") {
		t.Fatalf("output %q did not contain project team", output)
	}
	if !containsFold(output, "challenges") && !containsFold(output, "coordination") {
		t.Fatalf("output %q did not contain a corporate rewrite signal", output)
	}
	for _, banned := range []string{"dumbasses", "fucking", "stupid"} {
		if containsFold(output, banned) {
			t.Fatalf("output %q should not retain the harsh phrase %q", output, banned)
		}
	}
}

func TestCorporateizeHandlesEmptyInput(t *testing.T) {
	if got := Corporateize(""); got != "" {
		t.Fatalf("expected empty string, got %q", got)
	}
	if got := Corporateize("   \n\t  "); got != "" {
		t.Fatalf("expected whitespace-only input to be empty, got %q", got)
	}
	if got := Corporateize("!!!"); got != "" {
		t.Fatalf("expected punctuation-only input to be empty, got %q", got)
	}
}

func TestCorporateizePreservesTechnicalNames(t *testing.T) {
	input := "PM says SLA is broken and release-v2-final.txt is late"
	got := Corporateize(input)
	if !containsFold(got, "sla") && !containsFold(got, "release-v2-final.txt") {
		t.Fatalf("expected technical names to be preserved, got %q", got)
	}
}

func TestInverseCorporateize(t *testing.T) {
	input := "The project team appears to be experiencing significant challenges in meeting current expectations."
	output := InverseCorporateize(input)
	if output == input {
		t.Fatal("expected inverse transform to change the wording")
	}
	if !containsFold(output, "team") || !containsFold(output, "mess") {
		t.Fatalf("inverse output %q should preserve team context and blunt tone", output)
	}
}

func TestCorporateizeMatchesREADMEExampleIntent(t *testing.T) {
	input := `these dumbasses in the project team are totaly incompetent, the whole thing is a fucking mess, and the deadline is getting pushed because these clowns keep making stupid mistakes and dont even understand basic requirements.`
	output := Corporateize(input)
	for _, want := range []string{"project team", "significant", "challenge", "coordination", "requirements"} {
		if !containsFold(output, want) {
			t.Fatalf("output %q should include %q to match the corporate rewrite intent", output, want)
		}
	}
}

func TestCorporateizeCoversMilestoneExamples(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "harsh collective criticism",
			input: "these dumbasses are totaly incompetent and this is a fucking mess",
			want:  []string{"project team", "significant", "operational", "challenge"},
		},
		{
			name:  "misspelling and weak phrase cleanup",
			input: "the mangment is lazy as hell and dont understand basic requirements",
			want:  []string{"management", "execution efficiency", "basic requirements", "understand"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Corporateize(tt.input)
			for _, want := range tt.want {
				if !containsFold(got, want) {
					t.Fatalf("Corporateize(%q) = %q, want substring %q", tt.input, got, want)
				}
			}
			for _, banned := range []string{"dumbasses", "fucking", "stupid", "hell", "lazy"} {
				if containsFold(got, banned) {
					t.Fatalf("Corporateize(%q) = %q, should not retain harsh phrase %q", tt.input, got, banned)
				}
			}
		})
	}
}

func TestCorporateizeCoversMilestone15FixtureMatrix(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
		banned []string
	}{
		{
			name:  "basic transformation",
			input: "these dumbasses in the project team are totaly incompetent, the whole thing is a fucking mess, and the deadline is getting pushed because these clowns keep making stupid mistakes and dont even understand basic requirements.",
			want:  []string{"project team", "not meeting expected performance standards", "significant operational challenge", "requirements", "coordination"},
			banned: []string{"dumbasses", "fucking", "clowns", "stupid"},
		},
		{
			name:  "misspelling normalization",
			input: "totaly dont arent mangment are bad bad bad",
			want:  []string{"totally", "don't", "aren't", "management"},
			banned: []string{"totaly", "dont", "arent", "mangment"},
		},
		{
			name:  "punctuation and shouting",
			input: "THIS IS A DISASTER!!! THIS IS A FUCKING MESS!!!",
			want:  []string{"this is a disaster", "significant", "operational", "challenge"},
			banned: []string{"!!!", "fucking"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Corporateize(tt.input)
			for _, want := range tt.want {
				if !containsFold(got, want) {
					t.Fatalf("Corporateize(%q) = %q, want substring %q", tt.input, got, want)
				}
			}
			for _, banned := range tt.banned {
				if containsFold(got, banned) {
					t.Fatalf("Corporateize(%q) = %q, should not retain banned phrase %q", tt.input, got, banned)
				}
			}
		})
	}
}

