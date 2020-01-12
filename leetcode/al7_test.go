package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test7AlgSample(t *testing.T) {
	assert.Equal(t, 321, reverse(123))
	assert.Equal(t, -123, reverse(-321))
	assert.Equal(t, 21, reverse(120))
	assert.Equal(t, -21, reverse(-120))
}
