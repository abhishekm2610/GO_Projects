from typing import List


class Solution:
    def differenceOfDistinctValues(self, grid: List[List[int]]) -> List[List[int]]:
        ROWS = len(grid)
        COLS = len(grid[0])
        ans = [[0] * COLS for i in grid]

        for i in range(len(grid)):
            for j in range(len(grid[0])):
                leftDiag = set()
                r, c = i - 1, j - 1
                while r >= 0 and c >= 0:
                    leftDiag.add(grid[r][c])
                    r -= 1
                    c -= 1
                rightDiag = set()
                r, c = i + 1, j + 1
                while r < ROWS and c < COLS:
                    rightDiag.add(grid[r][c])
                    # rightDiag += grid[r][c]
                    r += 1
                    c += 1
                print(i, j, leftDiag, rightDiag)
                ans[i][j] = abs(len(leftDiag) - len(rightDiag))

        return ans


# target =
obj = Solution
# s = "leetcode"
grid = [[1, 2, 3], [3, 1, 5], [3, 2, 1], [1, 2, 3]]

print(obj.differenceOfDistinctValues(obj, grid))
