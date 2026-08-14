package transform

import (
	"regexp"
	"strings"
)

var whitespacePattern = regexp.MustCompile(`\s+`)
var repeatedPunctuationPattern = regexp.MustCompile(`[!?.,;:]{2,}`)
var spacingPattern = regexp.MustCompile(`\s+([,.;:!?])`)

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
	text = strings.NewReplacer(
		"totaly", "totally",
		"totallly", "totally",
		"dont", "don't",
		"arent", "aren't",
		"cant", "can't",
		"wont", "won't",
		"doesnt", "doesn't",
		"isnt", "isn't",
		"dumbasses", "team members",
		"clowns", "stakeholders",
		"stupid", "inconsistent",
		"fucking", "significant",
		"shit", "concerns",
		"hell", "challenges",
		"idiot", "professional",
		"moron", "team member",
		"useless", "underperforming",
		"disgrace", "opportunity for improvement",
		"mess", "challenge",
	).Replace(text)

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
	lower = strings.NewReplacer(
		"the project team", "team",
		"project team members", "team",
		"appears to be experiencing significant challenges", "is a mess",
		"significant challenges", "mess",
		"significant operational challenges", "mess",
		"delivery timeline requires alignment", "deadline is slipping",
		"overall situation reflects", "the whole thing is",
		"does not fully appreciate", "doesn't understand",
		"current expectations", "requirements",
		"stakeholders", "clowns",
		"inconsistent execution decisions", "bad decisions",
		"concerns", "issues",
		"team is", "team is",
	).Replace(lower)

	if len(lower) == 0 {
		return ""
	}
	if lower[len(lower)-1] != '.' {
		lower += "."
	}
	return strings.ToUpper(lower[:1]) + lower[1:]
}
