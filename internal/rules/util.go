// Package rules contains the interface for the rule engine and all supported
// rule types together with their rule evaluation logic.
package rules

import "github.com/hashicorp/go-set/v3"

// Ensure that all rule types implement the Rule interface.
var (
	_ Rule = &AllOfRule{}
	_ Rule = &AnyOfRule{}
	_ Rule = &MatchRule{}
	_ Rule = &MatchRegexRule{}
	_ Rule = &NotRule{}
	_ Rule = &OneOfRule{}
)

// emptyStringSet is a helper to create an empty string set. This is mainly
// used to avoid verbose type hints at the call sites because set.From returns
// a set.Set, but we actually work with the set.Collection interface.
func emptyStringSet() set.Collection[string] {
	return set.From([]string{})
}
