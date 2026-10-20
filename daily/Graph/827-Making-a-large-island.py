class Solution:
    def largestIsland(self, grid: list[list[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        def dfs(row, col, size, id_size):
            if row >= rows or row < 0 or col < 0 or col >= cols or grid[row][col] != 1:
                return
            grid[row][col] = id_size
            size[id_size] += 1
            dfs(row + 1, col, size, id_size)
            dfs(row - 1, col, size, id_size)
            dfs(row, col + 1, size, id_size)
            dfs(row, col - 1, size, id_size)
        size = defaultdict(int)
        counter = 2
        for row in range(rows):
            for col in range(cols):
                if grid[row][col] == 1:
                    dfs(row, col, size, counter)
                    counter += 1
        
        res = max(size.values(), default = 0)
        for row in range(rows):
            for col in range(cols):
                if grid[row][col] == 0:
                    connected = set()
                    if row > 0: connected.add(grid[row - 1][col])
                    if row < rows - 1: connected.add(grid[row + 1][col])
                    if col > 0: connected.add(grid[row][col - 1])
                    if col < cols - 1: connected.add(grid[row][col + 1])
                    current_island = 1 + sum(size[i] for i in connected)
                    res = max(res, current_island)
        return res