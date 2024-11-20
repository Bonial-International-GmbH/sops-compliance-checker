package rules

// OneOfRule asserts that exactly one of the nested rules matches.
type OneOfRule struct {
	meta  Meta
	rules []Rule
}

// OneOf creates a OneOfRule from zero or more rules.
func OneOf(rules ...Rule) *OneOfRule {
	return &OneOfRule{rules: rules}
}

// Kind implements Rule.
func (*OneOfRule) Kind() Kind {
	return KindOneOf
}

// Meta implements MetaRule.
func (r *OneOfRule) Meta() Meta {
	return r.meta
}

// SetMeta implements MetaRule.
func (r *OneOfRule) SetMeta(meta Meta) {
	r.meta = meta
}

// WithMeta implements MetaRule.
func (r *OneOfRule) WithMeta(meta Meta) Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements EvalRule.
func (r *OneOfRule) Eval(ctx *EvalContext) EvalResult {
	result := evalRules(ctx, r.rules)

	return EvalResult{
		Rule:      r,
		Success:   result.successCount == 1,
		Matched:   result.matched,
		Unmatched: ctx.TrustAnchors.Difference(result.matched),
		Nested:    result.results,
	}
}
