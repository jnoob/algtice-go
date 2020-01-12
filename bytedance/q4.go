package bytedance

const _0 = "0"
const _1 = "1"

func multiply(num1 string, num2 string) string {
	if num1 == _0 || num2 == _0 {
		return _0
	} else if num1 == _1 {
		return num2
	} else if num2 == _1 {
		return num1
	}

	digits1, digits2 := []rune(num1), []rune(num2)
	// 999 * 999 = 999 * 1000 - 999, so the digit count of the result must <= len(digits1) + len(digits2)
	// 100 * 100 = 10000, so the digit count must >= len(digits1) + len(digits2) - 1
	rDigits := make([]int, len(digits1)+len(digits2))
	for i := 1; i <= len(digits1); i++ {
		x := charToInt(digits1[len(digits1)-i])
		for j := 1; j <= len(digits2); j++ {
			y := charToInt(digits2[len(digits2)-j])
			r := x*y + rDigits[i+j-2]
			carry, d := r/10, r%10
			rDigits[i+j-2] = d
			rDigits[i+j-1] += carry
		}
	}
	var start, l int
	if rDigits[len(digits1)+len(digits2)-1] > 0 {
		l = len(digits1) + len(digits2)
	} else {
		l = len(digits1) + len(digits2) - 1
	}
	start = l - 1
	output := make([]rune, start+1)
	for i := start; i >= 0; i-- {
		output[l-i-1] = intToChar(rDigits[i])
	}
	return string(output)
}

func charToInt(c rune) int {
	return int(c - '0')
}

func intToChar(i int) rune {
	return rune('0' + i)
}
