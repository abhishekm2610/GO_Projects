package main

import (
	"fmt"
)

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxVowels(s string, k int) int {
	max := 0
	last := 0
	for j, _ := range s[0:k] {
		if isVowel(s[j]) {
			max++
		}
	}
	count := max
	for i := k; i < len(s); i++ {
		if max == k {
			return max
		}
		if isVowel(s[last]) {
			count--
		}
		if isVowel(s[i]) {
			count++
		}
		last++
		max = maxInt(max, count)
	}
	return max
}

func main() {
	// target := 4
	s := "abciiidef"
	k := 3
	fmt.Println(maxVowels(s, k))

}
