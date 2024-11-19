package rules

import "strings"

// OneOfRule asserts that exactly one of the nested rules matches.
type OneOfRule struct {
	meta  Meta
	rules []Rule
}

// OneOf creates a OneOfRule from zero or more rules.
func OneOf(rules ...Rule) *OneOfRule {
	return &OneOfRule{rules: rules}
}

// Describe implements Describe
func (r *OneOfRule) Describe() string {
	var sb strings.Builder
	describeRuleMeta(&sb, r.meta)
	sb.WriteString(r.DescribeSelf())
	sb.WriteString(":\n")
	describeRules(&sb, r.rules)
	return sb.String()
}

// DescribeSelf implements rule.DescribeRule.
func (r *OneOfRule) DescribeSelf() string {
	return "Must match exactly ONE of"
}

// Kind implements Describe
func (*OneOfRule) Kind() Kind {
	return KindOneOf
}

// Meta implements Meta
func (r *OneOfRule) Meta() Meta {
	return r.meta
}

// SetMeta implements Meta
func (r *OneOfRule) SetMeta(meta Meta) {
	r.meta = meta
}

// WithMeta implements Meta
func (r *OneOfRule) WithMeta(meta Meta) Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements Eval
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
