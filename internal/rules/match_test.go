package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchRule(t *testing.T) {
	r := Match("foo")
	assert.Equal(t, KindMatch, r.Kind())

	t.Run("match", func(t *testing.T) {
		res := evalRule(r, "foo")
		assert.True(t, res.Success)
		assert.Len(t, res.Unmatched.Slice(), 0)
	})

	t.Run("no match", func(t *testing.T) {
		res := evalRule(r, "foobar")
		assert.False(t, res.Success)
		assert.Len(t, res.Unmatched.Slice(), 1)

		res = evalRule(r, "bar")
		assert.False(t, res.Success)
		assert.Len(t, res.Unmatched.Slice(), 1)
	})

	t.Run("multiple match", func(t *testing.T) {
		res := evalRule(r, "foobar", "foo")
		assert.True(t, res.Success)
		assert.ElementsMatch(t, []string{"foo"}, res.Matched.Slice())
		assert.ElementsMatch(t, []string{"foobar"}, res.Unmatched.Slice())
	})

	t.Run("multiple no match", func(t *testing.T) {
		res := evalRule(r, "foobar", "bar")
		assert.False(t, res.Success)
		assert.Len(t, res.Matched.Slice(), 0)
		assert.ElementsMatch(t, []string{"foobar", "bar"}, res.Unmatched.Slice())
	})
}
