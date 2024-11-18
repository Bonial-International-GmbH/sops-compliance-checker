package rules

import "github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"

// AnyOfRule asserts that at least one of the nested rules matches.
type AnyOfRule struct {
	meta  rule.Meta
	rules []rule.Rule
}

// AnyOf creates an AnyOfRule from zero or more rules.
func AnyOf(rules ...rule.Rule) *AnyOfRule {
	return &AnyOfRule{rules: rules}
}

// Describe implements rule.DescribeRule.
func (r *AnyOfRule) Describe() string {
	return ""
}

// Kind implements rule.DescribeRule.
func (*AnyOfRule) Kind() rule.Kind {
	return rule.AnyOf
}

// Meta implements rule.MetaRule.
func (r *AnyOfRule) Meta() rule.Meta {
	return r.meta
}

// SetMeta implements rule.MetaRule.
func (r *AnyOfRule) SetMeta(meta rule.Meta) {
	r.meta = meta
}

// WithMeta implements rule.MetaRule.
func (r *AnyOfRule) WithMeta(meta rule.Meta) rule.Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements rule.EvalRule.
func (r *AnyOfRule) Eval(ctx *rule.EvalContext) rule.EvalResult {
	result := evalRules(ctx, r.rules)

	return rule.EvalResult{
		Rule:      r,
		Success:   result.successCount > 0,
		Matched:   result.matched,
		Unmatched: ctx.TrustAnchors.Difference(result.matched),
		Nested:    result.results,
	}
}
