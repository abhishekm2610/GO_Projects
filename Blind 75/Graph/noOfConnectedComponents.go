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

func dfs(node int, prev int, visited map[int]bool, dp map[int][]int) {

	visited[node] = true
	for _, i := range dp[node] {
		if i != prev {
			dfs(i, node, visited, dp)
		}

	}

}

func validTree(nodes int, edges [][]int) int {
	dp := make(map[int][]int, nodes)

	for i := 0; i < nodes; i++ {
		dp[i] = []int{}
	}

	for _, v := range edges {
		dp[v[0]] = append(dp[v[0]], v[1])
		// dp[v[1]] = append(dp[v[0]], v[0])

	}

	visited := make(map[int]bool)
	count := 0
	for i := 0; i < nodes; i++ {
		fmt.Println(i, visited)
		if _, p := visited[i]; !p {

			count++

			dfs(i, -1, visited, dp)
		}
	}

	fmt.Println(dp)
	return count
}

func main() {
	nums := [][]int{{0, 1}, {1, 2}, {3, 4}}
	target := 2
	fmt.Println(validTree(target, nums))
}
