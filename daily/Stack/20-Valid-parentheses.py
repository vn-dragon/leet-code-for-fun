class Solution:
    def isValid(self, s: str) -> bool:
        corresponding_bracket = { ')': '(', '}': '{', ']': '[' }
        stacks = []
        for c in s:
            if c in corresponding_bracket: 
                if stacks and corresponding_bracket[c] == stacks[-1]:
                    stacks.pop()
                else:
                    return False
            else:
                stacks.append(c)
        return True if not stacks else False