package btree

// my solution
// runtime: 0ms
// memory usage: 4.5MB
func maxDepth(root *TreeNode) int {
	maxDepth := 0
	maxDepthRecursively(root, 0, &maxDepth)
	return maxDepth
}

func maxDepthRecursively(node *TreeNode, depth int, maxDepth *int) {
	if node == nil {
		if *maxDepth < depth {
			*maxDepth = depth
		}
		return
	}
	maxDepthRecursively(node.Left, depth+1, maxDepth)
	maxDepthRecursively(node.Right, depth+1, maxDepth)
}

// standard solution
// runtime: 4ms
// memory usage: 4.4MB
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(lhs int, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}
