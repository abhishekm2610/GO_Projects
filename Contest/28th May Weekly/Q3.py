from typing import List


class Solution:
    def minimumCost(self, s: str) -> int:
        def flip(starting, end, string):
            ans = string[:starting]

            for i in range(starting, end):
                if string[i] == "1":
                    ans += "0"
                else:
                    ans += "1"

            return ans

        n = len(s)
        startVal = s[0]
        cost = 0
        for i in range(0, n):
            if s[i] != startVal:
                print(i, s[i])
                s = flip(i, n, s)
                cost += n - i
            print(s)

        return cost


# target =
obj = Solution
nums = "010101"

print(obj.minimumCost(obj, nums))
