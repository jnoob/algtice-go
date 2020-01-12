package bytedance

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}
	maxNonRepeat := 0
	appears, head := map[int32]int{}, -1
	for i, c := range s {
		_, appeared := appears[c]
		if appeared && appears[c] > head {
			nonRepeatCount := i - head - 1
			if nonRepeatCount > maxNonRepeat {
				maxNonRepeat = nonRepeatCount
			}
			head = appears[c]
		}
		appears[c] = i
	}
	nonRepeatCount := len(s) - head - 1
	if nonRepeatCount > maxNonRepeat {
		maxNonRepeat = nonRepeatCount
	}
	return maxNonRepeat
}
