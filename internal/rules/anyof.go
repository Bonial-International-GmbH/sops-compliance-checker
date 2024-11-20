package rules

// AnyOfRule asserts that at least one of the nested rules matches.
type AnyOfRule struct {
	meta  Meta
	rules []Rule
}

// AnyOf creates an AnyOfRule from zero or more rules.
func AnyOf(rules ...Rule) *AnyOfRule {
	return &AnyOfRule{rules: rules}
}

// Kind implements Rule.
func (*AnyOfRule) Kind() Kind {
	return KindAnyOf
}

// Meta implements MetaRule.
func (r *AnyOfRule) Meta() Meta {
	return r.meta
}

// SetMeta implements MetaRule.
func (r *AnyOfRule) SetMeta(meta Meta) {
	r.meta = meta
}

// WithMeta implements MetaRule.
func (r *AnyOfRule) WithMeta(meta Meta) Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements EvalRule.
func (r *AnyOfRule) Eval(ctx *EvalContext) EvalResult {
	result := evalRules(ctx, r.rules)

	return EvalResult{
		Rule:      r,
		Success:   result.successCount > 0,
		Matched:   result.matched,
		Unmatched: ctx.TrustAnchors.Difference(result.matched),
		Nested:    result.results,
	}
}
