class Solution:
    def closedIsland(self, grid: list[list[int]]) -> int:
        count = 0
        rows, cols = len(grid), len(grid[0])
        def dfs(row, col):
            if row < 0 or row >= rows or col < 0 or col >= cols or grid[row][col] == 1:
                return
            grid[row][col] = 1
            dfs(row - 1, col)
            dfs(row + 1, col)
            dfs(row, col - 1)
            dfs(row, col + 1)

        for row in range(rows):
            for col in range(cols):
                if (row == 0 or row == rows - 1 or col == 0 or col == cols - 1) and grid[row][col] == 0:
                    dfs(row, col)    
        for row in range(rows):
            for col in range(cols):
                if grid[row][col] == 0:
                    count += 1
                    dfs(row, col)
        return count