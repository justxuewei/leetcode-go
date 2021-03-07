package traversal

import (
	btree "github.com/xavier-niu/leetcode/binary_tree"
	ds "github.com/xavier-niu/leetcode/data_struct"
)

func InorderTraversalRecursively(root *btree.TreeNode) []int {
	output := make([]int, 0)
	inorderTraversalRecursively(root, &output)
	return output
}

func inorderTraversalRecursively(node *btree.TreeNode, output *[]int) {
	if node == nil { return }
	inorderTraversalRecursively(node.Left, output)
	*output = append(*output, node.Val)
	inorderTraversalRecursively(node.Right, output)
}

func inorderTraversalIteratively(root *btree.TreeNode) []int {
	if root == nil { return nil }

	var output []int
	crtNode := root
	stack := ds.NewStackWithoutCap()

	for crtNode != nil {
		stack.Push(crtNode)
		if crtNode.Left == nil {
			// pop node from stack
			var tmpNode *btree.TreeNode
			for {
				tmpNode, _ = stack.Pop().(*btree.TreeNode)
				if tmpNode == nil { return output }
				output = append(output, tmpNode.Val)
				if tmpNode.Right != nil {
					crtNode = tmpNode.Right
					break
				}
			}
		} else {
			crtNode = crtNode.Left
		}
	}

	return output
}
