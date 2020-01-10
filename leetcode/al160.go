package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodesA, nodesB := getNodes(headA), getNodes(headB)

	var joint *ListNode = nil
	for i := 1; i <= len(nodesA); i++ {
		if len(nodesB) < i {
			break
		}
		if nodesB[len(nodesB)-i] == nodesA[len(nodesA)-i] {
			joint = nodesA[len(nodesA)-i]
		}
	}
	return joint
}

func getNodes(head *ListNode) []*ListNode {
	var nodes []*ListNode
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	return nodes
}

func (root *ListNode) unfold() []int {
	vals := []int{}
	node := root
	for node != nil {
		vals = append(vals, node.Val)
		node = node.Next
	}
	return  vals
}
