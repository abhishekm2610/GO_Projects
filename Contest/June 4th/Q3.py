from typing import List


class Solution:
    def minimizedStringLength(self, s: str) -> int:
        ans = set()

        for i in s:
            ans.add(i)

        return len(ans)


# target =
obj = Solution
nums = "dddaaa"

print(obj.minimizedStringLength(obj, nums))
