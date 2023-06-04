from typing import List


class Solution:
    def semiOrderedPermutation(self, nums: List[int]) -> int:
        idx1 = nums.index(1)
        idxN = nums.index(len(nums))
        ans = 0
        if idx1 <= idxN:
            print(11)
            ans += idx1
            ans += len(nums) - idxN - 1

        else:
            print(idx1, idxN)
            ans += idx1 - 1
            ans += len(nums) - idxN - 1

        return ans


# target =
obj = Solution
s = [1]
# mon =

print(obj.semiOrderedPermutation(obj, s))
