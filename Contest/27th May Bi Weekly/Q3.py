from typing import List


class Solution:
    def maxStrength(self, nums: List[int]) -> int:
        # combs = 2 ** (len(nums) - 1)
        n = len(nums)
        maxVal = float("-inf")
        for i in range(2**n):
            binary = bin(i)[2:].zfill(n)
            # subsets = [nums[j] for j in range(n) if binary[j] == '1']

            ans = 1
            visit = False
            for j in range(n):
                if binary[j] == "1":
                    visit = True
                    ans *= nums[j]
            if visit == True:
                maxVal = max(maxVal, ans)

        return maxVal


# target =
obj = Solution
nums = [0]

print(obj.maxStrength(obj, nums))
