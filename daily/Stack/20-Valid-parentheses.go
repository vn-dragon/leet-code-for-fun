package daily

func isValid(s string) bool {
	corresponding := map[rune]rune{')': '(', '}': '{', ']': '['}
	stack := []rune{}

	for _, c := range s {
		if open, isClose := corresponding[c]; isClose {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}
