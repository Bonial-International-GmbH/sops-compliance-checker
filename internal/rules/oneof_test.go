package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneOf(t *testing.T) {
	r := OneOf(Match("foo"), Match("bar"))
	assert.Equal(t, KindOneOf, r.Kind())

	t.Run("exactly one", func(t *testing.T) {
		res := evalRule(r, "baz", "qux", "foo")
		assert.True(t, res.Success)
		assert.ElementsMatch(t, []string{"foo"}, res.Matched.Slice())
		assert.ElementsMatch(t, []string{"baz", "qux"}, res.Unmatched.Slice())
	})

	t.Run("none", func(t *testing.T) {
		res := evalRule(r, "baz", "qux")
		assert.False(t, res.Success)
		assert.Len(t, res.Matched.Slice(), 0)
		assert.Len(t, res.Unmatched.Slice(), 2)
	})

	t.Run("more than one", func(t *testing.T) {
		res := evalRule(r, "foo", "bar", "baz")
		assert.False(t, res.Success)
		assert.ElementsMatch(t, []string{"foo", "bar"}, res.Matched.Slice())
		assert.ElementsMatch(t, []string{"baz"}, res.Unmatched.Slice())
	})
}
