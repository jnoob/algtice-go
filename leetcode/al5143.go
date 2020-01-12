package leetcode

func decompressRLElist(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	} else {
		totalCount := 0
		for i := 0; i < len(nums); i += 2 {
			totalCount += nums[i]
		}
		list := make([]int, totalCount)
		pos := 0
		for i := 0; i < len(nums); i += 2 {
			for j := 0; j < nums[i]; j++ {
				list[pos] = nums[i+1]
				pos++
			}
		}
		return list
	}
}
