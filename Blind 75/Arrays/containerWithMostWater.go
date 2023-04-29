package main

import (
	"fmt"
	"math"
)

func maxArea(nums []int) int {
	// length := len(nums) - 1
	area := 0
	l, r := 0, len(nums)-1
	for l <= r {
		fmt.Println(l, r)
		currLength := r - l
		currMinHeight := int(math.Min(float64(nums[l]), float64(nums[r])))
		area = int(math.Max(float64(area), float64(currLength*currMinHeight)))

		if nums[l] < nums[r] {
			l++
		} else {
			r--
		}
	}
	return area
}

func main() {
	nums := []int{2, 3, 4, 5, 18, 17, 6}
	fmt.Println(maxArea(nums))

}
