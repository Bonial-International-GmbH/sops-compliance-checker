package rules

import "strings"

// AllOfRule asserts that all of the nested rules matches.
type AllOfRule struct {
	meta  Meta
	rules []Rule
}

// AllOf creates an AllOfRule from zero or more rules.
func AllOf(rules ...Rule) *AllOfRule {
	return &AllOfRule{rules: rules}
}

// Describe implements Describe
func (r *AllOfRule) Describe() string {
	var sb strings.Builder
	describeRuleMeta(&sb, r.meta)
	sb.WriteString(r.DescribeSelf())
	sb.WriteString(":\n")
	describeRules(&sb, r.rules)
	return sb.String()
}

// DescribeSelf implements rule.DescribeRule.
func (r *AllOfRule) DescribeSelf() string {
	return "Must match ALL of"
}

// Kind implements Describe
func (*AllOfRule) Kind() Kind {
	return KindAllOf
}

// Meta implements Meta
func (r *AllOfRule) Meta() Meta {
	return r.meta
}

// SetMeta implements Meta
func (r *AllOfRule) SetMeta(meta Meta) {
	r.meta = meta
}

// WithMeta implements Meta
func (r *AllOfRule) WithMeta(meta Meta) Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements Eval
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
