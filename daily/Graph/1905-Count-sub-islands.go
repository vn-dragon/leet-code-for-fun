package daily

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	rows, cols := len(grid2), len(grid2[0])
	var dfs func(row, col int) bool
	dfs = func(row, col int) bool {
		if row < 0 || row >= rows || col < 0 || col >= cols || grid2[row][col] == 0 {
			return true
		}

		grid2[row][col] = 0
		isSub := grid1[row][col] == 1

		up := dfs(row-1, col)
		down := dfs(row+1, col)
		left := dfs(row, col-1)
		right := dfs(row, col+1)
		return isSub && up && down && left && right
	}

	res := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid2[row][col] == 1 && dfs(row, col) {
				res++
			}
		}
	}
	return res
}
