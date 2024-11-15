package rules

import (
	"testing"

	"github.com/Bonial-International-GmbH/sops-compliance-checker/internal/engine"
	"github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"
	"github.com/stretchr/testify/assert"
)

func evalRule(r rule.Rule, trustAnchors ...string) rule.EvalResult {
	ctx := rule.NewEvalContext(trustAnchors)
	return r.Eval(ctx)
}

func TestCompile(t *testing.T) {
	t.Run("compiles fixture", func(t *testing.T) {
		compiled, err := Compile(configFixture.Rules)
		assert.NoError(t, err)
		assert.Equal(t, rulesFixture, compiled)
	})
}

func TestNestedRules(t *testing.T) {
	engine := engine.New(rulesFixture)

	t.Run("all trust anchors missing", func(t *testing.T) {
		trustAnchors := []string{}

		result := engine.Eval(trustAnchors)
		assert.False(t, result.Success)
		assert.Len(t, result.Matched.Slice(), 0)
		assert.Len(t, result.Unmatched.Slice(), 0)
	})

	t.Run("all trust anchors present", func(t *testing.T) {
		trustAnchors := []string{
			"age1u79ltfzz5k79ex4mpl3r76p2532xex4mpl3z7vttctudr6gedn6ex4mpl3",
			"arn:aws:kms:eu-central-1:123456789012:alias/team-foo",
			"arn:aws:kms:eu-west-1:123456789012:alias/team-foo",
			"arn:aws:kms:eu-central-1:123456789012:alias/production-cicd",
			"arn:aws:kms:eu-west-1:123456789012:alias/production-cicd",
		}

		result := engine.Eval(trustAnchors)
		assert.True(t, result.Success)
		assert.Len(t, result.Matched.Slice(), len(trustAnchors))
		assert.Len(t, result.Unmatched.Slice(), 0)
	})

	t.Run("excess trust anchor", func(t *testing.T) {
		trustAnchors := []string{
			"age1u79ltfzz5k79ex4mpl3r76p2532xex4mpl3z7vttctudr6gedn6ex4mpl3",
			"arn:aws:kms:eu-central-1:123456789012:alias/team-foo",
			"arn:aws:kms:eu-west-1:123456789012:alias/team-foo",
			"arn:aws:kms:eu-central-1:123456789012:alias/production-cicd",
			"arn:aws:kms:eu-west-1:123456789012:alias/production-cicd",
			"i don't belong here",
		}

		result := engine.Eval(trustAnchors)
		assert.True(t, result.Success)
		assert.Len(t, result.Matched.Slice(), len(trustAnchors)-1)
		assert.Equal(t, []string{"i don't belong here"}, result.Unmatched.Slice())
	})
}
