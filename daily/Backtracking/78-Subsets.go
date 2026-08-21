package daily

func subsets(nums []int) [][]int {
	var results [][]int
	var path []int

	var backtrack func(start int)
	backtrack = func(start int) {
		subset := make([]int, len(path))
		copy(subset, path)
		results = append(results, subset)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return results
}
