package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA23ASmaple(t *testing.T) {
	llNums := [][]int{
		{1, 4, 5},
		{1, 3, 4},
		{2, 6},
	}
	lls := []*ListNode{}
	for _, nums := range llNums {
		head, _ := buildLinkList(nums)
		fmt.Printf("%v\n", head.unfold())
		lls = append(lls, head)
	}
	root := mergeKLists(lls)
	arrays := root.unfold()
	expected := []int{1, 1, 2, 3, 4, 4, 5, 6}
	for i, e := range arrays {
		if e != expected[i] {
			fmt.Printf("EXPECTED:%v\n", expected)
			fmt.Printf("ACTUAL  :%v\n", arrays)
			assert.Fail(t, "failed!")
			break
		}
	}
}
