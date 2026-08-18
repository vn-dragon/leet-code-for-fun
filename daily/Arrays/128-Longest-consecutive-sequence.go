package daily

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLength := 0
	for num := range numSet {
		if !numSet[num-1] {
			current := 0
			for numSet[num+current] {
				current++
			}
			maxLength = max(maxLength, current)
		}
	}
	return maxLength
}
