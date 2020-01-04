package leetcode

import (
	"fmt"
	"math"
	"strconv"
)

func baseNeg2(N int) string {
	if N == 0{
		return "0"
	}

	var result string
	var sub float64
	n := N
	for n != 0 {
		fmt.Printf("n = %v, result = %v, sub = %v\n", n, result, sub)

		result = strconv.Itoa(absr(n, 2)) + result
		sub = float64(n) / -2
		n = int(math.Ceil(float64(n)/-2))
	}
	fmt.Printf("n = %v, result = %v, sub = %v\n", n, result, sub)
	return result
}

func absr(i int, j int) int {
	if i > 0 {
		return i % 2
	} else {
		return -(i % 2)
	}
}