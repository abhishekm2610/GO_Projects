package main

import (
	"fmt"
)

func predictPartyVictory(senate string) string {
	dp := make([]int, len(senate))
	dp2 := make(map[string]int)
	for i, v := range senate {
		dp[i] = 1
		dp2[string(v)]++

	}
	shutdownR := 0
	shutdownD := 0
	for dp2["D"] > 0 && dp2["R"] > 0 {

		for i, v := range senate {

			if string(v) == "R" && dp[i] == 1 {
				if shutdownR > 0 {
					dp[i] = 0
					shutdownR--
					continue
				}
				shutdownD++
				dp2["D"]--

			}
			if string(v) == "D" && dp[i] == 1 {
				if shutdownD > 0 {
					dp[i] = 0
					shutdownD--
					continue
				}

				shutdownR++
				dp2["R"]--

			}
		}

	}
	if dp2["D"] <= 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}
func main() {
	// target := 4
	nums := "DDRRR"
	fmt.Println(predictPartyVictory(nums))

}
