package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Bonial-International-GmbH/sops-check/pkg/config"
	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Run("allow unmatched", func(t *testing.T) {
		cfg := &config.Config{AllowUnmatched: true}

		output, err := runWithConfig(t, cfg)
		require.NoError(t, err)
		assert.Contains(t, output, "No issues found.")
	})

	t.Run("everything unmatched", func(t *testing.T) {
		cfg := &config.Config{AllowUnmatched: false}

		output, err := runWithConfig(t, cfg)
		require.Error(t, err)
		assert.Contains(t, output, "Found issues in")
		assert.Contains(t, output, "Unmatched trust anchors:")
	})

	t.Run("trust anchors not found", func(t *testing.T) {
		cfg := &config.Config{
			AllowUnmatched: false,
			Rules: []config.Rule{
				{
					AnyOf: []config.Rule{
						{
							Match: "this-is-trust-anchor-a",
						},
						{
							Match: "this-is-trust-anchor-b",
						},
					},
				},
			},
		}

		output, err := runWithConfig(t, cfg)
		require.Error(t, err)
		assert.Contains(t, output, "Expected trust anchor \"this-is-trust-anchor-a\" was not found.")
		assert.Contains(t, output, "Expected trust anchor \"this-is-trust-anchor-b\" was not found.")
	})

	t.Run("bad config file", func(t *testing.T) {
		cfg := &config.Config{
			AllowUnmatched: false,
			Rules: []config.Rule{
				{
					// A valid rule cannot have multiple conditions.
					Match:      "this-is-trust-anchor-a",
					MatchRegex: "^this-is-trust-anchor-[ab]$",
				},
			},
		}

		_, err := runWithConfig(t, cfg)
		require.Error(t, err)
		assert.ErrorContains(t, err, "failed to load config")
	})

	t.Run("nonexistent config file", func(t *testing.T) {
		_, err := runWithConfig(t, nil)
		require.Error(t, err)
		assert.Regexp(t, "config file \".*\" not found", err.Error())
	})
}

func runWithConfig(t *testing.T, cfg *config.Config) (string, error) {
	configPath := fmt.Sprintf("%s/.sops-check.yaml", t.TempDir())

	if cfg != nil {
		data, err := yaml.Marshal(cfg)
		require.NoError(t, err)
		require.NoError(t, os.WriteFile(configPath, data, 0o600))
	}

	var sb strings.Builder

	err := run(&sb, []string{"--config", configPath})

	return sb.String(), err
}
