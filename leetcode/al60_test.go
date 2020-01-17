package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA60A(t *testing.T) {
	var n, k int
	var expected, actual string

	expected, n, k = "213", 3, 3
	actual = getPermutation(n, k)
	assert.Equal(t, expected, actual)

	expected, n, k = "2314", 4, 9
	actual = getPermutation(n, k)
	assert.Equal(t, expected, actual)
}
