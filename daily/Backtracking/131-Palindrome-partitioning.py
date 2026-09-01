class Solution:
    def partition(self, s: str) -> List[List[str]]:
        res = []
        def is_palindrome(string) -> bool:
            left, right = 0, len(string) - 1
            while left < right:
                if string[left] != string[right]:
                    return False
                left += 1
                right -= 1
            return True
        def backtrack(start, path):
            if start == len(s):
                res.append(path[:])
                return
            for end in range(start + 1, len(s) + 1):
                sub_string = s[start:end]
                if is_palindrome(sub_string):
                    path.append(sub_string)
                    backtrack(end, path)
                    path.pop()
        backtrack(0, [])
        return res