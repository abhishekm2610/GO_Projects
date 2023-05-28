from typing import List


class Solution:
    def minCost(self, n: int, cuts: List[int]) -> int:
        cuts.sort()
        dp = {}
        tl = len(cuts)

        def dfs(l, r):
            if r - l == 1:
                return 0
            if (l, r) in dp:
                return dp[(l, r)]

            cost = float("inf")
            for c in cuts:
                if c > l and c < r:
                    cost = min(cost, dfs(l, c) + dfs(c, r) + (r - l))

            cost = 0 if cost == float("inf") else cost
            dp[(l, r)] = cost

            return cost

        return dfs(0, n)


# target =
obj = Solution
nums = [1, 3, 4, 5]
n = 7

print(obj.minCost(obj, n, nums))
