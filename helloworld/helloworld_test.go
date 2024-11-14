package HelloWorld

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	t.Parallel()
	assert.Equal(t, HelloWorld(), "Hello, World!")
}
