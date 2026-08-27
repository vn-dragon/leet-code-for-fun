package daily

import "sort"

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
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
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue
			}
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
