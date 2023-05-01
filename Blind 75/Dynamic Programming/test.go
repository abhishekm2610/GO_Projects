package main

import "fmt"

func main() {
	test1 := []int{}
	test2 := []int{1}

	fmt.Println(append(test2, test1...))
}
