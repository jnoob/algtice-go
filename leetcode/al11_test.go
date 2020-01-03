package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	output, indexs := maxArea(input)
	t.Errorf("%v get %v at %v", input, output, indexs)
}
