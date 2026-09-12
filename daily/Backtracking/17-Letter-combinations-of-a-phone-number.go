package daily

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}

	phone := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	path := []byte{}
	var backtrack func(index int)
	backtrack = func(index int) {
		if len(path) == len(digits) {
			res = append(res, string(path))
			return
		}
		choices := phone[digits[index]]
		for i := 0; i < len(choices); i++ {
			path = append(path, choices[i])
			backtrack(index + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}
