package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertSample(t *testing.T) {
	s := "LEETCODEISHIRING"
	n, expected := 3, "LCIRETOESIIGEDHN"
	result := convert(s, n)
	assert.EqualValues(t, expected, result, "[%v] --%v--> [%v] != [%v]", s, n, result, expected)

	n, expected = 4, "LDREOEIIECIHNTSG"
	result = convert(s, n)
	assert.EqualValues(t, expected, result, "[%v] --%v--> [%v] != [%v]", s, n, result, expected)
}

func TestConvertEdgeCases(t *testing.T) {
	s := ""
	n, expected := 3, ""
	result := convert(s, n)
	assert.EqualValues(t, expected, result, "[%v] --%v--> [%v] != [%v]", s, n, result, expected)

	s = "LEETCODEISHIRING"
	n, expected = 1, "LEETCODEISHIRING"
	result = convert(s, n)
	assert.EqualValues(t, expected, result, "[%v] --%v--> [%v] != [%v]", s, n, result, expected)

	s = "123456789A"
	n, expected = 2, "135792468A"
	result = convert(s, n)
	assert.EqualValues(t, expected, result, "[%v] --%v--> [%v] != [%v]", s, n, result, expected)
}
