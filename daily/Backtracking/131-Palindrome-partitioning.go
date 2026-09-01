package daily

func partition(s string) [][]string {
	var res [][]string
	var path []string

	isPal := func(sub string) bool {
		l, r := 0, len(sub)-1
		for l < r {
			if sub[l] != sub[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(s) {
			part := make([]string, len(path))
			copy(part, path)
			res = append(res, part)
			return
		}
		for end := start + 1; end <= len(s); end++ {
			sub := s[start:end]
			if isPal(sub) {
				path = append(path, sub)
				backtrack(end)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)
	return res
}
