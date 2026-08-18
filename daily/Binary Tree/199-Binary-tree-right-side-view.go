package daily

func rightSideView(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level_length := len(queue)
		for i := range level_length {
			node := queue[0]
			queue = queue[1:]
			if i == level_length-1 {
				result = append(result, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}
