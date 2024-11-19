package rules

import "github.com/Bonial-International-GmbH/sops-compliance-checker/pkg/config"

var configFixture = config.Config{
	Rules: []config.Rule{
		{
			Description: "Disaster recovery key must be present.",
			Match:       "age1u79ltfzz5k79ex4mpl3r76p2532xex4mpl3z7vttctudr6gedn6ex4mpl3",
		},
		{
			AnyOf: []config.Rule{
				{
					AllOf: []config.Rule{
						{
							Match: "arn:aws:kms:eu-central-1:123456789012:alias/team-foo",
						},
						{
							Match: "arn:aws:kms:eu-west-1:123456789012:alias/team-foo",
						},
					},
				},
				{
					AllOf: []config.Rule{
						{
							Match: "arn:aws:kms:eu-central-1:123456789012:alias/team-bar",
						},
						{
							Match: "arn:aws:kms:eu-west-1:123456789012:alias/team-bar",
						},
					},
				},
			},
		},
		{
			OneOf: []config.Rule{
				{
					AllOf: []config.Rule{
						{
							Match: "arn:aws:kms:eu-central-1:123456789012:alias/production-cicd",
						},
						{
							Match: "arn:aws:kms:eu-west-1:123456789012:alias/production-cicd",
						},
					},
				},
				{
					AllOf: []config.Rule{
						{
							Match: "arn:aws:kms:eu-central-1:123456789012:alias/staging-cicd",
						},
						{
							Match: "arn:aws:kms:eu-west-1:123456789012:alias/staging-cicd",
						},
					},
				},
			},
		},
	},
}

var rulesFixture = AllOf(
	Match("age1u79ltfzz5k79ex4mpl3r76p2532xex4mpl3z7vttctudr6gedn6ex4mpl3").
		WithMeta(Meta{Description: "Disaster recovery key must be present."}),
	AnyOf(
		AllOf(
			Match("arn:aws:kms:eu-central-1:123456789012:alias/team-foo"),
			Match("arn:aws:kms:eu-west-1:123456789012:alias/team-foo"),
		),
		AllOf(
			Match("arn:aws:kms:eu-central-1:123456789012:alias/team-bar"),
			Match("arn:aws:kms:eu-west-1:123456789012:alias/team-bar"),
		),
	),
	OneOf(
		AllOf(
			Match("arn:aws:kms:eu-central-1:123456789012:alias/production-cicd"),
			Match("arn:aws:kms:eu-west-1:123456789012:alias/production-cicd"),
		),
		AllOf(
			Match("arn:aws:kms:eu-central-1:123456789012:alias/staging-cicd"),
			Match("arn:aws:kms:eu-west-1:123456789012:alias/staging-cicd"),
		),
	),
)
