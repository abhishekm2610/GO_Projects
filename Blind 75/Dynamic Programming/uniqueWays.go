package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for v := range dp {
		dp[v] = make([]int, n+1)
	}
	dp[0][1] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m][n]
}

func main() {
	// target := 4
	m := 1
	n := 1
	fmt.Println(uniquePaths(m, n))

}
