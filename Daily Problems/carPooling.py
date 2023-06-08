from typing import List


class Solution:
    def carPooling(self, trips: List[List[int]], capacity: int) -> bool:
        # onboarded = {}
        # currentCap = 0
        # for trip in trips:
        #     passengers = trip[0]
        #     starting = trip[1]
        #     end = trip[2]

        #     for
        #     for i in range(trip[1], trip[2] + 1):
        #         if i not in onboarded:
        #             onboarded[i] = 0
        #         if onboarded[i] + trip[0] > capacity:
        #             return False
        #         else:
        #             onboarded[i] += trip[0]

        # return True

        # trips.sort(key=lambda x: x[1])
        off_boarding_map = {}  # 5: 2
        current_capacity = 0
        for trip in trips:
            num_passengers = trip[0]  # 2, 3
            start_station = trip[1]  # 1, 3
            end_station = trip[2]  # 5, 7

            for station in off_boarding_map:
                if station <= start_station:  # 5 < 3, 2 < 3
                    current_capacity -= off_boarding_map[station]
                    off_boarding_map[station] = 0

            if current_capacity + num_passengers > capacity:  # 0 + 2 !> 4
                return False
            else:
                current_capacity += num_passengers  # CC = 2
                if end_station not in off_boarding_map:
                    off_boarding_map[end_station] = 0
                off_boarding_map[end_station] += num_passengers  # 5: 2

            print(off_boarding_map)

        return True


# target =
obj = Solution
nums = [[8, 2, 3], [4, 1, 3], [1, 3, 6], [8, 4, 6], [4, 4, 8]]
n = 12
print(obj.carPooling(obj, nums, n))
