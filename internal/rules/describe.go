package rules

import (
	"fmt"
	"strings"

	"github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"
)

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

// describeRuleMeta writes a description of the rule metadata to a
// strings.Builder.
func describeRuleMeta(sb *strings.Builder, meta rule.Meta) {
	desc := strings.TrimSpace(meta.Description)

	if desc != "" {
		sb.WriteString(desc)
		sb.WriteString("\n\n")
	}

	url := strings.TrimSpace(meta.URL)

	if url != "" {
		sb.WriteString("Further information: ")
		sb.WriteString(url)
		sb.WriteString("\n\n")
	}
}

// describeRules writes a list of rule descriptions to a strings.Builder.
func describeRules(sb *strings.Builder, rules []rule.Rule) {
	for i, rule := range rules {
		fmt.Fprintf(sb, "  %d)\n", i+1)
		writeIndented(sb, rule.Describe(), 4)
	}
}
