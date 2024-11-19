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

func (b *formatBuffer) writeKindPrefix(kind Kind) {
	b.WriteRune('[')
	b.WriteString(string(kind))
	b.WriteString("] ")
}

func (b *formatBuffer) writeRuleMeta(meta Meta) {
	desc := strings.TrimSpace(meta.Description)

	if desc != "" {
		b.WriteString(desc)
		b.WriteString("\n\n")
	}

	url := strings.TrimSpace(meta.URL)

	if url != "" {
		b.WriteString("Further information: ")
		b.WriteString(url)
		b.WriteString("\n\n")
	}
}

func (b *formatBuffer) writeIndented(fn func(*formatBuffer)) {
	var buf formatBuffer
	fn(&buf)
	writeIndented(&b.Builder, buf.String(), 2)
}

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

func formatFailure(buf *formatBuffer, result *EvalResult) {
	buf.writeKindPrefix(result.Rule.Kind())
	buf.writeRuleMeta(result.Rule.Meta())

	successes, failures := result.partitionNested()

	switch r := result.Rule.(type) {
	case *MatchRule:
		fmt.Fprintf(buf, "Expected trust anchor %q was not found.\n", r.trustAnchor)
	case *MatchRegexRule:
		fmt.Fprintf(buf, "Trust anchor matching regular expression %q was not found.\n", r.pattern.String())
	case *NotRule:
		buf.WriteString("Expected nested rule to fail, but it did not:\n")

		if len(result.Nested) > 0 {
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

func formatUnexpectedSuccess(buf *formatBuffer, result *EvalResult) {
	buf.writeKindPrefix(result.Rule.Kind())
	buf.WriteString("Matched trust anchors:\n")

	trustAnchors := result.Matched.Slice()
	sort.Strings(trustAnchors)

	formatTrustAnchors(buf, result.Matched)
}

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
