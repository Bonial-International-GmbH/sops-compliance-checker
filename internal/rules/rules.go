// Package rules contains the interface for the rule engine and all supported
// rule types together with their rule evaluation logic.
package rules

// Meta describes metadata common to all available rules.
type Meta struct {
	// Description may contain the description of the rule. If the description
	// is not empty, it is used to enrich error messages presented to the user.
	Description string
	// URL may contain a link to some (internal) documentation that further
	// explains the purpose of a rule. If non-empty, it is used to enrich error
	// messages presented to the user.
	URL string
}

// Kind represents the kind of a rule.
type Kind string

const (
	// KindAllOf asserts that all of the nested rules match.
	KindAllOf Kind = "allOf"
	// AnyOf asserts that at least one of the nested rules matches.
	KindAnyOf Kind = "anyOf"
	// Match defines a string to match trust anchors against.
	KindMatch Kind = "match"
	// MatchRegex defines a regular expression to match trust anchors against.
	KindMatchRegex Kind = "matchRegex"
	// Not inverts the matching behaviour of a rule.
	KindNot Kind = "not"
	// OneOf asserts that exactly one of the nested rules matches.
	KindOneOf Kind = "oneOf"
)

// Rule is the interface implemented by all available rules.
type Rule interface {
	EvalRule
	MetaRule
	// Kind returns the kind of the rule.
	Kind() Kind
}

// EvalRule is a rule that can be evaluated.
type EvalRule interface {
	// Eval evaluates the rule using the provided EvalContext.
	Eval(ctx *EvalContext) EvalResult
}

// MetaRule provides setters and getters for rule metadata.
type MetaRule interface {
	// Meta returns the metadata associated with a rule.
	Meta() Meta
	// SetMeta sets the rule metadata.
	SetMeta(meta Meta)
}

// metaRule is used by all available rules as their implementation of MetaRule
// to reduce boilerplate.
type metaRule struct {
	meta Meta
}

// Meta implements MetaRule.
func (r *metaRule) Meta() Meta {
	return r.meta
}

// SetMeta implements MetaRule.
func (r *metaRule) SetMeta(meta Meta) {
	r.meta = meta
}

// Ensure that all rule types implement the Rule interface.
var (
	_ Rule = &AllOfRule{}
	_ Rule = &AnyOfRule{}
	_ Rule = &MatchRule{}
	_ Rule = &MatchRegexRule{}
	_ Rule = &NotRule{}
	_ Rule = &OneOfRule{}
)
