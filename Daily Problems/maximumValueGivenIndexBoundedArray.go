package main

import (
	"fmt"
)

func max(nums []int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
func maxVal(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getSum(index int, value int, n int) int {
	count := 0

	if value > index {
		count += (value + value - index) * (index + 1) / 2

	} else {
		count += (value+1)*value/2 + index - value + 1

	}

	if value >= n-index {
		count += (value + value - n + 1 + index) * (n - index) / 2

	} else {
		count += (value+1)*value/2 + n - index - value
	}

	return count - value
}

func maxValue(n int, index int, maxSum int) int {

	left, right := 1, maxSum
	for left < right {
		mid := (left + right + 1) / 2
		if getSum(index, mid, n) <= maxSum {
			left = mid

		} else {
			right = mid - 1

		}
	}

	return left
}

func main() {
	a, b, c := 4, 2, 6
	fmt.Println(maxValue(a, b, c))
}
