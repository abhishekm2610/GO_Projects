from typing import List


class Solution:
    def removeTrailingZeros(self, num: str) -> str:
        while num[-1] == "0":
            num = num[: len(num) - 1]

        return num


# target =
obj = Solution
s = "1"
# mon =

print(obj.removeTrailingZeros(obj, s))
