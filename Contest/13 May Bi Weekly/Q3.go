package main

import (
	"fmt"
	"math"
)

func maxInSlice(nums []int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
func minInSlice(nums []int) int {
	max := 0
	for _, v := range nums {
		if v < max {
			max = v
		}
	}
	return max
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//	func sumOfArray(arr []int) int {
//		ans := 0
//		for _, v := range arr {
//			ans = ans | v
//		}
//		return ans
//	}
func maximumOr(nums []int, k int) int64 {
	multipler := math.Pow(float64(2), float64(k))
	maxVal := 0
	preFix := make([]int, len(nums))
	sufFix := make([]int, len(nums))

	for i := 1; i < len(preFix); i++ {
		preFix[i] = preFix[i-1] | nums[i-1]
	}
	for i := len(sufFix) - 2; i >= 0; i-- {
		sufFix[i] = sufFix[i+1] | nums[i+1]
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println(preFix[i] | nums[i]*int(multipler) | sufFix[i])

		maxVal = max(maxVal, preFix[i]|nums[i]*int(multipler)|sufFix[i])
	}
	return int64(maxVal)
}

func main() {
	nums := []int{12, 9}
	target := 1
	fmt.Println(maximumOr(nums, target))
}
