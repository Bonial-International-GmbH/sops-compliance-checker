package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllOfRule(t *testing.T) {
	r := AllOf(Match("foo"), Match("bar"))
	assert.Equal(t, KindAllOf, r.Kind())

	t.Run("all match", func(t *testing.T) {
		res := evalRule(r, "foo", "bar")
		assert.True(t, res.Success)
		assert.Len(t, res.Matched.Slice(), 2)
		assert.Len(t, res.Unmatched.Slice(), 0)
	})

	t.Run("all match with excess", func(t *testing.T) {
		res := evalRule(r, "foo", "bar", "qux")
		assert.True(t, res.Success)
		assert.ElementsMatch(t, []string{"foo", "bar"}, res.Matched.Slice())
		assert.ElementsMatch(t, []string{"qux"}, res.Unmatched.Slice())
	})

	t.Run("one missing", func(t *testing.T) {
		res := evalRule(r, "foo", "qux")
		assert.False(t, res.Success)
		assert.ElementsMatch(t, []string{"foo"}, res.Matched.Slice())
		assert.ElementsMatch(t, []string{"qux"}, res.Unmatched.Slice())
	})
}
