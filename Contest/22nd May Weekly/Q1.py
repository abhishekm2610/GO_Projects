class Solution:
    def minLength(self, s: str) -> int:
        i = 0

        while ("AB" in s) or ("CD" in s):
            s = s.replace("AB", "")
            s = s.replace("CD", "")

        return len(s)


# target =
obj = Solution
s = "ABFCACDB"

print(obj.minLength(obj, s))
