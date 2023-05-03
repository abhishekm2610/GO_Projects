package main

import (
	"fmt"
)

// func dfs(n int, currList []int, candidates []int, memo map[int][][]int) [][]int {
// 	if _, present := memo[n]; present {
// 		return memo[n]
// 	}

// 	for _, v := range candidates {

//			if n-v == 0 {
//				currList = append(currList, v)
//				fmt.Println(currList)
//				memo[n] = append(memo[n], currList)
//				// memo[n] = append(memo[n], []int{v})
//			} else if n-v > 0 {
//				currList = append(currList, v)
//				memo[n] = dfs((n - v), currList, candidates, memo)
//				// return
//				// memo[n] = append(memo[n], dfs((n-v), candidates, memo))
//			}
//		}
//		return memo[n]
//	}

func dfs(i int, cur []int, total int, target int, memo *[][]int, candidates []int) {
	if total == target {
		tmp := []int{}
		tmp = append(tmp, cur...)
		*memo = append(*memo, tmp)
		return
	}
	if i >= len(candidates) || total > target {
		return
	}
	dfs(i+1, cur, total, target, memo, candidates)

	cur = append(cur, candidates[i])
	dfs(i, cur, total+candidates[i], target, memo, candidates)
	// cur = cur[:len(cur)-1]

	return
}

func combinationSum(candidates []int, target int) [][]int {
	memo := [][]int{}

	dfs(0, []int{}, 0, target, &memo, candidates)
	return memo
}

func main() {
	// n := []int{0, 1, 0, 3, 2, 3}
	target := 7
	nums := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(nums, target))

}
