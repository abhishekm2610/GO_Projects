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
func countUnique(nums []int) int {
	visited := make(map[int]int)
	count := 0
	for i := 0; i < len(nums); i++ {
		if _, present := visited[nums[i]]; !present {
			count++
			visited[nums[i]] = i
		}
	}
	return count
}
func distinctDifferenceArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		prefix := []int{}
		suffix := []int{}
		prefix = nums[0 : i+1]
		suffix = nums[i+1:]
		ans[i] = countUnique(prefix) - countUnique(suffix)
	}
	return ans
}
func main() {
	// target := 4
	nums := []int{3}
	fmt.Println(distinctDifferenceArray(nums))
}
