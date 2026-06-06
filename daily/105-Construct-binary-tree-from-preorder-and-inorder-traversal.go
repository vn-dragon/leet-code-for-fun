package daily

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	node := &TreeNode{Val: rootVal}

	// Find rootVal's position in inorder
	mid := 0
	for i, v := range inorder {
		if v == rootVal {
			mid = i
			break
		}
	}

	leftInorder := inorder[:mid]
	rightInorder := inorder[mid+1:]

	leftSize := len(leftInorder)
	leftPreorder := preorder[1 : 1+leftSize]
	rightPreorder := preorder[1+leftSize:]

	node.Left = buildTree(leftPreorder, leftInorder)
	node.Right = buildTree(rightPreorder, rightInorder)

	return node
}
