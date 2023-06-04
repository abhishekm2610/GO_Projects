package main

import "fmt"

func setZeroes(matrix [][]int) [][]int {
	rowflag := false
	colflag := false
	m := len(matrix)
	n := len(matrix[0])
	for i, v := range matrix {
		for j, _ := range v {
			if matrix[i][j] == 0 {
				if i == 0 {
					rowflag = true
				}
				if j == 0 {
					colflag = true
				}
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := range matrix[i] {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
	if rowflag {
		for j := range matrix[0] {
			matrix[0][j] = 0
		}

	}
	if colflag {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

	return matrix

}

func main() {

	nums := [][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 0}}
	fmt.Println(setZeroes(nums))
}
