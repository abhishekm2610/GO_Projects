from typing import List


class Solution:
    def buyChoco(self, prices: List[int], money: int) -> int:
        prices.sort()

        if prices[0] + prices[1] <= money:
            return money - prices[0] - prices[1]

        return money


# target =
obj = Solution
s = list([1, 1, 1])
mon = 3

print(obj.buyChoco(obj, s, mon))
