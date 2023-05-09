package main

import "fmt"

type Node struct {
	up     *Node
	down   *Node
	left   *Node
	right  *Node
	height int
	goal   [2]bool
}

func dfsv2(node *Node, visited []*Node, tryingToReach int) bool {
	if node.goal[tryingToReach] == true {
		return true
	}
	for _, v := range visited {
		if v == node {
			return false
		}
	}
	visited = append(visited, node)
	if node.up != nil && node.up.height <= node.height {
		if dfsv2(node.up, visited, tryingToReach) {
			node.goal[tryingToReach] = true
		}
	}
	if node.goal[tryingToReach] == true {
		return true
	}
	if node.down != nil && node.down.height <= node.height {
		if dfsv2(node.down, visited, tryingToReach) {
			node.goal[tryingToReach] = true
		}
	}
	if node.goal[tryingToReach] == true {
		return true
	}
	if node.left != nil && node.left.height <= node.height {
		if dfsv2(node.left, visited, tryingToReach) {
			node.goal[tryingToReach] = true
		}
	}
	if node.goal[tryingToReach] == true {
		return true
	}
	if node.right != nil && node.right.height <= node.height {
		if dfsv2(node.right, visited, tryingToReach) {
			node.goal[tryingToReach] = true
		}
	}
	if node.goal[tryingToReach] != true {
		return false
	} else {
		return true
	}
}

func pacificAtlanticv2(heights [][]int) [][]int {
	ans := [][]int{}
	dp := make([][]*Node, len(heights))
	m := len(heights)
	n := len(heights[0])
	for i := 0; i < m; i++ {
		dp[i] = []*Node{}
		for j := 0; j < n; j++ {
			currNode := Node{}
			currNode.height = heights[i][j]
			if i == 0 || j == 0 {
				currNode.goal[1] = true
			}
			if i == m-1 {
				currNode.goal[0] = true
			}
			if j == n-1 {
				currNode.goal[0] = true
			}
			dp[i] = append(dp[i], &currNode)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			currNode := dp[i][j]
			if i == 0 {
				currNode.up = nil
			}
			if i != 0 {
				currNode.up = dp[i-1][j]
			}
			if i == m-1 {
				currNode.down = nil
			}
			if i < m-1 {
				currNode.down = dp[i+1][j]
			}
			if j == 0 {
				currNode.left = nil
			}
			if j != 0 {
				currNode.left = dp[i][j-1]
			}
			if j == n-1 {
				currNode.right = nil
			}
			if j < n-1 {
				currNode.right = dp[i][j+1]
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			currNode := dp[i][j]
			if currNode.goal[1] == true && currNode.goal[0] == true {
				ans = append(ans, []int{i, j})
				continue
			} else {
				visited := []*Node{}
				if dfsv2(currNode, visited, 0) {
					if dfsv2(currNode, visited, 1) {
						ans = append(ans, []int{i, j})
					}
				}
			}
		}
	}
	return ans
}

func dfs(r int, c int, ROWS int, COLS int, visited map[[2]int]bool, heights [][]int, prevHeight int) {
	if _, present := visited[[2]int{r, c}]; present || r < 0 || c < 0 || r == ROWS || c == COLS || heights[r][c] < prevHeight {
		return
	}

	visited[[2]int{r, c}] = true
	dfs(r-1, c, ROWS, COLS, visited, heights, heights[r][c])
	dfs(r+1, c, ROWS, COLS, visited, heights, heights[r][c])
	dfs(r, c-1, ROWS, COLS, visited, heights, heights[r][c])
	dfs(r, c+1, ROWS, COLS, visited, heights, heights[r][c])

}

func pacificAtlantic(heights [][]int) [][]int {
	ans := [][]int{}
	ROWS := len(heights)
	COLS := len(heights[0])
	atlantic := make(map[[2]int]bool)
	pacific := make(map[[2]int]bool)

	for r := 0; r < ROWS; r++ {
		dfs(r, 0, ROWS, COLS, pacific, heights, heights[r][0])
		dfs(r, COLS-1, ROWS, COLS, atlantic, heights, heights[r][COLS-1])
	}
	for c := 0; c < COLS; c++ {
		dfs(0, c, ROWS, COLS, pacific, heights, heights[0][c])
		dfs(ROWS-1, c, ROWS, COLS, atlantic, heights, heights[ROWS-1][c])
	}
	for v, _ := range pacific {
		if atlantic[v] == true {
			ans = append(ans, []int{v[0], v[1]})
		}
	}
	return ans
}

func main() {
	nums := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
	fmt.Println(pacificAtlantic(nums))
}
