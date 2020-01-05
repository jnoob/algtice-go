package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	// N -> a node, Nl -> left child of the node, Nr -> right child of the node
	// v(N) -> the value of the node, maxv(N) -> the max value of the subtree which's root is N
	// for a node N, the thief can choose rob or not
	// if rob
	//		maxv(N) = v(N) + maxv(Nll) + maxv(Nlr) + maxv(Nrl) + maxv(Nrr)
	// else
	//		maxv(N) = maxv(Nl) + maxv(Nr)
	memt := map[*TreeNode]int{}
	return maxv(root, memt)
}

func maxv(node *TreeNode, memt map[*TreeNode]int) int {
	if node == nil {
		return 0
	}

	max, exists := memt[node]
	if exists {
		return max
	}

	robv := node.Val + maxv_(node.Left, memt) + maxv_(node.Right, memt)
	nrobv := maxv(node.Left, memt) + maxv(node.Right, memt)
	if robv > nrobv {
		max = robv
	} else {
		max = nrobv
	}
	memt[node] = max
	return max
}

func maxv_(node *TreeNode, memt map[*TreeNode]int) int {
	if node == nil {
		return 0
	}
	return maxv(node.Left, memt) + maxv(node.Right, memt)
}
