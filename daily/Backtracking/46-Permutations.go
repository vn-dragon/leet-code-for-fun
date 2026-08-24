package daily

func permute(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	path := []int{}

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			perm := make([]int, len(path))
			copy(perm, path)
			res = append(res, perm)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}
