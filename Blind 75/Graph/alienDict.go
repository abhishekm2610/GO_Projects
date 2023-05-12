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
func dfs(c string, ans *[]string, visited map[string]bool, dp map[string]map[string]bool) bool {

	if _, p := visited[c]; p {
		return visited[c]
	}
	visited[c] = true

	for nei, _ := range dp[c] {
		fmt.Println(nei)

		if dfs(nei, ans, visited, dp) {
			return true
		}
	}
	visited[c] = false
	*ans = append(*ans, c)
	return false
}
func alienOrder(words []string) string {
	dp := make(map[string]map[string]bool)
	for i := 0; i < len(words)-1; i++ {
		s1 := words[i]
		s2 := words[i+1]
		minLen := min(len(s1), len(s2))

		if s1[:minLen] == s2[:minLen] && len(s1) > len(s2) {
			return ""
		}
		for j := 0; j < minLen; j++ {
			if s1[j] != s2[j] {
				if len(dp[string(s1[j])]) == 0 {
					dp[string(s1[j])] = make(map[string]bool)

				}
				dp[string(s1[j])][string(s2[j])] = true
				break
			}
		}
	}
	visited := map[string]bool{}
	res := []string{}

	for c, _ := range dp {
		if dfs(c, &res, visited, dp) {
			return ""
		}
	}
	ans := ""
	for i := len(res) - 1; i >= 0; i-- {
		ans += res[i]
	}
	// var ans string
	// for i, _ := range dp {
	// 	ans += dfs(i, ans, visited, dp)
	// }
	fmt.Println(res)
	return ans
}
func main() {
	nums := []string{"wrt", "wrf", "er", "ett", "rftt"}
	fmt.Println(alienOrder(nums))
}
