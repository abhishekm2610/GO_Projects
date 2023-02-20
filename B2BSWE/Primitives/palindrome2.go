package main

import (
	"fmt"
	"log"
	"time"
)

func checkPalindrome(inputStr string) bool {
	lengthOfStr := len(inputStr)
	endingPtr := lengthOfStr - 1
	startingPtr := 0

	for startingPtr <= endingPtr {
		if inputStr[startingPtr] != inputStr[endingPtr] {
			return false
		}
		startingPtr++
		endingPtr--
	}
	return true

}
func main() {
	start := time.Now()
	inputStr := "12321"
	fmt.Println(checkPalindrome(inputStr))
	elapsed := time.Since(start)
	log.Printf("Runtime took %s", elapsed)
}
