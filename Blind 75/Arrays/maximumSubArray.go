package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	runningSum := 0
	sum := nums[0]
	value := 0
	for value < len(nums) {
		if runningSum < 0 {
			runningSum = 0
		}
		runningSum += nums[value]
		sum = int(math.Max(float64(sum), float64(runningSum)))
		value++
	}
	return sum
}

// func maxSubArray(nums []int) int {
// 	currSum := 0
// 	maxSum := -10

// 	for i := 0; i < len(nums); i++ {
// 		if currSum+nums[i] > currSum {
// 			currSum += nums[i]
// 		} else {
// 			currSum = nums[i]
// 		}
// 		if currSum > maxSum {
// 			maxSum = currSum
// 		}
// 	}
// 	return maxSum
// }

// func maxSubArray(nums []int) int {
// 	max := nums[0]
// 	runningSum := nums[0]
// 	maxList := []int{}
// 	for _, v := range nums {
// 		if v > max {
// 			max = v
// 		}
// 		runningSum += v
// 		if runningSum > max {
// 			maxList = append(maxList, runningSum)
// 			max =
// 		}

// 	}
// 	return max
// }

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))

}
