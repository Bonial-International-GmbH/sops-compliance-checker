package engine

import (
	"github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"
	"github.com/hashicorp/go-set/v3"
)

type Engine struct {
	rootRule rule.Rule
}

func New(rootRule rule.Rule) *Engine {
	return &Engine{rootRule}
}

func (e *Engine) Evaluate(trustAnchors []string) rule.EvalResult {
	ctx := rule.EvalContext{TrustAnchors: set.From(trustAnchors)}
	return e.rootRule.Evaluate(&ctx)
}
