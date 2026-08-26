class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        stacks, res = [], []
        def backtrack(open, close):
            if open == close == n:
                res.append("".join(stacks))
                return
            if open < n:
                stacks.append("(")
                backtrack(open + 1, close)
                stacks.pop()
            if close < open:
                stacks.append(")")
                backtrack(open, close + 1)
                stacks.pop()
        backtrack(0, 0)
        return res