package leetcode

func maxArea(height []int) (int, [2]int) {
	// brute force
	padding, max, index := 0, 0, [2]int{0, 0}
	for padding < len(height)-1 {
		endIndex := len(height) - 1 - padding
		for i := 0; i < endIndex; i++ {
			area := (endIndex - i) * lower(height[i], height[endIndex])
			if area > max {
				max = area
				index[0] = i
				index[1] = endIndex
			}
			// if @i is heigher than @end, its area must gt than the right indexs
			if height[i] >= height[endIndex] {
				break
			}
		}
		padding++
	}
	return max, index
}

func lower(h1 int, h2 int) int {
	if h1 > h2 {
		return h2
	}
	return h1
}
