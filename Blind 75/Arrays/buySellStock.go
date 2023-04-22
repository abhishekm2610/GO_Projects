package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	leftPointer := 0
	rightPointer := 1
	res := 0
	for rightPointer < len(prices) {
		if prices[rightPointer] < prices[leftPointer] {
			leftPointer = rightPointer
			rightPointer++
		} else {
			res = int(math.Max(float64(res), float64(prices[rightPointer]-prices[leftPointer])))
			rightPointer++
		}
	}
	return res
}

func main() {
	nums := []int{7, 11}
	fmt.Println(maxProfit(nums))

}
