package rules

import (
	"testing"

	"github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"
	"github.com/stretchr/testify/assert"
)

func TestNotRule(t *testing.T) {
	r := Not(Match("foo"))
	assert.Equal(t, rule.Not, r.Kind())

	t.Run("match", func(t *testing.T) {
		res := evalRule(r, "bar", "foobar")
		assert.True(t, res.Success)
		assert.Len(t, res.Matched.Slice(), 2)
		assert.Len(t, res.Unmatched.Slice(), 0)
	})

	t.Run("no match", func(t *testing.T) {
		res := evalRule(r, "foo")
		assert.False(t, res.Success)
		assert.Len(t, res.Unmatched.Slice(), 1)
	})

	t.Run("multiple match", func(t *testing.T) {
		res := evalRule(r, "foobar", "bar", "foo")
		assert.False(t, res.Success)
		assert.ElementsMatch(t, []string{"foobar", "bar"}, res.Matched.Slice())
		assert.ElementsMatch(t, []string{"foo"}, res.Unmatched.Slice())
	})
}
