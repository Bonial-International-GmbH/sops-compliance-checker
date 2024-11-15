package rules

import "github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"

// MatchRule asserts that a trust anchor exactly matches a user-defined string.
type MatchRule struct {
	meta        rule.Meta
	trustAnchor string
}

// Match create a MatchRule for the expected trust anchor.
func Match(trustAnchor string) *MatchRule {
	return &MatchRule{trustAnchor: trustAnchor}
}

// Describe implements rule.DescribeRule.
func (r *MatchRule) Describe() string {
	return ""
}

// Kind implements rule.DescribeRule.
func (*MatchRule) Kind() rule.Kind {
	return rule.Match
}

// Meta implements rule.MetadataRule.
func (r *MatchRule) Meta() rule.Meta {
	return r.meta
}

// SetMeta implements rule.MetadataRule.
func (r *MatchRule) SetMeta(meta rule.Meta) {
	r.meta = meta
}

// WithMeta implements rule.MetaRule.
func (r *MatchRule) WithMeta(meta rule.Meta) rule.Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements rule.EvalRule.
func (r *MatchRule) Eval(ctx *rule.EvalContext) rule.EvalResult {
	matched := emptyStringSet()

	for trustAnchor := range ctx.TrustAnchors.Items() {
		if r.trustAnchor == trustAnchor {
			matched.Insert(trustAnchor)
		}
	}

	return rule.EvalResult{
		Rule:      r,
		Success:   !matched.Empty(),
		Matched:   matched,
		Unmatched: ctx.TrustAnchors.Difference(matched),
	}
}
