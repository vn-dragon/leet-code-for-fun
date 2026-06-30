package daily

func twoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left <= right {
		sumTwoNums := numbers[left] + numbers[right]
		if sumTwoNums > target {
			right--
		} else if sumTwoNums < target {
			left++
		} else if sumTwoNums == target {
			return []int{left + 1, right + 1}
		}
	}
	return []int{left + 1, right + 1}
}
