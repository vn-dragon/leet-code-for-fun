package daily

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var dfs func(node *TreeNode, currentSum int) bool
	dfs = func(node *TreeNode, currentSum int) bool {
		if node == nil {
			return false
		}

		currentSum += node.Val
		if node.Left == nil && node.Right == nil {
			return currentSum == targetSum
		}
		left := dfs(node.Left, currentSum)
		right := dfs(node.Right, currentSum)
		return left || right
	}
	return dfs(root, 0)
}
