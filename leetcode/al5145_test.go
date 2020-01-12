package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA5145A(t *testing.T) {
	items := []int{6, 7, 8, 2, 7, 1, 3, 9, -1, 1, 4, -1, -1, -1, 5}
	node := buildBinaryTree(items)
	assert.Equal(t, 18, sumEvenGrandparent(node))
}
