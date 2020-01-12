package bytedance

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	initMap := map[int32]int{}
	for _, m := range s1 {
		_, exists := initMap[m]
		if exists {
			initMap[m]++
		} else {
			initMap[m] = 1
		}
	}
	charMap := map[int32][]int{}
	for k, v := range initMap {
		charMap[k] = make([]int, v+1)
		charMap[k][0] = v
	}
	firstMatch, matched, target := -1, 0, len(s1)
	for ic, c := range s2 {
		info, exists := charMap[c]
		if exists {
			if info[0] == 0 { // the char appeared too many time
				charFirstIndex := info[len(info)-1]
				discardPartial(charMap, s2[firstMatch:charFirstIndex+1])
				matched -= charFirstIndex - firstMatch + 1
				firstMatch = charFirstIndex + 1
			}
			info[info[0]] = ic
			info[0]--
			matched++
			if firstMatch == -1 {
				firstMatch = ic
			}
			if matched == target {
				break
			}
		} else if matched > 0 {
			renewRecord(charMap)
			matched = 0
			firstMatch = -1
		}
	}
	return matched == target
	//	for c -> s2:
	// 		if c in s1 and count(c int currents2sub) <= count(c in s1):
	//			put c into matchLoc and increase count(c in currents2sub), matched++
	//		if c not in s1:
	//			new round
	//		if c in s1 and count(c int currents2sub) > count(c in s1):
	//			find first c loc, discard the loc and the prev, maintain matched and the locs
	//		if matched == target:
	//			succeed

}

func renewRecord(charMap map[int32][]int) {
	for _, v := range charMap {
		v[0] = len(v) - 1
	}
}

func discardPartial(charMap map[int32][]int, discard string) {
	for _, s := range discard {
		for j := len(charMap[s]) - 2; j > charMap[s][0]; j-- {
			charMap[s][j+1] = charMap[s][j]
		}
		charMap[s][0]++
	}
}
