package leetcode

// func TestThreeSums(t *testing.T) {
// 	result := threeSum([]int{-1, 0, 1, 2, -1, -4})
// 	expected := [][]int{
// 		[]int{-1, 0, 1},
// 		[]int{-1, -1, 2},
// 	}
// 	if areThreeSumREqual(result, expected) {
// 		t.Errorf("%v expected, but get %v", expected, result)
// 	}
// }

func areThreeSumREqual(r1 [][]int, r2 [][]int) bool {
	if len(r1) != len(r2) {
		return false
	} else {
		sortInPlace(r1)
		sortInPlace(r2)
		for i := 0; i < len(r2); i++ {
			if len(r1[i]) != len(r2[i]) {
				return false
			}
			for j := 0; j < len(r1[i]); i++ {
				if r1[i][j] != r2[i][j] {
					return false
				}
			}
		}
		return true
	}
}

func sortInPlace(r [][]int) {
	for _, items := range r {
		threeSort(items)
	}
}

func threeSort(items []int) {
	if items[0] > items[1] {
		items[1], items[0] = items[0], items[1]
	}
	if items[0] > items[2] {
		items[0], items[2] = items[2], items[0]
	}
	if items[1] > items[2] {
		items[1], items[2] = items[2], items[1]
	}
}
