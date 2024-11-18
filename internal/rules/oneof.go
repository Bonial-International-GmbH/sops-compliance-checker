package rules

import (
	"github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"
)

// OneOfRule asserts that exactly one of the nested rules matches.
type OneOfRule struct {
	meta  rule.Meta
	rules []rule.Rule
}

// OneOf creates a OneOfRule from zero or more rules.
func OneOf(rules ...rule.Rule) *OneOfRule {
	return &OneOfRule{rules: rules}
}

// Describe implements rule.DescribeRule.
func (r *OneOfRule) Describe() string {
	return ""
}

// Kind implements rule.DescribeRule.
func (*OneOfRule) Kind() rule.Kind {
	return rule.OneOf
}

// Meta implements rule.MetaRule.
func (r *OneOfRule) Meta() rule.Meta {
	return r.meta
}

// SetMeta implements rule.MetaRule.
func (r *OneOfRule) SetMeta(meta rule.Meta) {
	r.meta = meta
}

// WithMeta implements rule.MetaRule.
func (r *OneOfRule) WithMeta(meta rule.Meta) rule.Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements rule.EvalRule.
func (r *OneOfRule) Eval(ctx *rule.EvalContext) rule.EvalResult {
	matched := emptyStringSet()
	successes := 0

	results := evalRules(ctx, r.rules, func(result *rule.EvalResult) {
		if result.Success {
			matched.InsertSet(result.Matched)
			successes++
		}
	})

	return rule.EvalResult{
		Rule:      r,
		Success:   successes == 1,
		Matched:   matched,
		Unmatched: ctx.TrustAnchors.Difference(matched),
		Nested:    results,
	}
}
