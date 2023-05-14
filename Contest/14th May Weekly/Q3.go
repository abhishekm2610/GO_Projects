package main

import "fmt"

func dfs(course int, visited map[int]bool, coursesMapping map[int][]int) bool {
	if len(coursesMapping[course]) == 0 {
		return true
	}
	if _, present := visited[course]; present {
		return false
	}
	visited[course] = true
	for _, dependancy := range coursesMapping[course] {
		return dfs(dependancy, visited, coursesMapping)
	}

	return true

}

func maxMoves(grid [][]int) int {
	dp := make([][]int, len(grid)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(grid[0])+1)
	}

	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols-1; j++ {
			// if grid[i][j+1] > grid[i][j]{

			// }
			if grid[i][j+1] > grid[i][j] {
				dp[i][j]++
			}
			if i < rows-1 && grid[i+1][j+1] > grid[i][j] {
				dp[i][j]++
			}
			if i > 0 && grid[i-1][j+1] > grid[i][j] {
				dp[i][j]++
			}
			// dp[i][j] = dp[i+1][j+1] + dp[i-1][j+1]
			// if i == rows-1 && j == cols-1 {
			// 	dp[i][j] = 0
			// 	continue
			// }
		}

	}
	fmt.Println(dp)
	return 0
}
func main() {
	// target := 10
	nums := [][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}
	fmt.Println(maxMoves(nums))
}
