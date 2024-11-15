package rule

import "github.com/hashicorp/go-set/v3"

type Meta struct {
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
	EvalRule
	DescribeRule
	MetaRule
}

type EvalRule interface {
	Eval(ctx *EvalContext) EvalResult
}

type MetaRule interface {
	Meta() Meta
	SetMeta(meta Meta)
	WithMeta(meta Meta) Rule
}

type DescribeRule interface {
	Describe() string
	Kind() Kind
}

type EvalContext struct {
	TrustAnchors set.Collection[string]
}

func NewEvalContext(trustAnchors []string) *EvalContext {
	return &EvalContext{TrustAnchors: set.From(trustAnchors)}
}

type EvalResult struct {
	Rule      Rule
	Success   bool
	Matched   set.Collection[string]
	Unmatched set.Collection[string]
	Nested    []EvalResult
}
