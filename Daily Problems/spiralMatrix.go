package main

import "fmt"

func pacificAtlantic(heights [][]int) [][]int {
	ans := [][]int{}
	dp := make([][]bool, len(heights))
	dp2 := make([][]bool, len(heights))

	for i, v := range heights {
		dp2[i], dp[i] = make([]bool, len(v)), make([]bool, len(v))
	}

	m := len(heights)
	n := len(heights[0])
	dp[0][0] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 || i == 0 {
				dp[i][j] = true
			}
			if (i < m-1) && dp[i][j] && heights[i+1][j] >= heights[i][j] {
				dp[i+1][j] = true

			}
			if (j < n-1) && dp[i][j] && heights[i][j+1] >= heights[i][j] {

				dp[i][j+1] = true

			}
			if i > 0 && dp[i][j] && (heights[i-1][j] >= heights[i][j]) {

				dp[i-1][j] = true

			}
			if j > 0 && dp[i][j] && (heights[i][j-1] >= heights[i][j]) {

				dp[i][j-1] = true

			}
		}
	}
	dp2[m-1][n-1] = true
	fmt.Println(dp2)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if j == n-1 || i == m-1 {
				dp2[i][j] = true
			}
			if i > 0 && dp2[i][j] && (heights[i-1][j] >= heights[i][j]) {

				dp2[i-1][j] = true

			}
			if j > 0 && dp2[i][j] && (heights[i][j-1] >= heights[i][j]) {

				dp2[i][j-1] = true

			}
			if (i < m-1) && dp2[i][j] && heights[i+1][j] >= heights[i][j] {
				dp2[i+1][j] = true

			}
			if (j < n-1) && dp2[i][j] && heights[i][j+1] >= heights[i][j] {

				dp2[i][j+1] = true

			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if dp2[i][j] && dp[i][j] {
				ans = append(ans, []int{i, j})
			}

		}
	}
	fmt.Println(dp)

	fmt.Println(dp2)

	return ans

}

func spiralOrder(matrix [][]int) []int {
	leftBorder := 0
	rightBorder := len(matrix[0]) - 1
	upperBorder := 0
	lowerBorder := len(matrix) - 1
	ans := []int{}
	for {
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
		if leftBorder > rightBorder {
			break
		}
		for i := rightBorder; i >= leftBorder; i-- {
			ans = append(ans, matrix[lowerBorder][i])
		}
		lowerBorder--
		if upperBorder > lowerBorder {
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
	// target := 10
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(nums))
}
