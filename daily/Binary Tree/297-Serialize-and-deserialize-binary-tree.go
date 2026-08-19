package daily

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var sb []string
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb = append(sb, "X")
			return
		}
		sb = append(sb, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(sb, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	idx := 0
	var build func() *TreeNode
	build = func() *TreeNode {
		val := vals[idx]
		idx++
		if val == "X" {
			return nil
		}
		num, _ := strconv.Atoi(val)
		node := &TreeNode{Val: num}
		node.Left = build()
		node.Right = build()
		return node
	}
	return build()
}
