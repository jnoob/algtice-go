package bytedance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQ2QLongestCommonPrefix(t *testing.T) {
	var list []string
	var expected, result string

	list, expected = []string{"flower", "flow", "flight"}, "fl"
	result = longestCommonPrefix(list)
	assert.Equal(t, expected, result)

	list, expected = []string{"flower2", "flow", "flower1"}, "flow"
	result = longestCommonPrefix(list)
	assert.Equal(t, expected, result)

	list, expected = []string{"", "flow", "flower1"}, ""
	result = longestCommonPrefix(list)
	assert.Equal(t, expected, result)

	list, expected = []string{}, ""
	result = longestCommonPrefix(list)
	assert.Equal(t, expected, result)

	list, expected = []string{"aaaaaaaa", "aaaaaa", "aaaaaaaa"}, "aaaaaa"
	result = longestCommonPrefix(list)
	assert.Equal(t, expected, result)

	list, expected = []string{"aaaaaa", "aaaaaa", "aaaaaa"}, "aaaaaa"
	result = longestCommonPrefix(list)
	assert.Equal(t, expected, result)

	list, expected = []string{"a", "a", "a"}, "a"
	result = longestCommonPrefix(list)
	assert.Equal(t, expected, result)

	list, expected = []string{"dog","racecar","car"}, ""
	result = longestCommonPrefix(list)
	assert.Equal(t, expected, result)
}

