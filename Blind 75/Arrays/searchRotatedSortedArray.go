package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		pivot := (left + right) / 2
		if nums[pivot] == target {
			return pivot
		}

		if nums[pivot] >= nums[left] {
			if target <= nums[pivot] && target >= nums[left] {
				right = pivot - 1
			} else {
				left = pivot + 1
			}
		} else {
			if target >= nums[pivot] && target <= nums[right] {
				left = pivot + 1

			} else {
				right = pivot - 1

			}

		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))

}
