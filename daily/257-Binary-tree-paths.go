package daily

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	results := []string{}
	path := []int{}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			parts := make([]string, len(path))
			for i, v := range path {
				parts[i] = strconv.Itoa(v)
			}
			results = append(results, strings.Join(parts, "->"))
		}

		dfs(node.Left)
		dfs(node.Right)

		path = path[:len(path)-1]
	}

	dfs(root)
	return results
}
