package bytedance

const empty int32 = ' '

func reverseWords(s string) string {
	start := -1
	locs := [][2]int{}
	for i, c := range s {
		if c != empty {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				locs = append(locs, [2]int{start, i - 1})
				start = -1
			}
		}
	}
	if start != -1 {
		if len(locs) == 0 {
			return s[start:]
		} else {
			locs = append(locs, [2]int{start, len(s) - 1})
		}
	}
	var r string
	for i := len(locs) - 1; i >= 0; i-- {
		r += s[locs[i][0] : locs[i][1]+1]
		if i != 0 {
			r += " "
		}
	}
	return r
}
