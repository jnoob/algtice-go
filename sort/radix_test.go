package sort

import "testing"

func TestRadixSort(t *testing.T) {
	_2dList := [][]int {
		{ 9, 23, 11, 90, 34, 66, 88},
		{ 8, 81, 12, 76, 21,  7, 71},
		{12, 33,  2,  1, 23, 42, 53},
		{11, 23,  4, 79, 80, 90, 33},
		{ 8, 44, 55,  7,  1, 98,  2},
	}

	RadixSort(_2dList)

	expected := [][]int {
		{ 8, 44, 55,  7,  1, 98,  2},
		{8, 81, 12, 76, 21, 7, 71},
		{ 9, 23, 11, 90, 34, 66, 88},
		{11, 23,  4, 79, 80, 90, 33},
		{12, 33,  2,  1, 23, 42, 53},
	}

	failed := false
	for i := 0; i < len(expected); i++ {
		for j:= 0; j < len(expected[i]); j++ {
			if expected[i][j] != _2dList[i][j] {
				failed = true
				t.Errorf("Expected: %v; Actual: %v!", expected, _2dList)
				break
			}
		}
		if failed {
			break
		}
	}
}