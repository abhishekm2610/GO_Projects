package main

import (
	"fmt"
)

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
func dfs(node int, visited map[int]bool, dp map[int][]int) bool {
	if len(dp[node]) == 0 {
		return false
	}
	if visited[node] {
		return false
	} else {
		for i, _ := range dp[node] {
			if !dfs(i, visited, dp) {
				return false
			}
		}
	}
	return true
}

func validTree(nodes int, edges [][]int) bool {
	dp := make(map[int][]int, nodes)

	for i := 0; i < nodes; i++ {
		dp[i] = []int{}
	}

	for _, v := range edges {
		dp[v[0]] = append(dp[v[0]], v[1])
	}

	for i, _ := range dp {

		visited := make(map[int]bool)
		visited[i] = true
		if !dfs(i, visited, dp) {
			return false
		}

	}
	fmt.Println(dp)
	return true
}

func main() {
	nums := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	target := 5
	fmt.Println(validTree(target, nums))
}
