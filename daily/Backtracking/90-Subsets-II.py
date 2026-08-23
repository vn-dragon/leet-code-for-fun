class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        results = []
        nums.sort()
        def backtrack(start, path):
            results.append(path[:])
            for i in range(start, len(nums)):
                if i > start and nums[i-1] == nums[i]:
                    continue
                path.append(nums[i])
                backtrack(i + 1, path)
                path.pop()
        backtrack(0, [])
        return results