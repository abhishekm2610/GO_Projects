package main

import (
	"fmt"
	"sort"
)

// func threeSum(nums []int) [][]int {
// 	ans := [][]int{}
// 	// seen := make(map[[3]int]int)
// 	sort.Ints(nums)
// 	hashMap := make(map[int][]int)

// 	for i, v := range nums {
// 		if _, isPresent := hashMap[v]; isPresent {
// 			hashMap[v] = append(hashMap[v], i)

// 		} else {
// 			hashMap[v] = []int{i}
// 		}

// 	}
// 	fmt.Println(hashMap)
// 	for i := 0; i < len(nums); i++ {
// 		used := make(map[int]int)
// 		for j := i + 1; j < len(nums); j++ {
// 			new := nums[j]
// 			sum := nums[i] + new
// 			if _, isPresent := hashMap[-1*sum]; isPresent {
// 				ans = append(ans, []int{nums[i], nums[j], -1 * sum})
// 				used[new] = 1

// 			} else {
// 				continue
// 			}

// 		}
// 	}

// 	// for i := 0; i < len(nums); i++ {
// 	// 	for j := i + 1; j < len(nums); j++ {
// 	// 		for k := j + 1; k < len(nums); k++ {
// 	// 			if i != j && j != k && k != i && nums[i]+nums[j]+nums[k] == 0 {
// 	// 				tmp := [3]int{nums[i], nums[j], nums[k]}
// 	// 				if _, isPresent := seen[tmp]; isPresent {
// 	// 					continue
// 	// 				} else {
// 	// 					seen[tmp] = 1
// 	// 					ans = append(ans, []int{nums[i], nums[j], nums[k]})
// 	// 				}
// 	// 			}
// 	// 		}
// 	// 	}
// 	// }
// 	return ans
// }

// func threeSum(nums []int) [][]int {
// 	ans := [][]int{}
// 	sort.Ints(nums)
// 	k := len(nums) - 1
// 	i := 0
// 	j := i + 1
// 	visited := make(map[int]int)
// 	for i < len(nums)-1 {
// 		if index, isPresent := visited[nums[i]]; isPresent && i != index {
// 			i++
// 			j = i + 1
// 			continue

// 		}

// 		if nums[j]+nums[k] == -1*nums[i] {
// 			if i != j && j != k && k != i {

// 				ans = append(ans, []int{nums[i], nums[j], nums[k]})
// 				visited[nums[i]] = i
// 				fmt.Println(i, j, k)

// 				currK := nums[k]
// 				for k > j {
// 					k--
// 					if nums[k] != currK {
// 						break
// 					}
// 				}

// 				currJ := nums[j]
// 				for j < k {
// 					j++
// 					if nums[j] != currJ {
// 						break
// 					}
// 				}
// 			}
// 		} else if nums[j]+nums[k] > -1*nums[i] {
// 			// k--
// 			currK := nums[k]
// 			for k > j {
// 				k--
// 				if nums[k] != currK {
// 					break
// 				}
// 			}

// 		} else {
// 			// j++
// 			currJ := nums[j]
// 			for j < k {
// 				j++
// 				if nums[j] != currJ {
// 					break
// 				}
// 			}
// 		}
// 		if j >= k {

// 			i++
// 			j = i + 1
// 			k = len(nums) - 1

// 			continue
// 		}

// 	}

// 	return ans
// }

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		r := len(nums) - 1
		l := i + 1

		for l < r {
			threeSum := nums[i] + nums[l] + nums[r]

			if threeSum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			} else if threeSum > 0 {
				r--
			} else {
				l++
			}

		}

	}

	return ans
}

func main() {
	nums := []int{-2, -3, 0, 0, -2}
	fmt.Println(threeSum(nums))

}
