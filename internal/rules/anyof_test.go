package rules

import (
	"testing"

	"github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"
	"github.com/stretchr/testify/assert"
)

func TestAnyOf(t *testing.T) {
	r := AnyOf(Match("foo"), Match("bar"))
	assert.Equal(t, rule.AnyOf, r.Kind())

	t.Run("first match wins", func(t *testing.T) {
		res := evalRule(r, "foo", "bar", "baz")
		assert.True(t, res.Success)
		assert.ElementsMatch(t, []string{"foo"}, res.Matched.Slice())
		assert.ElementsMatch(t, []string{"bar", "baz"}, res.Unmatched.Slice())
	})

	t.Run("match in the middle", func(t *testing.T) {
		res := evalRule(r, "baz", "qux", "foo", "bar", "qux")
		assert.True(t, res.Success)
		assert.ElementsMatch(t, []string{"foo"}, res.Matched.Slice())
		assert.ElementsMatch(t, []string{"bar", "baz", "qux"}, res.Unmatched.Slice())
	})

	t.Run("one matches", func(t *testing.T) {
		res := evalRule(r, "foo", "qux")
		assert.True(t, res.Success)
		assert.ElementsMatch(t, []string{"foo"}, res.Matched.Slice())
		assert.ElementsMatch(t, []string{"qux"}, res.Unmatched.Slice())
	})

	t.Run("none matches", func(t *testing.T) {
		res := evalRule(r, "baz", "qux")
		assert.False(t, res.Success)
		assert.Len(t, res.Matched.Slice(), 0)
		assert.Len(t, res.Unmatched.Slice(), 2)
	})
}
