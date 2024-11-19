package rules

import (
	"fmt"
	"strings"
)

// MatchRule asserts that a trust anchor exactly matches a user-defined string.
type MatchRule struct {
	meta        Meta
	trustAnchor string
}

// Match create a MatchRule for the expected trust anchor.
func Match(trustAnchor string) *MatchRule {
	return &MatchRule{trustAnchor: trustAnchor}
}

// Describe implements Describe
func (r *MatchRule) Describe() string {
	var sb strings.Builder
	describeRuleMeta(&sb, r.meta)
	fmt.Fprintf(&sb, "Must include trust anchor %q\n", r.trustAnchor)
	return sb.String()
}

// Kind implements Describe
func (*MatchRule) Kind() Kind {
	return KindMatch
}

// Meta implements Metadata
func (r *MatchRule) Meta() Meta {
	return r.meta
}

// SetMeta implements Metadata
func (r *MatchRule) SetMeta(meta Meta) {
	r.meta = meta
}

// WithMeta implements Meta
func (r *MatchRule) WithMeta(meta Meta) Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements Eval
func (r *MatchRule) Eval(ctx *EvalContext) EvalResult {
	matched := emptyStringSet()

	for trustAnchor := range ctx.TrustAnchors.Items() {
		if r.trustAnchor == trustAnchor {
			matched.Insert(trustAnchor)
		}
	}

	return EvalResult{
		Rule:      r,
		Success:   !matched.Empty(),
		Matched:   matched,
		Unmatched: ctx.TrustAnchors.Difference(matched),
	}
}
