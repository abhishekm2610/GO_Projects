package main

import (
	"fmt"
)

func countBitsv0(n int) []int {
	ans := []int{}
	for i := 0; i < n+1; i++ {
		tmp := i
		count := 0
		for tmp != 0 {
			if tmp%2 == 1 {
				count++
			}
			tmp /= 2
		}
		ans = append(ans, count)
	}
	return ans
}

func countBits(n int) []int {
	dp := make([]int, n+1)
	current2Power := 1
	dp[0] = 0
	for i := 1; i <= n; i++ {
		if 2*current2Power == i {
			current2Power = i

		}
		dp[i] = 1 + dp[i-current2Power]
	}
	return dp
}
func main() {

	num1 := 2
	fmt.Println(countBits(num1))

}
