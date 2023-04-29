package main

import (
	"fmt"
)

func getSum(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}

func getSum2(a int, b int) int {
	carry := a & b
	ans := a ^ b
	for carry != 0 {
		ans = (carry << 1) ^ ans
		carry = ans & carry

	}
	return ans
}

func main() {
	num1 := 2
	num2 := 3
	fmt.Println(getSum(num1, num2))

}
