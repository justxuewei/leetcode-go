package jzoffer1

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	} else if len(preorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}

	var mid int

	for i := range inorder {
		if inorder[i] == preorder[0] {
			mid = i
			break
		}
	}

	return &TreeNode{
		Val: preorder[0],
		Left: buildTree(preorder[1:mid+1], inorder[:mid]),
		Right: buildTree(preorder[mid+1:], inorder[mid+1:]),
	}
}
