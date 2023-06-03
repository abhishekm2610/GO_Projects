from typing import List


class Solution:
    def numOfMinutes(
        self, n: int, headID: int, manager: List[int], informTime: List[int]
    ) -> int:
        managers = {}

        for i, v in enumerate(manager):
            if v not in managers:
                managers[v] = []
            managers[v].append(i)
        queue = managers[headID]
        ans = informTime[headID]
        while len(queue) != 0:
            print(queue)
            mgr = queue[0]
            queue = queue[1:]

            if mgr in managers:
                possible = []
                for v in managers[mgr]:
                    queue.append(v)
                    possible.append(informTime[v])
                ans += max(possible)

        return ans


# target =
obj = Solution
n = 22
headID = 7
manager = [
    12,
    7,
    18,
    11,
    13,
    21,
    12,
    -1,
    6,
    5,
    14,
    13,
    14,
    9,
    20,
    13,
    11,
    1,
    1,
    2,
    3,
    19,
]
informTime = [
    0,
    540,
    347,
    586,
    0,
    748,
    824,
    486,
    0,
    777,
    0,
    202,
    653,
    713,
    454,
    0,
    0,
    0,
    574,
    735,
    721,
    772,
]

print(obj.numOfMinutes(obj, n, headID, manager, informTime))
