package btree

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil { return false }
	sums := pathSum(root)
	for _, v := range sums {
		if v == targetSum {
			return true
		}
	}
	return false
}

func pathSum(node *TreeNode) []int {
	if node == nil { return []int{} }
	left := pathSum(node.Left)
	right := pathSum(node.Right)

	// is leaf
	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}

	ret := append(left, right...)
	for i, _ := range ret {
		ret[i] += node.Val
	}
	return ret
}

func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil { return false }
	if root.Left == nil && root .Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum1(root.Left, targetSum-root.Val) || hasPathSum1(root.Right, targetSum-root.Val)
}
