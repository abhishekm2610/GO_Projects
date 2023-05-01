package main

import (
	"fmt"
	"math"
)

// func dfs(n int, coins []int, dp map[int]int) int {
// 	if n < 0 {
// 		return 0
// 	} else if n == 0 {
// 		return 1
// 	} else {
// 		if _, isPresent := dp[n]; isPresent {
// 			return dp[n]
// 		}

//		}
//		for _, coin := range coins {
//			tmp := dfs(coin, coins, dp)
//			if tmp < dp[n] {
//				dp[n] = tmp
//			}
//		}
//		return dfs(n, coins, dp)
//	}
func coinChange(coins []int, amount int) int {
	dp := make(map[int]int)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i-c >= 0 {
				dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-c])))
			}
		}
	}
	if dp[amount] != amount+1 {
		return dp[amount]
	} else {
		return -1
	}

}
func main() {
	n := []int{1, 2, 5}
	target := 11
	fmt.Println(coinChange(n, target))

}
