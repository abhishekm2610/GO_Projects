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

func dfs(current int, dp map[int]int) int {
	if current == 0 {
		return 1
	} else if current < 0 {
		return 0
	} else if _, isPresent := dp[current]; isPresent {
		return dp[current]
	}
	dp[current] = dfs(current-1, dp) + dfs(current-2, dp)
	return dp[current]
}
func climbStairs(n int) int {

	dp := make(map[int]int)

	return dfs(n-1, dp) + dfs(n-2, dp)
}

func main() {
	n := 4
	fmt.Println(climbStairs(n))

}
