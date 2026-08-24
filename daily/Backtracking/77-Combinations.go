package daily

func combine(n int, k int) [][]int {
	res := [][]int{}
	path := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path) == k {
			combination := make([]int, len(path))
			copy(combination, path)
			res = append(res, combination)
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(1)
	return res
}
