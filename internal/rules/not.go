package rules

import (
	"fmt"
	"strings"

	"github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"
)

// NotRule inverts the match result of a nested rule.
type NotRule struct {
	meta rule.Meta
	rule rule.Rule
}

// Not creates a NotRule around another one.
func Not(rule rule.Rule) *NotRule {
	return &NotRule{rule: rule}
}

// Describe implements rule.DescribeRule.
func (r *NotRule) Describe() string {
	var sb strings.Builder
	describeRuleMeta(&sb, r.meta)
	fmt.Fprintf(&sb, "Must NOT match:\n")
	writeIndented(&sb, r.rule.Describe(), 2)
	return sb.String()
}

// Kind implements rule.DescribeRule.
func (*NotRule) Kind() rule.Kind {
	return rule.Not
}

// Meta implements rule.MetaRule.
func (r *NotRule) Meta() rule.Meta {
	return r.meta
}

// SetMeta implements rule.MetaRule.
func (r *NotRule) SetMeta(meta rule.Meta) {
	r.meta = meta
}

// WithMeta implements rule.MetaRule.
func (r *NotRule) WithMeta(meta rule.Meta) rule.Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements rule.EvalRule.
func (r *NotRule) Eval(ctx *rule.EvalContext) rule.EvalResult {
	result := r.rule.Eval(ctx)

	// Invert the result.
	return rule.EvalResult{
		Rule:      r,
		Success:   !result.Success,
		Matched:   result.Unmatched,
		Unmatched: result.Matched,
		Nested:    []rule.EvalResult{result},
	}
}
