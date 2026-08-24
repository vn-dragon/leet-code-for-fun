class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        results = []
        used = [False] * len(nums)
        def backtrack(path):
            if len(path) == len(nums):
                results.append(path[:])
                return
            for i in range(len(nums)):
                if used[i]:
                    continue
                used[i] = True
                path.append(nums[i])
                backtrack(path)
                path.pop()
                used[i] = False
        backtrack([])
        return results