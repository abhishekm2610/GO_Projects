package main

import (
	"fmt"
	"math"
)

func maxVal(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	max := 1
	for index, _ := range dp {
		// curr := nums[index]
		for i := index - 1; i >= 0; i-- {
			if nums[i] < nums[index] {
				dp[index] = maxVal(dp[i]+1, dp[index])
				max = int(math.Max(float64(max), float64(dp[index])))
			}
		}
		if dp[index] == 0 {
			dp[index] = 1
		}
	}
	fmt.Println(dp)
	return max

}

func binarySearch(lis []int, target int) int {
	pivot := len(lis) / 2
	if lis[pivot] == target {
		return pivot
	} else if lis[pivot] > target {
		binarySearch(lis[0:pivot], target)
	} else {
		binarySearch(lis[pivot+1:], target)
	}
	return -1
}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	res := make([]int, len(obstacles))
	dp := []int{}

	// rdp := make(map[int][]int)
	for i, v := range obstacles {
		indx := binarySearch(dp, v)
		if indx == -1 {
			dp = append(dp, v)
		} else {
			dp[indx] = v
		}
		res[i] = indx + 1
	}
	// for i, v := range dp {
	// 	res[i] = v
	// }
	return res
}

// Solution 1: LIS + Binary Search
func upperBound(LIS []int, x int) int {
	var low, high int = 0, len(LIS) - 1
	var ans int = high
	for low <= high {
		mid := low + (high-low)/2
		if x < LIS[mid] {
			if mid < ans {
				ans = mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}

func longestObstacleCourseAtEachPositionans(obstacles []int) []int {
	LIS := []int{}
	ans := make([]int, len(obstacles))
	for i := 0; i < len(obstacles); i++ {
		if len(LIS) == 0 || obstacles[i] >= LIS[len(LIS)-1] {
			LIS = append(LIS, obstacles[i])
			ans[i] = len(LIS)
		} else {
			ans[i] = upperBound(LIS, obstacles[i]) + 1
			LIS[ans[i]-1] = obstacles[i]
		}
	}
	return ans
}

func main() {
	// target := 4
	nums := []int{3, 1, 5, 6, 4, 2}
	fmt.Println(longestObstacleCourseAtEachPosition(nums))
}
