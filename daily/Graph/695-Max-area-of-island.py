class Solution:
    def maxAreaOfIsland(self, grid: list[list[int]]) -> int:
        if not grid:
            return 0
        rows, cols = len(grid), len(grid[0])
        count = 0
        def dfs(row, col):
            if row >= rows or row < 0 or col >= cols or col < 0 or grid[row][col] != 1:
                return 0
            grid[row][col] = 0
            return 1 + dfs(row + 1, col) + dfs(row - 1, col) + dfs(row, col + 1) + dfs(row, col - 1)
        
        for row in range(rows):
            for col in range(cols):
                if grid[row][col] == 1:
                    count = max(count, dfs(row, col))
        return count