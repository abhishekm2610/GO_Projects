package main

import (
	"fmt"
)

func max(nums []int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
func maxVal(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func mostPoints(questions [][]int) int64 {
	dp := make([]int, len(questions)+1)

	totalQ := len(questions)
	ans := 0
	for i := totalQ - 1; i >= 0; i-- {
		if i+questions[i][1] >= totalQ-1 {
			dp[i] = maxVal(questions[i][0], dp[i+1])
			ans = maxVal(dp[i], ans)

			continue
		} else {
			// fmt.Println(i, questions[i][1], dp[i+questions[i][1]])
			dp[i] = maxVal(dp[i+questions[i][1]+1]+questions[i][0], dp[i+1])
			ans = maxVal(dp[i], ans)
		}
	}
	// fmt.Println(dp)

	return int64(ans)
}
func main() {
	nums := [][]int{{21, 5}, {92, 3}, {74, 2}, {39, 4}, {58, 2}, {5, 5}, {49, 4}, {65, 3}}
	fmt.Println(mostPoints(nums))
}
