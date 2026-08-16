package daily

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max_index, max_value := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max_value {
			max_value = nums[i]
			max_index = i
		}
	}
	node := &TreeNode{}
	node.Val = max_value
	node.Left = constructMaximumBinaryTree(nums[:max_index])
	node.Right = constructMaximumBinaryTree(nums[max_index+1:])
	return node
}
