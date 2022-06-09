package problems

import "math"

func isValidBST(root *TreeNode) bool {
	return validateNode(root, math.MinInt64, math.MaxInt64)
}

func validateNode(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	rootResult := true
	if node.Left != nil {
		rootResult = rootResult &&
			(node.Left.Val < node.Val) &&
			(validateNode(node.Left, min, node.Val))
	}
	if node.Right != nil {
		rootResult = rootResult &&
			(node.Right.Val > node.Val) &&
			(validateNode(node.Right, node.Val, max))
	}

	return rootResult
}
