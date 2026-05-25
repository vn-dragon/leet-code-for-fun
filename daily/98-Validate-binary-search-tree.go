package daily

import "math"

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, min, max int) bool
	dfs = func(root *TreeNode, minVal, maxVal int) bool {
		if root == nil {
			return true
		}

		if root.Val <= minVal || root.Val >= maxVal {
			return false
		}

		return dfs(root.Left, minVal, root.Val) && dfs(root.Right, root.Val, maxVal)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}
