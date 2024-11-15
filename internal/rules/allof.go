package rules

import "github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"

type AllOfRule struct {
	meta  rule.Meta
	rules []rule.Rule
}

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
	results := make([]rule.EvalResult, len(r.rules))
	matched := emptyStringSet()
	failures := 0

	for i, rule := range r.rules {
		result := rule.Eval(ctx)

		if result.Success {
			matched = matched.Union(result.Matched)
		} else {
			failures += 1
		}

		results[i] = result
	}

	return rule.EvalResult{
		Rule:      r,
		Success:   failures == 0,
		Matched:   matched,
		Unmatched: ctx.TrustAnchors.Difference(matched),
		Nested:    results,
	}
}
