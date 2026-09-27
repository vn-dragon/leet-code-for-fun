package daily

func solve(board [][]byte) {
	rows, cols := len(board), len(board[0])

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row >= rows || row < 0 || col >= cols || col < 0 || board[row][col] != 'O' {
			return
		}
		board[row][col] = 'T'
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if (row == 0 || row == rows-1 || col == 0 || col == cols-1) && board[row][col] == 'O' {
				dfs(row, col)
			}
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == 'O' {
				board[row][col] = 'X'
			} else if board[row][col] == 'T' {
				board[row][col] = 'O'
			}
		}
	}
}
