class Solution:
    def countSubIslands(self, grid1: list[list[int]], grid2: list[list[int]]) -> int:
        rows, cols = len(grid2), len(grid2[0])
        def dfs(row, col):
            if row < 0 or row >= rows or col < 0 or col >= cols or grid2[row][col] == 0:
                return True
            grid2[row][col] = 0
            is_sub = grid1[row][col] == 1
            up, down, left, right = (
                dfs(row - 1, col),
                dfs(row + 1, col),
                dfs(row, col - 1),
                dfs(row, col + 1),
            )
            return is_sub and up and down and left and right
        
        res = 0
        for row in range(rows):
            for col in range(cols):
                if grid2[row][col] == 1 and dfs(row, col):
                    res += 1
        return res