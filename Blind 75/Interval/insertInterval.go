package main

import "fmt"

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

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	} else if len(newInterval) == 0 {
		return intervals
	}
	ans := [][]int{}
	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			ans = append(ans, newInterval)
			for j := i; j < len(intervals); j++ {
				fmt.Println("oomb")

				ans = append(ans, intervals[j])
			}
			return ans
		} else if newInterval[0] > intervals[i][1] {
			ans = append(ans, intervals[i])
		} else {
			newInterval = []int{min(intervals[i][0], newInterval[0]), max(intervals[i][1], newInterval[1])}
		}
	}
	ans = append(ans, newInterval)
	return ans
}

func main() {
	dp := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	// dp := [][]int{{1, 3}, {6, 9}}

	newInterval := []int{4, 8}
	fmt.Println(insert(dp, newInterval))
}
