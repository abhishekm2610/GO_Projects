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

func doesValidArrayExist(derived []int) bool {
	original := []int{}

	for i := 0; i < len(derived); i++ {

	}
}
func main() {
	nums := []int{1, 1, 0}
	// target := 5
	// k := 4
	fmt.Println(doesValidArrayExist(nums))
}
