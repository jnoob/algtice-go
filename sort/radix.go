package sort

func RadixSort(list [][]int) {
	sortee := new2DArraySortee(list)
	colCount := len(list[0])

	for i := colCount - 1; i >= 0; i-- {
		sortee.MoveToCol(i)
		quickSortInternal(sortee, 0, sortee.GetLength()-1)
	}
}
