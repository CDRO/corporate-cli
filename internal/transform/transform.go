package transform

import (
	"regexp"
	"strings"
)

var whitespacePattern = regexp.MustCompile(`\s+`)
var repeatedPunctuationPattern = regexp.MustCompile(`[!?.,;:]{2,}`)
var spacingPattern = regexp.MustCompile(`\s+([,.;:!?])`)

var wholeWordReplacements = []struct {
	pattern    *regexp.Regexp
	replacement string
}{
	{pattern: regexp.MustCompile(`(?i)\btotaly\b`), replacement: "totally"},
	{pattern: regexp.MustCompile(`(?i)\btotallly\b`), replacement: "totally"},
	{pattern: regexp.MustCompile(`(?i)\bdont\b`), replacement: "don't"},
	{pattern: regexp.MustCompile(`(?i)\barent\b`), replacement: "aren't"},
	{pattern: regexp.MustCompile(`(?i)\bcant\b`), replacement: "can't"},
	{pattern: regexp.MustCompile(`(?i)\bwont\b`), replacement: "won't"},
	{pattern: regexp.MustCompile(`(?i)\bdoesnt\b`), replacement: "doesn't"},
	{pattern: regexp.MustCompile(`(?i)\bisnt\b`), replacement: "isn't"},
	{pattern: regexp.MustCompile(`(?i)\bdumbasses\b`), replacement: "team members"},
	{pattern: regexp.MustCompile(`(?i)\bclowns\b`), replacement: "stakeholders"},
	{pattern: regexp.MustCompile(`(?i)\bstupid\b`), replacement: "inconsistent"},
	{pattern: regexp.MustCompile(`(?i)\bfucking\b`), replacement: "significant"},
	{pattern: regexp.MustCompile(`(?i)\bshit\b`), replacement: "concerns"},
	{pattern: regexp.MustCompile(`(?i)\bhell\b`), replacement: "challenges"},
	{pattern: regexp.MustCompile(`(?i)\bidiot\b`), replacement: "professional"},
	{pattern: regexp.MustCompile(`(?i)\bmoron\b`), replacement: "team member"},
	{pattern: regexp.MustCompile(`(?i)\buseless\b`), replacement: "underperforming"},
	{pattern: regexp.MustCompile(`(?i)\bdisgrace\b`), replacement: "opportunity for improvement"},
	{pattern: regexp.MustCompile(`(?i)\bmess\b`), replacement: "challenge"},
}

func normalizeText(input string) string {
	text := strings.ReplaceAll(input, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	text = strings.TrimSpace(text)
	text = strings.NewReplacer(
		"’", "'",
		"‘", "'",
		"‛", "'",
		"´", "'",
	).Replace(text)

	text = whitespacePattern.ReplaceAllString(text, " ")
	for _, item := range wholeWordReplacements {
		text = item.pattern.ReplaceAllString(text, item.replacement)
	}

	text = repeatedPunctuationPattern.ReplaceAllStringFunc(text, func(match string) string {
		return string(match[len(match)-1])
	})
	text = spacingPattern.ReplaceAllString(text, "$1")

	return strings.TrimSpace(text)
}

// Corporateize rewrites text into a more corporate tone.
func Corporateize(input string) string {
	text := normalizeText(strings.TrimSpace(input))
	if text == "" {
		return ""
	}

	lower := strings.ToLower(text)
	lower = strings.NewReplacer(
		"these team members", "the project team",
		"team members", "project team members",
		"the project team members", "the project team",
		"significant challenge", "significant challenge",
		"significant challenges", "significant challenges",
		"overall situation", "overall situation",
		"whole thing is a challenge", "the overall situation reflects significant operational challenges",
		"whole thing is a mess", "the overall situation reflects significant operational challenges",
		"the deadline is getting pushed", "the delivery timeline requires alignment",
		"do not fully align with the current requirements", "do not fully align with the current requirements",
		"really bad", "substantially misaligned",
		"inconsistent decisions", "inconsistent execution decisions",
		"underperforming", "underperforming",
		"opportunity for improvement", "opportunity for improvement",
		"team members", "project team members",
		"project team members are", "the project team is",
		"don't understand", "do not fully appreciate",
		"doesn't understand", "does not fully appreciate",
		"not aligned with expectations", "not aligned with current expectations",
		"significant operational challenges", "significant operational challenges",
		"problematic", "challenging",
		"concerns", "concerns",
	).Replace(lower)

	if len(lower) > 0 {
		lower = strings.ToUpper(lower[:1]) + lower[1:]
	}
	return lower
}

// InverseCorporateize converts polished language back into a blunt, direct style.
func InverseCorporateize(input string) string {
	text := normalizeText(strings.TrimSpace(input))
	if text == "" {
		return ""
	}

	lower := strings.ToLower(text)
	replacements := []struct {
		old string
		new string
	}{
		{old: "the project team", new: "the team"},
		{old: "project team members", new: "the team"},
		{old: "team members", new: "the team"},
		{old: "appears to be experiencing significant operational challenges", new: "is a mess"},
		{old: "appears to be experiencing significant challenges", new: "is a mess"},
		{old: "appears to be experiencing", new: "is"},
		{old: "significant operational challenges", new: "a mess"},
		{old: "significant challenges", new: "a mess"},
		{old: "delivery timeline requires alignment", new: "the deadline is slipping"},
		{old: "overall situation reflects", new: "the whole thing is"},
		{old: "current expectations", new: "basic requirements"},
		{old: "does not fully appreciate", new: "doesn't understand"},
		{old: "do not fully appreciate", new: "don't understand"},
		{old: "stakeholders", new: "the people involved"},
		{old: "inconsistent execution decisions", new: "bad decisions"},
		{old: "opportunity for improvement", new: "we need to fix this"},
		{old: "requires immediate attention and improved coordination", new: "needs cleanup because nobody knows what the hell is going on"},
		{old: "requires alignment", new: "needs fixing"},
		{old: "not aligned with", new: "isn't aligned with"},
		{old: "project team", new: "team"},
	}

	for _, replacement := range replacements {
		lower = strings.ReplaceAll(lower, replacement.old, replacement.new)
	}

	lower = strings.ReplaceAll(lower, "  ", " ")
	lower = strings.TrimSpace(lower)
	if len(lower) == 0 {
		return ""
	}
	if lower[len(lower)-1] != '.' {
		lower += "."
	}
	return strings.ToUpper(lower[:1]) + lower[1:]
}
