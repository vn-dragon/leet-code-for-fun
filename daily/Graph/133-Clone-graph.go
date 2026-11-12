package daily

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	clones := make(map[*Node]*Node)
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n, ok := clones[node]; ok {
			return n
		}
		newNode := &Node{Val: node.Val}
		clones[node] = newNode

		for _, nn := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, dfs(nn))
		}
		return newNode
	}
	return dfs(node)
}
