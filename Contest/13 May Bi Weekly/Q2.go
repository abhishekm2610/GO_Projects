package main

import (
	"fmt"
)

func maxInSlice(nums []int) []int {
	max := -2
	index := 0
	for i, v := range nums {
		if v > max {
			max = v
			index = i
		}
	}
	return []int{index, max}
}
func minInSlice(nums []int) int {
	max := 0
	for _, v := range nums {
		if v < max {
			max = v
		}
	}
	return max
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func matrixSum(nums [][]int) int {

	rows := len(nums)
	// cols := len(nums[0])
	currMax := -2
	score := 0
	i := 0
	for i < rows {
		maxOfRow := maxInSlice(nums[i])
		nums[i][maxOfRow[0]] = -1
		if maxOfRow[1] > currMax {
			currMax = maxOfRow[1]
			nums[i][maxOfRow[0]] = -1

		}
		if i+1 == rows && currMax != -1 {
			score += currMax
			currMax = -2
			i = 0
		} else if currMax == -1 {
			break
		} else {
			i++

		}

	}
	return score
}

func main() {
	nums := [][]int{{0}}
	// target := 5
	fmt.Println(matrixSum(nums))
}
