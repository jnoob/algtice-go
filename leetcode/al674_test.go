package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA674A(t *testing.T) {
	var nums []int
	var expected, actual int

	nums, expected = []int{1, 3, 5, 4, 7}, 3
	actual = findLengthOfLCIS(nums)
	assert.Equal(t, expected, actual)

	nums, expected = []int{1}, 1
	actual = findLengthOfLCIS(nums)
	assert.Equal(t, expected, actual)

	nums, expected = []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 1
	actual = findLengthOfLCIS(nums)
	assert.Equal(t, expected, actual)
}
