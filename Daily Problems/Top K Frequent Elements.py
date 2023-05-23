from typing import List


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        dp = {}
        rdp = {}
        for i in range(len(nums)):
            if nums[i] not in dp:
                dp[nums[i]] = 1
            else:
                dp[nums[i]] += 1
            if dp[nums[i]] not in rdp:
                rdp[dp[nums[i]]] = []
            rdp[dp[nums[i]]].append(nums[i])
        ans = []
        rdp = dict(sorted(rdp.items(), reverse=True))
        for i in rdp:
            ans.extend(rdp[i])
            ans = list(dict.fromkeys(ans))
            if len(ans) >= k:
                break

        return ans


# target =
obj = Solution
s = [1, 2]
k = 2


print(obj.topKFrequent(obj, s, k))
