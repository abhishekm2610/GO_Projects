package main

import (
	"fmt"
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

func maximumOr(nums []int, k int) int64 {

}

func main() {
	nums := []int{12, 9}
	target := 1
	fmt.Println(maximumOr(nums, target))
}
