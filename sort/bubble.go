package sort

// BubbleSort is a O(n^2) time compelecity sorting algorithm
func BubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
