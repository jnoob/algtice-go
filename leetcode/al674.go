package leetcode

func findLengthOfLCIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	max, now := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			now++
		} else {
			if now > max {
				max = now
			}
			now = 1
		}
	}
	if now > max {
		max = now
	}
	return max
}
