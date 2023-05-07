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
func checkAdjacent(nums []int) int {
	curr := nums[0]
	count := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] == 0 {
			curr = nums[j]
			continue
		}
		if nums[j] == curr {
			count++
		} else {
			curr = nums[j]
		}
	}
	return count
}

func colorTheArray(n int, queries [][]int) []int {
	ans := []int{}
	colorArray := make([]int, n)
	adj := 0
	for _, q := range queries {
		currentColor := colorArray[q[0]]
		colorArray[q[0]] = q[1]
		if q[0] < n-1 && colorArray[q[0]+1] == colorArray[q[0]] {
			adj++
		}
		if q[0] >= 1 && colorArray[q[0]-1] == colorArray[q[0]] {

			adj++
		}
		if q[0] < n-1 && currentColor != 0 && colorArray[q[0]+1] == currentColor {
			adj--
		} else if q[0] >= 1 && currentColor != 0 && colorArray[q[0]-1] == currentColor {
			adj--
		}
		ans = append(ans, adj)
	}
	return ans
}
func main() {
	// target := 4
	nums := 4
	queries := [][]int{{0, 100000}}

	fmt.Println(colorTheArray(nums, queries))
}
