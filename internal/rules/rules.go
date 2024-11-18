// Package rules contains all supported rule types together with their rule
// evaluation logic.
package rules

import (
	"github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"
	"github.com/hashicorp/go-set/v3"
)

// Ensure that all rule types implement the rule.Rule interface.
var (
	_ rule.Rule = &AllOfRule{}
	_ rule.Rule = &AnyOfRule{}
	_ rule.Rule = &MatchRule{}
	_ rule.Rule = &NotRule{}
	_ rule.Rule = &OneOfRule{}
)

// emptyStringSet is a helper to create an empty string set. This is mainly
// used to avoid verbose type hints at the call sites because set.From returns
// a set.Set, but we actually work with the set.Collection interface.
func emptyStringSet() set.Collection[string] {
	return set.From([]string{})
}

// evalRulesResult is a helper type returned by evalRules.
type evalRulesResult struct {
	results      []rule.EvalResult
	matched      set.Collection[string]
	successCount int
}

// evalRules evaluates a slice of rules and collects the results along with the
// number of successes and a set of matched trust anchors.
func evalRules(ctx *rule.EvalContext, rules []rule.Rule) evalRulesResult {
	matched := emptyStringSet()
	successCount := 0
	results := make([]rule.EvalResult, len(rules))

	for i, rule := range rules {
		result := rule.Eval(ctx)

		if result.Success {
			matched.InsertSet(result.Matched)
			successCount++
		}

		results[i] = result
	}

	return evalRulesResult{results, matched, successCount}
}
