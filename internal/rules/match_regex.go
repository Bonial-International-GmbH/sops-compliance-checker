package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"
)

// MatchRegexRule asserts that trust anchors match a user-defined regular
// expression.
type MatchRegexRule struct {
	meta    rule.Meta
	pattern *regexp.Regexp
}

// MatchRegex creates a MatchRegexRule for the given pattern.
func MatchRegex(pattern *regexp.Regexp) *MatchRegexRule {
	return &MatchRegexRule{pattern: pattern}
}

// Describe implements rule.DescribeRule.
func (r *MatchRegexRule) Describe() string {
	var sb strings.Builder
	describeRuleMeta(&sb, r.meta)
	fmt.Fprintf(&sb, "Must match the pattern %q\n", r.pattern)
	return sb.String()
}

// Kind implements rule.DescribeRule.
func (*MatchRegexRule) Kind() rule.Kind {
	return rule.MatchRegex
}

// Meta implements rule.MetadataRule.
func (r *MatchRegexRule) Meta() rule.Meta {
	return r.meta
}

// SetMeta implements rule.MetadataRule.
func (r *MatchRegexRule) SetMeta(meta rule.Meta) {
	r.meta = meta
}

// WithMeta implements rule.MetaRule.
func (r *MatchRegexRule) WithMeta(meta rule.Meta) rule.Rule {
	r.SetMeta(meta)
	return r
}

// Eval implements rule.EvalRule.
func (r *MatchRegexRule) Eval(ctx *rule.EvalContext) rule.EvalResult {
	matched := emptyStringSet()

	for trustAnchor := range ctx.TrustAnchors.Items() {
		if r.pattern.MatchString(trustAnchor) {
			matched.Insert(trustAnchor)
		}
	}

	return rule.EvalResult{
		Rule:      r,
		Success:   !matched.Empty(),
		Matched:   matched,
		Unmatched: ctx.TrustAnchors.Difference(matched),
	}
}
