package main

import (
	"fmt"
	"math"
)

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	max := 1
	for index, _ := range dp {
		// curr := nums[index]
		for i := index - 1; i >= 0; i-- {
			if nums[i] < nums[index] {
				dp[index] = int(math.Max(float64(dp[i]+1), float64(dp[index])))
				max = int(math.Max(float64(max), float64(dp[index])))
			}
		}
		if dp[index] == 0 {
			dp[index] = 1
		}
	}
	fmt.Println(dp)
	return max

}
func main() {
	n := []int{0, 1, 0, 3, 2, 3}
	fmt.Println(lengthOfLIS(n))

}
