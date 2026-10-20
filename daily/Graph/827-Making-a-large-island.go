package daily

func largestIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	size := make(map[int]int)

	var dfs func(row, col, id int)
	dfs = func(row, col, id int) {
		if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] != 1 {
			return
		}
		grid[row][col] = id
		size[id]++
		dfs(row+1, col, id)
		dfs(row-1, col, id)
		dfs(row, col+1, id)
		dfs(row, col-1, id)
	}

	counter := 2
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {
				dfs(row, col, counter)
				counter++
			}
		}
	}

	res := 0
	for _, s := range size {
		if s > res {
			res = s
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 0 {
				connected := make(map[int]struct{})
				if row > 0 {
					connected[grid[row-1][col]] = struct{}{}
				}
				if row < rows-1 {
					connected[grid[row+1][col]] = struct{}{}
				}
				if col > 0 {
					connected[grid[row][col-1]] = struct{}{}
				}
				if col < cols-1 {
					connected[grid[row][col+1]] = struct{}{}
				}
				current := 1
				for id := range connected {
					current += size[id]
				}
				if current > res {
					res = current
				}
			}
		}
	}
	return res
}
