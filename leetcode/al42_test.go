package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA42A(t *testing.T) {
	var height []int
	var expected, actual int

	height, expected = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6
	actual = trap(height)
	assert.Equal(t, expected, actual)

	height, expected = []int{0, 1, 2, 3, 5, 9, 8, 7, 6, 5}, 0
	actual = trap(height)
	assert.Equal(t, expected, actual)

	height, expected = []int{0, 1, 2, 3, 5, 9}, 0
	actual = trap(height)
	assert.Equal(t, expected, actual)

	height, expected = []int{9, 8, 7, 6, 5}, 0
	actual = trap(height)
	assert.Equal(t, expected, actual)

	height, expected = []int{9, 0, 9}, 9
	actual = trap(height)
	assert.Equal(t, expected, actual)

	height, expected = []int{1, 0, 9}, 1
	actual = trap(height)
	assert.Equal(t, expected, actual)

	height, expected = []int{1, 9, 0, 1}, 1
	actual = trap(height)
	assert.Equal(t, expected, actual)

	height, expected = []int{0, 9, 8, 9, 9, 9, 9, 9, 9, 9}, 1
	actual = trap(height)
	assert.Equal(t, expected, actual)

	height, expected = []int{5, 4, 1, 2}, 1
	actual = trap(height)
	assert.Equal(t, expected, actual)

	height, expected = []int{5, 2, 1, 2, 1, 5}, 14
	actual = trap(height)
	assert.Equal(t, expected, actual)

	height, expected = []int{6, 4, 2, 0, 3, 2, 0, 3, 1, 4, 5, 3, 2, 7, 5, 3, 0, 1, 2, 1, 3, 4, 6, 8, 1, 3}, 83
	actual = trap(height)
	assert.Equal(t, expected, actual)
}
