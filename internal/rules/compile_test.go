package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompile(t *testing.T) {
	t.Run("compiles fixture", func(t *testing.T) {
		compiled, err := Compile(configFixture.Rules)
		assert.NoError(t, err)
		assert.Equal(t, rulesFixture, compiled)
	})
}
