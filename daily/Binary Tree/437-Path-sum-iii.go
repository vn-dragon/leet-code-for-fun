package daily

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return dfsPathSum(root, 0, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func dfsPathSum(node *TreeNode, currentSum int, targetSum int) int {
	if node == nil {
		return 0
	}

	currentSum += node.Val
	count := 0
	if currentSum == targetSum {
		count = 1
	}
	count += dfsPathSum(node.Left, currentSum, targetSum)
	count += dfsPathSum(node.Right, currentSum, targetSum)
	return count
}
