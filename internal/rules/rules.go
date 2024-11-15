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

func emptyStringSet() set.Collection[string] {
	return set.From([]string{})
}
