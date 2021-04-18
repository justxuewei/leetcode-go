package jzoffer

func isSymmetric(root *TreeNode) bool {
	if root == nil { return true }
	return nodeSymmetric(root.Left, root.Right)
}

func nodeSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil { return true }
	if left == nil || right == nil || left.Val != right.Val { return false }
	return nodeSymmetric(left.Left, right.Right) && nodeSymmetric(left.Right, right.Left)
}
