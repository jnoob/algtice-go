package basicalg

func QuickSort(candidates []int) []int {
	return quickSortInternal(candidates, 0, len(candidates) - 1)
}

func quickSortInternal(array []int, lo int, hi int) []int {
	if lo < hi {
		p := partition(array, lo, hi)
		quickSortInternal(array, lo, p - 1)
		quickSortInternal(array, p + 1, hi)
	}
	return array
}

func partition(array []int, lo int, hi int) int {
	pivot, i := array[hi], lo
	for j := lo; j <= hi; j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i += 1
		}
	}
	array[i], array[hi] = array[hi], array[i]
	return  i
}
