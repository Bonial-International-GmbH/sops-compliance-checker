package rules

import (
	"strings"

	"github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"
)

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
	var sb strings.Builder
	describeRuleMeta(&sb, r.meta)
	sb.WriteString("Must match ALL of:\n")
	describeRules(&sb, r.rules)
	return sb.String()
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
	result := evalRules(ctx, r.rules)

	return rule.EvalResult{
		Rule:      r,
		Success:   result.successCount == len(r.rules),
		Matched:   result.matched,
		Unmatched: ctx.TrustAnchors.Difference(result.matched),
		Nested:    result.results,
	}
}