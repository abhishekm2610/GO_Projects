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

func main() {
	// nums := []string{"5612624052M0130", "5378802576M6424", "5447619845F0171", "2941701174O9078"}
	target := 5
	k := 4
	fmt.Println(circularGameLosers(target, k))
}
