package miniurl_test

import (
	"testing"

	miniurl "github.com/laroma/miniurl"
	"github.com/stretchr/testify/assert"
)

func TestHashLength(t *testing.T) {
	const (
		input = "https://github.com/laroma/miniurl"
		expectedLength = 32
	)

	output := miniurl.Hash(input)
	assert.Len(t, output, expectedLength)
}

func TestHasIsDeterministic(t *testing.T) {
	const input = "https://github.com/laroma/miniurl"

	output1 := miniurl.Hash(input)
	output2 := miniurl.Hash(input)
	assert.Equal(t, output1, output2)
}