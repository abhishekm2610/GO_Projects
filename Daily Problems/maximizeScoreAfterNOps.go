package main

import (
	"fmt"
	"sort"
)

func maxInSlice(nums []int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
func minInSlice(nums []int) int {
	max := 0
	for _, v := range nums {
		if v < max {
			max = v
		}
	}
	return max
}

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
func circularGameLosers(n int, k int) []int {
	visited := make(map[int]int, n)
	whoHasBall := 1
	i := 1
	for i := 1; i <= n; i++ {
		visited[i] = 0
	}
	for true {

		fmt.Println(whoHasBall)
		visited[whoHasBall] += 1
		if visited[whoHasBall] == 2 {
			break
		}

		whoHasBall = whoHasBall + (i * k)
		i++
		if whoHasBall > n {
			if whoHasBall%n == 0 {
				// whoHasBall %= n

				fmt.Println(whoHasBall, n)

				whoHasBall = n
			} else {
				whoHasBall %= n

				// whoHasBall++

			}
		}

	}
	ans := []int{}
	for i, v := range visited {
		if v == 0 {
			ans = append(ans, i)

		}
	}
	sort.Ints(ans)
	fmt.Println(visited)
	return ans
}
func gcd(a int, b int) int {
	for a != b {
		if b < a {
			return gcd(a-b, b)
		} else {
			return gcd(a, b-a)
		}
	}
	return a
}

// func maxScore(nums []int) int {
// 	memo := make([]int, 1<<len(nums))
// 	fmt.Println(1<<len(nums), math.Pow(float64(2), float64(len(nums))))
// 	for i := 0; i < len(memo); i++ {
// 		memo[i] = -1
// 	}
// 	return backtrack(nums, 0, 0, memo)
// }

// func backtrack(nums []int, mask, pairsPicked int, memo []int) int {
// 	// If we have picked all the numbers from 'nums' array, we can't get more score
// 	if 2*pairsPicked == len(nums) {
// 		return 0
// 	}

// 	// If we already solved this sub-problem then return the stored result
// 	if memo[mask] != -1 {
// 		return memo[mask]
// 	}

// 	maxScore := 0

// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if ((mask>>i)&1) == 1 || ((mask>>j)&1) == 1 {
// 				fmt.Println(mask>>i, mask, i)
// 				continue
// 			}

// 			newMask := mask | (1 << i) | (1 << j)

// 			currScore := (pairsPicked + 1) * gcd(nums[i], nums[j])
// 			remainingScore := backtrack(nums, newMask, pairsPicked+1, memo)

// 			if currScore+remainingScore > maxScore {
// 				maxScore = currScore + remainingScore
// 			}
// 		}
// 	}

// 	memo[mask] = maxScore
// 	return maxScore
// }

func dfs(nums []int, mask int, pairs int, memo []int) int {
	if 2*pairs == len(nums) {
		return 0
	}
	if memo[mask] != -1 {
		return memo[mask]
	}
	maxScore := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (mask>>i)&1 == 1 || (mask>>j)&1 == 1 {
				continue
			}

			newmask := mask | (1 << i) | (1 << j)

			currScore := (pairs + 1) * gcd(nums[i], nums[j])
			remainScore := dfs(nums, newmask, pairs+1, memo)
			if currScore+remainScore > maxScore {
				maxScore = currScore + remainScore
			}
		}

	}
	memo[mask] = maxScore
	return maxScore
}
func maxScore(nums []int) int {
	memo := make([]int, 1<<len(nums))
	for i, _ := range memo {
		memo[i] = -1
	}
	mask := 0
	pairs := 0
	return dfs(nums, mask, pairs, memo)
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	// target := 5
	// k := 4
	fmt.Println(maxScore(nums))
}
