package rules

import (
	"fmt"
	"strings"
)

// NotRule inverts the match result of a nested
type NotRule struct {
	meta Meta
	rule Rule
}

// Not creates a NotRule around another one.
func Not(rule Rule) *NotRule {
	return &NotRule{rule: rule}
}

// Describe implements Describe
func (r *NotRule) Describe() string {
	var sb strings.Builder
	describeRuleMeta(&sb, r.meta)
	fmt.Fprintf(&sb, "Must NOT match:\n")
	writeIndented(&sb, r.rule.Describe(), 2)
	return sb.String()
}

// Kind implements Describe
func (*NotRule) Kind() Kind {
	return KindNot
}

// Meta implements Meta
func (r *NotRule) Meta() Meta {
	return r.meta
}

// SetMeta implements Meta
func (r *NotRule) SetMeta(meta Meta) {
	r.meta = meta
}

// WithMeta implements Meta
func (r *NotRule) WithMeta(meta Meta) Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements Eval
func (r *NotRule) Eval(ctx *EvalContext) EvalResult {
	result := r.rule.Eval(ctx)

	// Invert the result.
	return EvalResult{
		Rule:      r,
		Success:   !result.Success,
		Matched:   result.Unmatched,
		Unmatched: result.Matched,
		Nested:    []EvalResult{result},
	}
}
