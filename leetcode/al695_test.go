package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA695AMaxAreaOfIsland(t *testing.T) {
	var input [][]int
	var expected int
	//input = [][]int{
	//	{0,0,1,0,0,0,0,1,0,0,0,0,0},
	//{0,0,0,0,0,0,0,1,1,1,0,0,0},
	//{0,1,1,0,1,0,0,0,0,0,0,0,0},
	//{0,1,0,0,1,1,0,0,1,0,1,0,0},
	//{0,1,0,0,1,1,0,0,1,1,1,0,0},
	//{0,0,0,0,0,0,0,0,0,0,1,0,0},
	//{0,0,0,0,0,0,0,1,1,1,0,0,0},
	//{0,0,0,0,0,0,0,1,1,0,0,0,0},
	//}
	//expected = 6
	//assert.Equal(t, expected, maxAreaOfIsland(input))
	//
	//input = [][]int{
	//	{0,0,0,0,0,0,0,0,0,0,0,0,0},
	//	{0,0,0,0,0,0,0,0,0,0,0,0,0},
	//	{0,0,0,0,0,0,0,0,0,0,0,0,0},
	//	{0,0,0,0,0,0,0,0,0,0,0,0,0},
	//	{0,0,0,0,0,0,0,0,0,0,0,0,0},
	//}
	//expected = 0
	//assert.Equal(t, expected, maxAreaOfIsland(input))
	//
	//input = [][]int{
	//	{0,0,1,0,0,1,1,1,0,0,0,0,0},
	//	{0,0,0,1,1,1,0,1,1,1,0,0,0},
	//	{0,1,1,0,1,0,0,0,0,0,0,0,0},
	//	{0,1,0,0,1,1,0,0,1,0,1,0,0},
	//	{0,1,0,0,1,1,0,0,1,1,1,0,0},
	//	{0,0,0,0,0,0,0,0,0,0,1,0,0},
	//	{0,0,0,0,0,0,0,1,1,1,0,0,0},
	//	{0,0,0,0,0,0,0,1,1,0,0,0,0},
	//}
	//expected = 14
	//assert.Equal(t, expected, maxAreaOfIsland(input))
	//
	//
	//input = [][]int{
	//	{0,0,0,0,0,0,0,0,0,0,0,0,0},
	//	{0,0,1,1,1,0,0,1,1,1,0,0,0},
	//	{0,0,0,0,1,1,1,1,0,1,1,0,0},
	//	{0,1,1,0,1,0,0,0,0,0,0,0,0},
	//	{0,1,0,0,1,1,0,0,1,0,1,0,0},
	//	{0,1,0,0,1,1,0,0,1,1,1,0,0},
	//	{0,0,0,0,0,0,0,0,0,0,1,0,0},
	//	{0,0,0,0,0,0,0,1,1,1,0,0,0},
	//	{0,0,0,0,0,0,0,1,1,0,0,0,0},
	//}
	//expected = 17
	//assert.Equal(t, expected, maxAreaOfIsland(input))
	//
	//input = [][]int {
	//	{1,1,0,0,0},
	//	{1,1,0,0,0},
	//	{0,0,0,1,1},
	//	{0,0,0,1,1},
	//}
	//expected = 4
	//assert.Equal(t, expected, maxAreaOfIsland(input))

	input = [][]int {{1,1,0,0,1,0,1,1,1,1,1,0,0,0,1,1,0,1,0,1,1,1,1,0,0},{0,0,0,1,0,1,0,1,1,0,0,0,0,1,0,1,1,1,1,0,1,0,0,0,1},{1,0,1,0,1,1,0,0,1,1,1,0,0,0,1,0,0,1,1,1,1,1,1,0,1},{0,0,1,1,0,1,1,0,1,0,0,0,0,0,0,0,1,0,1,0,1,0,0,1,1},{1,1,0,1,1,1,1,1,0,0,1,0,0,1,1,1,1,1,1,0,0,0,0,0,1},{0,1,0,1,0,0,1,1,0,1,0,0,0,0,0,0,0,0,0,1,0,1,1,1,1},{1,1,1,1,0,0,0,1,0,0,1,0,0,1,0,1,1,1,1,1,0,0,0,0,0},{0,1,1,1,1,1,0,1,0,0,0,1,0,1,1,1,0,0,1,0,0,0,0,1,1},{1,1,0,0,1,1,0,1,0,1,1,0,0,0,0,1,1,1,1,0,1,1,0,0,0},{0,1,1,1,1,0,1,1,0,1,1,0,1,1,0,1,0,0,0,0,1,0,1,1,1},{0,1,1,1,0,1,0,0,0,0,0,1,0,1,0,1,0,0,0,0,0,0,0,1,0},{1,0,0,1,0,1,1,1,1,1,1,1,1,1,0,0,1,1,1,1,1,0,1,1,1},{1,1,1,1,1,1,1,1,0,0,1,1,0,1,1,1,1,1,0,1,1,0,1,1,1},{1,1,1,0,1,0,1,1,1,1,1,1,0,1,1,0,1,0,1,1,1,0,0,1,1},{0,0,0,1,1,0,1,1,1,0,1,1,1,1,1,1,0,1,1,0,0,0,0,1,0},{0,0,1,0,0,1,0,1,1,1,1,0,1,0,0,0,1,0,1,1,1,1,1,1,1},{0,0,1,1,1,1,0,1,1,0,0,1,1,1,0,0,1,0,1,0,1,1,0,1,1},{1,0,0,1,0,1,0,0,1,1,1,1,0,0,0,1,0,1,0,1,1,1,0,1,1},{0,0,0,0,1,0,0,1,0,1,0,0,0,1,1,1,0,1,0,0,0,1,0,0,0},{1,1,1,1,0,0,0,0,0,1,0,0,0,0,0,1,1,0,1,1,1,0,0,0,1},{1,0,0,1,1,1,1,1,1,1,0,1,0,0,1,1,1,1,1,1,0,1,0,1,1},{1,1,0,0,0,0,0,0,1,0,1,0,1,0,0,1,0,1,0,1,0,0,0,0,0},{1,0,0,0,0,1,1,0,1,1,0,1,0,1,0,1,1,1,0,1,0,1,0,0,0},{0,1,1,0,1,1,0,1,0,1,0,1,1,1,1,1,0,1,0,0,1,1,0,1,1},{1,1,0,1,1,0,1,1,1,0,1,0,1,0,0,1,0,1,0,0,0,1,0,0,0},{0,1,0,0,1,0,0,0,0,0,0,1,0,1,1,0,1,1,0,1,1,1,0,0,0},{0,1,1,1,1,0,0,0,0,0,1,0,1,0,1,0,1,0,1,0,1,1,1,0,1},{1,0,0,1,1,1,0,0,0,1,0,1,0,0,0,0,0,0,1,1,0,0,0,0,0},{1,0,0,0,1,0,1,0,0,0,0,1,0,0,0,0,1,1,1,0,1,0,0,0,0},{0,0,0,1,0,1,0,1,1,1,0,1,1,1,0,1,1,1,0,1,0,0,1,1,0},{1,0,1,1,1,0,1,1,1,0,1,1,1,0,0,0,0,0,1,1,0,0,1,0,1},{0,1,1,0,0,0,0,1,0,1,0,0,1,0,1,1,1,1,0,0,0,1,1,0,0},{0,1,0,0,1,1,0,1,1,0,1,0,1,0,0,0,0,1,1,0,0,1,1,1,1},{0,0,1,0,1,1,1,1,0,1,1,1,0,1,0,1,1,1,1,1,1,1,1,0,0},{1,0,1,0,1,1,0,1,1,0,0,0,1,1,1,1,1,1,1,0,1,1,1,1,1},{1,0,1,0,1,1,1,0,1,0,1,0,1,1,0,0,1,1,0,0,1,1,0,0,1},{0,1,0,0,0,1,0,1,1,1,1,1,0,0,0,1,0,0,0,0,1,0,1,0,1},{0,0,1,0,1,0,1,1,0,0,0,0,0,0,0,0,1,0,0,1,1,1,0,0,1},{1,1,0,1,0,0,1,1,0,0,1,1,1,0,0,1,1,1,1,0,0,0,0,1,0},{1,0,0,0,0,0,0,1,0,0,0,1,0,1,0,1,0,1,1,1,0,0,0,0,0},{1,0,0,0,1,0,0,1,0,0,0,0,1,1,1,1,1,1,0,0,1,0,0,0,1},{0,1,0,1,0,1,1,0,0,1,1,1,0,1,0,0,0,1,1,0,0,1,1,1,0},{1,1,1,1,0,1,0,0,1,0,1,1,0,0,1,1,0,1,0,1,0,0,1,0,1},{1,0,0,1,0,1,0,1,0,0,1,0,1,0,0,0,1,1,0,1,1,1,0,0,0},{1,1,1,0,1,0,0,0,1,1,1,0,0,0,0,0,0,0,1,0,0,0,0,0,0},{1,1,1,0,0,0,1,0,0,1,0,1,1,0,1,1,1,1,0,0,1,1,0,0,1},{0,1,1,0,1,0,0,1,0,0,0,0,1,0,1,1,1,1,1,0,1,1,1,0,1},{0,1,1,1,1,1,1,1,1,1,0,1,0,0,1,1,1,0,0,0,0,1,1,0,1},{0,0,1,0,1,1,0,0,1,1,1,0,1,0,0,0,1,0,1,0,1,0,1,1,0},{1,0,0,1,1,0,1,0,0,1,0,0,1,1,0,1,0,0,1,0,0,1,0,1,0}}
	expected = 177
	assert.Equal(t, expected, maxAreaOfIsland(input))
}
