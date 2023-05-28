from typing import List


class Solution:
    def maxIncreasingCells(self, mat: List[List[int]]) -> int:
        ROWS = len(mat)
        COLS = len(mat[0])

        def dfs(r, c, cost):
            # if r == ROWS - 1 or c == COLS - 1 or r == 0 or c == 0:
            #     return 1

            if r > 0 and mat[r - 1][c] > mat[r][c]:
                cost = max(cost, dfs(r - 1, c, cost + 1))
            if c > 0 and mat[r][c - 1] > mat[r][c]:
                cost = max(cost, dfs(r, c - 1, cost + 1))
            if c < COLS - 1 and mat[r][c + 1] > mat[r][c]:
                cost = max(cost, dfs(r, c + 1, cost + 1))
            if r < ROWS - 1 and mat[r + 1][c] > mat[r][c]:
                cost = max(cost, dfs(r + 1, c, cost + 1))

            return cost

        maxVal = 0
        for i in range(ROWS):
            for j in range(COLS):
                maxVal = max(maxVal, dfs(i, j, 1))

        return maxVal


# target =
obj = Solution
nums = [[3, 1, 6], [-9, 5, 7]]

print(obj.maxIncreasingCells(obj, nums))
