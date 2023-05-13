package main

import (
	"fmt"
)

func maxInSlice(nums []int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
func minInSlice(nums []int) int {
	max := 0
	for _, v := range nums {
		if v < max {
			max = v
		}
	}
	return max
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func countSeniors(details []string) int {
	count := 0
	for _, s := range details {
		age := s[11:13]
		fmt.Println(age)
		if age[0] == '7' || age[0] == '8' || age[0] == '9' {
			count++
		} else if age[0] == '6' && age[1] != '0' {
			count++
		}

	}
	return count
}

func main() {
	nums := []string{"5612624052M0130", "5378802576M6424", "5447619845F0171", "2941701174O9078"}
	// target := 5
	fmt.Println(countSeniors(nums))
}
