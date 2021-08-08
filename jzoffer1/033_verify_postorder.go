package jzoffer1

func verifyPostorder(postorder []int) bool {
	if len(postorder) < 2 {
		return true
	}

	leftRoot := -1

	if postorder[len(postorder)-1] < postorder[len(postorder)-2] {
		for i := len(postorder) - 2; i >= 0; i-- {
			if postorder[i] < postorder[len(postorder)-1] {
				leftRoot = i
				break
			}
		}
	} else {
		leftRoot = len(postorder) - 2
	}

	if leftRoot >= 0 && leftRoot < len(postorder) {
		for i := leftRoot; i >= 0; i-- {
			if postorder[i] > postorder[len(postorder)-1] {
				return false
			}
		}
		return verifyPostorder(postorder[:leftRoot+1]) && verifyPostorder(postorder[leftRoot+1:len(postorder)-1])
	}

	return verifyPostorder(postorder[:len(postorder)-1])
}
