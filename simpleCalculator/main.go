package main

import "fmt"

func addition(op1 float64, op2 float64) float64 {
	return op1 + op2
}

func subtraction(op1 float64, op2 float64) float64 {
	return op1 - op2
}

func multiply(op1 float64, op2 float64) float64 {
	return op1 * op2
}
func divide(op1 float64, op2 float64) float64 {
	return op1 / op2
}
func main() {
	op1 := 10.00
	var op2 float64 = 0

	fmt.Println(addition(op1, op2))
	fmt.Println(subtraction(op1, op2))
	fmt.Println(multiply(op1, op2))
	fmt.Println(divide(op1, op2))
}
