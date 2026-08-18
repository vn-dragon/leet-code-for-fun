package daily

import "math"

func maxPathSum(root *TreeNode) int {
	bestSum := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left < 0 {
			left = 0
		}
		if right < 0 {
			right = 0
		}

		bestSum = max(bestSum, node.Val+left+right)
		return node.Val + max(left, right)
	}
	dfs(root)
	return bestSum
}
