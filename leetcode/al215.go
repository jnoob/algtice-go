package leetcode

func findKthLargest(nums []int, k int) int {
	minHeap := make([]int, k)
	copy(minHeap, nums[0:k])
	buildMinHeapA(minHeap)
	for i := k; i < len(nums); i++ {
		if nums[i] > minHeap[0] {
			minHeap[0] = nums[i]
			minHeapifyA(minHeap, 0, len(minHeap))
		}
	}
	return minHeap[0]
}

func buildMinHeapA(nums []int) {
	for i := len(nums) / 2; i >= 0; i-- {
		minHeapifyA(nums, i, len(nums))
	}
}

func minHeapifyA(nums []int, i int, length int) {
	minIndex := i
	left, right := leftChild(i), rightChild(i)
	if left >= length {
		return
	} else if nums[left] < nums[minIndex] {
		minIndex = left
	}

	if right < length && nums[right] < nums[minIndex] {
		minIndex = right
	}
	if minIndex != i {
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
		minHeapifyA(nums, minIndex, length)
	}
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}
