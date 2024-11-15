package engine

import (
	"github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"
)

type Engine struct {
	rootRule rule.EvalRule
}

func New(rootRule rule.EvalRule) *Engine {
	return &Engine{rootRule}
}

func (e *Engine) Eval(trustAnchors []string) rule.EvalResult {
	ctx := rule.NewEvalContext(trustAnchors)
	return e.rootRule.Eval(ctx)
}
