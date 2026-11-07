class Solution:
    def wallsAndGates(self, rooms: list[list[int]]) -> None:
        """
        Do not return anything, modify rooms in-place instead.
        """
        rows, cols = len(rooms), len(rooms[0])
        queue = deque()
        for row in range(rows):
            for col in range(cols):
                if rooms[row][col] == 0:
                    queue.append((row, col))
        
        INF = 2147483647
        while queue:
            for _ in range(len(queue)):
                r, c = queue.popleft()
                for dr, dc in [(1,0), (-1,0), (0,1), (0,-1)]:
                    nr, nc = r + dr, c + dc
                    if 0 <= nr < rows and 0 <= nc < cols and rooms[nr][nc] == INF:
                        rooms[nr][nc] = rooms[r][c] + 1
                        queue.append((nr, nc)) 