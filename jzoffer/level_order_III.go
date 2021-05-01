package jzoffer

func levelOrder1(root *TreeNode) [][]int {
	if root == nil { return nil }
	var queue []*TreeNode
	queue = append(queue, root)
	ret := make([][]int, 0)
	levelret := make([]int, 1)
	levelremaining := 1

	last := root

	for len(queue) > 0 {
		node := queue[0]
		if len(ret) % 2 == 0 {
			levelret[len(levelret)-levelremaining] = node.Val
		} else {
			levelret[levelremaining-1] = node.Val
		}
		queue = queue[1:]
		levelremaining--
		// levelret = append(levelret, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if last == node {
			ret = append(ret, levelret)
			levelret = make([]int, len(queue))
			levelremaining = len(queue)
			if len(queue) > 0 {
				last = queue[len(queue)-1]
			}
		}
	}

	return ret
}
