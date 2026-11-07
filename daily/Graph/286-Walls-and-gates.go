package daily

func wallsAndGates(rooms [][]int) {
	rows, cols := len(rooms), len(rooms[0])
	const INF = 2147483647
	queue := [][2]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if rooms[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		r, c := cell[0], cell[1]
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && rooms[nr][nc] == INF {
				rooms[nr][nc] = rooms[r][c] + 1
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}
}
