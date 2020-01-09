package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test91AlgNumDecodingSample(t *testing.T) {
	s, expected := "227", 2
	result := numDecodings(s)
	assert.Equal(t, expected, result, "%v -> %v != %v", s, result, expected)

	//s, expected = "226", 3
	//result = numDecodings(s)
	//assert.Equal(t, expected, result, "%v -> %v != %v", s, result, expected)
	//
	//s, expected = "12", 2
	//result = numDecodings(s)
	//assert.Equal(t, expected, result, "%v -> %v != %v", s, result, expected)
	//
	//s, expected = "0", 0
	//result = numDecodings(s)
	//assert.Equal(t, expected, result, "%v -> %v != %v", s, result, expected)
	//
	//s, expected = "10", 1
	//result = numDecodings(s)
	//assert.Equal(t, expected, result, "%v -> %v != %v", s, result, expected)
	//
	//s, expected = "101", 1
	//result = numDecodings(s)
	//assert.Equal(t, expected, result, "%v -> %v != %v", s, result, expected)
}
