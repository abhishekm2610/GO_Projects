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
func checkAdjacent(nums []int) int {
	curr := nums[0]
	count := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] == 0 {
			curr = nums[j]
			continue
		}
		if nums[j] == curr {
			count++
		} else {
			curr = nums[j]
		}
	}
	return count
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func minIncrements(n int, cost []int) int {
	dp := make(map[int]*Node)
	for i := 1; i <= n; i++ {
		tmpNode := Node{Val: cost[i-1], Right: nil, Left: nil}
		dp[i] = &tmpNode
	}

	for i, v := range dp {
		if (2*i)+1 <= n {
			v.Left = dp[2*i]
			v.Right = dp[2*i+1]
		}
	}
	moves := 0
	for i, _ := range dp {
		if dp[i].Left == nil {
			continue
		}
		fmt.Println(dp[i].Val, dp[i].Left.Val, dp[i].Right.Val)

		moves += max(dp[i].Left.Val, dp[i].Right.Val) - min(dp[i].Left.Val, dp[i].Right.Val)
	}
	return moves
}
func main() {
	// target := 4
	nums := 15
	queries := []int{764, 1460, 2664, 764, 2725, 4556, 5305, 8829, 5064, 5929, 7660, 6321, 4830, 7055, 3761}

	fmt.Println(minIncrements(nums, queries))
}
