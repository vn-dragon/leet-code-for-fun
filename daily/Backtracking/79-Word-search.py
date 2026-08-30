class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        rows, cols = len(board), len(board[0])
        def backtrack(row, col, index):
            if index == len(word):
                return True
            if row >= rows or row < 0 or col >= cols or col < 0 or board[row][col] != word[index]:
                return False

            tmp = board[row][col]
            board[row][col] = "#"
            res = (
                backtrack(row + 1, col, index + 1) or 
                backtrack(row - 1, col, index + 1) or 
                backtrack(row, col - 1, index + 1) or 
                backtrack(row, col + 1, index + 1)
            )
            board[row][col] = tmp
            return res

        for r in range(rows):
            for c in range(cols):
                if backtrack(r, c, 0):
                    return True
        return False