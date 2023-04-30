package main

import (
	"fmt"
)

func missingNumberUsingXOR(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = nums[i] ^ ans
	}
	curr := len(nums)
	for curr >= 0 {
		ans = curr ^ ans
		curr--
	}
	return ans
}

func missingNumber(nums []int) int {
	ans := len(nums)
	for i := 0; i < len(nums); i++ {
		ans += i - nums[i]
	}

	return ans
}
func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))

}
