from typing import List


class Solution:
    def countNegatives(self, grid: List[List[int]]) -> int:
        count = 0

        ROWS = len(grid) - 1
        COLS = len(grid[0]) - 1

        r = 0
        c = COLS

        while r <= ROWS and c >= 0:
            print(grid[r][c], r, c)
            if grid[r][c] < 0:
                count += ROWS - r
                count += 1
                c -= 1
                continue

            r += 1

        return count


# target =
obj = Solution
nums = [[4, 3, 2, -1], [3, 2, 1, -1], [1, 1, -1, -2], [-1, -1, -2, -3]]
# n = 7

print(obj.countNegatives(obj, nums))
