package btree

import (
	ds "github.com/xavier-niu/leetcode/data_struct"
)

func InorderTraversalRecursively(root *TreeNode) []int {
	output := make([]int, 0)
	inorderTraversalRecursively(root, &output)
	return output
}

func inorderTraversalRecursively(node *TreeNode, output *[]int) {
	if node == nil { return }
	inorderTraversalRecursively(node.Left, output)
	*output = append(*output, node.Val)
	inorderTraversalRecursively(node.Right, output)
}

func inorderTraversalIteratively(root *TreeNode) []int {
	if root == nil { return nil }

	var output []int
	crtNode := root
	stack := ds.NewStackWithoutCap()

	for crtNode != nil {
		stack.Push(crtNode)
		if crtNode.Left == nil {
			// pop node from stack
			var tmpNode *TreeNode
			for {
				topElem, err := stack.Pop()
				if err != nil { return output }
				tmpNode, _ = topElem.(*TreeNode)
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
