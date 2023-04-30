package main

import (
	"fmt"
)

func reverseBitsv0(num uint32) uint32 {
	var ans uint32
	ans = 0
	bits := 1
	for bits <= 32 {
		ans = ans << 1

		ans += (num & 1)
		num = num >> 1
		bits++
	}
	return ans
}

func reverseBits(num uint32) uint32 {
	var ans uint32
	ans = 0
	i := 0
	for i < 32 {
		bit := num & 1
		ans = ans | (bit << (31 - i))
		num = num >> 1
		i++
	}
	return ans
}
func main() {
	var n uint32
	n = 4294967293
	fmt.Println(reverseBits(n))

}
