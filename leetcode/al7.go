package leetcode

import (
	"math"
)

func reverse(x int) int {
	neg, num := false, x
	if x < 0 {
		neg, num = true, -x
	}
	bits := []int{}
	for num >= 1 {
		div, mod := num/10, num%10
		num = div
		bits = append(bits, mod)
	}
	//fmt.Printf("%v -> bits: %v\n",x, bits)
	result := 0
	for i := 1; i <= len(bits); i++ {
		bit := bits[len(bits)-i]
		if bit > 0 {
			result += bit * int(math.Pow10(i-1))
		}
		//fmt.Printf("bit=%v -> result=%v\n",bit, result)
	}
	if neg {
		return -result
	} else {
		return result
	}
}
