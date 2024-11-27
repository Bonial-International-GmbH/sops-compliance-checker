package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseArgs(t *testing.T) {
	t.Run("no args", func(t *testing.T) {
		args, err := ParseArgs(nil)
		require.NoError(t, err)

		expected := &Args{
			ConfigPath: Defaults.ConfigPath,
			CheckPath:  Defaults.CheckPath,
		}

		assert.Equal(t, expected, args)
	})

	t.Run("args", func(t *testing.T) {
		args, err := ParseArgs([]string{"--config", "the-config.yaml"})
		require.NoError(t, err)

		expected := &Args{
			ConfigPath: "the-config.yaml",
			CheckPath:  Defaults.CheckPath,
		}

		assert.Equal(t, expected, args)
	})

	t.Run("invalid args", func(t *testing.T) {
		_, err := ParseArgs([]string{"--nonexistent"})
		require.Error(t, err)
	})
}
