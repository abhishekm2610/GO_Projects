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
	fmt.Println(nums)
	// multipler := math.Pow(float64(2), float64(k))
	power := 0
	l := 0
	r := 0
	for l < len(nums) {
		if r == len(nums) {
			l++
			r = len(nums) - 1
			continue
		}
		minPower := nums[l]
		maxPower := nums[r]
		chars := int(math.Pow(float64(2), float64(r-l-1)))
		if l == r {
			power += (maxPower * maxPower * minPower)
			power %= ((int(math.Pow(float64(10), float64(9)))) + 7)

		} else {
			power += ((maxPower * maxPower * minPower) * chars)
			power %= ((int(math.Pow(float64(10), float64(9)))) + 7)

		}
		r++

	}
	return (power)
}

func main() {
	nums := []int{658, 489, 777, 2418, 1893, 130, 2448, 178, 1128, 2149, 1059, 1495, 1166, 608, 2006, 713, 1906, 2108, 680, 1348, 860, 1620, 146, 2447, 1895, 1083, 1465, 2351, 1359, 1187, 906, 533, 1943, 1814, 1808, 2065, 1744, 254, 1988, 1889, 1206}
	// target := 5
	fmt.Println(sumOfPower(nums))
}

// 567530242
