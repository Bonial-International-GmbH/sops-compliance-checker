package rules

// MatchRule asserts that a trust anchor exactly matches a user-defined string.
type MatchRule struct {
	meta        Meta
	trustAnchor string
}

// Match create a MatchRule for the expected trust anchor.
func Match(trustAnchor string) *MatchRule {
	return &MatchRule{trustAnchor: trustAnchor}
}

// Kind implements Rule.
func (*MatchRule) Kind() Kind {
	return KindMatch
}

// Meta implements MetaRule.
func (r *MatchRule) Meta() Meta {
	return r.meta
}

// SetMeta implements MetaRule.
func (r *MatchRule) SetMeta(meta Meta) {
	r.meta = meta
}

// WithMeta implements MetaRule.
func (r *MatchRule) WithMeta(meta Meta) Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements EvalRule.
func (r *MatchRule) Eval(ctx *EvalContext) EvalResult {
	matched := emptyStringSet()

	for trustAnchor := range ctx.TrustAnchors.Items() {
		if r.trustAnchor == trustAnchor {
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
