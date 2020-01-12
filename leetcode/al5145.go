package leetcode

func sumEvenGrandparent(root *TreeNode) int {
	if root == nil {
		return 0
	}
	levelNodes := []*TreeNode{root}
	val := 0
	for len(levelNodes) > 0 {
		nextLevel := []*TreeNode{}
		for _, node := range levelNodes {
			if node.Val%2 == 0 {
				val += ccVal(node)
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		levelNodes = nextLevel
	}
	return val
}

func ccVal(node *TreeNode) int {
	return cVal(node.Left) + cVal(node.Right)
}

func cVal(node *TreeNode) int {
	if node == nil {
		return 0
	}
	val := 0
	if node.Right != nil {
		val += node.Right.Val
	}
	if node.Left != nil {
		val += node.Left.Val
	}
	return val
}
