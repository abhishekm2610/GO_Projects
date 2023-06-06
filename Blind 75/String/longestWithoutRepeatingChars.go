package main

import "fmt"

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}
	slow := 0
	fast := 1
	n := len(s)
	cache := make(map[byte]bool)
	cache[s[0]] = true
	maxLen := 1
	for fast < n {
		if _, p := cache[s[fast]]; p {
			maxLen = maxInt(maxLen, len(cache))
			for s[slow] != s[fast] {
				delete(cache, s[slow])
				slow++
			}
			slow++
		} else {
			cache[s[fast]] = true
		}
		fast++
	}
	maxLen = maxInt(maxLen, len(cache))

	return maxLen

}

func main() {
	ip := "ba"
	fmt.Println(lengthOfLongestSubstring(ip))
}
