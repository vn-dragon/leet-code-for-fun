package daily

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	results := [][]int{}
	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		levelLength := len(queue)
		levels := []int{}

		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:] // popleft

			levels = append(levels, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if level%2 != 0 {
			left, right := 0, len(levels)-1
			for left < right {
				levels[left], levels[right] = levels[right], levels[left]
				left++
				right--
			}
		}

		results = append(results, levels)
		level++
	}
	return results
}
