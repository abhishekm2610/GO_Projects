package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix) - 1

	for j := 0; j <= n/2; j++ {
		topLeftBorder := []int{j, j}
		topRightBorder := []int{j, n - j}
		bottomLeftBorder := []int{n - j, j}
		bottomRightBorder := []int{n - j, n - j}
		old, new := 0, 0
		// fmt.Println(matrix[topLeftBorder[0]][topLeftBorder[1]])
		// fmt.Println(matrix[topRightBorder[0]][topRightBorder[1]])

		// fmt.Println(matrix[bottomLeftBorder[0]][bottomLeftBorder[1]])

		// fmt.Println(matrix[bottomRightBorder[0]][bottomRightBorder[1]])

		for i := 0; i < n-2*j; i++ {

			new = matrix[topLeftBorder[0]][topLeftBorder[1]]

			x, y := topRightBorder[0], topRightBorder[1]
			old = matrix[x][y]
			matrix[x][y] = new
			new = old

			x, y = bottomRightBorder[0], bottomRightBorder[1]
			old = matrix[x][y]
			matrix[x][y] = new
			new = old

			x, y = bottomLeftBorder[0], bottomLeftBorder[1]
			old = matrix[x][y]
			matrix[x][y] = new
			new = old

			x, y = topLeftBorder[0], topLeftBorder[1]
			old = matrix[x][y]
			matrix[x][y] = new
			new = old

			topLeftBorder[1]++
			topRightBorder[0]++
			bottomLeftBorder[0]--
			bottomRightBorder[1]--

		}
	}

}

func main() {

	nums := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(nums)
	fmt.Println(nums)
}
