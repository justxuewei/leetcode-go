package jzoffer

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 { return true }
	boundary := len(postorder)-2
	for boundary>=0 && postorder[boundary]>postorder[len(postorder)-1] { boundary-- }
	for i:=0; i<=boundary;i++ {
		if postorder[i] > postorder[len(postorder)-1] {
			return false
		}
	}
	return verifyPostorder(postorder[:boundary+1]) && verifyPostorder(postorder[boundary+1:len(postorder)-1])
}
