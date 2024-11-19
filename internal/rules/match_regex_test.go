package rules

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchRegexRule(t *testing.T) {
	r := MatchRegex(regexp.MustCompile("^(foo|bar)"))
	assert.Equal(t, KindMatchRegex, r.Kind())

	t.Run("match", func(t *testing.T) {
		res := evalRule(r, "foo")
		assert.True(t, res.Success)
		assert.Len(t, res.Unmatched.Slice(), 0)
	})

	t.Run("match prefix", func(t *testing.T) {
		res := evalRule(r, "foobar")
		assert.True(t, res.Success)
		assert.Len(t, res.Unmatched.Slice(), 0)

		res = evalRule(r, "bar")
		assert.True(t, res.Success)
		assert.Len(t, res.Unmatched.Slice(), 0)

		res = evalRule(r, "bazfoo")
		assert.False(t, res.Success)
		assert.Len(t, res.Unmatched.Slice(), 1)
	})

	t.Run("multiple matches", func(t *testing.T) {
		res := evalRule(r, "foobar", "foo", "bazfoo", "barfoo")
		assert.True(t, res.Success)
		assert.ElementsMatch(t, []string{"foo", "foobar", "barfoo"}, res.Matched.Slice())
		assert.ElementsMatch(t, []string{"bazfoo"}, res.Unmatched.Slice())
	})

	t.Run("no match", func(t *testing.T) {
		res := evalRule(r, "qux", "bazbar")
		assert.False(t, res.Success)
		assert.Len(t, res.Matched.Slice(), 0)
		assert.ElementsMatch(t, []string{"qux", "bazbar"}, res.Unmatched.Slice())
	})
}
