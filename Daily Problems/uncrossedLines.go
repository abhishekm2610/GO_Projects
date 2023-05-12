package main

import (
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxUncrossedLines(text1 []int, text2 []int) int {
	dp := make([][]int, len(text1)+1)
	for v := range dp {
		dp[v] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]

}

//	func maxUncrossedLines(nums1 []int, nums2 []int) int {
//		left := 0
//		right := 0
//		ans := 0
//		for left < len(nums1) && right < len(nums2) {
//			if nums1[left] == nums2[right] && right >= left {
//				fmt.Println(left, right)
//				ans++
//				right++
//				left++
//			} else {
//				right++
//			}
//		}
//		return ans
//	}
func main() {
	nums1 := []int{2, 5, 1, 2, 5}
	nums2 := []int{10, 5, 2, 1, 5, 2}

	fmt.Println(maxUncrossedLines(nums1, nums2))
}
