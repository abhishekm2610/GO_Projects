package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	maxProd := nums[0]
	minProd := nums[0]
	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		tmp := maxProd

		maxProd = int(math.Max(float64(nums[i]), float64(maxProd*nums[i])))
		maxProd = int(math.Max(float64(maxProd), float64(minProd*nums[i])))
		minProd = int(math.Min(float64(nums[i]), float64(minProd*nums[i])))
		minProd = int(math.Min(float64(minProd), float64(tmp*nums[i])))
		ans = int(math.Max(float64(ans), float64(maxProd)))
		// maxProd = int(math.Max(float64(maxProd), float64(totalProd/nums[j])))
		// maxProd = int(math.Max(float64(maxProd), float64(totalProd/(nums[i]*nums[j]))))
	}
	return ans
}

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))

}
