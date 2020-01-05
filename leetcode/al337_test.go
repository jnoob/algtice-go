package leetcode

import (
	"fmt"
	"testing"
)

func Test337AlgRobSmple(t *testing.T) {
	seq := []int{3,2,3,-1,3,-1,1}
	//seq := []int{3,4,5,1,3,-1,1}
	tree := buildBinaryTree(seq)
	printBinaryTree(tree)
	fmt.Printf("MAX = %v\v", rob(tree))
}

func printBinaryTree(root *TreeNode) {
	thisLevel := []*TreeNode{root}

	fmt.Print("\n")
	for thisLevel != nil {
		allNil := true
		nextLevel := []*TreeNode{}
		for _, node := range thisLevel {
			if node == nil {
				fmt.Print("-, ")
				nextLevel = append(nextLevel, nil, nil)
			} else {
				fmt.Printf("%v, ", node.Val)
				nextLevel = append(nextLevel, node.Left, node.Right)
				if node.Left != nil || node.Right != nil {
					allNil = false
				}
			}
		}

		if allNil {
			thisLevel = nil
		} else {
			thisLevel = nextLevel
		}
		fmt.Print("\n")
	}
}

func buildBinaryTree(seq []int) *TreeNode {
	root := &TreeNode{
		Val:   seq[0],
	}
	i := 0
	nodes := []*TreeNode{root}
	for nodes != nil {
		nextLevel := []*TreeNode{}
		for j, node := range nodes {
			if node == nil {
				nextLevel = append(nextLevel, nil, nil)
			} else {
				li := l(i+j)
				if li >= len(seq) {
					nodes = nil
					break
				}
				nl := buildTreeNode(seq[li])
				node.Left = nl
				nextLevel = append(nextLevel, nl)

				ri := r(i+j)
				if ri >= len(seq) {
					nodes = nil
					break
				}
				nr := buildTreeNode(seq[ri])
				node.Right = nr
				nextLevel = append(nextLevel, nr)
			}
		}
		i += len(nodes)
		if nodes != nil {
			nodes = nextLevel
		}
	}
	return root
}

func buildTreeNode(val int) *TreeNode {
	if val == -1 {
		return nil
	} else {
		return &TreeNode{Val:val}
	}
}

func l(i int) int {
	return 2 * i + 1
}

func r(i int) int {
	return 2*i + 2
}