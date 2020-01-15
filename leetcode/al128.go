package leetcode

func longestConsecutive(nums []int) int {
	seqLists := make(map[int][]bool)
	seqRange := make(map[int][]int)

	for _, num := range nums {
		var seq, pos int
		if num >= 0 {
			seq, pos = num/10, num%10
		} else {
			seq, pos = (num+1)/10-1, num%10
			if pos != 0 {
				pos += 10
			}
		}

		if _, exist := seqLists[seq]; exist {
			seqLists[seq][pos] = true
			if seqRange[seq][0] > pos {
				seqRange[seq][0] = pos
			}
			if seqRange[seq][1] < pos {
				seqRange[seq][1] = pos
			}
		} else {
			seqLists[seq] = make([]bool, 10)
			seqLists[seq][pos] = true
			seqRange[seq] = make([]int, 2)
			seqRange[seq][0] = pos
			seqRange[seq][1] = pos
		}
	}

	testedSeqs := make(map[int]bool)
	maxCL := 0
	for entry, _ := range seqLists {
		if _, tested := testedSeqs[entry]; tested {
			continue
		} else {
			seq := entry
			for seqRange[seq][0] == 0 {
				if _, exist := seqLists[seq-1]; exist && seqRange[seq-1][1] == 9 {
					seq = seq - 1
				} else {
					break
				}
			}
			count := 0
			for {
				testedSeqs[seq] = true
				for i := seqRange[seq][0]; i <= seqRange[seq][1]; i++ {
					if seqLists[seq][i] == true {
						count++
					} else {
						if count > maxCL {
							maxCL = count
						}
						count = 0
					}
				}
				forward := false
				if seqRange[seq][1] == 9 {
					if _, exists := seqLists[seq+1]; exists && seqRange[seq+1][0] == 0 {
						forward = true
					}
				}
				if forward {
					seq++
				} else {
					if count > maxCL {
						maxCL = count
					}
					break
				}
			}
		}
	}
	return maxCL
}
