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

func dfs(node int, prev int, visited map[int]bool, dp map[int][]int) bool {
	if visited[node] {
		return false
	} else {
		visited[node] = true
		for _, i := range dp[node] {
			if i != prev {
				if !dfs(i, node, visited, dp) {
					// fmt.Println(dp[node], node, visited)
					return false
				}
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
		dp[v[1]] = append(dp[v[0]], v[0])

	}

	visited := make(map[int]bool)

	if !dfs(0, -1, visited, dp) {
		return false
	}

	fmt.Println(dp)
	return true
}

func main() {
	nums := [][]int{{0, 1}, {1, 2}, {1, 3}, {1, 4}, {2, 3}}
	target := 5
	fmt.Println(validTree(target, nums))
}
