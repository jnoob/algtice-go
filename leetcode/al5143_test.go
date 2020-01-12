package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA5143A(t *testing.T) {
	var nums, expected []int

	nums, expected = []int{1, 2, 3, 4}, []int{2, 4, 4, 4}
	assert.Equal(t, expected, decompressRLElist(nums))

	nums, expected = []int{4, 8}, []int{8, 8, 8, 8}
	assert.Equal(t, expected, decompressRLElist(nums))

	nums, expected = []int{4, 8, 2, 6}, []int{8, 8, 8, 8, 6, 6}
	assert.Equal(t, expected, decompressRLElist(nums))
}
