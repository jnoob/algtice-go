package sort

func QuickSort(candidates []int) []int {
	sortee := newArraySortee(candidates)
	quickSortInternal(sortee, 0, sortee.GetLength() - 1)
	return sortee.Return().([]int)
}

func quickSortInternal(list Sortee, lo int, hi int) {
	if lo < hi {
		p := partition(list, lo, hi)
		quickSortInternal(list, lo, p-1)
		quickSortInternal(list, p+1, hi)
	}
}

func partition(list Sortee, lo int, hi int) int {
	pivot, i := list.Get(hi), lo-1
	for j := lo; j < hi; j++ {
		if list.Get(j) <= pivot {
			i++
			list.Swap(i, j)
		}
	}
	list.Swap(i+1, hi)
	return i+1
}
