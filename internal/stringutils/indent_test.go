package stringutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndent(t *testing.T) {
	t.Run("no indent", func(t *testing.T) {
		assert.Equal(t, "foo", Indent("foo", 0, true))
	})

	t.Run("empty string", func(t *testing.T) {
		assert.Equal(t, "", Indent("", 2, true))
	})

	t.Run("multiline string", func(t *testing.T) {
		given := `foo
bar

  baz`
		expected := `  foo
  bar

    baz`

		assert.Equal(t, expected, Indent(given, 2, true))
	})

	t.Run("skip indent first", func(t *testing.T) {
		given := `foo
bar

  baz`
		expected := `foo
  bar

    baz`

		assert.Equal(t, expected, Indent(given, 2, false))
	})

	t.Run("trailing newline", func(t *testing.T) {
		given := "foo\nbar\n"
		expected := "  foo\n  bar\n"

		assert.Equal(t, expected, Indent(given, 2, true))
	})
}
