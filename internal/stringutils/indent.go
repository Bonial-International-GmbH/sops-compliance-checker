package stringutils

import "strings"

// Indent indents a string by `count` spaces.
func Indent(s string, count int, indentFirst bool) string {
	if count == 0 || s == "" {
		return s
	}

	var sb strings.Builder

	lines := strings.SplitAfter(s, "\n")

	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	indent := strings.Repeat(" ", count)

	for i, line := range lines {
		if line != "\n" && line != "\r\n" && (i != 0 || indentFirst) {
			// Only indent non-empty lines.
			sb.WriteString(indent)
		}

		sb.WriteString(line)
	}

	return sb.String()
}
