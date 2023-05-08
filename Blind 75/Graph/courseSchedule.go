// package main

// import (
// 	"fmt"
// )

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// type Node struct {
// 	Val    int
// 	PreReq []*Node
// }

// func dfs(visited map[int]bool, dep int, dp map[int][]int) bool {
// 	if _, present := visited[dep]; present {
// 		return false
// 	} else {
// 		if len(dp[dep]) == 0 {
// 			return true
// 		}
// 		visited[dep] = true

// 		for _, dependancy := range dp[dep] {
// 			if !dfs(visited, dependancy, dp) {
// 				return false
// 			}
// 		}
// 		delete(visited, dep)
// 		dp[dep] = []int{}
// 		return true
// 	}
// }
// func canFinish(numCourses int, prerequisites [][]int) bool {
// 	dp := make(map[int][]int)

// 	for i := 0; i <= numCourses; i++ {
// 		dp[i] = []int{}
// 	}

// 	for _, v := range prerequisites {
// 		if v[0] == v[1] {
// 			return false
// 		}
// 		dp[v[0]] = append(dp[v[0]], v[1])
// 	}
// 	visited := make(map[int]bool)

// 	for i := 0; i < numCourses; i++ {
// 		if !dfs(visited, i, dp) {
// 			return false
// 		}

// 	}
// 	return true
// }

// func main() {
// 	target := 10
// 	nums := [][]int{{5, 8}, {3, 5}, {1, 9}, {4, 5}, {0, 2}, {1, 9}, {7, 8}, {4, 9}}
// 	fmt.Println(canFinish(target, nums))
// }

package main

import "fmt"

func dfs(course int, visited map[int]bool, coursesMapping map[int][]int) bool {
	if len(coursesMapping[course]) == 0 {
		return true
	}
	if _, present := visited[course]; present {
		return false
	}
	visited[course] = true
	for _, dependancy := range coursesMapping[course] {
		return dfs(dependancy, visited, coursesMapping)
	}

	return true

}
func canFinish(numCourses int, prerequisites [][]int) bool {
	coursesMapping := make(map[int][]int)

	for i := 0; i < numCourses; i++ {
		coursesMapping[i] = []int{}
	}

	for _, mapping := range prerequisites {
		coursesMapping[mapping[0]] = append(coursesMapping[mapping[0]], mapping[1])
	}

	for i := 0; i < numCourses; i++ {
		visited := make(map[int]bool)

		if dfs(i, visited, coursesMapping) == false {
			return false
		}
	}
	fmt.Println(coursesMapping)
	return true
}

func main() {
	target := 10
	nums := [][]int{{5, 8}, {3, 5}, {1, 9}, {4, 5}, {0, 2}, {1, 9}, {7, 8}, {4, 9}}
	fmt.Println(canFinish(target, nums))
}
