package bytedance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQ3QCheckInclusion(t *testing.T) {
	var s1, s2 string

	s1, s2 = "ab", "eidbaooo"

	assert.True(t, checkInclusion(s1, s2))

	s1, s2 = "ab", "eidboaoo"
	assert.False(t, checkInclusion(s1, s2))

	s1, s2 = "a", "eidboaoo"
	assert.True(t, checkInclusion(s1, s2))

	s1, s2 = "sddsda", "sad"
	assert.False(t, checkInclusion(s1, s2))

	s1, s2 = "aad", "sadba"
	assert.False(t, checkInclusion(s1, s2))

	s1, s2 = "baccd", "mbaccnbaccabdggh"
	assert.True(t, checkInclusion(s1, s2))

	s1, s2 = "abc", "bbbca"
	assert.True(t, checkInclusion(s1, s2))
}
