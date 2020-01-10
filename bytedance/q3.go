package bytedance

func checkInclusion(s1 string, s2 string) bool {
	checkMap := map[int32]int{}
	for _, m := range s1 {
		_, exists := checkMap[m]
		if exists {
			checkMap[m] += 1
		} else {
			checkMap[m] = 1
		}
	}
	charCount := len(s1)

}