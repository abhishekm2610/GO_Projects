package main

import (
	"fmt"
)

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i]-48 == 0 {
			dp[i] = 0
		} else {
			dp[i] = dp[i+1]
		}
		if i+1 < len(s) && (s[i]-48 == 1 || (s[i]-48 == 2 && s[i+1]-48 <= 6)) {
			dp[i] += dp[i+2]
		}
	}
	return dp[0]
}

func main() {
	// target := 4
	nums := "11106"
	fmt.Println(numDecodings(nums))

}
