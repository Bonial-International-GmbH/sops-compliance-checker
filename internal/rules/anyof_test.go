package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyOf(t *testing.T) {
	r := AnyOf(Match("foo"), Match("bar"))
	assert.Equal(t, KindAnyOf, r.Kind())

	t.Run("multiple matches", func(t *testing.T) {
		res := evalRule(r, "baz", "qux", "foo", "bar", "qux")
		assert.True(t, res.Success)
		assert.ElementsMatch(t, []string{"foo", "bar"}, res.Matched.Slice())
		assert.ElementsMatch(t, []string{"baz", "qux"}, res.Unmatched.Slice())
	})

	t.Run("one match", func(t *testing.T) {
		res := evalRule(r, "foo", "baz", "qux")
		assert.True(t, res.Success)
		assert.ElementsMatch(t, []string{"foo"}, res.Matched.Slice())
		assert.ElementsMatch(t, []string{"baz", "qux"}, res.Unmatched.Slice())
	})

	t.Run("no match", func(t *testing.T) {
		res := evalRule(r, "baz", "qux")
		assert.False(t, res.Success)
		assert.Len(t, res.Matched.Slice(), 0)
		assert.Len(t, res.Unmatched.Slice(), 2)
	})
}
