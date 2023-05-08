package main

import (
	"fmt"
)

func maxVal(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func diagonalSum(mat [][]int) int {
	m := len(mat)
	sum := 0
	for i := 0; i < m; i++ {

		sum += mat[i][m-1-i] + mat[i][i]

	}
	if m%2 == 0 {
		sum -= mat[m/2][m/2]
	}
	return sum
}

func main() {
	// target := 4
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(diagonalSum(nums))
}
