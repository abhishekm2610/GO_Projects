package main

import "fmt"

func spiralOrder(matrix [][]int) []int {

	rightBorder := len(matrix[0]) - 1
	leftBorder := 0
	upperBorder := 0
	lowerBorder := len(matrix) - 1
	ans := []int{}
	for true {

		for i := leftBorder; i <= rightBorder; i++ {
			ans = append(ans, matrix[upperBorder][i])
		}
		upperBorder++
		if upperBorder > lowerBorder {
			break
		}

		for i := upperBorder; i <= lowerBorder; i++ {
			ans = append(ans, matrix[i][rightBorder])
		}
		rightBorder--
		if rightBorder < leftBorder {
			break
		}
		for i := rightBorder; i >= leftBorder; i-- {
			ans = append(ans, matrix[lowerBorder][i])
		}
		lowerBorder--
		if lowerBorder < upperBorder {
			break
		}
		for i := lowerBorder; i >= upperBorder; i-- {
			ans = append(ans, matrix[i][leftBorder])

		}
		leftBorder++
		if leftBorder > rightBorder {
			break
		}

	}
	return ans

}

func main() {

	nums := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(nums))
}
