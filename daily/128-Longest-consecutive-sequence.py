class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        elements = set(nums)
        maxLength = 0
        for num in elements:
            if num - 1 not in elements:
                current = 0
                while num + current in elements:
                    current += 1
                maxLength = max(maxLength, current)

        return maxLength