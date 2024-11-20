package rules

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/go-set/v3"
)

// formatBuffer is a helper type for formatting EvalResults.
type formatBuffer struct {
	strings.Builder
}

// writeIndented passes a *formatBuffer to fn which will indent every line
// written to it by 2 spaces.
func (b *formatBuffer) writeIndented(fn func(*formatBuffer)) {
	// Pass a temporary buffer to the closure to capture the written bytes.
	var buf formatBuffer
	fn(&buf)

	// Indent the captured bytes and write them to the underlying
	// strings.Builder.
	writeIndented(&b.Builder, buf.String(), 2)
}

// writeIndentedList iterates the list of results and invokes fn for each
// result, passing a *formatBuffer which will indent every line written to it
// by 2 spaces.
func (b *formatBuffer) writeIndentedList(results []EvalResult, fn func(*formatBuffer, *EvalResult)) {
	b.writeIndented(func(buf *formatBuffer) {
		for i, result := range results {
			buf.WriteRune('\n')
			fmt.Fprintf(buf, "%d)\n", i+1)
			buf.writeIndented(func(buf *formatBuffer) {
				fn(buf, &result)
			})
		}
	})
}

// formatFailure formats a failed result and writes the human readable
// representation to buf. The caller must ensure that result.Success is false.
func formatFailure(buf *formatBuffer, result *EvalResult) {
	result = flattenResult(result)

	formatRuleKind(buf, result.Rule.Kind())
	formatRuleMeta(buf, result.Rule.Meta())

	successes, failures := result.partitionNested()

	switch r := result.Rule.(type) {
	case *MatchRule:
		fmt.Fprintf(buf, "Expected trust anchor %q was not found.\n", r.trustAnchor)
	case *MatchRegexRule:
		fmt.Fprintf(buf, "Trust anchor matching regular expression %q was not found.\n", r.pattern.String())
	case *NotRule:
		buf.WriteString("Expected nested rule to fail, but it did not:\n")

		if len(result.Nested) > 0 {
			buf.WriteRune('\n')
			buf.writeIndented(func(buf *formatBuffer) {
				formatUnexpectedSuccess(buf, &result.Nested[0])
			})
		}
	case *AllOfRule:
		buf.WriteString("Expected ALL of the nested rules to match, but found ")

		if len(failures) == 1 {
			buf.WriteString("one failure:\n")
		} else {
			fmt.Fprintf(buf, "%d failures:\n", len(failures))
		}

		buf.writeIndentedList(failures, formatFailure)
	case *AnyOfRule:
		buf.WriteString("Expected ANY of the nested rule to match, but none did:\n")
		buf.writeIndentedList(failures, formatFailure)
	case *OneOfRule:
		buf.WriteString("Expected EXACTLY ONE nested rule to match, but ")

		if len(successes) == 0 {
			buf.WriteString("none did:\n")
			buf.writeIndentedList(failures, formatFailure)
		} else {
			fmt.Fprintf(buf, "found %d:\n", len(successes))
			buf.writeIndentedList(successes, formatUnexpectedSuccess)
		}
	}
}

// formatUnexpectedSuccess format a result that was expected to fail, but
// succeeded unexpectedly, and writes the human readable representation to buf.
func formatUnexpectedSuccess(buf *formatBuffer, result *EvalResult) {
	formatRuleKind(buf, result.Rule.Kind())
	formatRuleMeta(buf, result.Rule.Meta())

	trustAnchors := result.Matched.Slice()
	sort.Strings(trustAnchors)

	buf.WriteString("Matched trust anchors:\n")
	formatTrustAnchors(buf, result.Matched)
}

// formatRuleKind writes the formatted rule kind to buf.
func formatRuleKind(buf *formatBuffer, kind Kind) {
	buf.WriteRune('[')
	buf.WriteString(string(kind))
	buf.WriteString("] ")
}

// formatRuleKind writes formatted rule metadata to buf, if any.
func formatRuleMeta(buf *formatBuffer, meta Meta) {
	desc := strings.TrimSpace(meta.Description)

	if desc != "" {
		buf.WriteString(desc)
		buf.WriteString("\n\n")
	}

	url := strings.TrimSpace(meta.URL)

	if url != "" {
		buf.WriteString("Further information: ")
		buf.WriteString(url)
		buf.WriteString("\n\n")
	}
}

// formatTrustAnchors produces a sorted and properly indented list of trust
// anchors and writes it to buf.
func formatTrustAnchors(buf *formatBuffer, items set.Collection[string]) {
	trustAnchors := items.Slice()
	sort.Strings(trustAnchors)

	for _, trustAnchor := range trustAnchors {
		buf.writeIndented(func(buf *formatBuffer) {
			buf.WriteString("- ")
			buf.WriteString(trustAnchor)
		})
		buf.WriteRune('\n')
	}
}

// flattenResult flattens results of compound rules (allOf, anyOf, oneOf) into
// their first nested result, if there's only one. This avoids unnecessary
// nesting in the human readable output to make it less verbose.
func flattenResult(result *EvalResult) *EvalResult {
	switch result.Rule.(type) {
	case *AllOfRule, *AnyOfRule, *OneOfRule:
		if len(result.Nested) == 1 {
			return &result.Nested[0]
		}
	}

	return result
}

// writeIndented writes a string indented by `count` spaces to a strings.Builder.
func writeIndented(sb *strings.Builder, s string, count int) {
	if count == 0 || s == "" {
		return
	}

	lines := strings.SplitAfter(s, "\n")

	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	indent := strings.Repeat(" ", count)

	for _, line := range lines {
		if line != "\n" && line != "\r\n" {
			// Only indent non-empty lines.
			sb.WriteString(indent)
		}

		sb.WriteString(line)
	}
}
