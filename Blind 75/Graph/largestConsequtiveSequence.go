package main

import (
	"fmt"
	"sort"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	ans := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 0 {
			continue
		}
		if nums[i]-nums[i-1] == 1 {
			count++
			ans = max(ans, count)
		} else {
			count = 1
		}
	}
	return ans
}
func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}
