package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNil(t *testing.T) {
	// Basic Types is not nil
	assert.False(t, IsNil(0))

	// null pointer should be true
	var a *int
	assert.True(t, IsNil(a))
}
