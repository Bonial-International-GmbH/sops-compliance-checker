package rules

import "regexp"

// MatchRegexRule asserts that trust anchors match a user-defined regular
// expression.
type MatchRegexRule struct {
	meta    Meta
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

// Meta implements MetaRule.
func (r *MatchRegexRule) Meta() Meta {
	return r.meta
}

// SetMeta implements MetaRule.
func (r *MatchRegexRule) SetMeta(meta Meta) {
	r.meta = meta
}

// WithMeta implements MetaRule.
func (r *MatchRegexRule) WithMeta(meta Meta) Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements EvalRule.
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
