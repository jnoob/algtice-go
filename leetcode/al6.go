package leetcode

func convert(s string, numRows int) string {
	// numRows == 1 make the underlying assumption failed, so it need to handle specially
	if numRows <= 1 {
		return s
	}
	// if i = 0 or numRows - 1, each group contains one ele
	// else each group contains two eles, @i and 2(numRows -1) - i
	output := make([]byte, 0, len(s))
	groupLen := 2 * numRows - 2
	output = printZRow(s, groupLen, []int {0}, output)
	for i := 1; i < numRows - 1; i++ {
		output = printZRow(s, groupLen, []int {i, groupLen - i}, output)
	}
	if numRows > 1 {
		output = printZRow(s, groupLen, []int {numRows - 1}, output)
	}
	return string(output)
}

func printZRow(s string, groupLen int, locs []int, output []byte) []byte {
	n := 0
	for {
		for _, loc := range locs {
			i := loc + n * groupLen
			if i < len(s) {
				output = append(output, s[i])
			} else {
				return output
			}
		}
		n++
	}
}