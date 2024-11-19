package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatResult(t *testing.T) {
	rootRule := rulesFixture

	t.Run("disaster recovery key missing", func(t *testing.T) {
		trustAnchors := []string{
			"arn:aws:kms:eu-central-1:123456789012:alias/team-foo",
			"arn:aws:kms:eu-west-1:123456789012:alias/team-foo",
			"arn:aws:kms:eu-central-1:123456789012:alias/production-cicd",
			"arn:aws:kms:eu-west-1:123456789012:alias/production-cicd",
		}

		result := evalRule(rootRule, trustAnchors...)
		assert.False(t, result.Success)

		expected := `[allOf] Expected ALL of the nested rules to match, but found one failure:

  1)
    [match] Disaster recovery key must be present.

    Expected trust anchor "age1u79ltfzz5k79ex4mpl3r76p2532xex4mpl3z7vttctudr6gedn6ex4mpl3" was not found.
`

		assert.Equal(t, expected, result.Format())
	})

	t.Run("team key missing", func(t *testing.T) {
		trustAnchors := []string{
			"age1u79ltfzz5k79ex4mpl3r76p2532xex4mpl3z7vttctudr6gedn6ex4mpl3",
			"arn:aws:kms:eu-central-1:123456789012:alias/team-foo",
			"arn:aws:kms:eu-central-1:123456789012:alias/production-cicd",
			"arn:aws:kms:eu-west-1:123456789012:alias/production-cicd",
		}

		result := evalRule(rootRule, trustAnchors...)
		assert.False(t, result.Success)

		expected := `[allOf] Expected ALL of the nested rules to match, but found one failure:

  1)
    [anyOf] Expected ANY of the nested rule to match, but none did:

      1)
        [allOf] Expected ALL of the nested rules to match, but found one failure:

          1)
            [match] Expected trust anchor "arn:aws:kms:eu-west-1:123456789012:alias/team-foo" was not found.

      2)
        [allOf] Expected ALL of the nested rules to match, but found 2 failures:

          1)
            [match] Expected trust anchor "arn:aws:kms:eu-central-1:123456789012:alias/team-bar" was not found.

          2)
            [match] Expected trust anchor "arn:aws:kms:eu-west-1:123456789012:alias/team-bar" was not found.
`

		assert.Equal(t, expected, result.Format())
	})

	t.Run("team key missing (2)", func(t *testing.T) {
		trustAnchors := []string{
			"age1u79ltfzz5k79ex4mpl3r76p2532xex4mpl3z7vttctudr6gedn6ex4mpl3",
			"arn:aws:kms:eu-central-1:123456789012:alias/team-foo",
			"arn:aws:kms:eu-west-1:123456789012:alias/team-foo",
			"arn:aws:kms:eu-central-1:123456789012:alias/team-bar",
			"arn:aws:kms:eu-central-1:123456789012:alias/production-cicd",
			"arn:aws:kms:eu-west-1:123456789012:alias/production-cicd",
		}

		result := evalRule(rootRule, trustAnchors...)
		assert.True(t, result.Success)
		assert.Equal(t, "", result.Format())
		assert.ElementsMatch(
			t,
			[]string{"arn:aws:kms:eu-central-1:123456789012:alias/team-bar"},
			result.Unmatched.Slice(),
		)
	})

	t.Run("keys for multiple environments", func(t *testing.T) {
		trustAnchors := []string{
			"age1u79ltfzz5k79ex4mpl3r76p2532xex4mpl3z7vttctudr6gedn6ex4mpl3",
			"arn:aws:kms:eu-central-1:123456789012:alias/team-foo",
			"arn:aws:kms:eu-west-1:123456789012:alias/team-foo",
			"arn:aws:kms:eu-central-1:123456789012:alias/production-cicd",
			"arn:aws:kms:eu-west-1:123456789012:alias/production-cicd",
			"arn:aws:kms:eu-central-1:123456789012:alias/staging-cicd",
			"arn:aws:kms:eu-west-1:123456789012:alias/staging-cicd",
		}

		result := evalRule(rootRule, trustAnchors...)
		assert.False(t, result.Success)

		expected := `[allOf] Expected ALL of the nested rules to match, but found one failure:

  1)
    [oneOf] Expected EXACTLY ONE nested rule to match, but found 2:

      1)
        [allOf] Matched trust anchors:
          - arn:aws:kms:eu-central-1:123456789012:alias/production-cicd
          - arn:aws:kms:eu-west-1:123456789012:alias/production-cicd

      2)
        [allOf] Matched trust anchors:
          - arn:aws:kms:eu-central-1:123456789012:alias/staging-cicd
          - arn:aws:kms:eu-west-1:123456789012:alias/staging-cicd
`

		assert.Equal(t, expected, result.Format())
	})
}
