package bytedance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQ6QSimplifyPath(t *testing.T) {
	var input, expected, actual string

	input, expected = "/home/", "/home"
	actual = simplifyPath(input)
	assert.Equal(t, expected, actual)

	input, expected = "/../", "/"
	actual = simplifyPath(input)
	assert.Equal(t, expected, actual)

	input, expected = "/a/./b/../../c/", "/c"
	actual = simplifyPath(input)
	assert.Equal(t, expected, actual)

	input, expected = "/home//foo/", "/home/foo"
	actual = simplifyPath(input)
	assert.Equal(t, expected, actual)

	input, expected = "/a/../../b/../c//.//", "/c"
	actual = simplifyPath(input)
	assert.Equal(t, expected, actual)

	input, expected = "/a//b////c/d//././/..", "/a/b/c"
	actual = simplifyPath(input)
	assert.Equal(t, expected, actual)

	input, expected = "", "/"
	actual = simplifyPath(input)
	assert.Equal(t, expected, actual)
}
