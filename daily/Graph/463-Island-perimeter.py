class Solution:
    def islandPerimeterC2(self, grid: list[list[int]]) -> int:
        count = 0
        rows, cols = len(grid), len(grid[0])
        visited = set()
        def dfs(row, col):
            if row < 0 or row >= rows or col < 0 or col >= cols or grid[row][col] == 0:
                return 1
            if (row, col) in visited:
                return 0
            visited.add((row, col))
            top = dfs(row - 1, col)
            bottom = dfs(row + 1, col)
            left = dfs(row, col - 1)
            right = dfs(row, col + 1)
            return top + bottom + left + right
        for row in range(rows):
            for col in range(cols):
                if grid[row][col] == 1:
                    count += dfs(row, col)
        return count
    def islandPerimeter(self, grid: list[list[int]]) -> int:
        res = 0
        rows, cols = len(grid), len(grid[0])
        for row in range(rows):
            for col in range(cols):
                if grid[row][col] == 1:
                    res += 4
                    if row > 0 and grid[row - 1][col] == 1:
                        res -= 2
                    if col > 0 and grid[row][col - 1] == 1:
                        res -= 2
        return res