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
	if parent[res] != res {
		parent[res] = parent[parent[res]]
		res = parent[res]
	}
	return res
}

func union(n1 int, n2 int, parent []int, rank []int) int {

	p1, p2 := find(n1, parent), find(n2, parent)

	if p1 == p2 {
		return 0
	}

	if rank[p1] > rank[p2] {
		parent[p2] = p1
		rank[p1] += rank[p2]
	} else {
		parent[p1] = p2
		rank[p2] += rank[p1]

	}
	return 1
}
func validTree(nodes int, edges [][]int) int {
	parent := make([]int, nodes)
	for i, _ := range parent {
		parent[i] = i
	}

	rank := make([]int, nodes)
	for i, _ := range rank {
		rank[i] = 1
	}

	// dp := make(map[int][]int, nodes)

	// for i := 0; i < nodes; i++ {
	// 	dp[i] = []int{}
	// }
	ans := nodes
	for _, v := range edges {
		fmt.Println(v)
		ans -= union(v[0], v[1], parent, rank)

	}
	return ans
}

func findv2(node int, parent []int) int {

	res := node

	for parent[res] != res {
		parent[res] = parent[parent[res]]
		res = parent[res]
	}
	return res

}

func unionV2(n1 int, n2 int, parent []int, rank []int) int {
	p1, p2 := find(n1, parent), find(n2, parent)
	if p1 == p2 {
		return 0
	}
	if rank[p1] > rank[p2] {
		parent[p2] = p1
		rank[p1] += rank[p2]
	} else {
		parent[p1] = p2
		rank[p2] += rank[p1]
	}
	return 1
}
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)

	parent := make([]int, n)
	rank := make([]int, n)
	for i, _ := range isConnected {
		parent[i] = i
		rank[i] = 1
	}
	ans := n
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && isConnected[i][j] == 1 {

				ans -= unionV2(i, j, parent, rank)
			}
		}
	}
	return ans
}
func main() {
	nums := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	// target := 5
	fmt.Println(findCircleNum(nums))
}
