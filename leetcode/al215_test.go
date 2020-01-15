package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA215A(t *testing.T) {
	var nums []int
	var k, expcted, actual int

	nums, k, expcted = []int{3, 2, 1, 5, 6, 4}, 2, 5
	actual = findKthLargest(nums, k)
	assert.Equal(t, expcted, actual)

	nums, k, expcted = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4
	actual = findKthLargest(nums, k)
	assert.Equal(t, expcted, actual)
}
