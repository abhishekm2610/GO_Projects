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

func canJump(nums []int) bool {
	maxReached := 0
	for i, _ := range nums {
		fmt.Println(i+nums[i], maxReached)
		if i <= maxReached {
			maxReached = max(i+nums[i], maxReached)

		}
	}
	if maxReached >= len(nums)-1 {
		return true
	}

	return false
}

func main() {
	// target := 4
	nums := []int{0}
	fmt.Println(canJump(nums))

}
