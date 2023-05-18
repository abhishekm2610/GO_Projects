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

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
	minSoFar := intervals[0][1]
	ans := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= minSoFar {
			minSoFar = intervals[i][1]
		} else {
			if intervals[i][1] < minSoFar {
				minSoFar = intervals[i][1]
			}
			ans++
		}

	}
	return ans
}

func main() {
	dp := [][]int{{99, 100}, {1, 3}, {1, 2}, {2, 6}, {8, 10}, {12, 16}}
	// dp := [][]int{{1, 3}, {6, 9}}

	// newInterval := []int{4, 8}
	fmt.Println(eraseOverlapIntervals(dp))
}
