package daily

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row >= rows || row < 0 || col >= cols || col < 0 || grid[row][col] != 1 {
			return 0
		}
		grid[row][col] = 0
		return 1 + dfs(row+1, col) + dfs(row-1, col) + dfs(row, col+1) + dfs(row, col-1)
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {
				count = max(count, dfs(row, col))
			}
		}
	}
	return count
}
