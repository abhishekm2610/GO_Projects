package main

import (
	"fmt"
)

// func productExceptSelf(nums []int) []int {
// 	ans := make([]int, len(nums))
// 	prefixList := make([]int, len(nums))
// 	copy(prefixList, nums)
// 	suffixList := make([]int, len(nums))
// 	copy(suffixList, nums)

//		prod := 1
//		for i, _ := range prefixList {
//			prefixList[i] = prefixList[i] * prod
//			prod = prefixList[i]
//		}
//		prod = 1
//		for i := len(nums) - 1; i >= 0; i-- {
//			suffixList[i] = suffixList[i] * prod
//			prod = suffixList[i]
//		}
//		for i, _ := range nums {
//			preVal := 1
//			if i != 0 {
//				preVal = prefixList[i-1]
//			}
//			postVal := 1
//			if i+1 < len(nums) {
//				postVal = suffixList[i+1]
//			}
//			ans[i] = preVal * postVal
//		}
//		return ans
//	}
// func productExceptSelf(nums []int) []int {
// 	ans := make([]int, len(nums))
// 	for i := range ans {
// 		ans[i] = 1
// 	}
// 	ans[0] = nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		ans[i] = nums[i-1] * ans[i-1]
// 	}
// 	suffix := 1
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		ans[i] *= suffix
// 		suffix *= nums[i]
// 	}
// 	return ans

// }

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = 1
	}
	prefix := 1
	suffix := 1

	for i := 0; i < len(nums); i++ {
		ans[i] *= prefix
		prefix *= nums[i]
		ans[len(nums)-1-i] *= suffix
		suffix *= nums[len(nums)-1-i]
	}
	return ans

}
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))

}
