package rules

import (
	"fmt"

	"github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"
	"github.com/Bonial-International-GmbH/sops-compliance-checker/pkg/config"
)

// Ensure that all rule types implement the rule.Rule interface.
var (
	_ rule.Rule = &AllOfRule{}
	_ rule.Rule = &AnyOfRule{}
	_ rule.Rule = &MatchRule{}
	_ rule.Rule = &NotRule{}
	_ rule.Rule = &OneOfRule{}
)

// Compile takes a slice of rule configurations and compiles it into a single
// rule that can be evaluated.
func Compile(rules []config.Rule) (root rule.Rule, err error) {
	compiled, err := compileRules(rules)
	if err != nil {
		return nil, err
	}

	return AllOf(compiled...), nil
}

func compileRules(rules []config.Rule) ([]rule.Rule, error) {
	compiled := make([]rule.Rule, len(rules))

	for i, rule := range rules {
		compiledRule, err := compileRule(rule)
		if err != nil {
			return nil, err
		}

		compiled[i] = compiledRule
	}

	return compiled, nil
}

func compileRule(config config.Rule) (rule.Rule, error) {
	compiled, err := compileRuleInner(config)
	if err != nil {
		return nil, err
	}

	compiled.SetMeta(rule.Meta{
		Description: config.Description,
		URL:         config.URL,
	})

	return compiled, nil
}

func compileRuleInner(rule config.Rule) (rule.Rule, error) {
	if rule.Match != "" {
		return Match(rule.Match), nil
	}

	if rule.Not != nil {
		inner, err := compileRule(*rule.Not)
		if err != nil {
			return nil, err
		}

		return Not(inner), nil
	}

	if len(rule.AllOf) > 0 {
		rules, err := compileRules(rule.AllOf)
		if err != nil {
			return nil, err
		}

		return AllOf(rules...), nil
	}

	if len(rule.AnyOf) > 0 {
		rules, err := compileRules(rule.AnyOf)
		if err != nil {
			return nil, err
		}

		return AnyOf(rules...), nil
	}

	if len(rule.OneOf) > 0 {
		rules, err := compileRules(rule.OneOf)
		if err != nil {
			return nil, err
		}

		return OneOf(rules...), nil
	}

	return nil, fmt.Errorf("rule %v has no conditions", rule)
}
