package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

func hammingWeightv1(num uint32) int {
	ans := 0
	// var mask uint32
	// mask = 00000000000000000000000000000001
	for num != 0 {
		ans += int(num % 2)
		num = num >> 1
	}
	return ans
}
func main() {
	var num1 uint32
	num1 = 00000000000000000000000000001011
	fmt.Println(hammingWeight(num1))

}
