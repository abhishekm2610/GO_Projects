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

// func longestConsecutive(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	sort.Ints(nums)
// 	ans := 1
// 	count := 1
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i]-nums[i-1] == 0 {
// 			continue
// 		}
// 		if nums[i]-nums[i-1] == 1 {
// 			count++
// 			ans = max(ans, count)
// 		} else {
// 			count = 1
// 		}
// 	}
// 	return ans
// }
// func main() {
// 	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
// 	fmt.Println(longestConsecutive(nums))
// }

//Using Graph - sucks

// type Node struct {
// 	val  int
// 	next *Node
// 	prev *Node
// }

// func longestConsecutive(nums []int) int {
// 	dp := make(map[int]*Node)

// 	for _, v := range nums {
// 		if _, present := dp[v]; !present {
// 			newNode := Node{}
// 			newNode.val = v
// 			newNode.next = nil
// 			newNode.prev = nil
// 			dp[v] = &newNode
// 		}
// 	}

// 	for _, v := range nums {
// 		if _, present := dp[v-1]; present {
// 			dp[v].prev = dp[v-1]
// 			dp[v-1].next = dp[v]
// 		}
// 		if _, present := dp[v+1]; present {
// 			dp[v].next = dp[v+1]
// 			dp[v+1].prev = dp[v]
// 		}
// 	}
// 	visited := make(map[*Node]int)
// 	maxVal := 0

// 	for _, v := range dp {
// 		if _, present := visited[v]; present {
// 			continue
// 		} else if v.prev == nil {
// 			count := 0
// 			currNode := v
// 			for currNode != nil {
// 				count++
// 				visited[currNode] = 1
// 				currNode = currNode.next
// 			}
// 			maxVal = max(maxVal, count)
// 		}

// 	}
// 	return maxVal
// }

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}
	maxVal := 0
	for i, v := range nums {
		if !set[v-1] {
			curr := 1
			count := 1
			for set[v+curr] {
				count++
				curr++
			}
			maxVal = max(maxVal, count)
		}
		if len(nums)-i < maxVal {
			break
		}
	}
	return maxVal
}

func main() {
	nums := []int{0, 1}
	fmt.Println(longestConsecutive(nums))
}
