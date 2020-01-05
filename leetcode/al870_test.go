package leetcode

import (
	"fmt"
	"testing"
)

func Test870AlgAdvCountSample(t *testing.T) {
	a, b := []int{2, 7, 11, 15}, []int{1, 10, 4, 11}
	r := advantageCount(a, b)
	fmt.Printf("%v\n", r)

	a, b = []int{12, 24, 8, 32}, []int{13, 25, 32, 11}
	r = advantageCount(a, b)
	fmt.Printf("%v\n", r)
}
