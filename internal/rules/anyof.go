package rules

import "strings"

// AnyOfRule asserts that at least one of the nested rules matches.
type AnyOfRule struct {
	meta  Meta
	rules []Rule
}

// AnyOf creates an AnyOfRule from zero or more rules.
func AnyOf(rules ...Rule) *AnyOfRule {
	return &AnyOfRule{rules: rules}
}

// Describe implements Describe
func (r *AnyOfRule) Describe() string {
	var sb strings.Builder
	describeRuleMeta(&sb, r.meta)
	sb.WriteString(r.DescribeSelf())
	sb.WriteString(":\n")
	describeRules(&sb, r.rules)
	return sb.String()
}

// DescribeSelf implements rule.DescribeRule.
func (r *AnyOfRule) DescribeSelf() string {
	return "Must match ANY of"
}

// Kind implements Describe
func (*AnyOfRule) Kind() Kind {
	return KindAnyOf
}

// Meta implements Meta
func (r *AnyOfRule) Meta() Meta {
	return r.meta
}

// SetMeta implements Meta
func (r *AnyOfRule) SetMeta(meta Meta) {
	r.meta = meta
}

// WithMeta implements Meta
func (r *AnyOfRule) WithMeta(meta Meta) Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements Eval
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
