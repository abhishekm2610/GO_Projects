package main

import (
	"fmt"
	"math"
)

func longestCommonSubsequencev1(text1 string, text2 string) int {
	if text1 == text2 {
		return len(text1)
	}
	dp := make(map[string]int)
	// bits := int(math.Pow(float64(2), float64(len(text1)))) - 1

	for i := 1; i < (1 << len(text1)); i++ {
		fmt.Println("IN HERE")
		bits := i
		var sub string
		j := 0

		for bits > 0 {
			if (bits & 1) > 0 {
				sub += string(text1[j])
			}
			j++
			bits = bits >> 1

		}

		dp[sub] = len(sub)
	}

	max := 0
	for i := 1; i < (1 << len(text2)); i++ {
		bits := i
		var sub string
		j := 0

		for bits > 0 {
			if (bits & 1) > 0 {
				sub += string(text2[j])
			}
			j++
			bits = bits >> 1

		}
		if _, isPresent := dp[sub]; isPresent {
			max = int(math.Max(float64(max), float64(dp[sub])))

		}
	}

	// for i := 0; i < len(text1); i++ {
	// 	tmp := string(text1[i])
	// 	dp[tmp] = len(tmp)
	// 	for j := i + 1; j < len(text1); j++ {
	// 		tmp += string(text1[j])
	// 		dp[tmp] = len(tmp)
	// 	}
	// }
	// max := 0
	// for i := 0; i < len(text2); i++ {
	// 	tmp := string(text2[i])
	// 	dp[tmp] = len(tmp)

	// 	for j := i + 1; j < len(text2); j++ {
	// 		tmp += string(text2[j])
	// 		if _, isPresent := dp[tmp]; isPresent {
	// 			max = dp[tmp]
	// 		}
	// 	}
	// }

	// fmt.Println(dp)
	// return max
	fmt.Println(dp)
	return max
}

func longestCommonSubsequencev2(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for v := range dp {
		dp[v] = make([]int, len(text2)+1)
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i+1][j]), float64(dp[i][j+1])))
			}
		}
	}
	return dp[0][0]

}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for v := range dp {
		dp[v] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}
	return dp[len(text1)][len(text2)]

}

func main() {
	n1, n2 := "abcde", "ace"
	fmt.Println(longestCommonSubsequence(n1, n2))

}
