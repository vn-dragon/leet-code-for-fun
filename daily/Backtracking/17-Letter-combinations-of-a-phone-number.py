class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        res = []
        if len(digits) == 0:
            return res
        phone_numbers = {
            "2": "abc",
            "3": "def",
            "4": "ghi",
            "5": "jkl",
            "6": "mno",
            "7": "pqrs",
            "8": "tuv",
            "9": "wxyz"
        }
        def backtrack(path, index):
            if len(path) == len(digits):
                letters = "".join(path)
                res.append(letters)
                return
            choices = phone_numbers[digits[index]]
            for i in range(len(choices)):
                path.append(choices[i])
                backtrack(path, index + 1)
                path.pop()
        backtrack([], 0)
        return res