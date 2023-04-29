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
func countBits(n int) []int {
	ans := []int{}
	for i := 0; i < n+1; i++ {
		tmp := i
		count := 0
		for tmp != 0 {
			if tmp%2 == 1 {
				count++
			}
			tmp /= 2
		}
		ans = append(ans, count)
	}
	return ans
}
func main() {

	num1 := 2
	fmt.Println(countBits(num1))

}
