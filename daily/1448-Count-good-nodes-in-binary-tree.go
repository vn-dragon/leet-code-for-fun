package daily

var res int = 0

func goodNodes(root *TreeNode) int {
	res = 0
	dfsGoodNodes(root, root.Val)
	return res
}

func dfsGoodNodes(node *TreeNode, max_val int) {
	if node == nil {
		return
	}
	if node.Val >= max_val {
		res += 1
		max_val = node.Val
	}
	dfsGoodNodes(node.Left, max_val)
	dfsGoodNodes(node.Right, max_val)
}
