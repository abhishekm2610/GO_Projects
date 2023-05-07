package main

import (
	"fmt"
	"math"
	"sort"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	count := 0
	r := len(nums) - 1
	for i, left := range nums {
		for i <= r && (left+nums[r] > target) {
			r--
		}
		if i <= r {
			count += int(math.Pow(float64(2), float64(r-i)))
			count %= ((int(math.Pow(float64(10), float64(9)))) + 7)
		}
	}
	return count % ((int(math.Pow(float64(10), float64(9)))) + 7)
}
func main() {
	// target := 4
	nums := []int{27, 21, 14, 2, 15, 1, 19, 8, 12, 24, 21, 8, 12, 10, 11, 30, 15, 18, 28, 14, 26, 9, 2, 24, 23, 11, 7, 12, 9, 17, 30, 9, 28, 2, 14, 22, 19, 19, 27, 6, 15, 12, 29, 2, 30, 11, 20, 30, 21, 20, 2, 22, 6, 14, 13, 19, 21, 10, 18, 30, 2, 20, 28, 22}
	target := 31
	fmt.Println(numSubseq(nums, target))

}
