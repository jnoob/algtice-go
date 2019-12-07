package leetcode

type twoIndexList struct {
	indexs [][2]int
}

func (list *twoIndexList) append(twoNums [2]int) {
	list.indexs = append(list.indexs, twoNums)
}

func threeSum(nums []int) [][]int {
	var sum3ZeroItems [][]int
	sum2Items := get2Sums(nums)
	for i := 0; i < len(nums); i++ {
		list, contains := sum2Items[-nums[i]]
		if contains {
			for _, index := range list.indexs {
				if index[0] != i && index[1] != i {
					sum3ZeroItems = append(sum3ZeroItems, []int {nums[i], nums[index[0]], nums[index[1]]})
				}
			}
		}
	}
	return sum3ZeroItems
}

func get2Sums(nums []int) map[int]twoIndexList {
	sumMap := map[int]twoIndexList{}
	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum2 := nums[i] + nums[j]
			list, contains := sumMap[sum2]
			if contains {
				list.append([2]int {i, j})
			} else {
				list = twoIndexList{indexs:[][2]int {{i, j}}}
				sumMap[sum2] = list
			}
		}
	}
	return sumMap
}