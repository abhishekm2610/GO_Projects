package main

import (
	"fmt"
	"sort"
)

type KthLargest struct {
	minheap []int
	k       int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	var obj KthLargest
	obj.minheap = nums

	obj.k = k

	for len(obj.minheap) > k {
		obj.minheap = obj.minheap[1:]
	}
	return obj
}

func (this *KthLargest) Add(val int) int {
	this.minheap = push(this.minheap, val)
	if len(this.minheap) > this.k {
		this.minheap = this.minheap[1:]
	}
	return this.minheap[0]
}

func push(queue []int, newQueue int) []int {
	index := len(queue)
	for index > 0 && queue[index-1] > newQueue {
		index--
	}
	//first
	if index == 0 {
		queue = append([]int{newQueue}, queue...)
		return queue
	}
	//last
	if index == len(queue) {
		queue = append(queue, newQueue)
		return queue
	}
	//not first and last index
	queue = append(queue[:index+1], queue[index:]...)
	queue[index] = newQueue
	return queue
}

func main() {

	obj := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
	fmt.Println(obj.Add(4))

	// nums := [][]int{{21, 5}, {92, 3}, {74, 2}, {39, 4}, {58, 2}, {5, 5}, {49, 4}, {65, 3}}
	// fmt.Println(mostPoints(nums))
}
