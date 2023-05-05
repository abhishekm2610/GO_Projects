package main

import (
	"fmt"
	"math"
)

func robold(nums []int) int {
	left := 0
	right := 0
	for i := 0; i < len(nums); i++ {
		sum := int(math.Max(float64(nums[i]+left), float64(right)))
		left = right
		right = sum
	}
	return right
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return int(math.Max(float64(robold(nums[1:])), float64(robold(nums[:len(nums)-1]))))

}
func main() {
	// target := 4
	nums := []int{2, 3, 2}
	fmt.Println(rob(nums))

}
