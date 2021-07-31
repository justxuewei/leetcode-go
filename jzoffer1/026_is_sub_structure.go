package jzoffer1

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return isSameTree(A, B) ||
		isSubStructure(A.Left, B) ||
		isSubStructure(A.Right, B)
}

func isSameTree(A, B *TreeNode) bool {
	if B != nil && (A == nil || A.Val != B.Val) {
		return false
	}
	if B == nil {
		return true
	}
	return isSameTree(A.Left, B.Left) &&
		isSameTree(A.Right, B.Right)
}
