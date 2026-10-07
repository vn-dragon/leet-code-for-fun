class Solution:
    def pacificAtlantic(self, heights: list[list[int]]) -> list[list[int]]:
        res = []
        rows, cols = len(heights), len(heights[0])
        pacific, atlantic = set(), set()
        def dfs(row, col, prev_height, visited):
            if row < 0 or row >= rows or col < 0 or col >= cols or (row, col) in visited or heights[row][col] < prev_height:
                return
            visited.add((row, col))
            dfs(row, col + 1, heights[row][col], visited)
            dfs(row, col - 1, heights[row][col], visited)
            dfs(row + 1, col, heights[row][col], visited)
            dfs(row - 1, col, heights[row][col], visited)
        for row in range(rows):
            dfs(row, 0, heights[row][0], pacific)
            dfs(row, cols - 1, heights[row][cols - 1], atlantic)
        for col in range(cols):
            dfs(0, col, heights[0][col], pacific)
            dfs(rows - 1, col, heights[rows - 1][col], atlantic)
        for row in range(rows):
            for col in range(cols):
                if (row, col) in atlantic and (row, col) in pacific:
                    res.append((row, col))
        return res