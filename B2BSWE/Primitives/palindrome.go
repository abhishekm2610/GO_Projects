package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func palindromeChecker(op1 int) bool {
	if op1 < 0 {
		return false
	}
	tempVar := 0
	original := op1

	for op1 > 10 {
		digit := op1 % 10
		tempVar = (tempVar * 10) + digit
		op1 = op1 / 10
	}
	tempVar = (tempVar * 10) + op1
	fmt.Println(original, tempVar)
	if tempVar == original {
		return true
	}
	return false
}

func main() {
	fmt.Print("Enter number to check if its palindrome: ")
	readerObj := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := readerObj.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\r\n")
	intVar, err := strconv.Atoi(input)
	fmt.Println(palindromeChecker(intVar))
}
