package leetcode

func advantageCount(A []int, B []int) []int {
	tmpB := make([]int, 0, len(B))
	for _, e:= range B{
		tmpB = append(tmpB, e)
	}
	sortedA, sortedB := quickSort(A, 0, len(A) - 1), quickSort(tmpB, 0, len(B)-1)
	advMap := map[int][]int{}

	bIndex, bTail := 0, len(B) - 1
	for _, eA := range sortedA {
		if eA > sortedB[bIndex] {
			addEle(advMap, eA, sortedB[bIndex])
			bIndex++
		} else {
			addEle(advMap, eA, sortedB[bTail])
			bTail--
		}
	}
	//fmt.Printf("advMap: %v\n", advMap)
	//fmt.Printf("B:%v\n", B)
	result := A[0:]
	for i, eB := range B {
		a := takeEle(advMap, eB)
		//fmt.Printf("%v -> %v\n", eB, a)
		result[i] = a
	}
	return result
}

func addEle(advMap map[int][]int, a int, b int) {
	list, exists := advMap[b]
	if exists {
		advMap[b] = append(list, a)
	} else {
		advMap[b] = []int{a}
	}
}

func takeEle(advMap map[int][]int, b int) int {
	list, _ := advMap[b]
	if len(list) == 1 {
		a := list[0]
		return a
	} else  {
		a := list[0]
		advMap[b] = list[1:]
		return a
	}

}

func quickSort(eles []int, start int, end int) []int {
	if start < end {
		eles, pivotIndex := partition(eles, start, end)
		eles = quickSort(eles, start, pivotIndex-1)
		eles = quickSort(eles, pivotIndex+1, end)
	}
	return eles
}

func partition(eles []int, start int, end int) ([]int, int) {
	pivot := eles[end]
	ltTail := start - 1
	for j := start; j < end; j++ {
		if eles[j] < pivot {
			ltTail++
			eles[ltTail], eles[j] = eles[j], eles[ltTail]
		}
	}
	eles[ltTail+1], eles[end] = eles[end], eles[ltTail+1]
	return eles, ltTail+1
}
