package basicalg

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	case1 := []int {5, 4, 3, 2, 1}
	toSort := case1[:]
	case1Sorted := QuickSort(toSort)
	if toSort[0] == 5 {
		t.Error("QuickSort should be a in-place sort algorithm!")
	}

	case1Expected := []int {1,2,3,4,5}

	if !isEqual(case1Sorted, case1Expected) {
		t.Errorf("QuickSort should return %v, but %v returned!", case1Expected, case1Sorted)
	}

	case2 := []int {9,4,22,56,2,12,89,45}
	toSort = case2[:]
	case2Sorted := QuickSort(toSort)
	case2Expected := []int {2, 4, 9, 12, 22, 45, 56, 89}
	if !isEqual(case2Sorted, case2Expected) {
		t.Errorf("QuickSort should return %v, but %v returned!",case2Expected, case2Sorted)
	}
}

func isEqual(test []int, std []int) bool {
	if len(test) != len(std){
		return  false
	}
	for i := 0; i < len(test); i++ {
		if test[i] != std[i] {
			return false
		}
	}
	return  true
}