package transform

import (
	"regexp"
	"strings"
)

var whitespacePattern = regexp.MustCompile(`\s+`)
var repeatedPunctuationPattern = regexp.MustCompile(`[!?.,;:]{2,}`)
var spacingPattern = regexp.MustCompile(`\s+([,.;:!?])`)
var uppercaseShoutPattern = regexp.MustCompile(`(?m)\b[A-Z]{3,}\b`)

var wholeWordReplacements = []struct {
	pattern    *regexp.Regexp
	replacement string
}{
	{pattern: regexp.MustCompile(`(?i)\btotaly\b`), replacement: "totally"},
	{pattern: regexp.MustCompile(`(?i)\btotallly\b`), replacement: "totally"},
	{pattern: regexp.MustCompile(`(?i)\bmangment\b`), replacement: "management"},
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
	{pattern: regexp.MustCompile(`(?i)\blazy\b`), replacement: "demonstrating limited execution efficiency"},
}

func normalizeText(input string) string {
	text := strings.ReplaceAll(input, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	text = strings.TrimSpace(text)
	if text == "" {
		return ""
	}
	text = strings.NewReplacer(
		"’", "'",
		"‘", "'",
		"‛", "'",
		"´", "'",
	).Replace(text)

	text = whitespacePattern.ReplaceAllString(text, " ")
	text = collapseRepeatedWords(text)
	text = strings.NewReplacer(
		"!!!", "!",
		"???", "?",
		"??", "?",
		"!!", "!",
		"!!! ?", "!?",
		"?? !", "?!",
	).Replace(text)
	text = repeatedPunctuationPattern.ReplaceAllStringFunc(text, func(match string) string {
		return string(match[len(match)-1])
	})
	text = spacingPattern.ReplaceAllString(text, "$1")
	text = uppercaseShoutPattern.ReplaceAllStringFunc(text, func(match string) string {
		return strings.ToLower(match)
	})
	for _, item := range wholeWordReplacements {
		text = item.pattern.ReplaceAllString(text, item.replacement)
	}

	text = strings.TrimSpace(text)
	if text == "" {
		return ""
	}
	if strings.Trim(text, "!?.,;: \n\t") == "" {
		return ""
	}
	return text
}

func collapseRepeatedWords(text string) string {
	parts := strings.Fields(text)
	if len(parts) < 3 {
		return text
	}

	out := make([]string, 0, len(parts))
	for i := 0; i < len(parts); i++ {
		if i+2 < len(parts) && strings.EqualFold(parts[i], parts[i+1]) && strings.EqualFold(parts[i+1], parts[i+2]) {
			out = append(out, parts[i])
			i += 2
			continue
		}
		out = append(out, parts[i])
	}
	return strings.Join(out, " ")
}

// Corporateize rewrites text into a more corporate tone.
func Corporateize(input string) string {
	text := normalizeText(strings.TrimSpace(input))
	if text == "" {
		return ""
	}

	lower := strings.ToLower(text)
	lower = strings.NewReplacer(
		"release-v2-final.txt", "release-v2-final.txt",
		"sla", "SLA",
		"pm", "PM",
		"cto", "CTO",
		"qa", "QA",
		"adr", "ADR",
	).Replace(lower)

	lower = strings.NewReplacer(
		"these dumbasses", "the project team",
		"dumbasses", "the project team",
		"these clowns", "these stakeholders",
		"clowns", "stakeholders",
		"mangment", "management",
		"lazy as hell", "demonstrating limited execution efficiency",
		"lazy", "demonstrating limited execution efficiency",
		"fucking mess", "significant operational challenges",
		"fucking", "significant",
		"stupid mistakes", "avoidable errors",
		"incompetent", "not meeting expected performance standards",
		"dont understand", "do not fully appreciate",
		"doesnt understand", "does not fully appreciate",
		"the whole thing is a mess", "the current situation presents significant operational challenges",
		"this is a significant challenge", "the current situation presents significant operational challenges",
		"the deadline is getting pushed", "the delivery timeline requires improved coordination",
		"deadline is getting pushed", "the delivery timeline requires improved coordination",
		"the whole thing is a challenge", "the current situation presents significant operational challenges",
		"whole thing is a challenge", "the current situation presents significant operational challenges",
		"the project team in the project team", "the project team",
		"project team in the project team", "the project team",
		"the project team are", "the project team is",
		"the project team members", "the project team",
		"team members", "project team members",
		"project team members are", "the project team is",
		"these team members", "the project team",
		"do not fully appreciate basic requirements", "do not fully appreciate the basic requirements",
		"significant challenge", "significant operational challenge",
		"significant challenges", "significant operational challenges",
		"opportunity for improvement", "an opportunity for improvement",
		"challenges", "operational challenges",
		"challenge", "operational challenge",
	).Replace(lower)

	lower = strings.NewReplacer(
		"the the ", "the ",
		"  ", " ",
		" ,", ",",
		" .", ".",
		"project team members members", "project team members",
		"the project team the project team", "the project team",
	).Replace(lower)

	lower = strings.TrimSpace(lower)
	if lower == "" {
		return ""
	}
	if lower[len(lower)-1] != '.' {
		lower += "."
	}
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
