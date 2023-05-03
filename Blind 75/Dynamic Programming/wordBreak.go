package main

import (
	"fmt"
)

func wordBreakv1(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for v := range wordDict {
		dict[string(wordDict[v])] = true
	}
	dp := make(map[int]int)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = -1
	}
	dp[0] = 1

	for i := 0; i < len(s); i++ {
		sub := ""
		for j := i; j < len(s); j++ {
			sub += string(s[j])
			if _, isPresent := dict[sub]; isPresent {
				fmt.Println(j, i)
				if dp[j-len(sub)+1] == 1 {
					dp[j+1] = 1
					// break
				}
				// if dp[j] != 0 {
				// 	if dp[dp[j]] != 0 {
				// 		dp[j] = i

				// 	}
				// }
			}
		}
	}
	fmt.Println(dp)
	if dp[len(s)] == 1 {
		return true
	} else {
		// nextIndex := dp[len(s)-1]
		// for dp[nextIndex] > 0 {
		// 	nextIndex = dp[nextIndex-1]
		// }
		// if dp[nextIndex] == 0 {
		// 	return true
		// } else {
		return false
		// }
	}

}

func wordBreak(s string, wordDict []string) bool {
	dp := make(map[int]bool)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = false
	}
	dp[len(s)] = true
	for i := len(s) - 1; i >= 0; i-- {
		for _, word := range wordDict {
			if (i+len(word) <= len(s)) && s[i:i+len(word)] == word {
				dp[i] = dp[i+len(word)]
			}
			if dp[i] == true {
				break
			}
		}
	}
	return dp[0]
}
func main() {
	// n := []int{0, 1, 0, 3, 2, 3}
	s := "aaaaaaa"
	wordDict := []string{"aaaa", "aaa"}
	fmt.Println(wordBreak(s, wordDict))

}
