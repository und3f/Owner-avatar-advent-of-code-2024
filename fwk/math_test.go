package fwk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDigits(t *testing.T) {
	assert.Equal(t, 1, CountDigits(0))
	assert.Equal(t, 1, CountDigits(1))
	assert.Equal(t, 1, CountDigits(-1))

	assert.Equal(t, 2, CountDigits(10))
	assert.Equal(t, 2, CountDigits(-10))

	assert.Equal(t, 3, CountDigits(999))
	assert.Equal(t, 3, CountDigits(-999))
}
