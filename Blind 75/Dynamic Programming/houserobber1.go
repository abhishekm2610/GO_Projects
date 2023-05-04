package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {
	left := 0
	right := 0

	for i := 0; i < len(nums); i++ {
		sum := int(math.Max(float64(nums[i]+left), float64(right)))
		left = right
		right = sum
	}
	return right

}
func main() {
	// target := 4
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))

}
