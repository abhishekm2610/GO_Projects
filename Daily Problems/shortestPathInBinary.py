from typing import List


class Solution:
    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        dp = {}
        N = len(grid)
        queue = []
        visited = set()
        queue.append((0, 0))
        cost = {}

        for i in range(N):
            for j in range(N):
                cost[(i, j)] = float("inf")

        if grid[0][0] == 1 or grid[N - 1][N - 1] == 1:
            return -1
        cost[(0, 0)] = 1

        while len(queue) != 0:
            node = queue[0]
            x, y = node

            queue = queue[1:]
            if x == -1 or y == -1 or x == N or y == N:
                continue
            if node in visited:
                continue

            visited.add((x, y))

            if x != 0 and grid[x - 1][y] == 0:
                queue.append((x - 1, y))
                cost[(x - 1, y)] = min(cost[(x - 1, y)], cost[(x, y)] + 1)
            if x != 0 and y != 0 and grid[x - 1][y - 1] == 0:
                queue.append((x - 1, y - 1))
                cost[(x - 1, y - 1)] = min(cost[(x - 1, y - 1)], cost[(x, y)] + 1)
            if y != 0 and grid[x][y - 1] == 0:
                queue.append((x, y - 1))
                cost[(x, y - 1)] = min(cost[(x, y - 1)], cost[(x, y)] + 1)
            if x != N - 1 and grid[x + 1][y] == 0:
                queue.append((x + 1, y))
                cost[(x + 1, y)] = min(cost[(x + 1, y)], cost[(x, y)] + 1)
            if x != N - 1 and y != N - 1 and grid[x + 1][y + 1] == 0:
                queue.append((x + 1, y + 1))
                cost[(x + 1, y + 1)] = min(cost[(x + 1, y + 1)], cost[(x, y)] + 1)
            if y != N - 1 and grid[x][y + 1] == 0:
                queue.append((x, y + 1))
                cost[(x, y + 1)] = min(cost[(x, y + 1)], cost[(x, y)] + 1)
            if y != N - 1 and x != 0 and grid[x - 1][y + 1] == 0:
                queue.append((x - 1, y + 1))
                cost[(x - 1, y + 1)] = min(cost[(x - 1, y + 1)], cost[(x, y)] + 1)
            if x != N - 1 and y != 0 and grid[x + 1][y - 1] == 0:
                queue.append((x + 1, y - 1))
                cost[(x + 1, y - 1)] = min(cost[(x + 1, y - 1)], cost[(x, y)] + 1)
        if cost[(N - 1, N - 1)] != float("inf"):
            return cost[(N - 1, N - 1)]

        else:
            return -1


# target =
obj = Solution
nums = [
    [0, 1, 0, 1, 0],
    [1, 0, 0, 0, 1],
    [0, 0, 1, 1, 1],
    [0, 0, 0, 0, 0],
    [1, 0, 1, 0, 0],
]
# n = 7

print(obj.shortestPathBinaryMatrix(obj, nums))
