package daily

func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)
	for i := range pacific {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}

	var dfs func(row, col, prevHeight int, visited [][]bool)
	dfs = func(row, col, prevHeight int, visited [][]bool) {
		if row < 0 || row >= rows || col < 0 || col >= cols ||
			visited[row][col] || heights[row][col] < prevHeight {
			return
		}
		visited[row][col] = true
		h := heights[row][col]
		dfs(row, col+1, h, visited)
		dfs(row, col-1, h, visited)
		dfs(row+1, col, h, visited)
		dfs(row-1, col, h, visited)
	}

	for row := 0; row < rows; row++ {
		dfs(row, 0, heights[row][0], pacific)
		dfs(row, cols-1, heights[row][cols-1], atlantic)
	}
	for col := 0; col < cols; col++ {
		dfs(0, col, heights[0][col], pacific)
		dfs(rows-1, col, heights[rows-1][col], atlantic)
	}

	res := [][]int{}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if pacific[row][col] && atlantic[row][col] {
				res = append(res, []int{row, col})
			}
		}
	}
	return res
}
