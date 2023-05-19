package main

import (
	"fmt"
)

func max(nums []int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
func maxVal(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBipartite(graph [][]int) bool {
	visited := make(map[int]int)
	for i := 0; i < len(graph); i++ {
		visited[i] = -1
	}
	var dfs func(node int, prev int) bool
	dfs = func(node int, prev int) bool {
		var cur int
		if prev == 0 {
			cur = 1
		} else {
			cur = 0
		}
		for _, edge := range graph[node] {
			if visited[edge] == visited[node] {
				return false
			} else if visited[edge] == -1 {
				visited[edge] = cur
				if !dfs(edge, cur) {
					return false
				}
			}
		}
		return true
	}

	for i, _ := range graph {
		if visited[i] == -1 {
			visited[i] = 0
		}
		if !dfs(i, visited[i]) {
			return false
		}
	}
	return true
}
func main() {
	nums := [][]int{{1}, {0, 3}, {3}, {1, 2}}
	fmt.Println(isBipartite(nums))
}
