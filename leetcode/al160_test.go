package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test160AlgSample(t *testing.T) {
	ah, at := buildLinkList([]int{4, 1})
	bh, bt := buildLinkList([]int{5,0,1})
	ch, _ := buildLinkList([]int{8, 4, 5})

	at.Next = ch
	bt.Next = ch

	n := getIntersectionNode(ah, bh)
	assert.Equal(t, ch, n)
}

func buildLinkList(eles []int) (head, tail *ListNode) {
	var prev *ListNode
	for _, ele := range eles {
		tail = new(ListNode)
		tail.Val = ele
		if prev == nil {
			head = tail
			prev = tail
		} else {
			prev.Next = tail
			prev = tail
		}
	}
	return head, tail
}