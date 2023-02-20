package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func checkPalindrome(digits int) bool {
	if digits < 0 {
		return false
	}
	exponent := math.Floor(math.Log10(float64(digits)))
	numOfDigits := exponent + 1
	mask := math.Pow(10, exponent)
	for i := 0; i < int(numOfDigits)/2; i++ {
		firstDigit := digits / int(mask)
		lastDigit := digits % 10
		if firstDigit != lastDigit {
			return false
		}
		digits = digits % int(mask)
		digits = digits / 10
		mask = mask / 100
	}

	return true
}
func main() {
	start := time.Now()

	digits := 12321
	fmt.Println(checkPalindrome(digits))
	elapsed := time.Since(start)
	log.Printf("Runtime took %s", elapsed)
}
