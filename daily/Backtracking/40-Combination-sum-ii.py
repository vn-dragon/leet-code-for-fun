class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        results = []
        candidates.sort()
        def backtrack(start, path, remaining):
            if remaining == 0:
                results.append(path[:])
                return
            for i in range(start, len(candidates)):
                if candidates[i] > remaining:
                    return
                if i > start and candidates[i - 1] == candidates[i]:
                    continue
                path.append(candidates[i])
                backtrack(i + 1, path, remaining - candidates[i])
                path.pop()
        backtrack(0, [], target)
        return results