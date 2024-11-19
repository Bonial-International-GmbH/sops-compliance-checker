package rules

import "github.com/hashicorp/go-set/v3"

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

// partitionNested partitions nested results into success and failue.
func (r *EvalResult) partitionNested() (successes, failures []EvalResult) {
	for _, result := range r.Nested {
		if result.Success {
			successes = append(successes, result)
		} else {
			failures = append(failures, result)
		}
	}

	return
}

// Format formats the EvalResult as a human readable string.
func (r *EvalResult) Format() string {
	var buf formatBuffer

	if !r.Success {
		formatFailure(&buf, r)
	}

	if !r.Unmatched.Empty() {
		if !r.Success {
			// Leave some space between the failure output and the unmatched
			// trust anchors.
			buf.WriteRune('\n')
		}

		buf.WriteString("Unmatched trust anchors:\n")
		formatTrustAnchors(&buf, r.Unmatched)
	}

	return buf.String()
}
