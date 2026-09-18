class Solution:
    def solveNQueens(self, n: int) -> list[list[str]]:
        res = []
        cols, diag1, diag2 = set(), set(), set()
        board = [["."] * n for _ in range(n)]

        def backtrack(row):
            if row == n:
                current = []
                for r in board:
                    current.append("".join(r))
                res.append(current)
                return

            for col in range(n):
                if col in cols or (row - col) in diag1 or (row + col) in diag2:
                    continue
                board[row][col] = "Q"
                cols.add(col)
                diag1.add(row - col)
                diag2.add(row + col)
                backtrack(row + 1)
                board[row][col] = "."
                cols.remove(col)
                diag1.remove(row - col)
                diag2.remove(row + col)
        backtrack(0)
        return res