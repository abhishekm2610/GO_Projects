package main

import (
	"fmt"
)

func dfs(i int, cur []int, total int, target int, memo *[][]int, candidates []int) {
	if total > target {
		return
	}
	if total == target {
		tmp := []int{}
		tmp = append(tmp, cur...)
		fmt.Println(cur)
		*memo = append(*memo, tmp)
		return
	}

	for j := i; j < len(candidates); j++ {
		cur = append(cur, candidates[j])
		dfs(j, cur, total+candidates[j], target, memo, candidates)
		// cur = cur[:len(cur)-1]

		dfs(j+1, cur, total, target, memo, candidates)

	}

	// dfs(i+1, cur, total, target, memo, candidates)
	// cur = append(cur, candidates[i])
	// dfs(i, cur, total+candidates[i], target, memo, candidates)
	return
}

func combinationSum4(candidates []int, target int) int {
	memo := [][]int{}

	dfs(0, []int{}, 0, target, &memo, candidates)
	fmt.Println(memo)
	return len(memo)
}

func main() {
	target := 4
	nums := []int{1, 2, 3}
	fmt.Println(combinationSum4(nums, target))

}
