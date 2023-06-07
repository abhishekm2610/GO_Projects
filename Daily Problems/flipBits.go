package main

import "fmt"

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minFlips(a int, b int, c int) int {

	bitMask := 1
	// maxLen := maxInt(a, b)
	flips := 0
	for a > 0 || b > 0 {

		ai := a & bitMask
		bi := b & bitMask
		ci := c & bitMask
		// fmt.Println(ai, bi, ci, carry)
		if ci == 0 {
			flips += ai + bi
		} else {
			if ai == bi {
				flips += 1
			}
		}
		a = a >> 1
		b = b >> 1
		c = c >> 1

		// bitMask = bitMask << 1

	}
	return flips
}

func main() {

	a := 1
	b := 2
	c := 3

	fmt.Println(minFlips(a, b, c))

}
