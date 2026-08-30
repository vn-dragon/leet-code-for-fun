package daily

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	var backtrack func(row, col, index int) bool
	backtrack = func(row, col, index int) bool {
		if index == len(word) {
			return true
		}
		if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != word[index] {
			return false
		}

		tmp := board[row][col]
		board[row][col] = '#'
		found := backtrack(row, col+1, index+1) ||
			backtrack(row, col-1, index+1) ||
			backtrack(row+1, col, index+1) ||
			backtrack(row-1, col, index+1)

		board[row][col] = tmp
		return found
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if backtrack(r, c, 0) {
				return true
			}
		}
	}
	return false
}
