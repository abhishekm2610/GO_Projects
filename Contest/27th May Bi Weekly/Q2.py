from typing import List


class Solution:
    def minExtraChar(self, s: str, dictionary: List[str]) -> int:
        dp = [-1] * (len(s) + 1)

        dp[0] = 0

        for i in range(1, len(s) + 1):
            dp[i] = dp[i - 1] + 1
            for j in range(i - 1, -1, -1):
                if s[j:i] in dictionary:
                    dp[i] = min(dp[i], dp[j])

        return dp[len(s)]


# target =
obj = Solution
s = "leetcode"
mon = [
    "na",
    "i",
    "edd",
    "wobow",
    "kecv",
    "b",
    "n",
    "or",
    "jj",
    "zul",
    "vk",
    "yeb",
    "qnfac",
    "azv",
    "grtjba",
    "yswmjn",
    "xowio",
    "u",
    "xi",
    "pcmatm",
    "maqnv",
]

print(obj.minExtraChar(obj, s, mon))
