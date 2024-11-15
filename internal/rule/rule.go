package rule

import "github.com/hashicorp/go-set/v3"

type Metadata struct {
	Description string
	URL         string
}

type Kind string

const (
	AllOf Kind = "allOf"
	AnyOf Kind = "anyOf"
	OneOf Kind = "oneOf"
	Match Kind = "match"
	Not   Kind = "not"
)

type Rule interface {
	Evaluate(ctx *EvalContext) EvalResult

	Metadata() Metadata

	Kind() Kind
}

type EvalContext struct {
	TrustAnchors set.Collection[string]
}

type EvalResult struct {
	Rule      Rule
	Success   bool
	Matched   set.Collection[string]
	Unmatched set.Collection[string]
	Nested    []EvalResult
}
