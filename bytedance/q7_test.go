package bytedance

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestQ7QRestoreIpAddress(t *testing.T) {
	var input string
	var expected []string

	input, expected = "25525511135", []string{"255.255.11.135", "255.255.111.35"}
	assert.True(t, arrayMatch(expected, restoreIpAddresses(input)))

	input, expected = "11111", []string{"1.1.1.11", "11.1.1.1", "1.11.1.1", "1.1.11.1"}
	assert.True(t, arrayMatch(expected, restoreIpAddresses(input)))

	input, expected = "100013", []string{"10.0.0.13","100.0.1.3"}
	assert.True(t, arrayMatch(expected, restoreIpAddresses(input)))

	input, expected = "0000", []string{"0.0.0.0"}
	assert.True(t, arrayMatch(expected, restoreIpAddresses(input)))
}

func arrayMatch(a1 []string, a2 []string) bool {
	r := true
	if len(a1) != len(a2) {
		r = false
	} else {
		sort.Strings(a1)
		sort.Strings(a2)
		for i, l := range a1 {
			if l != a2[i] {
				r = false
				break
			}
		}
	}
	if !r {
		fmt.Printf("[1]->%v\n", a1)
		fmt.Printf("[2]->%v\n", a2)
	}
	return r
}
