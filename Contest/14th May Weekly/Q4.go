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

func dfsv2(node int, visited map[int]bool, adj map[int][]int) {
	if visited[node] {
		return
	}
	visited[node] = true
	for _, ad := range adj[node] {
		visited[ad] = true
		dfs(ad, visited, adj)
	}
	return

}
func countCompleteComponents(n int, edges [][]int) int {
	adj := make(map[int][]int)
	for i := 0; i < n; i++ {
		adj[i] = []int{}
	}
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	visited := make(map[int]bool)
	count := 0
	for i, _ := range adj {
		if !visited[i] {
			count++
			dfsv2(i, visited, adj)
		}
	}
	return count
}

func main() {
	// target := 10
	nums := [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}
	n := 6
	fmt.Println(countCompleteComponents(n, nums))
}
