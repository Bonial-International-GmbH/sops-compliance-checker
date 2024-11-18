package rules

import (
	"testing"

	"github.com/Bonial-International-GmbH/sops-compliance-checker/internal/rule"
	"github.com/stretchr/testify/assert"
)

func evalRule(r rule.Rule, trustAnchors ...string) rule.EvalResult {
	ctx := rule.NewEvalContext(trustAnchors)
	return r.Eval(ctx)
}

func TestNestedRules(t *testing.T) {
	rootRule := rulesFixture

	t.Run("no trust anchors", func(t *testing.T) {
		result := evalRule(rootRule)
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

		result := evalRule(rootRule, trustAnchors...)
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

		result := evalRule(rootRule, trustAnchors...)
		assert.True(t, result.Success)
		assert.Len(t, result.Matched.Slice(), len(trustAnchors)-1)
		assert.Equal(t, []string{"i don't belong here"}, result.Unmatched.Slice())
	})
}

func TestDescribeNestedRules(t *testing.T) {
	desc := rulesFixture.Describe()

	expected := `Must match ALL of:
  1)
    Disaster recovery key must be present.

    Must include trust anchor "age1u79ltfzz5k79ex4mpl3r76p2532xex4mpl3z7vttctudr6gedn6ex4mpl3"
  2)
    Must match ANY of:
      1)
        Must match ALL of:
          1)
            Must include trust anchor "arn:aws:kms:eu-central-1:123456789012:alias/team-foo"
          2)
            Must include trust anchor "arn:aws:kms:eu-west-1:123456789012:alias/team-foo"
      2)
        Must match ALL of:
          1)
            Must include trust anchor "arn:aws:kms:eu-central-1:123456789012:alias/team-bar"
          2)
            Must include trust anchor "arn:aws:kms:eu-west-1:123456789012:alias/team-bar"
  3)
    Must match exactly ONE of:
      1)
        Must match ALL of:
          1)
            Must include trust anchor "arn:aws:kms:eu-central-1:123456789012:alias/production-cicd"
          2)
            Must include trust anchor "arn:aws:kms:eu-west-1:123456789012:alias/production-cicd"
      2)
        Must match ALL of:
          1)
            Must include trust anchor "arn:aws:kms:eu-central-1:123456789012:alias/staging-cicd"
          2)
            Must include trust anchor "arn:aws:kms:eu-west-1:123456789012:alias/staging-cicd"
`

	assert.Equal(t, expected, desc)
}
