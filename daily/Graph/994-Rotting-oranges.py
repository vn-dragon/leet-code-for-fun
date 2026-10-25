class Solution:
    def orangesRotting(self, grid: list[list[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        queue = deque()
        fresh = 0
        for row in range(rows):
            for col in range(cols):
                if grid[row][col] == 2:
                    queue.append((row, col))
                elif grid[row][col] == 1:
                    fresh += 1
        
        minutes = 0
        while queue and fresh > 0:
            for _ in range(len(queue)):
                row, col = queue.popleft()
                for r, c in ([(1, 0), (-1, 0), (0, 1), (0, -1)]):
                    new_row, new_col = row + r, col + c
                    if 0 <= new_row < rows and 0 <= new_col < cols and grid[new_row][new_col] == 1:
                        grid[new_row][new_col] = 2
                        queue.append((new_row, new_col))
                        fresh -= 1
            minutes += 1
        return minutes if fresh == 0 else -1
