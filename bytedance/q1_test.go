package bytedance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQ1QLongestNonRepeat(t *testing.T) {
	var s string
	var expected int

	s, expected = "abcabcbb", 3
	assert.Equal(t, expected, lengthOfLongestSubstring(s))
	s, expected = "bbbbb", 1
	assert.Equal(t, expected, lengthOfLongestSubstring(s))
	s, expected = "pwwkew", 3
	assert.Equal(t, expected, lengthOfLongestSubstring(s))

	s, expected = "", 0
	assert.Equal(t, expected, lengthOfLongestSubstring(s))
	s, expected = "", 0
	assert.Equal(t, expected, lengthOfLongestSubstring(s))

	s, expected = "   ", 1
	assert.Equal(t, expected, lengthOfLongestSubstring(s))

	s, expected = "  3 ", 2
	assert.Equal(t, expected, lengthOfLongestSubstring(s))

	//s, expected = "  世界 ", 3
	//assert.Equal(t, expected, lengthOfLongestSubstring(s))

	s, expected = "au", 2
	assert.Equal(t, expected, lengthOfLongestSubstring(s))

	s, expected = "u", 1
	assert.Equal(t, expected, lengthOfLongestSubstring(s))
}
