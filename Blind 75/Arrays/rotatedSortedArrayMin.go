package main

import (
	"fmt"
	"math"
)

func findMin(nums []int) int {

	left := 0
	right := len(nums) - 1
	min := nums[0]
	for left <= right {

		if nums[left] < nums[right] {
			min = int(math.Min(float64(nums[left]), float64(min)))
			break
		}

		pivot := (left + right) / 2
		if nums[pivot] < min {
			min = nums[pivot]
		}

		if nums[pivot] >= nums[left] {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}
	return min
}

func main() {
	nums := []int{11, 12, 13, 1}
	fmt.Println(findMin(nums))

}
