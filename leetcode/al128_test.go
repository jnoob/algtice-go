package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA128A(t *testing.T) {
	var nums []int
	var expected, actual int

	//nums, expected = []int{1, 3, 5, 7, 8, 10, 11, 101, 103, 105, 107, 108, 110, 10001, 10003, 10005, 10007, 10009, 100011}, 2
	//actual = longestConsecutive(nums)
	//assert.Equal(t, expected, actual)
	//
	//nums, expected = []int{1, 2, 3, 5, 7, 8, 9, 10, 15, 101, 103, 105, 107, 108, 110, 10001, 10003, 10005, 10007, 10009, 100011}, 4
	//actual = longestConsecutive(nums)
	//assert.Equal(t, expected, actual)
	//
	//nums, expected = []int{100, 4, 200, 1, 3, 2}, 4
	//actual = longestConsecutive(nums)
	//assert.Equal(t, expected, actual)
	//
	//nums, expected = []int{}, 0
	//actual = longestConsecutive(nums)
	//assert.Equal(t, expected, actual)

	//nums, expected = []int{0, -1}, 2
	//actual = longestConsecutive(nums)
	//assert.Equal(t, expected, actual)

	nums, expected = []int{-21, -20, -19, -18, -17, -16, -15, -14, -13, -12, -11, -10, -9, -8, -7, -5, -4, -3, 0}, 15
	actual = longestConsecutive(nums)
	assert.Equal(t, expected, actual)
}
