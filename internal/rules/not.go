package rules

// NotRule inverts the match result of a nested
type NotRule struct {
	meta Meta
	rule Rule
}

// Not creates a NotRule around another one.
func Not(rule Rule) *NotRule {
	return &NotRule{rule: rule}
}

// Kind implements Rule.
func (*NotRule) Kind() Kind {
	return KindNot
}

// Meta implements MetaRule.
func (r *NotRule) Meta() Meta {
	return r.meta
}

// SetMeta implements MetaRule.
func (r *NotRule) SetMeta(meta Meta) {
	r.meta = meta
}

// WithMeta implements MetaRule.
func (r *NotRule) WithMeta(meta Meta) Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements EvalRule.
func (r *NotRule) Eval(ctx *EvalContext) EvalResult {
	result := r.rule.Eval(ctx)

	// Invert the result.
	return EvalResult{
		Rule:      r,
		Success:   !result.Success,
		Matched:   result.Unmatched,
		Unmatched: result.Matched,
		Nested:    []EvalResult{result},
	}
}
