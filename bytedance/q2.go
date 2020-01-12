package bytedance

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	} else {
		minL := getMinLength(strs)
		commonEnd, terminate := minL-1, false
		for i := 0; i < minL; i++ {
			for j := 1; j < len(strs); j++ {
				if strs[j][i] != strs[j-1][i] {
					commonEnd = i - 1
					terminate = true
					break
				}
			}
			if terminate {
				break
			}
		}
		if commonEnd >= 0 {
			return strs[0][0 : commonEnd+1]
		} else {
			return ""
		}
	}
}

func getMinLength(strs []string) int {
	minL := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minL {
			minL = len(strs[i])
		}
	}
	return minL
}
