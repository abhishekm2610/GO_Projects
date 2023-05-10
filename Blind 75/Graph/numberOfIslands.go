package main

import "fmt"

func dfs(i int, j int, rows int, cols int, visited map[[2]int]int, grid [][]byte) {
	if _, present := visited[[2]int{i, j}]; present || grid[i][j] == '0' {
		return
	}
	// i == rows || j == cols || j == 0 || i == 0 {
	visited[[2]int{i, j}] = 1

	if i != 0 {
		dfs(i-1, j, rows, cols, visited, grid)
	}
	if i != rows-1 {
		dfs(i+1, j, rows, cols, visited, grid)
	}
	if j != 0 {
		dfs(i, j-1, rows, cols, visited, grid)

	}
	if j != cols-1 {
		dfs(i, j+1, rows, cols, visited, grid)

	}
}
func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])

	visited := make(map[[2]int]int)
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if _, present := visited[[2]int{i, j}]; present {
				continue
			} else if grid[i][j] == '1' {
				// visited[[2]int{i, j}] = 1
				// fmt.Println(visited, i, j)

				count++
				dfs(i, j, rows, cols, visited, grid)
				// fmt.Println(i, j, visited)

			}
		}
	}
	return count
}

func main() {
	nums := [][]byte{{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(nums))
}
