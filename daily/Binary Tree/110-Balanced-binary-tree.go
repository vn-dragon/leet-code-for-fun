package daily

func isBalanced(root *TreeNode) bool {
	var dfs func(*TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}
		leftHeight, leftBalance := dfs(node.Left)
		rightHeight, rightBalance := dfs(node.Right)

		diff := leftHeight - rightHeight
		if diff < 0 {
			diff = -diff
		}
		currentBalance := leftBalance && rightBalance && diff <= 1
		currentHeight := 1 + max(leftHeight, rightHeight)
		return currentHeight, currentBalance
	}
	_, balanced := dfs(root)
	return balanced
}
