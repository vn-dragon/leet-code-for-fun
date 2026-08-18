package daily

func kthSmallestTree(root *TreeNode, k int) int {
	var results []int
	dfskthSmallestTree(root, &results)
	return results[k-1]
}

func dfskthSmallestTree(node *TreeNode, results *[]int) {
	if node == nil {
		return
	}
	dfskthSmallestTree(node.Left, results)
	*results = append(*results, node.Val)
	dfskthSmallestTree(node.Right, results)
}
