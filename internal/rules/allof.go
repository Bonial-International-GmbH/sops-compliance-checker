package rules

// AllOfRule asserts that all of the nested rules matches.
type AllOfRule struct {
	meta  Meta
	rules []Rule
}

// AllOf creates an AllOfRule from zero or more rules.
func AllOf(rules ...Rule) *AllOfRule {
	return &AllOfRule{rules: rules}
}

// Kind implements Rule
func (*AllOfRule) Kind() Kind {
	return KindAllOf
}

// Meta implements MetaRule
func (r *AllOfRule) Meta() Meta {
	return r.meta
}

// SetMeta implements MetaRule
func (r *AllOfRule) SetMeta(meta Meta) {
	r.meta = meta
}

// WithMeta implements MetaRule
func (r *AllOfRule) WithMeta(meta Meta) Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements EvalRule
func (r *AllOfRule) Eval(ctx *EvalContext) EvalResult {
	result := evalRules(ctx, r.rules)

	return EvalResult{
		Rule:      r,
		Success:   result.successCount == len(r.rules),
		Matched:   result.matched,
		Unmatched: ctx.TrustAnchors.Difference(result.matched),
		Nested:    result.results,
	}
}
