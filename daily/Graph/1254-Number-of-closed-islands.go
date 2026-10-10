package daily

func closedIsland(grid [][]int) int {
	count := 0
	rows, cols := len(grid), len(grid[0])

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] == 1 {
			return
		}
		grid[row][col] = 1
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if (row == 0 || row == rows-1 || col == 0 || col == cols-1) && grid[row][col] == 0 {
				dfs(row, col)
			}
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 0 {
				count++
				dfs(row, col)
			}
		}
	}
	return count
}
