package rules

// MatchRule asserts that a trust anchor exactly matches a user-defined string.
type MatchRule struct {
	metaRule
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

// Eval implements Rule.
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
