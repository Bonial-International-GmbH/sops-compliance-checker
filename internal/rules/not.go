package rules

// NotRule inverts the match result of a nested
type NotRule struct {
	metaRule
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
