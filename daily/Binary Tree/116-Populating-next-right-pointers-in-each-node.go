package daily

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	var bfs func(levelNodes []*Node)
	bfs = func(levelNodes []*Node) {
		if len(levelNodes) == 0 {
			return
		}

		nextLevel := []*Node{}
		for _, n := range levelNodes {
			if n.Left != nil {
				nextLevel = append(nextLevel, n.Left)
			}
			if n.Right != nil {
				nextLevel = append(nextLevel, n.Right)
			}
		}

		for i := 0; i < len(nextLevel)-1; i++ {
			nextLevel[i].Next = nextLevel[i+1]
		}

		bfs(nextLevel)
	}

	if root == nil {
		return nil
	}
	bfs([]*Node{root})
	return root
}
