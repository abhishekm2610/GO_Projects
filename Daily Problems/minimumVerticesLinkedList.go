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

func find(node int, parent []int) int {
	res := node

	for res != parent[res] {
		parent[res] = parent[parent[res]]
		res = parent[res]
	}
	return res
}

func union(e1 int, e2 int, parent []int, rank []int) {

	p1, p2 := find(e1, parent), find(e2, parent)

	if p1 == p2 {
		return
	}

	if rank[p1] > rank[p2] {
		parent[p2] = p1
		rank[p1] += rank[p2]
	} else {
		parent[p1] = p2
		rank[p2] += rank[p1]
	}

	return

}
func findSmallestSetOfVertices(n int, edges [][]int) []int {

	dp := make(map[int][]int)

	for i := 0; i < n; i++ {
		dp[i] = []int{}
	}
	for _, e := range edges {
		dp[e[1]] = append(dp[e[1]], e[0])
	}
	ans := []int{}
	for i, _ := range dp {
		if len(dp[i]) == 0 {
			ans = append(ans, i)
		}
	}

	return ans

}

func main() {
	nums := [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}
	target := 6
	fmt.Println(findSmallestSetOfVertices(target, nums))
}
