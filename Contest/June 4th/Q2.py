from typing import List


class Solution:
    def matrixSumQueries(self, n: int, queries: List[List[int]]) -> int:
        grid = [[0] * n for i in range(n)]
        print(grid)
        ans = 0
        prev = 0
        visited_i = {}
        visited_j = {}
        for i in range(n):
            visited_j[i] = 0
            visited_i[i] = 0
        for q in queries:
            if q[0] == 0:
                ans += q[2] * n
                ans -= visited_i[q[1]]
                visited_i[q[1]] = q[2]
                for j in visited_j:
                    visited_j[j] = q[2]
            else:
                ans += q[2] * n
                ans -= visited_j[q[1]]
                visited_j[q[1]] = q[2]
                for i in visited_i:
                    visited_i[i] = q[2]

            print(visited_i, visited_j)

            # ans += q[2] * n
            # ans -= visited_i[q[1]]
            # prev = q[2]

        # ans = 0
        # for i in range(n):
        #     ans += sum(grid[i])

        return ans


# target =
obj = Solution
# s = "leetcode"
grid = [[0, 0, 4], [0, 1, 2], [1, 0, 1], [0, 2, 3], [1, 2, 1]]
print(obj.matrixSumQueries(obj, 3, grid))
