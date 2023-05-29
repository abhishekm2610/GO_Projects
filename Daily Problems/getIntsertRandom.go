package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) []int {
	n := len(nums) / 3
	// n++
	sort.Ints(nums)
	fmt.Println(n, nums)
	ans := []int{}

	for i := 0; i <= len(nums)-n-1; i++ {
		prev := nums[i]
		fmt.Println(i, prev)
		if nums[i+n] == prev {
			ans = append(ans, prev)
			for j := i + n; i < len(nums)-n-1; j++ {
				if nums[j] != prev {
					prev = nums[j]
					i = j
					break
				}
			}
		}

	}
	return ans
}

func main() {
	nums := []int{1, 2}

	fmt.Println(majorityElement(nums))
}
