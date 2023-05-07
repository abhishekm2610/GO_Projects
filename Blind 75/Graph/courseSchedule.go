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

type Node struct {
	Val    int
	PreReq *Node
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	dp := make(map[int]*Node)

	for i := 0; i <= numCourses; i++ {
		newNode := Node{Val: i, PreReq: nil}
		dp[i] = &newNode
	}

	for _, v := range prerequisites {
		if v[0] == v[1] {
			return false
		}
		dp[v[0]].PreReq = dp[v[1]]
	}
	visited := make(map[*Node]bool)
	for i, v := range dp {
		if v.PreReq == nil {
			visited = make(map[*Node]bool)
			continue
		}
		if _, present := visited[dp[i]]; present {
			return false
		} else {
			visited[dp[i]] = true
			visited[v.PreReq] = true
		}
	}
	return true
}

func main() {
	target := 10
	nums := [][]int{{5, 8}, {3, 5}, {1, 9}, {4, 5}, {0, 2}, {1, 9}, {7, 8}, {4, 9}}
	fmt.Println(canFinish(target, nums))
}
