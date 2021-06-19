package btree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

func compare(lhs *TreeNode, rhs *TreeNode) bool {
	if lhs == nil && rhs == nil {
		return true
	} else if lhs != nil && rhs != nil {
		return lhs.Val == rhs.Val && compare(lhs.Left, rhs.Right) && compare(lhs.Right, rhs.Left)
	} else {
		return false
	}
}
