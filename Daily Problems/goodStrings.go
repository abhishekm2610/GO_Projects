package main

import (
	"fmt"
	"math"
	"strings"
)

func dfs(currStr string, count *int, low int, high int, zero int, one int, memo map[string]int) int {
	if _, p := memo[currStr]; p {
		return memo[currStr]
	}
	if len(currStr) > high {
		// fmt.Println(currStr)
		return *count
	}

	if len(currStr) >= low && len(currStr) <= high {

		*count++

	}
	dfs(currStr+strings.Repeat("0", zero), count, low, high, zero, one, memo)
	dfs(currStr+strings.Repeat("1", one), count, low, high, zero, one, memo)
	*count %= ((int(math.Pow(float64(10), float64(9)))) + 7)
	memo[currStr] = *count
	return *count
}
func countGoodStrings(low int, high int, zero int, one int) int {
	// dp := []int{}
	// highCount := 0
	// lowCount := 0
	// if low != high {
	// 	highCount = int(math.Pow(float64(2), float64(high)))
	// 	lowCount = int(math.Pow(float64(2), float64(low)))
	// } else {
	// 	highCount = int(math.Pow(float64(2), float64(high)))
	// }
	ans := ""
	count := 0
	memo := make(map[string]int)

	// numberOfCombs := highCount - lowCount

	return dfs(ans, &count, low, high, zero, one, memo) % ((int(math.Pow(float64(10), float64(9)))) + 7)
}
func main() {
	low := 200
	high := 200
	zero := 10
	one := 1
	fmt.Println(countGoodStrings(low, high, zero, one))
}
