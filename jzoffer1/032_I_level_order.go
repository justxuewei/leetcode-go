package jzoffer1

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	var queue []*TreeNode
	queue = append(queue, root)

	var node *TreeNode
	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:]
		ret = append(ret, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return ret
}
