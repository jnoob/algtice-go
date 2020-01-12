package bytedance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQ5QReverseWords(t *testing.T) {
	var input, expected, actual string

	input, expected = "the sky is blue", "blue is sky the"
	actual = reverseWords(input)
	assert.Equal(t, expected, actual)

	input, expected = "  hello world!  ", "world! hello"
	actual = reverseWords(input)
	assert.Equal(t, expected, actual)

	input, expected = "    ", ""
	actual = reverseWords(input)
	assert.Equal(t, expected, actual)

	input, expected = "a good   example", "example good a"
	actual = reverseWords(input)
	assert.Equal(t, expected, actual)

	input, expected = " 1", "1"
	actual = reverseWords(input)
	assert.Equal(t, expected, actual)
}
