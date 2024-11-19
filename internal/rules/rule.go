package rules

import "github.com/hashicorp/go-set/v3"

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
	// Kind returns the kind of the rule.
	Kind() Kind
}

// EvalContext encapsulates data needed during rule evaluation, like the trust
// anchors found within a given SOPS file.
type EvalContext struct {
	// TrustAnchors is a set of trust anchors found in a SOPS file.
	TrustAnchors set.Collection[string]
}

// NewEvalContext creates a new EvalContext from a list of trust anchors.
func NewEvalContext(trustAnchors []string) *EvalContext {
	return &EvalContext{TrustAnchors: set.From(trustAnchors)}
}

// EvalResult represents the result of a rule evaluation.
type EvalResult struct {
	// Rule is the rule that produced this result.
	Rule Rule
	// Success indicates whether the rule was matched by the input or not.
	Success bool
	// Matched contains trust anchors that were matched during rule evaluation,
	// if any. This may even contain trust anchors if rule evaluation failed,
	// indicating partial matches.
	Matched set.Collection[string]
	// Unmatched contains all trust anchors not matched during rule evaluation.
	Unmatched set.Collection[string]
	// Nested contains the results of any nested rules that had to be evaluated
	// in order to produce the result. This allows identifying the exact nested
	// rules that led to evaluation success (or failure).
	Nested []EvalResult
}
