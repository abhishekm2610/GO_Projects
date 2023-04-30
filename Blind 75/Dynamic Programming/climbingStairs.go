package main

import (
	"fmt"
)

func climbStairsdP(n int) int {

	left := 1
	right := 1
	n -= 2
	for n >= 0 {
		tmp := left + right
		right = left
		left = tmp
		n--
	}
	return left
}

func dfs(current int, length int, dp map[int]int) int {
	if current == length {
		return 1
	} else if current > length {
		return 0
	} else if _, isPresent := dp[current]; isPresent {
		return dp[current]
	}
	dp[current] = dfs(current+1, length, dp) + dfs(current+2, length, dp)
	return dp[current]
}
func climbStairs(n int) int {

	ans := 0
	dp := make(map[int]int)
	ans = dfs(1, n, dp) + dfs(2, n, dp)

	return ans
}

func main() {
	n := 2
	fmt.Println(climbStairs(n))

}
