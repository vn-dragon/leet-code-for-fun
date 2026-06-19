package daily

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		return balanceNode(root.Left, root.Right)
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func balanceNode(leftNode *TreeNode, rightNode *TreeNode) *TreeNode {
	if rightNode == nil {
		return leftNode
	}
	rightNode.Left = balanceNode(leftNode, rightNode.Left)
	return rightNode
}
