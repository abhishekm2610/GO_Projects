package main

import (
	"fmt"
	"log"
	"time"
)

func convertToBool(number int) int {
	output := 0

	for number != 0 {
		output = output << 1
		if (number & 1) == 1 {
			output = output | 1
		}
		number = number >> 1
	}
	return output
}
func main() {
	start := time.Now()
	intNumber := 12345678
	binaryOp := convertToBool(intNumber)
	fmt.Println(binaryOp)
	elapsed := time.Since(start)
	log.Printf("Runtime took %s", elapsed)

}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"
// )

// func convertToBool(number int) uint64 {

// 	var multipleFactor uint64
// 	multipleFactor = 0
// 	zeroFlag := false
// 	for number > 0 {
// 		if number%2 == 1 && zeroFlag == false {
// 			zeroFlag = true
// 		}
// 		if zeroFlag {
// 			multipleFactor = (uint64(2) * uint64(multipleFactor)) + uint64(number%2)
// 		}
// 		number /= 2
// 	}
// 	return multipleFactor
// }
// func main() {
// 	start := time.Now()
// 	intNumber := 12345678
// 	binaryOp := convertToBool(intNumber)
// 	fmt.Println(binaryOp)
// 	elapsed := time.Since(start)
// 	log.Printf("Runtime took %s", elapsed)

// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func convertToBool(number int) int {
// 	binaryAns := 0
// 	for number > 0 {
// 		binaryAns *= 10
// 		binaryAns += number % 2
// 		number /= 2
// 	}
// 	return binaryAns
// }
// func binaryToInt(binary int) int {
// 	intAns := 0
// 	multipleFactor := 0
// 	for binary > 0 {
// 		intAns += (binary % 10) * int(math.Pow(2, float64(multipleFactor)))
// 		binary /= 10
// 		multipleFactor += 1
// 	}
// 	return intAns
// }
// func main() {
// 	intNumber := 12345678
// 	binaryOp := convertToBool(intNumber)
// 	fmt.Println(binaryToInt(binaryOp))

// }

// package main

// import "fmt"

// func convertToBool(number int) int {
// 	binaryAns := 0
// 	mask := 1
// 	for number > 0 {
// 		SignificantNum := mask * (number % 2)
// 		SignificantNum += binaryAns
// 		binaryAns = SignificantNum
// 		number /= 2
// 		mask *= 10
// 	}
// 	return binaryAns
// }
// func main() {
// 	intNumber := 9090
// 	fmt.Println(convertToBool(intNumber))
// }
