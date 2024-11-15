package rules

import "github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"

type AnyOfRule struct {
	meta  rule.Meta
	rules []rule.Rule
}

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
	results := make([]rule.EvalResult, len(r.rules))

	for i, rule := range r.rules {
		results[i] = rule.Eval(ctx)
	}

	for _, result := range results {
		if result.Success {
			return rule.EvalResult{
				Rule:      r,
				Success:   true,
				Matched:   result.Matched,
				Unmatched: result.Unmatched,
				Nested:    results,
			}
		}
	}

	return rule.EvalResult{
		Rule:      r,
		Success:   false,
		Matched:   emptyStringSet(),
		Unmatched: ctx.TrustAnchors,
		Nested:    results,
	}
}
