class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        first, last = 0, len(numbers) - 1
        while first < last:
            currentSum = numbers[first] + numbers[last]
            if currentSum < target:
                first += 1
            elif currentSum > target:
                last -= 1
            else:
                return [first + 1, last + 1]
        return []