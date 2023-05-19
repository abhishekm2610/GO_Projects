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
	// s1 := make(map[int]bool)
	// color := make(map[int]bool)
	visited := make(map[int]int)
	for i := 0; i < len(graph); i++ {
		visited[i] = -1
	}
	fmt.Println(visited)

	for i, _ := range graph {
		if !visited[i] {
			for _, edge := range graph[i] {
				nei := graph[edge]
				for _, n := range nei {
					if s2[n] {
						return false
					}
					if n != i {
						visited[i] = true

						s2[n] = true
					}
				}
			}
		}
		s2 = make(map[int]bool)
	}
	return true
}
func main() {
	nums := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	fmt.Println(isBipartite(nums))
}
