package main

import "fmt"

// func twoSum(nums []int, target int) []int {
// 	left := 0
// 	right := 1
// 	for left < len(nums)-1 {
// 		right := left + 1
// 		for right < len(nums) {
// 			if nums[left]+nums[right] == target {
// 				return ([]int{left, right})

// 			} else {
// 				right++
// 			}
// 		}
// 		left++
// 	}
// 	return ([]int{left, right})

// }

func twoSum(nums []int, target int) []int {
	visited := make(map[int]int)

	for index, value := range nums {
		if requiredIndex, isPresent := visited[target-value]; isPresent {
			return []int{requiredIndex, index}
		} else {
			visited[value] = index
		}
	}
	return []int{}
}

func main() {
	target := 9
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, target))
}
