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
	DescribeRule
	MetaRule
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
	// WithMeta sets the rule metadata and returns the updated rule. This
	// method is useful in situations where a fluent interface would make
	// things more ergonomic.
	WithMeta(meta Meta) Rule
}

// DescribeRule provides methods for a rule to describe itself.
type DescribeRule interface {
	// Describe returns a human readable string, describing the rule.
	Describe() string
	// Describe returns a human readable string, describing the rule.
	DescribeSelf() string
	// Kind returns the kind of the rule.
	Kind() Kind
}
