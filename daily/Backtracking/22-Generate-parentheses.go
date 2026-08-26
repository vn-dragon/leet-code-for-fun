package daily

func generateParenthesis(n int) []string {
	var res []string
	var stacks []byte

	var backtrack func(open, closed int)
	backtrack = func(open, closed int) {
		if open == n && closed == n {
			res = append(res, string(stacks))
			return
		}

		if open < n {
			stacks = append(stacks, '(')
			backtrack(open+1, closed)
			stacks = stacks[:len(stacks)-1]
		}
		if closed < open {
			stacks = append(stacks, ')')
			backtrack(open, closed+1)
			stacks = stacks[:len(stacks)-1]
		}
	}
	backtrack(0, 0)
	return res
}
