class Solution:
    def dfs(self, numbers, target, runningSum, curr, partial):
        if curr == len(numbers) - 1:
            if target == runningSum:
                print(numbers, partial)

                if numbers == partial:
                    return True
                else:
                    return False
        else:
            partial.append(numbers[curr + 1])
            print(numbers, partial)

            self.dfs(
                self, numbers, target, runningSum + numbers[curr + 1], curr + 1, partial
            )

    def punishmentNumber(self, n: int) -> int:
        ans = 0

        for i in range(1, n + 1):
            square = i * i

            substr = list(int(i) for i in str(square))

            for j in range(0, len(substr)):
                if self.dfs(self, substr, i, sum(substr[: j + 1]), j, substr[: j + 1]):
                    ans += square
                    break

            # if i % 10 == 0 or i % 100 == 0:
            #     ans += square

            # print(i, sum(substr))

            # if sum(substr) == i:
            #     ans += square

            # i
            # # j = 0
            # for j in range(1, len(substr)):
            #     if int(substr[:j]) + int(substr[j:]) == i:
            #         break

        return ans


# target =
obj = Solution
s = 10

print(obj.punishmentNumber(obj, s))
