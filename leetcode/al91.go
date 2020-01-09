package leetcode

import "strconv"

func numDecodings(s string) int {
	// a[0...n]
	//	=> 1 (a[0]) + count(a[1...n-1])
	//	=> if a[0...1] can be decode, count + 1
	memt := map[int]int{}
	result := countDecodeWays(s, 0, memt)
	return result
}

func countDecodeWays(s string, index int, memt map[int]int) int {
	count, exists := memt[index]
	if exists {
		return count
	}
	// if start with 0, invalid decoding
	if s[index] == '0' {
		count = 0
	} else if len(s)-1 == index {
		count = 1
	} else {
		count = countDecodeWays(s, index+1, memt)
		if v, _ := strconv.Atoi(s[index : index+2]); v < 27 {
			if len(s)-2 == index {
				count += 1
			} else {
				count += countDecodeWays(s, index+2, memt)
			}
		}
	}
	memt[index] = count
	return count
}
