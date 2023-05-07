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
	Val       int
	Neighbors []*Node
}

func bfs() {

}

func cloneGraphv0(node *Node) *Node {
	if node == nil {
		return nil
	}

	queue := []*Node{node}
	visited := make(map[int]*Node)

	for len(queue) > 0 {
		curr := queue[0]
		if curr == nil {
			queue = queue[1:]
			continue
		}
		if _, present := visited[curr.Val]; !present {
			newNode := Node{Val: curr.Val, Neighbors: curr.Neighbors}
			visited[curr.Val] = &newNode
		}
		if len(curr.Neighbors) == 0 {
			queue = queue[1:]
			continue
		}

		queue = queue[1:]
		for _, nei := range curr.Neighbors {
			if _, present := visited[nei.Val]; !present {
				queue = append(queue, nei)
			}
		}

	}

	queue = []*Node{node}

	for _, v := range visited {
		curr := v
		newNei := []*Node{}
		for _, nei := range curr.Neighbors {
			newNei = append(newNei, visited[nei.Val])

		}
		visited[curr.Val].Neighbors = newNei

	}
	fmt.Println(visited)
	return visited[node.Val]
}
func dfs(node *Node, dp map[*Node]*Node) *Node {
	if node != nil {
		currNode := &Node{Val: node.Val}
		dp[node] = currNode
		for _, nei := range node.Neighbors {
			if _, present := dp[nei]; present {
				currNode.Neighbors = append(currNode.Neighbors, dp[nei])
			} else {
				currNode.Neighbors = append(currNode.Neighbors, dfs(nei, dp))
			}
		}
		return currNode

	} else {
		return nil

	}
}
func cloneGraph(node *Node) *Node {

	dp := make(map[*Node]*Node)

	dfs(node, dp)
	return dp[node]
}
