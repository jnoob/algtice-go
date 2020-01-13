package leetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, stop := 0, len(nums)-1
	mid := (start + stop) / 2
	for {
		if nums[mid] == target {
			return mid
		} else if nums[start] == target {
			return start
		} else if nums[stop] == target{
			return stop
		} else if target < nums[mid] {
			if nums[mid] > nums[start] {
				if target > nums[start] {
					// --X--+---
					//          ------
					stop = mid-1
					tmp := (start + stop)/2
					if tmp == mid {
						tmp--
					}
					mid = tmp
					if mid <= start {
						return -1
					}
				} else if nums[stop] < nums[mid] {
					// ----+---
					//         ---X---
					start = mid+1
					tmp := (start + stop)/2
					if tmp == mid {
						tmp++
					}
					mid = tmp
					if mid >= stop {
						return -1
					}
				} else if nums[start] > nums[mid] {
					// ------
					//       --X--+----
					stop = mid-1
					tmp := (start + stop)/2
					if tmp == mid {
						tmp--
					}
					mid = tmp
					if mid <= start {
						return -1
					}
				} else {
					return -1
				}
			} else if target < nums[mid] {
				if nums[stop] > target {

					start = mid+1
					tmp := (start+stop)/2
					if tmp == mid {
						tmp++
					}
					mid = tmp
					if mid >= stop {
						return -1
					}
				} else if  {
					// ----+--X-
					//         ------
				}
			}
		}
	}
}