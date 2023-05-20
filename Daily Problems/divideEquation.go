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

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	dp := make(map[string]map[string]float64)
	for _, v := range equations {
		dp[v[0]] = make(map[string]float64)
		dp[v[1]] = make(map[string]float64)

	}
	for i, v := range equations {
		dp[v[0]][v[1]] = values[i]
		dp[v[1]][v[0]] = 1 / values[i]
	}

	var dfs func(s string, d string, visited map[string]bool) float64

	dfs = func(s string, d string, visited map[string]bool) float64 {

		_, sp := dp[s]
		_, desp := dp[d]

		if !sp || !desp {
			return -1
		}
		if len(dp[s]) == 0 {
			return -1.0
		}
		if _, p := dp[s][d]; p {
			return dp[s][d]
		}
		for i, _ := range dp[s] {
			fmt.Println(s, dp[s], i, d, visited)

			if !visited[i] {
				visited[i] = true
				tmp := dfs(i, d, visited)
				if tmp == -1.0 {
					continue
				} else {
					return tmp * dp[s][i]
				}
			}

		}

		return -1.0
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		s, d := q[0], q[1]

		visited := make(map[string]bool)
		ans[i] = dfs(s, d, visited)
		fmt.Println("----")

	}
	return ans
}

func main() {
	nums := [][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}}
	values := []float64{3.0, 4.0, 5.0, 6.0}
	q := [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}}

	fmt.Println(calcEquation(nums, values, q))
}
