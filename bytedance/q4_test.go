package bytedance

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestQ4QMultiply(t *testing.T) {
	var num1, num2, expected, result string
	num1, num2, expected = "2", "3", "6"
	result = multiply(num1, num2)
	assert.Equal(t, expected, result)

	num1, num2, expected = "123", "456", "56088"
	result = multiply(num1, num2)
	assert.Equal(t, expected, result)

	num1, num2, expected = "123", "0", "0"
	result = multiply(num1, num2)
	assert.Equal(t, expected, result)

	num1, num2, expected = strings.Repeat("9", 109), "1"+strings.Repeat("0", 108), strings.Repeat("9", 109)+strings.Repeat("0", 108)
	result = multiply(num1, num2)
	assert.Equal(t, expected, result)

	num1, num2, expected = "99999", "1", "99999"
	result = multiply(num1, num2)
	assert.Equal(t, expected, result)

}
