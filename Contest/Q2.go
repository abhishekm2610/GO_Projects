package main

import (
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

type FrequencyTracker struct {
	memo map[int]int
}

func Constructor() FrequencyTracker {
	var obj FrequencyTracker
	obj.memo = make(map[int]int)
	return obj
}

func (this *FrequencyTracker) Add(number int) {
	if _, present := this.memo[number]; present {
		this.memo[number]++
	} else {
		this.memo[number] = 1
	}
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if _, present := this.memo[number]; present {
		this.memo[number]--
		if this.memo[number] <0{
			this.memo[number] =0
		}
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	for _, v := range this.memo {
		if v == frequency {
			return true
		}
	}
	return false
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
func main() {
	// target := 4
	nums := []int{3}
	fmt.Println(distinctDifferenceArray(nums))
}


["FrequencyTracker","deleteOne","deleteOne","deleteOne","deleteOne","hasFrequency","add","deleteOne","hasFrequency","deleteOne","deleteOne","deleteOne","hasFrequency","add","hasFrequency","deleteOne","hasFrequency","hasFrequency","add","hasFrequency","add","deleteOne","hasFrequency","add","hasFrequency","hasFrequency","add"]
[[],[9],[8],[1],[4],[1],[10],[5],[1],[10],[9],[10],[1],[4],[1],[4],[1],[1],[10],[1],[2],[1],[2],[7],[1],[1],[6]]


[null,null,null,null,null,false,null,null,true,null,null,null,false,null,true,null,false,false,null,false,null,null,false,null,true,true,null]

[null,null,null,null,null,false,null,null,true,null,null,null,false,null,true,null,false,false,null,true,null,null,false,null,true,true,null]