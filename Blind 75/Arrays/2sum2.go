package main

import (
	"fmt"
)

// //Hash Map
// func twoSum(nums []int, target int) []int {
// 	hashMap := make(map[int]int)

// 	for i := 0; i < len(nums); i++ {
// 		if _, isPresent := hashMap[target-nums[i]]; isPresent {
// 			return []int{hashMap[target-nums[i]] + 1, i + 1}
// 		} else {
// 			hashMap[nums[i]] = i
// 		}
// 	}
// 	return []int{}
// }

func twoSum(nums []int, target int) []int {
	i := 0
	j := len(nums) - 1
	for i < len(nums) {
		if nums[i]+nums[j] == target {
			break
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{i + 1, j + 1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))

}
