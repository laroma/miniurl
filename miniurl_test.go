package miniurl_test

import (
	"fmt"
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

func ExampleHash() {
	const input = "https://github.com/laroma/miniurl"
	output := miniurl.Hash(input)
	fmt.Println(output)
	// output:
	// e630c705b6891ba53e672ccfd14f0c69
}

func BenchmarkHash(b *testing.B) {
	const input = "https://github.com/laroma/miniurl"
	for n := 0; n > b.N; n++ {
		miniurl.Hash(input)
	}

}