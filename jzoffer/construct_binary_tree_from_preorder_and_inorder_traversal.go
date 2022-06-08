package jzoffer

func buildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	inorderMid := -1
	for i, v := range inorder {
		if v == preorder[0] {
			inorderMid = i
			break
		}
	}
	left := buildTree(preorder[1:inorderMid+1], inorder[:inorderMid+1])
	right := buildTree(preorder[inorderMid+1:], inorder[inorderMid+1:])
	root.Left, root.Right = left, right
	return root
}
