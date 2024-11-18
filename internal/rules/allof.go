package rules

import "github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"

// AllOfRule asserts that all of the nested rules matches.
type AllOfRule struct {
	meta  rule.Meta
	rules []rule.Rule
}

// AllOf creates an AllOfRule from zero or more rules.
func AllOf(rules ...rule.Rule) *AllOfRule {
	return &AllOfRule{rules: rules}
}

// Describe implements rule.DescribeRule.
func (r *AllOfRule) Describe() string {
	return ""
}

// Kind implements rule.DescribeRule.
func (*AllOfRule) Kind() rule.Kind {
	return rule.AllOf
}

// Meta implements rule.MetaRule.
func (r *AllOfRule) Meta() rule.Meta {
	return r.meta
}

// SetMeta implements rule.MetaRule.
func (r *AllOfRule) SetMeta(meta rule.Meta) {
	r.meta = meta
}

// WithMeta implements rule.MetaRule.
func (r *AllOfRule) WithMeta(meta rule.Meta) rule.Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements rule.EvalRule.
func (r *AllOfRule) Eval(ctx *rule.EvalContext) rule.EvalResult {
	matched := emptyStringSet()
	failures := 0

	results := evalRules(ctx, r.rules, func(result *rule.EvalResult) {
		if result.Success {
			matched.InsertSet(result.Matched)
		} else {
			failures++
		}
	})

	return rule.EvalResult{
		Rule:      r,
		Success:   failures == 0,
		Matched:   matched,
		Unmatched: ctx.TrustAnchors.Difference(matched),
		Nested:    results,
	}
}
