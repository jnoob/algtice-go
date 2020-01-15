package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return result
	} else if (nums[0] == 0 && nums[2] == 0) || (nums[len(nums)-1] == 0 && nums[len(nums)-3] == 0) {
		result = append(result, []int{0, 0, 0})
	} else {
		firstPositive, lastNegative := -1, -1
		i := len(nums) / 2
		for firstPositive == -1 {
			if nums[i] > 0 {
				if nums[i-1] <= 0 {
					firstPositive = i
				} else {
					i = i / 2
				}
			} else if nums[i] == 0 {
				i = i + 1
			} else {
				i = (len(nums) - 1 + i) / 2
			}
		}
		i = firstPositive - 1
		for lastNegative == -1 {
			if nums[i] == 0 {
				i--
			} else {
				lastNegative = i
			}
		}
		if firstPositive-lastNegative > 3 {
			result = append(result, []int{0, 0, 0})
		}

	}
	return nil
}

//type twoIndexList struct {
//	indexs [][2]int
//}
//
//func (list *twoIndexList) append(twoNums [2]int) {
//	list.indexs = append(list.indexs, twoNums)
//}
//
//func threeSum(nums []int) [][]int {
//	var sum3ZeroItems [][]int
//	sum2Items := get2Sums(nums)
//	for i := 0; i < len(nums); i++ {
//		list, contains := sum2Items[-nums[i]]
//		if contains {
//			for _, index := range list.indexs {
//				if index[0] != i && index[1] != i {
//					sum3ZeroItems = append(sum3ZeroItems, []int{nums[i], nums[index[0]], nums[index[1]]})
//				}
//			}
//		}
//	}
//	return sum3ZeroItems
//}
//
//func get2Sums(nums []int) map[int]twoIndexList {
//	sumMap := map[int]twoIndexList{}
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			sum2 := nums[i] + nums[j]
//			list, contains := sumMap[sum2]
//			if contains {
//				list.append([2]int{i, j})
//			} else {
//				list = twoIndexList{indexs: [][2]int{{i, j}}}
//				sumMap[sum2] = list
//			}
//		}
//	}
//	return sumMap
//}
