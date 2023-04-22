package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	visited := make(map[int]int)

	for index, value := range nums {
		if _, isPresent := visited[value]; isPresent {
			return true
		} else {
			visited[value] = index
		}
	}
	return false
}

func main() {
	nums := []int{7, 11, 11}
	fmt.Println(containsDuplicate(nums))

}
