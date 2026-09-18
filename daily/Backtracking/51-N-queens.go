package daily

func solveNQueens(n int) [][]string {
	res := [][]string{}
	cols := map[int]bool{}
	diag1 := map[int]bool{}
	diag2 := map[int]bool{}

	// board[row][col], each cell is a byte '.' or 'Q'
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			current := make([]string, 0, n)
			for _, r := range board {
				current = append(current, string(r))
			}
			res = append(res, current)
			return
		}

		for col := 0; col < n; col++ {
			if cols[col] || diag1[row-col] || diag2[row+col] {
				continue
			}
			board[row][col] = 'Q'
			cols[col] = true
			diag1[row-col] = true
			diag2[row+col] = true

			backtrack(row + 1)

			board[row][col] = '.'
			delete(cols, col)
			delete(diag1, row-col)
			delete(diag2, row+col)
		}
	}

	backtrack(0)
	return res
}
