package leetcode

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	peaks := make([]int, 0)
	up := true
	for i := 1; i < len(height); i++ {
		if up {
			if height[i] < height[i-1] {
				peaks = append(peaks, i-1)
				up = false
			}
		} else if height[i-1] < height[i] {
			up = true
		}
	}
	if up {
		peaks = append(peaks, len(height)-1)
	}
	//fmt.Printf("peaks:%v\n", peaks)
	if len(peaks) < 2 {
		return 0
	} else {
		ranges, ceils := []int{peaks[0]}, make([]int, 0)
		for i := 1; i < len(peaks); {
			if height[peaks[i]] >= height[peaks[i-1]] {
				ranges = append(ranges, peaks[i])
				ceils = append(ceils, height[peaks[i-1]])
				i++
			} else {
				ceil, ceilIndex := -1, -1
				for j := len(peaks) - 1; j > i; j-- {
					if height[peaks[j]] > height[peaks[i-1]] {
						ceil = height[peaks[j]]
						ceilIndex = j
					} else if height[peaks[j]] >= height[peaks[i]] && height[peaks[j]] > ceil {
						ceil = height[peaks[j]]
						ceilIndex = j
					}
				}
				if ceil > 0 {
					// cell should be max height of the right side of i
					ranges = append(ranges, peaks[ceilIndex])
					if ceil > height[peaks[i-1]] {
						ceils = append(ceils, height[peaks[i-1]])
					} else {
						ceils = append(ceils, ceil)
					}
					i = ceilIndex + 1
				} else {
					ranges = append(ranges, peaks[i])
					ceils = append(ceils, height[peaks[i]])
					i++
				}
			}
		}
		//fmt.Printf("ranges:%v\n", ranges)
		//fmt.Printf("ceils:%v\n", ceils)
		total := 0
		for i := 0; i < len(ceils); i++ {
			ceil := ceils[i]
			for j := ranges[i] + 1; j < ranges[i+1]; j++ {
				capa := ceil - height[j]
				if capa > 0 {
					total += capa
				}
			}
		}
		return total
	}
}
