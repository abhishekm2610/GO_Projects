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
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{}
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			if intervals[i][1] >= end {
				end = intervals[i][1]

			}
		} else {
			ans = append(ans, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
	}
	ans = append(ans, []int{start, end})

	return ans

}

func main() {
	dp := [][]int{{99, 100}, {1, 3}, {2, 6}, {8, 10}, {12, 16}}
	// dp := [][]int{{1, 3}, {6, 9}}

	// newInterval := []int{4, 8}
	fmt.Println(merge(dp))
}
