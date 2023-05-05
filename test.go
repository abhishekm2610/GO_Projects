package main

import (
	"fmt"
	"sort"
)

func main() {
	test1 := []int{}
	test2 := []int{1}

	fmt.Println(append(test2, test1...))
}

func findDifference(nums1 []int, nums2 []int) [][]int {

	sort.Ints(nums1)
	sort.Ints(nums2)
	ans := make([][]int, 2)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				continue
			}
			if nums2[j] > nums1[i] {
				ans[0] = append(ans[0], nums1[i])
			} else {
				ans[1] = append(ans[1], nums2[j])
			}
		}
	}
	return ans
}
