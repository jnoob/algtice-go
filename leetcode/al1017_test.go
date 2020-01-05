package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1017BaseNeg2SampleCases(t *testing.T) {
	expected, N := "110", 2
	result := baseNeg2(N)
	assert.Equalf(t, expected, result, "%v-->%v != %v", N, result, expected)

	//expected, N = "111", 3
	//result = baseNeg2(N)
	//assert.Equalf(t, expected, result, "%v-->%v != %v", N, result, expected)
	//
	//expected, N = "100", 4
	//result = baseNeg2(N)
	//assert.Equalf(t, expected, result, "%v-->%v != %v", N, result, expected)
}
