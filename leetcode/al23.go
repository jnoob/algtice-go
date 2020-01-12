package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	lists = removeNils(lists)
	if len(lists) == 0 {
		return nil
	}
	buildMinHeap(lists)
	root, length := popMin(lists, len(lists))
	node := root
	for length > 0 {
		node.Next, length = popMin(lists, length)
		node = node.Next
	}
	return root
}

func removeNils(lists []*ListNode) []*ListNode {
	rendered := []*ListNode{}
	for _, e := range lists {
		if e != nil {
			rendered = append(rendered, e)
		}
	}
	return rendered
}

func popMin(lists []*ListNode, length int) (*ListNode, int) {
	min := lists[0]
	if min.Next != nil {
		lists[0] = min.Next
		siftDown(lists, length)
	} else {
		length--
		if length > 0 {
			lists[0] = lists[length]
			siftDown(lists, length)
		}
	}
	return min, length
}

func siftDown(list []*ListNode, length int) {
	i, minIndex := 0, 0
	for {
		left, right := leftChild(i), rightChild(i)
		if left >= length {
			break
		} else if list[left].Val < list[minIndex].Val {
			minIndex = left
		}
		if right < length && list[right].Val < list[minIndex].Val {
			minIndex = right
		}
		if minIndex != i {
			list[i], list[minIndex] = list[minIndex], list[i]
			i = minIndex
		} else {
			break
		}
	}
}

func buildMinHeap(lists []*ListNode) {
	length := len(lists)
	for i := length / 2; i >= 0; i-- {
		minHeapify(lists, i, length)
	}
}

func minHeapify(lists []*ListNode, i, length int) {
	minIndex := i
	left, right := leftChild(i), rightChild(i)
	if left >= length {
		return
	} else if lists[left].Val < lists[minIndex].Val {
		minIndex = left
	}

	if right < length && lists[right].Val < lists[minIndex].Val {
		minIndex = right
	}
	if minIndex != i {
		lists[i], lists[minIndex] = lists[minIndex], lists[i]
		minHeapify(lists, minIndex, length)
	}
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}
