package main

import (
	"fmt"
	"math"
	"sort"
)

func maxInSlice(nums []int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
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

func sumOfPower(nums []int) int {
	sort.Ints(nums)
	res := 0
	pre := 0
	mod := int(math.Pow(float64(10), float64(9))) + 7
	for _, x := range nums {
		res = (res + (x*x%mod)*x%mod + (x*x%mod)*pre%mod) % mod
		pre = (pre*2 + x) % mod
	}

	return res
}

func main() {
	nums := []int{658, 489, 777, 2418, 1893, 130, 2448, 178, 1128, 2149, 1059, 1495, 1166, 608, 2006, 713, 1906, 2108, 680, 1348, 860, 1620, 146, 2447, 1895, 1083, 1465, 2351, 1359, 1187, 906, 533, 1943, 1814, 1808, 2065, 1744, 254, 1988, 1889, 1206}
	// target := 5
	// nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sumOfPower(nums))
}

// 567530242
