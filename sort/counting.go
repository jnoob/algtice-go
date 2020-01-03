package sort

// CountingSort is sort with counting
func CountingSort(array []int, k int) []int {
	counts := make([]int, k+1)
	for i := 0; i < len(array); i++ {
		counts[array[i]]++
	}
	for i := 1; i < k+1; i++ {
		counts[i] += counts[i-1]
	}

	target := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		target[counts[array[i]]] = array[i]
		counts[array[i]]--
	}
	return target
}
