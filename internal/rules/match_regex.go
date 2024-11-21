package rules

import "regexp"

// MatchRegexRule asserts that trust anchors match a user-defined regular
// expression.
type MatchRegexRule struct {
	metaRule
	pattern *regexp.Regexp
}

// MatchRegex creates a MatchRegexRule for the given pattern.
func MatchRegex(pattern *regexp.Regexp) *MatchRegexRule {
	return &MatchRegexRule{pattern: pattern}
}

// Kind implements Rule.
func (*MatchRegexRule) Kind() Kind {
	return KindMatchRegex
}

// Eval implements Rule.
func (r *MatchRegexRule) Eval(ctx *EvalContext) EvalResult {
	matched := emptyStringSet()

	for trustAnchor := range ctx.TrustAnchors.Items() {
		if r.pattern.MatchString(trustAnchor) {
			matched.Insert(trustAnchor)
		}
	}

	return EvalResult{
		Rule:      r,
		Success:   !matched.Empty(),
		Matched:   matched,
		Unmatched: ctx.TrustAnchors.Difference(matched),
	}
}
