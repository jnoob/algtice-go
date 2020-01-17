package leetcode

func getPermutation(n int, k int) string {
	bits := make([]int, n)
	for i := 0; i < n; i++ {
		bits[i] = i + 1
	}
	//fmt.Printf("bits=%v\n", bits)
	r := getBitsPermutation(bits, 0, k-1)
	//fmt.Printf("%v\n", r)
	result := ""
	for _, c := range r {
		result += string('0' + c)
	}
	return result
}

func getBitsPermutation(bits []int, m int, k int) []int {
	if m == len(bits)-1 || k <= 0 {
		return bits
	}
	levelCount := 1
	for i := 1; i < len(bits)-m; i++ {
		levelCount *= i //(n-1)!
	}
	kLevel, kLeft := k/levelCount, k%levelCount
	//fmt.Printf("m=%v,k=%v,lc=%v,level=%v,left=%v\n", m, k, levelCount, kLevel, kLeft)
	for j := m + kLevel; j > m; j-- {
		bits[j], bits[j-1] = bits[j-1], bits[j]
		//fmt.Printf("exchange [%v]%v<-->[%v]%v\n", j-1, bits[j], j, bits[j-1])
		//fmt.Printf("bits=%v\n", bits)
	}
	return getBitsPermutation(bits, m+1, kLeft)
}
