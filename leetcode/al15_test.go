package leetcode

import (
	"algtice/basicalg"
	"testing"
)

func TestThreeSums(t *testing.T) {
	result := threeSum([]int{ -1, 0, 1, 2, -1, -4})
}

func resultCompare(r1 [][]int, r2 [][]int) bool {
	if len(r1) != len(r2) {
		return false
	} else {

	}
}

func sortInPlace(r [][]int) {
	for _, items := range r {
		threeSort(items)
	}
	
}

func threeSort(items []int) {
	if items[0] > items[1] {
		items[1], items[0] = items[0], items[1]
	}
	if items[0] > items[2] {
		items[0], items[2] = items[2], items[0]
	}
	if items[1] > items[2] {
		items[1], items[2] = items[2], items[1]
	}
}