package daily

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	path := []int{}

	var backtrack func(start, remaining int)
	backtrack = func(start, remaining int) {
		if remaining == 0 {
			combo := make([]int, len(path))
			copy(combo, path)
			res = append(res, combo)
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > remaining {
				return
			}
			if i > start && candidates[i-1] == candidates[i] {
				continue
			}
			path = append(path, candidates[i])
			backtrack(i+1, remaining-candidates[i])
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target)
	return res
}
